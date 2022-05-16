package repository

import (
	"database/sql"
	"eform-gateway/lib"
	models "eform-gateway/models/access_places"
	"log"

	"github.com/DATA-DOG/go-sqlmock"
	// "github.com/google/uuid"
)

var ap = &models.AccessPlaces{
	ID:          1,
	Name:        "Rumah Sakit",
	Icon:        "img/rumah_sakit.png",
	Description: "Rumah Sakit",
	UpdatedAt:   lib.GetTimeNow("timestime"),
	CreatedAt:   lib.GetTimeNow("timestime"),
}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	return db, mock
}
