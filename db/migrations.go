package db

import (
	"golang-api-boilerplate/db/migrations"

	"github.com/go-gormigrate/gormigrate/v2"
)

var Migrations = []*gormigrate.Migration{

	migrations.CreateUsersTable,
	migrations.CreateAdminsTable,
}
