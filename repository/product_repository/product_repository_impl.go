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

type productRepositoryImpl struct {
	Db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepositoryImpl{Db: db}
}

func (repo *productRepositoryImpl) Insert(cp *product_entity.ProductEntity) (errMdl model.ErrorModel) {
	err := repo.Db.Create(cp).Error
	if err != nil {
		errMdl = model.GenerateUnknownError(err)
	}
	return
}

func (repo *productRepositoryImpl) Update(cp *product_entity.ProductEntity) (errMdl model.ErrorModel) {
	err := repo.Db.Save(cp).Error
	if err != nil {
		errMdl = model.GenerateUnknownError(err)
	}
	return
}

func (repo *productRepositoryImpl) List(dtoList dto.GetListRequest, searchParam []dto.SearchByParam, ctxModel *common.ContextModel) (result []interface{}, errMdl model.ErrorModel) {
	for i := 0; i < len(searchParam); i++ {
		searchParam[i].SearchKey = "p." + searchParam[i].SearchKey
	}
	searchParam = append(searchParam, dto.SearchByParam{
		SearchKey:      "p.company_id",
		SearchOperator: "eq",
		SearchValue:    strconv.Itoa(int(ctxModel.AuthAccessTokenModel.CompanyID)),
	})
	dtoList.OrderBy = "p." + dtoList.OrderBy
	query := "SELECT p.id, p.code, p.name, " +
		"p.created_at, p.updated_at, " +
		"cd.id, cd.code, cd.name " +
		" FROM product p " +
		"LEFT JOIN company_division cd ON cd.id = p.division_id "

	return repository.GetListDataDefault(repo.Db, query, nil, dtoList, searchParam,
		func(rows *sql.Rows) (interface{}, error) {
			var temp product_entity.ProductDetailEntity
			err := rows.Scan(&temp.ID, &temp.Code, &temp.Name,
				&temp.CreatedAt, &temp.UpdatedAt,
				&temp.DivisionID, &temp.DivisionCode, &temp.DivisionName)
			return temp, err
		})

}

func (repo *productRepositoryImpl) Count(searchParam []dto.SearchByParam, ctxModel *common.ContextModel) (result int64, errMdl model.ErrorModel) {
	searchParam = append(searchParam, dto.SearchByParam{
		SearchKey:      "p.company_id",
		SearchOperator: "eq",
		SearchValue:    strconv.Itoa(int(ctxModel.AuthAccessTokenModel.CompanyID)),
	})
	query := "SELECT COUNT(0) FROM product p "

	return repository.GetCountDataDefault(repo.Db, query, nil, searchParam)

}
func (repo *productRepositoryImpl) View(id int64) (result product_entity.ProductDetailEntity, errMdl model.ErrorModel) {
	query := "SELECT p.id, p.code, p.name, " +
		"p.created_at, p.updated_at, " +
		"p.selling_price, p.buying_price, p.uom_1, " +
		"p.uom_2, p.conv_1_to_2, " +
		"cd.id, cd.code, cd.name, " +
		"pc.id, pc.code, pc.name, " +
		"pg.id, pg.code, pg.name " +
		"FROM product p " +
		"LEFT JOIN company_division cd ON p.division_id = cd.id " +
		"LEFT JOIN product_category pc ON p.category_id = pc.id " +
		"LEFT JOIN product_group_hierarchy pg ON p.group_id = pg.id " +
		"WHERE p.id = $1 "

	err := repo.Db.Raw(query, id).Row().Scan(
		&result.ID, &result.Code, &result.Name,
		&result.CreatedAt, &result.UpdatedAt,
		&result.SellingPrice, &result.BuyingPrice, &result.Uom1,
		&result.Uom2, &result.Conv1To2,
		&result.DivisionID, &result.DivisionCode, &result.DivisionName,
		&result.CategoryID, &result.CategoryCode, &result.CategoryName,
		&result.GroupID, &result.GroupCode, &result.GroupName)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		errMdl = model.GenerateUnknownError(err)
		return
	}

	return
}

func (repo *productRepositoryImpl) FetchData(entity product_entity.ProductEntity) (result product_entity.ProductEntity, errMdl model.ErrorModel) {
	err := repo.Db.Where(&entity).First(&result).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		errMdl = model.GenerateInternalDBServerError(err)
	}
	return
}
