package regional_repository

import (
	"go-master-data/entity/regional_entity"
	"go-master-data/model"
)

type SubDistrictRepository interface {
	Insert(district *regional_entity.SubDistrict) model.ErrorModel
	GetByCode(code string) (regional_entity.SubDistrict, model.ErrorModel)
	//Update(district regional_entity.SubDistrict) model.ErrorModel
	//View(district regional_entity.SubDistrict) (regional_entity.SubDistrict, model.ErrorModel)
}
