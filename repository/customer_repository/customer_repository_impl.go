package customer_repository

import (
	"database/sql"
	"errors"
	"go-master-data/common"
	"go-master-data/dto"
	"go-master-data/entity/customer_entity"
	"go-master-data/model"
	"go-master-data/repository"
	"strconv"

	"gorm.io/gorm"
)

type productRepositoryImpl struct {
	Db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &productRepositoryImpl{Db: db}
}

func (repo *productRepositoryImpl) Insert(cp *customer_entity.CustomerEntity) (errMdl model.ErrorModel) {
	err := repo.Db.Create(cp).Error
	if err != nil {
		errMdl = model.GenerateUnknownError(err)
	}
	return
}

func (repo *productRepositoryImpl) Update(cp *customer_entity.CustomerEntity) (errMdl model.ErrorModel) {
	err := repo.Db.Save(cp).Error
	if err != nil {
		errMdl = model.GenerateUnknownError(err)
	}
	return
}

func (repo *productRepositoryImpl) List(dtoList dto.GetListRequest, searchParam []dto.SearchByParam, ctxModel *common.ContextModel) (result []interface{}, errMdl model.ErrorModel) {
	for i := 0; i < len(searchParam); i++ {
		searchParam[i].SearchKey = "c." + searchParam[i].SearchKey
	}
	searchParam = append(searchParam, dto.SearchByParam{
		SearchKey:      "c.company_id",
		SearchOperator: "eq",
		SearchValue:    strconv.Itoa(int(ctxModel.AuthAccessTokenModel.CompanyID)),
	}, dto.SearchByParam{
		SearchKey:      "c.branch_id",
		SearchOperator: "eq",
		SearchValue:    strconv.Itoa(int(ctxModel.AuthAccessTokenModel.BranchID)),
	})
	dtoList.OrderBy = "c." + dtoList.OrderBy
	query := "SELECT c.id, c.code, c.name, " +
		"c.phone, c.created_at, c.updated_at " +
		" FROM customer c "

	return repository.GetListDataDefault(repo.Db, query, nil, dtoList, searchParam,
		func(rows *sql.Rows) (interface{}, error) {
			var temp customer_entity.CustomerEntity
			err := rows.Scan(&temp.ID, &temp.Code, &temp.Name,
				&temp.Phone, &temp.CreatedAt, &temp.UpdatedAt)
			return temp, err
		})

}

func (repo *productRepositoryImpl) Count(searchParam []dto.SearchByParam, ctxModel *common.ContextModel) (result int64, errMdl model.ErrorModel) {
	searchParam = append(searchParam, dto.SearchByParam{
		SearchKey:      "c.company_id",
		SearchOperator: "eq",
		SearchValue:    strconv.Itoa(int(ctxModel.AuthAccessTokenModel.CompanyID)),
	}, dto.SearchByParam{
		SearchKey:      "c.branch_id",
		SearchOperator: "eq",
		SearchValue:    strconv.Itoa(int(ctxModel.AuthAccessTokenModel.BranchID)),
	})
	query := "SELECT COUNT(0) FROM customer c "

	return repository.GetCountDataDefault(repo.Db, query, nil, searchParam)

}
func (repo *productRepositoryImpl) View(id int64) (result customer_entity.CustomerDetailEntity, errMdl model.ErrorModel) {
	query := "SELECT cust.id, cust.code, cust.name, " +
		"cust.phone, cust.email, cust.address, " +
		"cust.created_at, cust.updated_at, " +
		"c.id, c.code, c.name, " +
		"p.id, p.code, p.name, " +
		"d.id, d.code, d.name, " +
		"sd.id, sd.code, sd.name, " +
		"uv.id, uv.code, uv.name " +
		"FROM customer cust " +
		"LEFT JOIN country c ON cust.country_id = c.id " +
		"LEFT JOIN province p ON cust.province_id = p.id " +
		"LEFT JOIN district d ON cust.district_id = d.id " +
		"LEFT JOIN sub_district sd ON cust.sub_district_id = sd.id " +
		"LEFT JOIN urban_village uv ON cust.urban_village_id = uv.id " +
		"WHERE cust.id = $1 "

	err := repo.Db.Raw(query, id).Row().Scan(
		&result.ID, &result.Code, &result.Name,
		&result.Phone, &result.Email, &result.Address,
		&result.CreatedAt, &result.UpdatedAt,
		&result.CountryID, &result.CountryCode, &result.CountryName,
		&result.ProvinceID, &result.ProvinceCode, &result.ProvinceName,
		&result.DistrictID, &result.DistrictCode, &result.DistrictName,
		&result.SubDistrictID, &result.SubDistrictCode, &result.SubDistrictName,
		&result.UrbanVillageID, &result.UrbanVillageCode, &result.UrbanVillageName,
	)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		errMdl = model.GenerateUnknownError(err)
		return
	}

	return
}

func (repo *productRepositoryImpl) FetchData(entity customer_entity.CustomerEntity) (result customer_entity.CustomerEntity, errMdl model.ErrorModel) {
	err := repo.Db.Where(&entity).First(&result).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		errMdl = model.GenerateInternalDBServerError(err)
	}
	return
}
