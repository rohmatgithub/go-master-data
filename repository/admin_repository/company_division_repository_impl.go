package admin_repository

import (
	"database/sql"
	"errors"
	"go-master-data/dto"
	"go-master-data/entity/admin_entity"
	"go-master-data/model"
	"go-master-data/repository"
	"gorm.io/gorm"
)

type companyDivisionRepositoryImpl struct {
	Db *gorm.DB
}

func NewCompanyDivisionRepository(db *gorm.DB) CompanyDivisionRepository {
	return &companyDivisionRepositoryImpl{Db: db}
}

func (repo *companyDivisionRepositoryImpl) Insert(cp *admin_entity.CompanyDivisionEntity) (errMdl model.ErrorModel) {
	err := repo.Db.Create(cp).Error
	if err != nil {
		errMdl = model.GenerateUnknownError(err)
	}
	return
}

func (repo *companyDivisionRepositoryImpl) Update(cp *admin_entity.CompanyDivisionEntity) (errMdl model.ErrorModel) {
	err := repo.Db.Save(cp).Error
	if err != nil {
		errMdl = model.GenerateUnknownError(err)
	}
	return
}

func (repo *companyDivisionRepositoryImpl) List(dtoList dto.GetListRequest, searchParam []dto.SearchByParam) (result []interface{}, errMdl model.ErrorModel) {
	query := "SELECT id, code, name FROM company_division "

	return repository.GetListDataDefault(repo.Db, query, nil, dtoList, searchParam,
		func(rows *sql.Rows) (interface{}, error) {
			var temp admin_entity.CompanyDivisionEntity
			err := rows.Scan(&temp.ID, &temp.Code, &temp.Name)
			return temp, err
		})

}

func (repo *companyDivisionRepositoryImpl) View(id int64) (result admin_entity.CompanyDivisionDetailEntity, errMdl model.ErrorModel) {
	query := "SELECT id, code, name, " +
		"created_at, updated_at FROM company_division WHERE id = $1 "

	err := repo.Db.Raw(query, id).Row().Scan(
		&result.ID, &result.Code, &result.Name,
		&result.CreatedAt, &result.UpdatedAt)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		errMdl = model.GenerateUnknownError(err)
		return
	}

	return
}

func (repo *companyDivisionRepositoryImpl) FetchData(entity admin_entity.CompanyDivisionEntity) (result admin_entity.CompanyDivisionEntity, errMdl model.ErrorModel) {
	err := repo.Db.Where(&entity).First(&result).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		errMdl = model.GenerateInternalDBServerError(err)
	}
	return
}
