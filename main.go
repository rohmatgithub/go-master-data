package main

import (
	"fmt"
	"go-master-data/common"
	"go-master-data/config"
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
}
