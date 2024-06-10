package model

import "gorm.io/gorm"

type OperationType struct {
	gorm.Model
	ID          int
	Description string
}

func (op *OperationType) NewOperationType() OperationType {
	return OperationType{ID: op.ID, Description: op.Description}
}
