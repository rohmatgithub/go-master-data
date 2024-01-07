package regional_entity

import "go-master-data/entity"

type Province struct {
	ID   int64
	Code string
	Name string
	entity.AbstractEntity
}

func (Province) TableName() string {
	return "province"
}
