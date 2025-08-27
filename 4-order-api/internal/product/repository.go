package product

import (
	"go/courses/pkg/db"

	"gorm.io/gorm/clause"
)

type ProductRepository struct {
	Database *db.Db
}

type ProductRepositoryDeps struct {
	Database *db.Db
}

func NewProductRepository(deps ProductRepositoryDeps) *ProductRepository {
	return &ProductRepository{
		Database: deps.Database,
	}
}

func (repo *ProductRepository) GetAll(limit int, page int) ([]Product, error) {
	var products []Product

	result := repo.Database.DB.Limit(limit).Offset((page - 1) * limit).Find(&products)

	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}

func (repo *ProductRepository) GetById(id uint) (*Product, error) {
	var product Product

	result := repo.Database.DB.First(&product, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &product, nil
}

func (repo *ProductRepository) Create(product *Product) (*Product, error) {
	result := repo.Database.DB.Create(product)

	if result.Error != nil {
		return nil, result.Error
	}

	return product, nil
}

func (repo *ProductRepository) Update(product *Product) (*Product, error) {
	result := repo.Database.DB.Clauses(clause.Returning{}).Updates(product)

	if result.Error != nil {
		return nil, result.Error
	}

	return product, nil
}

func (repo *ProductRepository) Delete(id uint) error {
	result := repo.Database.DB.Delete(&Product{}, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
