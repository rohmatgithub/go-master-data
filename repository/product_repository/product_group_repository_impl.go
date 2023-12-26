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

type productGroupRepositoryImpl struct {
	Db *gorm.DB
}

func NewProductGroupRepository(db *gorm.DB) ProductGroupRepository {
	return &productGroupRepositoryImpl{Db: db}
}

func (repo *productGroupRepositoryImpl) Insert(cp *product_entity.ProductGroupEntity) (errMdl model.ErrorModel) {
	err := repo.Db.Create(cp).Error
	if err != nil {
		errMdl = model.GenerateUnknownError(err)
	}
	return
}

func (repo *productGroupRepositoryImpl) Update(cp *product_entity.ProductGroupEntity) (errMdl model.ErrorModel) {
	err := repo.Db.Save(cp).Error
	if err != nil {
		errMdl = model.GenerateUnknownError(err)
	}
	return
}

func (repo *productGroupRepositoryImpl) List(dtoList dto.GetListRequest, searchParam []dto.SearchByParam) (result []interface{}, errMdl model.ErrorModel) {
	query := "SELECT id, code, name, " +
		"level, parent_id FROM product_group "

	return repository.GetListDataDefault(repo.Db, query, nil, dtoList, searchParam,
		func(rows *sql.Rows) (interface{}, error) {
			var temp product_entity.ProductGroupEntity
			err := rows.Scan(&temp.ID, &temp.Code, &temp.Name,
				&temp.Level, &temp.ParentID)
			return temp, err
		})

}

func (repo *productGroupRepositoryImpl) View(id int64) (result product_entity.ProductGroupDetailEntity, errMdl model.ErrorModel) {
	query := "SELECT id, code, name, " +
		"level, parent_id, " +
		"created_at, updated_at FROM product_group WHERE id = $1 "

	err := repo.Db.Raw(query, id).Row().Scan(
		&result.ID, &result.Code, &result.Name,
		&result.Level, &result.Name,
		&result.CreatedAt, &result.UpdatedAt)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		errMdl = model.GenerateUnknownError(err)
		return
	}

	return
}

func (repo *productGroupRepositoryImpl) FetchData(entity product_entity.ProductGroupEntity) (result product_entity.ProductGroupEntity, errMdl model.ErrorModel) {
	err := repo.Db.Where(&entity).First(&result).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		errMdl = model.GenerateInternalDBServerError(err)
	}
	return
}
