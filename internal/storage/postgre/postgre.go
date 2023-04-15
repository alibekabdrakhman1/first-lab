package postgre

import (
	"fmt"
	"github.com/alibekabdrakhman1/first-lab/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Dial(cfg *configs.Config) (*gorm.DB, error) {
	fmt.Println(cfg.Database.PgUrl)
	db, err := gorm.Open(postgres.Open(cfg.Database.PgUrl), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
