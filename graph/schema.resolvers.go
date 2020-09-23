package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/nekishdev/graphql-example-catalog/db"
	"github.com/nekishdev/graphql-example-catalog/gorm_models"
	"github.com/nekishdev/graphql-example-catalog/graph/generated"
	"github.com/nekishdev/graphql-example-catalog/graph/model"
)

func (r *mutationResolver) CreateProduct(ctx context.Context, input model.ProductFillable) (*gorm_models.Product, error) {
	product := &gorm_models.Product{
		Name:        input.Name,
		Description: input.Description,
		ImageSrc:    input.ImageSrc,
		Price:       input.Price,
		CategoryID:  input.CategoryID,
	}
	if input.Properties != nil {
		var productProperties []gorm_models.ProductPropertyValue
		for _, pv := range input.Properties {
			productProperties = append(productProperties, gorm_models.ProductPropertyValue{
				PropertyCode: pv.Code,
				Value:        pv.Value,
			})
		}
		product.Properties = productProperties
	}
	err := db.GetDB().
		Create(&product).
		Error

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (r *mutationResolver) UpdateProduct(ctx context.Context, id uint, input model.ProductFillable) (*gorm_models.Product, error) {
	var (
		product          gorm_models.Product
		propertyValues   []gorm_models.ProductPropertyValue
		propertyValueMap map[string]string
		err              error
	)

	db.GetDB().LogMode(true)
	db.GetDB().Transaction(func(tx *gorm.DB) error {
		if err = tx.First(&product, id).Error; err != nil {
			return err
		}

		if err = db.GetDB().Model(&product).Update(input).Error; err != nil {
			return fmt.Errorf("failed to update product with id \"%d\": %s", id, err)
		}

		if err = db.GetDB().Model(&product).Association("Properties").Find(&propertyValues).Error; err != nil {
			return err
		}

		propertyValueMap = make(map[string]string)
		for _, pv := range propertyValues {
			propertyValueMap[pv.PropertyCode] = pv.Value
		}

		for _, pv := range input.Properties {
			_, exists := propertyValueMap[pv.Code]
			if exists {
				err = db.GetDB().
					Model(&gorm_models.ProductPropertyValue{}).
					Where(gorm_models.ProductPropertyValue{
						ProductID:    product.ID,
						PropertyCode: pv.Code,
					}).
					Update(
						gorm_models.ProductPropertyValue{
							Value: pv.Value,
						},
					).
					Error
				if err != nil {
					return fmt.Errorf(
						"failed to update property with code \"%s\" to product with id \"%d\". %s",
						pv.Code,
						product.ID,
						err.Error(),
					)
				}
			} else {
				err = db.GetDB().
					Model(&product).
					Association("Properties").
					Append(
						gorm_models.ProductPropertyValue{
							PropertyCode: pv.Code,
							Value:        pv.Value,
						},
					).
					Error
				if err != nil {
					return fmt.Errorf(
						"failed to append property with code \"%s\" to product with id \"%d\". %s",
						pv.Code,
						product.ID,
						err.Error(),
					)
				}
			}
		}

		return nil
	})

	return &product, nil
}

func (r *mutationResolver) DeleteProduct(ctx context.Context, id uint) (bool, error) {
	err := db.GetDB().Select("Properties").Delete(gorm_models.Product{ID: id}).Error
	if err != nil {
		return false, fmt.Errorf("failed to delete product with id \"%d\": %s", id, err)
	}
	return true, nil
}

func (r *mutationResolver) CreateCategory(ctx context.Context, input model.CreateCategoryInput) (*gorm_models.Category, error) {
	category := &gorm_models.Category{
		Name:        input.Name,
		Description: input.Description,
		ImageSrc:    input.ImageSrc,
	}

	err := db.GetDB().
		Create(&category).
		Error

	if err != nil {
		return nil, err
	}

	return category, nil
}

func (r *mutationResolver) CreateProductProperty(ctx context.Context, input model.CreateProductProperty) (*gorm_models.ProductProperty, error) {
	property := &gorm_models.ProductProperty{
		Name:     input.Name,
		Code:     input.Code,
		Required: input.Required,
	}

	err := db.GetDB().
		Create(&property).
		Error

	if err != nil {
		return nil, err
	}

	return property, nil
}

func (r *queryResolver) Products(ctx context.Context, limit int, offset int, filter *model.ProductFilter) ([]*gorm_models.Product, error) {
	var (
		err      error
		products []*gorm_models.Product
	)

	dbq := db.GetDB().
		LogMode(true).
		Model(gorm_models.Product{}).
		Limit(limit).
		Offset(offset).
		Preload("Category").
		Preload("Properties.Property")

	if filter != nil {
		if dbq, err = filter.PrepareForDB(dbq); err != nil {
			return nil, fmt.Errorf("failed to prepare filter: %s", err)
		}
	}

	err = dbq.Find(&products).Error

	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r *queryResolver) Product(ctx context.Context, id uint) (*gorm_models.Product, error) {
	var product gorm_models.Product

	err := db.GetDB().
		Preload("Category").
		Preload("Properties.Property").
		First(&product, id).
		Error

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *queryResolver) Categories(ctx context.Context, limit int, offset int) ([]*gorm_models.Category, error) {
	var categories []*gorm_models.Category

	err := db.GetDB().
		Model(gorm_models.Category{}).
		Limit(limit).
		Offset(offset).
		Find(&categories).
		Error

	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (r *queryResolver) Category(ctx context.Context, id uint) (*gorm_models.Category, error) {
	var category gorm_models.Category

	err := db.GetDB().First(&category, id).Error

	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *queryResolver) ProductProperties(ctx context.Context) ([]*gorm_models.ProductProperty, error) {
	var properties []*gorm_models.ProductProperty

	err := db.GetDB().
		Model(gorm_models.ProductProperty{}).
		Find(&properties).
		Error

	if err != nil {
		return nil, err
	}

	return properties, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
