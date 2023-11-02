package regional_repository

import (
	"database/sql"
	"errors"
	"go-master-data/entity/regional_entity"
	"go-master-data/model"
	"gorm.io/gorm"
)

func NewUrbanVillageRepository(db *gorm.DB) UrbanVillageRepository {
	return &urbanVillageRepositoryImpl{Db: db}
}

type urbanVillageRepositoryImpl struct {
	Db *gorm.DB
}

func (repo *urbanVillageRepositoryImpl) Insert(urbanVillage *regional_entity.UrbanVillage) (errMdl model.ErrorModel) {
	err := repo.Db.Create(urbanVillage).Error
	if err != nil {
		errMdl = model.GenerateUnknownError(err)
	}
	return
}

func (repo *urbanVillageRepositoryImpl) GetByCode(code string) (result regional_entity.UrbanVillage, errMdl model.ErrorModel) {
	err := repo.Db.Where("code = ?", code).Find(&result).Error
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		errMdl = model.GenerateUnknownError(err)
	}
	return
}
