package service

import "songdb/pkg/models"
import "gorm.io/gorm"

type AlbumService interface {
  Create(*models.AlbumDto) (string, error)
  ReadAll() ([]models.Album, error)
  ReadOne(string) (*models.Album, error)
  Update(string, *models.AlbumDto) (error)
  Delete(string) (error)
}

func NewAlbumService(db *gorm.DB) AlbumService {
  return AlbumServiceImpl {db: db}
}
