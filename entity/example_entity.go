package entity

type ExampleEntity struct {
	AbstractEntity
	Code      string
	Name      string
	Address1  string `gorm:"column:address_1"`
	ForeignID int64
}

type ExampleDetailEntity struct {
	ExampleEntity
}

func (ExampleEntity) TableName() string {
	return "example_entity"
}
