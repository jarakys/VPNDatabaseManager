package Managers

import "gorm.io/gorm"

type DbRollbacker interface {
	Commit() error
	Rollback() error
	Perform(operation operation)
}

type dbRollbackerImpl struct {
	db *gorm.DB
}

func (dbR *dbRollbackerImpl) Commit() error {
	return dbR.db.Commit().Error
}

func (dbR *dbRollbackerImpl) Rollback() error {
	return dbR.db.Rollback().Error
}

type operation func() error

func (d *dbRollbackerImpl) Perform(operation operation) {
	d.db.Begin()
	err := operation()
	if err != nil {
		d.db.Rollback()
	}
	d.db.Commit()
}
