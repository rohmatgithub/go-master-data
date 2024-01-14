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

type companyRepositoryImpl struct {
	Db *gorm.DB
}

func NewCompanyRepository(db *gorm.DB) CompanyRepository {
	return &companyRepositoryImpl{Db: db}
}

func (repo *companyRepositoryImpl) Insert(cp *admin_entity.CompanyEntity) (errMdl model.ErrorModel) {
	err := repo.Db.Create(cp).Error
	if err != nil {
		errMdl = model.GenerateUnknownError(err)
	}
	return
}

func (repo *companyRepositoryImpl) Update(cp *admin_entity.CompanyEntity) (errMdl model.ErrorModel) {
	err := repo.Db.Save(cp).Error
	if err != nil {
		errMdl = model.GenerateUnknownError(err)
	}
	return
}

func (repo *companyRepositoryImpl) List(dtoList dto.GetListRequest, searchParam []dto.SearchByParam) (result []interface{}, errMdl model.ErrorModel) {
	query := "SELECT c.id, c.code, cp.name, cp.address_1, " +
		"c.created_at, c.updated_at " +
		"FROM company c " +
		"LEFT JOIN company_profile cp ON c.company_profile_id = cp.id "

	return repository.GetListDataDefault(repo.Db, query, nil, dtoList, searchParam,
		func(rows *sql.Rows) (interface{}, error) {
			var temp admin_entity.CompanyDetailEntity
			err := rows.Scan(&temp.ID, &temp.Code, &temp.CompanyProfile.Name, &temp.CompanyProfile.Address1,
				&temp.CreatedAt, &temp.UpdatedAt)
			return temp, err
		})

}

func (repo *companyRepositoryImpl) Count(searchParam []dto.SearchByParam) (result int64, errMdl model.ErrorModel) {
	query := "SELECT COUNT(0) FROM company "

	return repository.GetCountDataDefault(repo.Db, query, nil, searchParam)

}
func (repo *companyRepositoryImpl) View(id int64) (result admin_entity.CompanyDetailEntity, errMdl model.ErrorModel) {
	query := "SELECT c.id, c.code, cp.name, cp.address_1, cp.npwp, " +
		"c.created_at, cp.updated_at, c.created_by, c.updated_by, " +
		"c.company_profile_id " +
		"FROM company c " +
		"LEFT JOIN company_profile cp ON c.company_profile_id = cp.id " +
		"WHERE c.id = $1 "

	err := repo.Db.Raw(query, id).Row().Scan(
		&result.ID, &result.Code, &result.CompanyProfile.Name, &result.CompanyProfile.Address1, &result.CompanyProfile.NPWP,
		&result.CreatedAt, &result.UpdatedAt, &result.CreatedBy, &result.UpdatedBy,
		&result.CompanyProfileID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		errMdl = model.GenerateUnknownError(err)
		return
	}

	return
}

func (repo *companyRepositoryImpl) FetchData(entity admin_entity.CompanyEntity) (result admin_entity.CompanyEntity, errMdl model.ErrorModel) {
	err := repo.Db.Where(&entity).First(&result).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		errMdl = model.GenerateInternalDBServerError(err)
	}
	return
}
