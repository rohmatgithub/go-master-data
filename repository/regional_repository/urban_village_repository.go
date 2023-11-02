package regional_repository

import (
	"go-master-data/entity/regional_entity"
	"go-master-data/model"
)

type UrbanVillageRepository interface {
	Insert(urbanVillage *regional_entity.UrbanVillage) model.ErrorModel
	GetByCode(code string) (regional_entity.UrbanVillage, model.ErrorModel)
	//Update(urbanVillage regional_entity.UrbanVillage) model.ErrorModel
	//View(urbanVillage regional_entity.UrbanVillage) (regional_entity.UrbanVillage, model.ErrorModel)
}
