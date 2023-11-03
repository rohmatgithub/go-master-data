package entity

import (
	"time"
)

type AbstractEntity struct {
	ID        int64 `gorm:"primaryKey"`
	CreatedBy int64
	UpdatedBy int64
	CreatedAt time.Time
	UpdatedAt time.Time
	Deleted   bool
}

type AbstractEntitySql struct {
}
