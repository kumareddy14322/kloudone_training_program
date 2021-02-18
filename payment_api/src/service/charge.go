package Service

import (
	_ "github.com/go-sql-driver/mysql"
	Config "github.com/kumareddy14322/kloudone_training_program/payment_api/src/config"
	"github.com/kumareddy14322/kloudone_training_program/payment_api/src/models"
)

func SavePayment(charge *models.Charge) (err error) {
	if err = Config.DB.Create(charge).Error; err != nil {
		return err
	}
	return nil

}
