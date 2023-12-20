package regional_entity

import "go-master-data/entity"

type Country struct {
	ID   int64
	Code string
	Name string
	entity.AbstractEntity
}

func (Country) TableName() string {
	return "country"
}
