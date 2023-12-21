package example_repository

import (
	"database/sql"
	"errors"
	"go-master-data/dto"
	"go-master-data/entity"
	"go-master-data/model"
	"go-master-data/repository"
	"gorm.io/gorm"
)

type exampleRepositoryImpl struct {
	Db *gorm.DB
}

func NewExampleRepository(db *gorm.DB) ExampleRepository {
	return &exampleRepositoryImpl{Db: db}
}

func (repo *exampleRepositoryImpl) Insert(cp *entity.ExampleEntity) (errMdl model.ErrorModel) {
	err := repo.Db.Create(cp).Error
	if err != nil {
		errMdl = model.GenerateUnknownError(err)
	}
	return
}

func (repo *exampleRepositoryImpl) Update(cp *entity.ExampleEntity) (errMdl model.ErrorModel) {
	err := repo.Db.Save(cp).Error
	if err != nil {
		errMdl = model.GenerateUnknownError(err)
	}
	return
}

func (repo *exampleRepositoryImpl) List(dtoList dto.GetListRequest, searchParam []dto.SearchByParam) (result []interface{}, errMdl model.ErrorModel) {
	query := "SELECT id, npwp, name, address_1 FROM example_entity "

	return repository.GetListDataDefault(repo.Db, query, nil, dtoList, searchParam,
		func(rows *sql.Rows) (interface{}, error) {
			var temp entity.ExampleEntity
			err := rows.Scan(&temp.ID, &temp.Code, &temp.Name, &temp.Address1)
			return temp, err
		})

}

func (repo *exampleRepositoryImpl) View(id int64) (result entity.ExampleDetailEntity, errMdl model.ErrorModel) {
	query := "SELECT * FROM example_entity WHERE cp.id = $1 "

	err := repo.Db.Raw(query, id).Row().Scan(
		&result.ID, &result.Code, &result.Name, &result.Address1,
		&result.CreatedAt, &result.UpdatedAt, &result.CreatedBy, &result.UpdatedBy)
	if err != nil {
		errMdl = model.GenerateUnknownError(err)
		return
	}

	return
}

func (repo *exampleRepositoryImpl) FetchData(entity entity.ExampleEntity) (result entity.ExampleEntity, errMdl model.ErrorModel) {
	err := repo.Db.Where(&entity).First(&result).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		errMdl = model.GenerateInternalDBServerError(err)
	}
	return
}
