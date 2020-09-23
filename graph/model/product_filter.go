package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/nekishdev/graphql-example-catalog/gorm_models"
	"github.com/nekishdev/graphql-example-catalog/helpers"
)

var (
	ConditionOperators = []string{"=", ">", ">=", "<", "<=", "<>"}
	ConditionBy        = []string{"Field", "Property"}
)

type ProductFilterCondition struct {
	By       string      `json:"By"`
	Field    string      `json:"Field"`
	Operator string      `json:"Operator"`
	Value    interface{} `json:"Value"`
}

func (condition ProductFilterCondition) Validate() error {
	if !helpers.Contains(condition.By, ConditionBy) {
		return fmt.Errorf("invalid condition value in \"By\"")
	}

	if !helpers.Contains(condition.Operator, ConditionOperators) {
		return fmt.Errorf("invalid condition operator")
	}

	return nil
}

type ProductFilter struct {
	Conditions []ProductFilterCondition
}

func (filter ProductFilter) PrepareForDB(dbq *gorm.DB) (*gorm.DB, error) {

	var (
		err   error
		_dbq  = dbq
		scope = dbq.NewScope(gorm_models.Product{})
	)

	for _, condition := range filter.Conditions {
		if err = condition.Validate(); err != nil {
			return nil, fmt.Errorf("condition validation error: %s", err)
		}
		switch condition.By {
		case "Field":
			field, ok := scope.FieldByName(condition.Field)
			if !ok {
				return nil, fmt.Errorf("could not found field \"%s\"", condition.Field)
			}
			where := fmt.Sprintf("%s %s ?", field.DBName, condition.Operator)
			_dbq = _dbq.Where(where, condition.Value)
			break
		case "Property":
			pvAlias := "pv_" + condition.Field
			_dbq = _dbq.
				Joins(fmt.Sprintf(
					"inner join %s %s on (products.ID = %s.product_id and %s.property_code = '%s')",
					"product_property_values",
					pvAlias,
					pvAlias,
					pvAlias,
					condition.Field,
				)).
				Where(fmt.Sprintf("%s.value = ?", pvAlias), condition.Value)
			break
		}
	}

	return _dbq, nil
}
