package Managers

import (
	"github.com/jarakys/VPNDatabaseManager/Errors"
	"gorm.io/gorm"
)

type operation func() error

type DbRollbacker interface {
	Commit() error
	Rollback() error
	Perform(operation operation)
}

type dbRollbackerImpl struct {
	db *gorm.DB
}

func (dbR *dbRollbackerImpl) Commit() error {
	if err := dbR.db.Commit().Error; err != nil {
		return Errors.ErrorDbWrapperUtils(err)
	}
	return nil
}

func (dbR *dbRollbackerImpl) Rollback() error {
	if err := dbR.db.Rollback().Error; err != nil {
		return Errors.ErrorDbWrapperUtils(err)
	}
	return nil
}

func (d *dbRollbackerImpl) Perform(operation operation) {
	d.db.Begin()
	err := operation()
	if err != nil {
		d.db.Rollback()
	}
	d.db.Commit()
}
