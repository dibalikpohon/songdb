package service

import (
  "database/sql"
	"songdb/pkg/models"
)

type AlbumService interface {
  Create(*models.AlbumDto) (string, error)
  ReadAll() ([]models.Album, error)
  ReadOne(string) (*models.Album, error)
  Update(string, *models.AlbumDto) (error)
  Delete(string) (error)
}

func NewAlbumService(db *sql.DB) AlbumService {
  return AlbumServiceImpl {db: db}
}
