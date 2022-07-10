package access_places

import (
	"database/sql"
	"infolelang/lib"
	models "infolelang/models/access_places"
	"log"
	"math/rand"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	// "github.com/google/uuid"
)

var (
	randInt = rand.NewSource(time.Now().UnixNano())
	ap      = &models.AccessPlacesRequest{
		ID:          randInt.Int63(),
		Name:        "Rumah Sakit",
		Icon:        "img/rumah_sakit.png",
		Description: "Rumah Sakit",
	}
	date = lib.GetTimeNow("timestime")
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	return db, mock
}

func TestGetAll(t *testing.T) {
	db, mock := NewMock()
	repo := &AccessPlaceRepository{dbRaw: db}
	defer func() {
		repo.dbRaw.Close()
	}()

	query := "SELECT id, name, icon, description, updated_at, created_at"
	rows := sqlmock.NewRows([]string{"id", "name", "icon", "description", "updated_at", "created_at"}).
		AddRow(ap.ID, ap.Name, ap.Icon, ap.Description, date, date).
		AddRow(ap.ID, ap.Name, ap.Icon, ap.Description, date, date)
	mock.ExpectQuery(query).WillReturnRows(rows)

	ap, err := repo.GetAll()
	assert.NotNil(t, ap)
	assert.NoError(t, err)
}

func TestGetAllError(t *testing.T) {
	db, mock := NewMock()
	repo := &AccessPlaceRepository{dbRaw: db}
	defer func() {
		repo.dbRaw.Close()
	}()

	query := "SELECT id, name, icon, description, updated_at, created_at"
	rows := sqlmock.NewRows([]string{"id", "name", "icon", "description", "updated_at", "created_at"})
	mock.ExpectQuery(query).WillReturnRows(rows)

	ap, err := repo.GetAll()
	assert.Empty(t, ap)
	assert.Error(t, err)
	assert.Len(t, ap, 1)
}

func TestGetOne(t *testing.T) {
	db, mock := NewMock()
	repo := &AccessPlaceRepository{dbRaw: db}
	defer func() {
		repo.dbRaw.Close()
	}()

	query := "SELECT id, name, icon, description, updated_at, created_at where id = \\?"
	rows := sqlmock.NewRows([]string{"id", "name", "icon", "description", "updated_at", "created_at"})
	mock.ExpectQuery(query).WithArgs(ap.ID).WillReturnRows(rows)

	ap, err := repo.GetOne(ap.ID)
	assert.NotNil(t, ap)
	assert.NoError(t, err)
}

func TestGetOneError(t *testing.T) {
	db, mock := NewMock()
	repo := &AccessPlaceRepository{dbRaw: db}
	defer func() {
		repo.dbRaw.Close()
	}()

	query := "SELECT id, name, icon, description, updated_at, created_at"
	rows := sqlmock.NewRows([]string{"id", "name", "icon", "description", "updated_at", "created_at"})
	mock.ExpectQuery(query).WithArgs(ap.ID).WillReturnRows(rows)

	ap, err := repo.GetOne(ap.ID)
	assert.Empty(t, ap)
	assert.Error(t, err)
}
func TestStore(t *testing.T) {
	db, mock := NewMock()
	repo := &AccessPlaceRepository{dbRaw: db}
	defer func() {
		repo.dbRaw.Close()
	}()

	query := "INSERT into access_places (name, icon, description, created_at) VALUES (?,?,?,?)"
	mock.ExpectQuery(query).WithArgs(ap.Name, ap.Icon, ap.Description, date)

	_, err := repo.Store(ap)
	assert.NoError(t, err)
}

func TestStoreError(t *testing.T) {
	db, mock := NewMock()
	repo := &AccessPlaceRepository{dbRaw: db}
	defer func() {
		repo.dbRaw.Close()
	}()

	query := "INSERT into access_places (name, icon, description, created_at) VALUES (?,?,?,?)"
	mock.ExpectQuery(query).WithArgs(ap.Name, ap.Icon, ap.Description, date)

	_, err := repo.Store(ap)
	assert.Error(t, err)
}
func TestUpdate(t *testing.T) {
	db, mock := NewMock()
	repo := &AccessPlaceRepository{dbRaw: db}
	defer func() {
		repo.dbRaw.Close()
	}()

	query := "UPDATE access_places SET name = \\?, icon = \\?, description = \\?, updated_at = \\? WHERE id = \\?"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(ap.Name, ap.Icon, ap.Description, date).WillReturnResult(sqlmock.NewResult(0, 1))

	_, err := repo.Update(ap)
	assert.NoError(t, err)
}

func TestUpdateError(t *testing.T) {
	db, mock := NewMock()
	repo := &AccessPlaceRepository{dbRaw: db}
	defer func() {
		repo.dbRaw.Close()
	}()

	query := "UPDATE access_places SET name = \\?, icon = \\?, description = \\?, updated_at = \\? WHERE id = \\?"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(ap.Name, ap.Icon, ap.Description, date).WillReturnResult(sqlmock.NewResult(0, 0))

	_, err := repo.Update(ap)
	assert.Error(t, err)
}
func TestDelete(t *testing.T) {
	db, mock := NewMock()
	repo := &AccessPlaceRepository{dbRaw: db}
	defer func() {
		repo.dbRaw.Close()
	}()

	query := "DELETE FROM access_places WHERE id = \\?"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(ap.ID).WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.Delete(ap.ID)
	assert.NoError(t, err)
}

func TestDeleteError(t *testing.T) {
	db, mock := NewMock()
	repo := &AccessPlaceRepository{dbRaw: db}
	defer func() {
		repo.dbRaw.Close()
	}()

	query := "DELETE FROM access_places WHERE id = \\?"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(ap.ID).WillReturnResult(sqlmock.NewResult(0, 0))

	err := repo.Delete(ap.ID)
	assert.Error(t, err)
}
