package regional_repository

import (
	"database/sql"
	"errors"
	"go-master-data/entity/regional_entity"
	"go-master-data/model"
	"gorm.io/gorm"
)

func NewSubDistrictRepository(db *gorm.DB) SubDistrictRepository {
	return &subDistrictRepositoryImpl{Db: db}
}

type subDistrictRepositoryImpl struct {
	Db *gorm.DB
}

func (repo *subDistrictRepositoryImpl) Insert(subDistrict *regional_entity.SubDistrict) (errMdl model.ErrorModel) {
	err := repo.Db.Create(subDistrict).Error
	if err != nil {
		errMdl = model.GenerateUnknownError(err)
	}
	return
}

func (repo *subDistrictRepositoryImpl) GetByCode(code string) (result regional_entity.SubDistrict, errMdl model.ErrorModel) {
	err := repo.Db.Where("code = ?", code).Find(&result).Error
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		errMdl = model.GenerateUnknownError(err)
	}
	return
}
