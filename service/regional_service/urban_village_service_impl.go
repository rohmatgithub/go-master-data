package regional_service

import (
	"encoding/csv"
	"fmt"
	"go-master-data/constanta"
	"go-master-data/dto"
	"go-master-data/dto/regional_dto"
	"go-master-data/entity"
	"go-master-data/entity/regional_entity"
	"go-master-data/model"
	"go-master-data/repository/regional_repository"
	"go-master-data/service"
	"io"
	"os"
	"strconv"
	"time"
)

type urbanVillageServiceImpl struct {
	UrbanVillageRepo regional_repository.UrbanVillageRepository
	SubDistrictRepo  regional_repository.SubDistrictRepository
}

func NewUrbanVillageService(subDistrictRepo regional_repository.SubDistrictRepository, urbanVillageRepo regional_repository.UrbanVillageRepository) UrbanVillageService {
	return &urbanVillageServiceImpl{
		UrbanVillageRepo: urbanVillageRepo,
		SubDistrictRepo:  subDistrictRepo,
	}
}
func (uv *urbanVillageServiceImpl) Insert(request regional_dto.UrbanVillageRequest) (response regional_dto.UrbanVillageResponse, errMdl model.ErrorModel) {

	return
}

func (uv *urbanVillageServiceImpl) Import(pathFile string) (errMdl model.ErrorModel) {
	// open file
	f, err := os.Open(pathFile)
	if err != nil {
		errMdl = model.GenerateUnknownError(err)
		return
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			errMdl = model.GenerateUnknownError(err)
			return
		}
		// do something with read line
		fmt.Printf("%+v\n", rec)
		code := rec[0]
		parentCode := rec[1]
		name := rec[2]
		if code == "code" {
			continue
		}

		var parent regional_entity.SubDistrict
		parent, errMdl = uv.SubDistrictRepo.GetByCode(parentCode)
		if errMdl.Error != nil {
			return
		}

		// insert
		timeNow := time.Now()
		repoInsert := regional_entity.UrbanVillage{
			ParentID: parent.ID,
			Code:     code,
			Name:     name,
			AbstractEntity: entity.AbstractEntity{
				CreatedBy: 1,
				UpdatedBy: 1,
				CreatedAt: timeNow,
				UpdatedAt: timeNow,
				Deleted:   false,
			},
		}

		errMdl = uv.UrbanVillageRepo.Insert(&repoInsert)
		if errMdl.Error != nil {
			return
		}
	}

	return
}

func (uv *urbanVillageServiceImpl) List(dtoList dto.GetListRequest, searchParam []dto.SearchByParam) (out dto.Payload, errMdl model.ErrorModel) {
	parentID := 0
	for _, param := range searchParam {
		if param.SearchKey == "parent_id" {
			parentID, _ = strconv.Atoi(param.SearchValue)
			break
		}
	}
	if parentID == 0 {
		errMdl = model.GenerateEmptyFieldError(constanta.ParentID)
		return
	}
	resultDB, errMdl := uv.UrbanVillageRepo.List(dtoList, searchParam)
	if errMdl.Error != nil {
		return
	}

	var result []regional_dto.UrbanVillageListResponse
	for _, temp := range resultDB {
		data := temp.(regional_entity.UrbanVillage)
		result = append(result, regional_dto.UrbanVillageListResponse{
			ID:       data.ID,
			ParentID: data.ParentID,
			Code:     data.Code,
			Name:     data.Name,
		})
	}

	out.Data = result

	out.Status.Message = service.ListI18NMessage(constanta.LanguageEn)
	return
}
