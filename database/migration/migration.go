package migration

import (
	"github.com/sultannaufal/clean-architecture/database"
	"github.com/sultannaufal/clean-architecture/internal/model"
)

var tables = []interface{}{
	&model.User{},
}

func Migrate() {
	conn := database.GetConnection()
	conn.AutoMigrate(tables...)
}
