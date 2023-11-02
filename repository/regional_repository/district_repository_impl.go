package regional_repository

import (
	"database/sql"
	"errors"
	"go-master-data/entity/regional_entity"
	"go-master-data/model"
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
