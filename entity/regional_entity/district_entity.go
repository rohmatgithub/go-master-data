package regional_entity

import (
	"go-master-data/entity"
)

type District struct {
	ID       int64
	ParentID int64
	Code     string
	Name     string
	entity.AbstractEntity
}

func (District) TableName() string {
	return "district"
}
