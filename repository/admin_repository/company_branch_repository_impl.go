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

type companyBranchRepositoryImpl struct {
	Db *gorm.DB
}

func NewCompanyBranchRepository(db *gorm.DB) CompanyBranchRepository {
	return &companyBranchRepositoryImpl{Db: db}
}

func (repo *companyBranchRepositoryImpl) Insert(cp *admin_entity.CompanyBranchEntity) (errMdl model.ErrorModel) {
	err := repo.Db.Create(cp).Error
	if err != nil {
		errMdl = model.GenerateUnknownError(err)
	}
	return
}

func (repo *companyBranchRepositoryImpl) Update(cp *admin_entity.CompanyBranchEntity) (errMdl model.ErrorModel) {
	err := repo.Db.Save(cp).Error
	if err != nil {
		errMdl = model.GenerateUnknownError(err)
	}
	return
}

func (repo *companyBranchRepositoryImpl) List(dtoList dto.GetListRequest, searchParam []dto.SearchByParam) (result []interface{}, errMdl model.ErrorModel) {
	query := "SELECT cb.id, cb.code, cp.name, cp.address_1 " +
		"FROM company_branch cb " +
		"LEFT JOIN company_profile cp ON cb.company_profile_id = cp.id "

	return repository.GetListDataDefault(repo.Db, query, nil, dtoList, searchParam,
		func(rows *sql.Rows) (interface{}, error) {
			var temp admin_entity.CompanyBranchDetailEntity
			err := rows.Scan(&temp.ID, &temp.Code, &temp.CompanyProfile.Name, &temp.CompanyProfile.Address1)
			return temp, err
		})

}

func (repo *companyBranchRepositoryImpl) View(id int64) (result admin_entity.CompanyBranchDetailEntity, errMdl model.ErrorModel) {
	query := "SELECT cb.id, cb.code, cp.name, cp.address_1, " +
		"cb.created_at, cb.updated_at, cb.created_by, cb.updated_by, " +
		"c.id, c.code, ccp.name " +
		"FROM company_branch cb " +
		"LEFT JOIN company_profile cp ON cb.company_profile_id = cp.id " +
		"LEFT JOIN company c ON cb.company_id = c.id " +
		"LEFT JOIN company_profile ccp ON c.company_profile_id = ccp.id " +
		"WHERE cb.id = $1 "

	err := repo.Db.Raw(query, id).Row().Scan(
		&result.ID, &result.Code, &result.CompanyProfile.Name, &result.CompanyProfile.Address1,
		&result.CreatedAt, &result.UpdatedAt, &result.CreatedBy, &result.UpdatedBy,
		&result.CompanyID, &result.CompanyCode, &result.CompanyName)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		errMdl = model.GenerateUnknownError(err)
		return
	}

	return
}

func (repo *companyBranchRepositoryImpl) FetchData(entity admin_entity.CompanyBranchEntity) (result admin_entity.CompanyBranchEntity, errMdl model.ErrorModel) {
	err := repo.Db.Where(&entity).First(&result).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		errMdl = model.GenerateInternalDBServerError(err)
	}
	return
}
