package VPNDatabaseManager

import (
	"github.com/jarakys/VPNDatabaseManager/Configs"
	"github.com/jarakys/VPNDatabaseManager/Errors"
	"github.com/jarakys/VPNDatabaseManager/Models"
)

func Start(databasename string) error {
	db, err := Configs.SQLDatabaseConnection("roblox.db")
	if err != nil {
		return Errors.ErrorDbWrapperUtils(err)
	}
	err = db.Table(Models.ConfigItemModel{}.TableName()).AutoMigrate(&Models.ConfigItemModel{})
	if err != nil {
		return Errors.ErrorDbWrapperUtils(err)
	}
	return nil
}
