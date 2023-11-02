package regional_service

import (
	"encoding/csv"
	"fmt"
	"go-master-data/dto"
	"go-master-data/entity"
	"go-master-data/entity/regional_entity"
	"go-master-data/model"
	"go-master-data/repository/regional_repository"
	"io"
	"os"
	"time"
)

type urbanVillageServiceImpl struct {
	UrbanVillageRepo regional_repository.UrbanVillageRepository
	SubDistrictRepo  regional_repository.SubDistrictRepository
}

func NewUrbanVillageService(urbanVillageRepo regional_repository.UrbanVillageRepository, subDistrictRepo regional_repository.SubDistrictRepository) UrbanVillageService {
	return &urbanVillageServiceImpl{
		UrbanVillageRepo: urbanVillageRepo,
		SubDistrictRepo:  subDistrictRepo,
	}
}
func (service *urbanVillageServiceImpl) Insert(request dto.UrbanVillageRequest) (response dto.UrbanVillageResponse, errMdl model.ErrorModel) {

	return
}

func (service *urbanVillageServiceImpl) Import(pathFile string) (errMdl model.ErrorModel) {
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
		parent, errMdl = service.SubDistrictRepo.GetByCode(parentCode)
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

		errMdl = service.UrbanVillageRepo.Insert(&repoInsert)
		if errMdl.Error != nil {
			return
		}
	}

	return
}
