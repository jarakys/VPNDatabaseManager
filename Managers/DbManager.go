package Managers

import (
	"github.com/jarakys/VPNDatabaseManager/Errors"
	"github.com/jarakys/VPNDatabaseManager/Models"
	"gorm.io/gorm"
)

type DbManager interface {
	Create(userId string, config string, name string, vpnType string, date int64) (string, error)
	Delete(userId string, vpnType string) error
	Get(userId string, vpnType string) (string, error)
	GetDate(userId string, vpnType string) (int64, error)
}

func NewDbManager(db *gorm.DB) DbManager {
	return &dbManagerImpl{db: db}
}

type dbManagerImpl struct {
	db *gorm.DB
}

func (d *dbManagerImpl) Create(userId string, config string, name string, vpnType string, date int64) (string, error) {
	err := d.db.Create(Models.ConfigItemModel{
		UserId: userId,
		Config: config,
		Name:   name,
		Type:   vpnType,
		Date:   date,
	}).Error
	if err != nil {
		return "", Errors.ErrorDbWrapperUtils(err)
	}
	return config, err
}

func (d *dbManagerImpl) Delete(userId string, vpnType string) error {
	err := d.db.Delete(&Models.ConfigItemModel{}, "user_id = ? and type = ?", userId, vpnType).Error
	if err != nil {
		return Errors.ErrorDbWrapperUtils(err)
	}
	return err
}

func (d *dbManagerImpl) Get(userId string, vpnType string) (string, error) {
	var item Models.ConfigItemModel
	err := d.db.First(&item, "user_id = ? and type = ?", userId, vpnType).Error
	if err != nil {
		return "", Errors.ErrorDbWrapperUtils(err)
	}
	return item.Config, err
}

func (d *dbManagerImpl) GetDate(userId string, vpnType string) (int64, error) {
	var item Models.ConfigItemModel
	err := d.db.First(&item, "user_id = ? and type = ?", userId, vpnType).Error
	if err != nil {
		return 0, Errors.ErrorDbWrapperUtils(err)
	}
	return item.Date, err
}
