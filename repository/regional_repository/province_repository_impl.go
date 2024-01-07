package regional_repository

import (
	"database/sql"
	"go-master-data/dto"
	"go-master-data/entity/regional_entity"
	"go-master-data/model"
	"go-master-data/repository"
	"gorm.io/gorm"
)

func NewProvinceRepository(db *gorm.DB) ProvinceRepository {
	return &provinceRepositoryImpl{Db: db}
}

type provinceRepositoryImpl struct {
	Db *gorm.DB
}

func (repo *provinceRepositoryImpl) List(dtoList dto.GetListRequest, searchParam []dto.SearchByParam) (result []interface{}, errMdl model.ErrorModel) {
	query := "SELECT id, code, name FROM province "

	return repository.GetListDataDefault(repo.Db, query, nil, dtoList, searchParam,
		func(rows *sql.Rows) (interface{}, error) {
			var temp regional_entity.Province
			err := rows.Scan(&temp.ID, &temp.Code, &temp.Name)
			return temp, err
		})

}
