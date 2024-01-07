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

type companyProfileRepositoryImpl struct {
	Db *gorm.DB
}

func NewCompanyProfileRepository(db *gorm.DB) CompanyProfileRepository {
	return &companyProfileRepositoryImpl{Db: db}
}

func (repo *companyProfileRepositoryImpl) Insert(cp *admin_entity.CompanyProfileEntity) (errMdl model.ErrorModel) {
	err := repo.Db.Create(cp).Error
	if err != nil {
		errMdl = model.GenerateUnknownError(err)
	}
	return
}

func (repo *companyProfileRepositoryImpl) Update(cp *admin_entity.CompanyProfileEntity) (errMdl model.ErrorModel) {
	err := repo.Db.Save(cp).Error
	if err != nil {
		errMdl = model.GenerateUnknownError(err)
	}
	return
}

func (repo *companyProfileRepositoryImpl) List(dtoList dto.GetListRequest, searchParam []dto.SearchByParam) (result []interface{}, errMdl model.ErrorModel) {
	query := "SELECT id, npwp, name, address_1, created_at, updated_at FROM company_profile "

	return repository.GetListDataDefault(repo.Db, query, nil, dtoList, searchParam,
		func(rows *sql.Rows) (interface{}, error) {
			var temp admin_entity.CompanyProfileEntity
			err := rows.Scan(&temp.ID, &temp.NPWP, &temp.Name, &temp.Address1,
				&temp.CreatedAt, &temp.UpdatedAt)
			return temp, err
		})

}

func (repo *companyProfileRepositoryImpl) Count(searchParam []dto.SearchByParam) (result int64, errMdl model.ErrorModel) {
	query := "SELECT COUNT(0) FROM company_profile "

	return repository.GetCountDataDefault(repo.Db, query, nil, searchParam)

}

func (repo *companyProfileRepositoryImpl) View(id int64) (result admin_entity.CompanyProfileDetailEntity, errMdl model.ErrorModel) {
	query := "SELECT cp.id, cp.npwp, cp.name, cp.address_1, " +
		"cp.address_2, c.id, c.code, c.name, " +
		"d.id, d.code, d.name, " +
		"sd.id, sd.code, sd.name, " +
		"uv.id, uv.code, uv.name, " +
		"cp.created_at, cp.updated_at, cp.created_by, cp.updated_by " +
		//"uc.name, up.name " +
		"FROM company_profile cp " +
		"LEFT JOIN country c ON c.id = cp.country_id " +
		"LEFT JOIN district d ON d.id = cp.district_id " +
		"LEFT JOIN sub_district sd ON sd.id = cp.sub_district_id " +
		"LEFT JOIN urban_village uv ON uv.id = cp.urban_village_id " +
		"WHERE cp.id = $1 "

	var (
		cID, dID, sdID, uvID         sql.NullInt64
		cCode, dCode, sdCode, uvCode sql.NullString
		cName, dName, sdName, uvName sql.NullString
	)
	err := repo.Db.Raw(query, id).Row().Scan(
		&result.ID, &result.NPWP, &result.Name, &result.Address1,
		&result.Address2, &cID, &cCode, &cName,
		&dID, &dCode, &dName,
		&sdID, &sdCode, &sdName,
		&uvID, &uvCode, &uvName,
		&result.CreatedAt, &result.UpdatedAt, &result.CreatedBy, &result.UpdatedBy)
	if err != nil {
		errMdl = model.GenerateUnknownError(err)
		return
	}

	result.CountryID = cID.Int64
	result.CountryCode = cCode.String
	result.CountryName = cName.String

	result.DistrictID = dID.Int64
	result.DistrictCode = dCode.String
	result.DistrictName = dName.String

	result.SubDistrictID = sdID.Int64
	result.SubDistrictCode = sdCode.String
	result.SubDistrictName = sdName.String

	result.UrbanVillageID = uvID.Int64
	result.UrbanVillageCode = uvCode.String
	result.UrbanVillageName = uvName.String
	return
}

func (repo *companyProfileRepositoryImpl) FetchData(entity admin_entity.CompanyProfileEntity) (result admin_entity.CompanyProfileEntity, errMdl model.ErrorModel) {
	err := repo.Db.Where(&entity).First(&result).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		errMdl = model.GenerateInternalDBServerError(err)
	}
	return
}
