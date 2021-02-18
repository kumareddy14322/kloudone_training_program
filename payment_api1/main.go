package main

import (
	"github.com/kumareddy14322/kloudone_training_program/payment_api1/app"
)

func main() {
	application := &app.App{}
	application.Initialize()
	application.Run(":8080")
}
