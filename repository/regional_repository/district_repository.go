package regional_repository

import (
	"go-master-data/entity/regional_entity"
	"go-master-data/model"
)

type DistrictRepository interface {
	Insert(district *regional_entity.District) model.ErrorModel
	GetByCode(code string) (regional_entity.District, model.ErrorModel)
	//Update(district regional_entity.District) model.ErrorModel
	//View(district regional_entity.District) (regional_entity.SubDistrict, model.ErrorModel)
}
