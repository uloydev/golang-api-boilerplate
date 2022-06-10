package initialize

import (
	"golang-api-boilerplate/app/controller"
)

var InitFunctions = []InitFunc{
	controller.InitializeUserController,
	controller.InitializeAdminController,
}
