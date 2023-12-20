package main

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"go-master-data/common"
	"go-master-data/config"
	"go-master-data/controller/restapi"
	"go-master-data/dto"
	"golang.org/x/text/language"
	"os"
	"runtime"
)

func main() {
	var arguments = "development"
	args := os.Args
	if len(args) > 1 {
		arguments = args[1]
	}
	_, f, l, _ := runtime.Caller(0)

	fmt.Println(f, l)
	config.GenerateConfiguration(arguments)
	dto.GenerateValidOperator()
	common.Validation = common.NewGoValidator()
	loadBundleI18N()

	err := common.SetServerAttribute()
	if err != nil {
		fmt.Println("ERROR common server attribute : ", err)
		os.Exit(3)
	}
	err = common.MigrateSchema(common.ConnectionDB, config.ApplicationConfiguration.GetSqlMigrateDirPath(), config.ApplicationConfiguration.GetPostgresqlConfig().DefaultSchema)
	if err != nil {
		fmt.Println("ERROR migrate sql : ", err)
		os.Exit(3)
	}

	err = restapi.Router()
	if err != nil {
		fmt.Println("ERROR router : ", err)
		os.Exit(3)
	}
}

func loadBundleI18N() {
	prefixPath := config.ApplicationConfiguration.GetLanguageDirectoryPath()
	var err error
	//------------ error bundle
	common.CommonBundle = i18n.NewBundle(language.Indonesian)
	common.CommonBundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	_, err = common.CommonBundle.LoadMessageFile(prefixPath + "/common/en-US.json")
	readError(err)

	_, err = common.CommonBundle.LoadMessageFile(prefixPath + "/common/id-ID.json")
	readError(err)

	//------------ constanta bundle
	common.ConstantaBundle = i18n.NewBundle(language.Indonesian)
	common.ConstantaBundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	_, err = common.ConstantaBundle.LoadMessageFile(prefixPath + "/constanta/en-US.json")
	readError(err)

	_, err = common.ConstantaBundle.LoadMessageFile(prefixPath + "/constanta/id-ID.json")
	readError(err)

	//------------ constanta bundle
	common.ErrorBundle = i18n.NewBundle(language.Indonesian)
	common.ErrorBundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	_, err = common.ErrorBundle.LoadMessageFile(prefixPath + "/error/en-US.json")
	readError(err)

	_, err = common.ErrorBundle.LoadMessageFile(prefixPath + "/error/id-ID.json")
	readError(err)
}

func readError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
