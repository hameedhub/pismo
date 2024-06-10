package repository

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/hameedhub/pismo/internal/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOperationTypeRepository_GetById(t *testing.T) {
	db, mock, gormDb := SetupMock(t)
	defer db.Close()

	id := 1
	expected := model.OperationType{ID: id, Description: "PURCHASE"}
	rows := sqlmock.NewRows([]string{"id", "description"}).AddRow(expected.ID, expected.Description)
	mock.ExpectQuery("^SELECT \\* FROM \"operation_types\" WHERE \"operation_types\"\\.\"id\" = \\$1 AND \"operation_types\"\\.\"deleted_at\" IS NULL ORDER BY \"operation_types\"\\.\"id\" LIMIT \\$2").WithArgs(id, 1).WillReturnRows(rows)
	repo := NewOperationTypeRepository(gormDb)
	opt, err := repo.GetById(id)
	assert.Nil(t, err)
	assert.Equal(t, expected, opt)
}

func TestOperationTypeRepository_GetAll(t *testing.T) {

}
