package regional_entity

import "go-master-data/entity"

type UrbanVillage struct {
	ParentID int64
	Code     string
	Name     string
	entity.AbstractEntity
}

func (UrbanVillage) TableName() string {
	return "urban_village"
}
