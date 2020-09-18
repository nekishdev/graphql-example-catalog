package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/nekishdev/graphql-example-catalog/db"
	"github.com/nekishdev/graphql-example-catalog/gorm_models"
	"github.com/nekishdev/graphql-example-catalog/graph/generated"
	"github.com/nekishdev/graphql-example-catalog/graph/model"
)

func (r *mutationResolver) CreateProduct(ctx context.Context, input model.CreateProductInput) (*gorm_models.Product, error) {
	product := &gorm_models.Product{
		Name:        input.Name,
		Description: input.Description,
		ImageSrc:    input.ImageSrc,
		Price:       input.Price,
	}
	if input.CategoryID != nil {
		product.CategoryID = *input.CategoryID
	}
	if input.Properties != nil {
		var productProperties []gorm_models.ProductProperty
		for _, pv := range input.Properties {
			if pv != nil {
				productProperties = append(productProperties, gorm_models.ProductProperty{
					PropertyID: pv.PropertyID,
					Value:      pv.Value,
				})
			}
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

func (r *mutationResolver) CreateProductProperty(ctx context.Context, input model.CreateProductProperty) (*gorm_models.Property, error) {
	property := &gorm_models.Property{
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

func (r *queryResolver) Products(ctx context.Context, categoryID *int, limit int, offset int) ([]*gorm_models.Product, error) {
	var products []*gorm_models.Product
	var _categoryId uint = 0
	if categoryID != nil {
		_categoryId = uint(*categoryID)
	}

	err := db.GetDB().
		Model(gorm_models.Product{}).
		Where(
			gorm_models.Product{
				CategoryID: _categoryId,
			}).
		Limit(limit).
		Offset(offset).
		Preload("Category").
		Preload("Properties.Property").
		Find(&products).
		Error

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

func (r *queryResolver) ProductProperties(ctx context.Context) ([]*gorm_models.Property, error) {
	var properties []*gorm_models.Property

	err := db.GetDB().
		Model(gorm_models.Property{}).
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
