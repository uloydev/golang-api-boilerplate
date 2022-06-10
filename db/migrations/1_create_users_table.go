package migrations

import (
	"golang-api-boilerplate/app/entity"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var CreateUsersTable *gormigrate.Migration = &gormigrate.Migration{
	ID: "create_users_table",
	Migrate: func(dbConn *gorm.DB) error {
		return dbConn.AutoMigrate(&entity.User{})
	},
	Rollback: func(dbConn *gorm.DB) error {
		return dbConn.Migrator().DropTable("users")
	},
}
