package regional_entity

import (
	"go-master-data/entity"
)

type SubDistrict struct {
	ParentID int64
	Code     string
	Name     string
	entity.AbstractEntity
}

func (SubDistrict) TableName() string {
	return "sub_district"
}
