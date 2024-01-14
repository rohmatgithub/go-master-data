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
	query := "SELECT cd.id, cd.code, cd.name, " +
		"cd.created_at, cd.updated_at " +
		"FROM company_division cd "

	return repository.GetListDataDefault(repo.Db, query, nil, dtoList, searchParam,
		func(rows *sql.Rows) (interface{}, error) {
			var temp admin_entity.CompanyDivisionEntity
			err := rows.Scan(&temp.ID, &temp.Code, &temp.Name,
				&temp.CreatedAt, &temp.UpdatedAt)
			return temp, err
		})

}

func (repo *companyDivisionRepositoryImpl) Count(searchParam []dto.SearchByParam) (result int64, errMdl model.ErrorModel) {
	query := "SELECT COUNT(0) FROM company_division "

	return repository.GetCountDataDefault(repo.Db, query, nil, searchParam)

}

func (repo *companyDivisionRepositoryImpl) View(id int64) (result admin_entity.CompanyDivisionDetailEntity, errMdl model.ErrorModel) {
	query := "SELECT cd.id, cd.code, cd.name, " +
		"cd.created_at, cd.updated_at, " +
		"cd.company_id, c.code, cp.name " +
		"FROM company_division cd " +
		"LEFT JOIN company c ON cd.company_id = c.id " +
		"LEFT JOIN company_profile cp ON c.company_profile_id = cp.id " +
		"WHERE cd.id = $1 "

	err := repo.Db.Raw(query, id).Row().Scan(
		&result.ID, &result.Code, &result.Name,
		&result.CreatedAt, &result.UpdatedAt,
		&result.CompanyID, &result.CompanyCode, &result.CompanyName)
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
