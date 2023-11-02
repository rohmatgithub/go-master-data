package test

import (
	"go-master-data/common"
	"go-master-data/repository/regional_repository"
	"go-master-data/service/regional_service"
	"testing"
)

type structInterface interface {
	Find(str string) string
}

type structImpl1 struct {
}

func (s structImpl1) Find(str string) string {
	return "struct impl 1 : " + str
}

type structImpl2 struct {
}

func (s structImpl2) Find(str string) string {
	return "struct impl 1 : " + str
}

func TestImportCsv(t *testing.T) {
	db, err := common.ConnectDB("user=postgres password=root dbname=microservice sslmode=disable host=localhost port=5432", "master_data")
	if err != nil {
		t.Fatal(err)
	}

	//repo := regional_repository.NewSubDistrictRepository(db)
	//districtRepo := regional_repository.NewDistrictRepository(db)
	subDistrictRepo := regional_repository.NewSubDistrictRepository(db)
	urbanVillageRepo := regional_repository.NewUrbanVillageRepository(db)
	subDistrictService := regional_service.NewUrbanVillageService(urbanVillageRepo, subDistrictRepo)
	errMdl := subDistrictService.Import("C:\\Users\\NEXSOFT\\Documents\\Kuliah\\skripsi-2\\source\\desa.csv")
	if errMdl.CausedBy != nil {
		t.Fatal(errMdl.CausedBy)
	}
}
