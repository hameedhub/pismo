package seeder

import (
	"fmt"
	"github.com/hameedhub/pismo/internal/model"
	"gorm.io/gorm"
)

func truncateTable(db *gorm.DB, tableName string) error {
	hasTable := db.Migrator().HasTable(tableName)
	if hasTable {
		err := db.Exec(fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY;", tableName)).Error
		if err != nil {
			return fmt.Errorf("error truncating table '%s': %w", tableName, err)
		}
	}
	return nil
}

func OperationTypes(db *gorm.DB) error {
	if err := truncateTable(db, `operation_types`); err != nil {
		return err
	}

	operationTypes := []model.OperationType{
		{ID: model.PURCHASE, Description: "PURCHASE"},
		{ID: model.INSTALLMENT_PURCHASE, Description: "INSTALLMENT PURCHASE"},
		{ID: model.WITHDRAWAL, Description: "WITHDRAWAL"},
		{ID: model.PAYMENT, Description: "PAYMENT"},
	}

	db = db.Set("gorm:save_associations", false).Set("gorm:association_autoupdate", false).Set("gorm:association_autocreate", false)

	for _, operationType := range operationTypes {
		err := db.Where(model.OperationType{Description: operationType.Description}).Attrs(operationType).FirstOrCreate(&operationType).Error
		if err != nil {
			return err
		}
	}
	return nil
}
