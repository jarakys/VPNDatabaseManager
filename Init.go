package VPNDatabaseManager

import (
	"github.com/jarakys/VPNDatabaseManager/Configs"
	"github.com/jarakys/VPNDatabaseManager/Errors"
	"github.com/jarakys/VPNDatabaseManager/Models"
	"gorm.io/gorm"
)

func Init(databasename string) (*gorm.DB, error) {
	db, err := Configs.SQLDatabaseConnection(databasename)
	if err != nil {
		return nil, Errors.ErrorDbWrapperUtils(err)
	}
	err = db.Table(Models.ConfigItemModel{}.TableName()).AutoMigrate(&Models.ConfigItemModel{})
	if err != nil {
		return nil, Errors.ErrorDbWrapperUtils(err)
	}
	return db, nil
}
