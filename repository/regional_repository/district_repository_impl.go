package regional_repository

import (
	"database/sql"
	"errors"
	"go-master-data/dto"
	"go-master-data/entity/regional_entity"
	"go-master-data/model"
	"go-master-data/repository"
	"gorm.io/gorm"
)

func NewDistrictRepository(db *gorm.DB) DistrictRepository {
	return &districtRepositoryImpl{Db: db}
}

type districtRepositoryImpl struct {
	Db *gorm.DB
}

func (repo *districtRepositoryImpl) Insert(district *regional_entity.District) (errMdl model.ErrorModel) {
	err := repo.Db.Create(district).Error
	if err != nil {
		errMdl = model.GenerateUnknownError(err)
	}
	return
}

func (repo *districtRepositoryImpl) GetByCode(code string) (result regional_entity.District, errMdl model.ErrorModel) {
	err := repo.Db.Where("code = ?", code).Find(&result).Error
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		errMdl = model.GenerateUnknownError(err)
	}
	return
}

func (repo *districtRepositoryImpl) List(dtoList dto.GetListRequest, searchParam []dto.SearchByParam) (result []interface{}, errMdl model.ErrorModel) {
	query := "SELECT id, code, name, parent_id FROM district "

	return repository.GetListDataDefault(repo.Db, query, nil, dtoList, searchParam,
		func(rows *sql.Rows) (interface{}, error) {
			var temp regional_entity.District
			err := rows.Scan(&temp.ID, &temp.Code, &temp.Name, &temp.ParentID)
			return temp, err
		})

}
