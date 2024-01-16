package product_repository

import (
	"database/sql"
	"errors"
	"go-master-data/common"
	"go-master-data/dto"
	"go-master-data/entity/product_entity"
	"go-master-data/model"
	"go-master-data/repository"
	"strconv"

	"gorm.io/gorm"
)

type productGroupRepositoryImpl struct {
	Db *gorm.DB
}

func NewProductGroupRepository(db *gorm.DB) ProductGroupRepository {
	return &productGroupRepositoryImpl{Db: db}
}

func (repo *productGroupRepositoryImpl) Insert(cp *product_entity.ProductGroupEntity) (errMdl model.ErrorModel) {
	err := repo.Db.Create(cp).Error
	if err != nil {
		errMdl = model.GenerateUnknownError(err)
	}
	return
}

func (repo *productGroupRepositoryImpl) Update(cp *product_entity.ProductGroupEntity) (errMdl model.ErrorModel) {
	err := repo.Db.Save(cp).Error
	if err != nil {
		errMdl = model.GenerateUnknownError(err)
	}
	return
}

func (repo *productGroupRepositoryImpl) List(dtoList dto.GetListRequest, searchParam []dto.SearchByParam, ctxModel *common.ContextModel) (result []interface{}, errMdl model.ErrorModel) {
	for i := 0; i < len(searchParam); i++ {
		searchParam[i].SearchKey = "pg." + searchParam[i].SearchKey
	}
	searchParam = append(searchParam, dto.SearchByParam{
		SearchKey:      "pg.company_id",
		SearchOperator: "eq",
		SearchValue:    strconv.Itoa(int(ctxModel.AuthAccessTokenModel.CompanyID)),
	})

	dtoList.OrderBy = "pg." + dtoList.OrderBy
	query := "SELECT pg.id, pg.code, pg.name, pg.level, " +
		"pg.created_at, pg.updated_at, pg.parent_id, " +
		"cd.id, cd.code, cd.name " +
		"FROM product_group_hierarchy pg " +
		"LEFT JOIN company_division cd ON pg.division_id = cd.id "

	return repository.GetListDataDefault(repo.Db, query, nil, dtoList, searchParam,
		func(rows *sql.Rows) (interface{}, error) {
			var temp product_entity.ProductGroupDetailEntity
			err := rows.Scan(&temp.ID, &temp.Code, &temp.Name, &temp.Level,
				&temp.CreatedAt, &temp.UpdatedAt, &temp.ParentID,
				&temp.DivisionID, &temp.DivisionCode, &temp.DivisionName)
			return temp, err
		})

}

func (repo *productGroupRepositoryImpl) Count(searchParam []dto.SearchByParam, ctxModel *common.ContextModel) (result int64, errMdl model.ErrorModel) {
	searchParam = append(searchParam, dto.SearchByParam{
		SearchKey:      "company_id",
		SearchOperator: "eq",
		SearchValue:    strconv.Itoa(int(ctxModel.AuthAccessTokenModel.CompanyID)),
	})
	query := "SELECT COUNT(0) FROM product_group_hierarchy "

	return repository.GetCountDataDefault(repo.Db, query, nil, searchParam)

}

func (repo *productGroupRepositoryImpl) View(id int64) (result product_entity.ProductGroupDetailEntity, errMdl model.ErrorModel) {
	query := "SELECT pg.id, pg.code, pg.name, pg.level, " +
		"pg.created_at, pg.updated_at, pg.parent_id, " +
		"parent.code, parent.name, " +
		"cd.id, cd.code, cd.name " +
		"FROM product_group_hierarchy pg " +
		"LEFT JOIN product_group_hierarchy parent ON pg.parent_id = parent.id " +
		"LEFT JOIN company_division cd ON pg.division_id = cd.id " +
		"WHERE pg.id = $1 "

	err := repo.Db.Raw(query, id).Row().Scan(
		&result.ID, &result.Code, &result.Name, &result.Level,
		&result.CreatedAt, &result.UpdatedAt, &result.ParentID,
		&result.ParentCode, &result.ParentName,
		&result.DivisionID, &result.DivisionCode, &result.DivisionName)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		errMdl = model.GenerateUnknownError(err)
		return
	}

	return
}

func (repo *productGroupRepositoryImpl) FetchData(entity product_entity.ProductGroupEntity) (result product_entity.ProductGroupEntity, errMdl model.ErrorModel) {
	err := repo.Db.Where(&entity).First(&result).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		errMdl = model.GenerateInternalDBServerError(err)
	}
	return
}
