package main

import (
	"fmt"

	"github.com/kumareddy14322/kloudone_training_program/payment_api/src/models"
	"github.com/kumareddy14322/kloudone_training_program/payment_api/src/router"

	"github.com/jinzhu/gorm"
	
	Config "github.com/kumareddy14322/kloudone_training_program/payment_api/src/config"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("error", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&models.Charge{})
	r := router.ChargeRouter()
	r.Run()

}
