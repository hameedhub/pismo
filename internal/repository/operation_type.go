package repository

import (
	"github.com/hameedhub/pismo/internal/model"
	"gorm.io/gorm"
)

type OperationTypeRepository interface {
	GetById(id int) (model.OperationType, error)
	GetAll() ([]model.OperationType, error)
}

type operationTypeRepository struct {
	db *gorm.DB
}

func (o operationTypeRepository) GetById(id int) (model.OperationType, error) {
	var opt model.OperationType
	err := o.db.First(&opt, id).Error
	return opt, err
}

func (o operationTypeRepository) GetAll() ([]model.OperationType, error) {
	var list []model.OperationType
	err := o.db.Find(list).Error
	return list, err
}

func NewOperationTypeRepository(db *gorm.DB) OperationTypeRepository {
	return &operationTypeRepository{db: db}
}
