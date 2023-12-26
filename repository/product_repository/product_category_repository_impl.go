package product_repository

import (
	"database/sql"
	"errors"
	"go-master-data/dto"
	"go-master-data/entity/product_entity"
	"go-master-data/model"
	"go-master-data/repository"
	"gorm.io/gorm"
)

type productCategoryRepositoryImpl struct {
	Db *gorm.DB
}

func NewProductCategoryRepository(db *gorm.DB) ProductCategoryRepository {
	return &productCategoryRepositoryImpl{Db: db}
}

func (repo *productCategoryRepositoryImpl) Insert(cp *product_entity.ProductCategoryEntity) (errMdl model.ErrorModel) {
	err := repo.Db.Create(cp).Error
	if err != nil {
		errMdl = model.GenerateUnknownError(err)
	}
	return
}

func (repo *productCategoryRepositoryImpl) Update(cp *product_entity.ProductCategoryEntity) (errMdl model.ErrorModel) {
	err := repo.Db.Save(cp).Error
	if err != nil {
		errMdl = model.GenerateUnknownError(err)
	}
	return
}

func (repo *productCategoryRepositoryImpl) List(dtoList dto.GetListRequest, searchParam []dto.SearchByParam) (result []interface{}, errMdl model.ErrorModel) {
	query := "SELECT id, code, name FROM product_category "

	return repository.GetListDataDefault(repo.Db, query, nil, dtoList, searchParam,
		func(rows *sql.Rows) (interface{}, error) {
			var temp product_entity.ProductCategoryEntity
			err := rows.Scan(&temp.ID, &temp.Code, &temp.Name)
			return temp, err
		})

}

func (repo *productCategoryRepositoryImpl) View(id int64) (result product_entity.ProductCategoryDetailEntity, errMdl model.ErrorModel) {
	query := "SELECT id, code, name, " +
		"created_at, updated_at FROM product_category WHERE id = $1 "

	err := repo.Db.Raw(query, id).Row().Scan(
		&result.ID, &result.Code, &result.Name,
		&result.CreatedAt, &result.UpdatedAt)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		errMdl = model.GenerateUnknownError(err)
		return
	}

	return
}

func (repo *productCategoryRepositoryImpl) FetchData(entity product_entity.ProductCategoryEntity) (result product_entity.ProductCategoryEntity, errMdl model.ErrorModel) {
	err := repo.Db.Where(&entity).First(&result).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		errMdl = model.GenerateInternalDBServerError(err)
	}
	return
}
