package db

import (
  "database/sql"
	"songdb/pkg/models"
)

type AlbumDb interface {
  Create(*models.AlbumDto) (string, error)
  ReadAll() ([]models.Album, error)
  ReadOne(string) (*models.Album, error)
  Update(string, *models.AlbumDto) (error)
  Delete(string) (error)
}

func NewAlbumDb(db *sql.DB) AlbumDb {
  return &AlbumDbImpl {db: db}
}
