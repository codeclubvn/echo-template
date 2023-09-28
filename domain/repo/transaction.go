package repo

import (
	"gorm.io/gorm"
	"trial_backend/infra"
)

type TX struct {
	db infra.Database
}

type TxFn func(*TX) error

func WithTransaction(db *infra.Database, fn TxFn) error {
	return db.Transaction(func(tx *gorm.DB) error {
		return fn(&TX{
			db: infra.Database{
				DB: tx,
			},
		})
	})
}

func GetTX(tx *TX, db infra.Database) {
	if tx == nil {
		tx = &TX{db: db}
	}
}
