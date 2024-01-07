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

type subDistrictServiceImpl struct {
	SubDistrictRepo regional_repository.SubDistrictRepository
	DistrictRepo    regional_repository.DistrictRepository
}

func NewSubDistrictService(districtRepo regional_repository.DistrictRepository, subDistrictRepo regional_repository.SubDistrictRepository) SubDistrictService {
	return &subDistrictServiceImpl{
		SubDistrictRepo: subDistrictRepo,
		DistrictRepo:    districtRepo,
	}
}
func (sd *subDistrictServiceImpl) Insert(request regional_dto.SubDistrictRequest) (response regional_dto.SubDistrictResponse, errMdl model.ErrorModel) {

	return
}

func (sd *subDistrictServiceImpl) Import(pathFile string) (errMdl model.ErrorModel) {
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

		var parent regional_entity.District
		parent, errMdl = sd.DistrictRepo.GetByCode(parentCode)
		if errMdl.Error != nil {
			return
		}

		// insert
		timeNow := time.Now()
		repoInsert := regional_entity.SubDistrict{
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

		errMdl = sd.SubDistrictRepo.Insert(&repoInsert)
		if errMdl.Error != nil {
			return
		}
	}

	return
}

func (sd *subDistrictServiceImpl) List(dtoList dto.GetListRequest, searchParam []dto.SearchByParam) (out dto.Payload, errMdl model.ErrorModel) {
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
	resultDB, errMdl := sd.SubDistrictRepo.List(dtoList, searchParam)
	if errMdl.Error != nil {
		return
	}

	var result []regional_dto.SubDistrictListResponse
	for _, temp := range resultDB {
		data := temp.(regional_entity.SubDistrict)
		result = append(result, regional_dto.SubDistrictListResponse{
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
