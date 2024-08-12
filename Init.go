package VPNDatabaseManager

import (
	"github.com/jarakys/VPNDatabaseManager/Configs"
	"github.com/jarakys/VPNDatabaseManager/Models"
)

func Start(databasename string) error {
	db, err := Configs.SQLDatabaseConnection("roblox.db")
	if err != nil {
		return err
	}
	db.Table(Models.ConfigItemModel{}.TableName()).AutoMigrate(&Models.ConfigItemModel{})
	return nil
}
