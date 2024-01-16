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
	query := "SELECT pc.id, pc.code, pc.name, pc.created_at, pc.updated_at FROM product_category pc "

	return repository.GetListDataDefault(repo.Db, query, nil, dtoList, searchParam,
		func(rows *sql.Rows) (interface{}, error) {
			var temp product_entity.ProductCategoryEntity
			err := rows.Scan(&temp.ID, &temp.Code, &temp.Name, &temp.CreatedAt, &temp.UpdatedAt)
			return temp, err
		})

}

func (repo *productCategoryRepositoryImpl) Count(searchParam []dto.SearchByParam) (result int64, errMdl model.ErrorModel) {
	query := "SELECT COUNT(0) FROM product_category "

	return repository.GetCountDataDefault(repo.Db, query, nil, searchParam)

}
func (repo *productCategoryRepositoryImpl) View(id int64) (result product_entity.ProductCategoryDetailEntity, errMdl model.ErrorModel) {
	query := "SELECT pc.id, pc.code, pc.name, " +
		"pc.created_at, pc.updated_at, " +
		"cd.id, cd.code, cd.name " +
		"FROM product_category pc " +
		"LEFT JOIN company_division cd ON pc.division_id = cd.id " +
		"WHERE pc.id = $1 "

	err := repo.Db.Raw(query, id).Row().Scan(
		&result.ID, &result.Code, &result.Name,
		&result.CreatedAt, &result.UpdatedAt,
		&result.DivisionID, &result.DivisionCode, &result.DivisionName)
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
