package service

import "songdb/pkg/models"
import "gorm.io/gorm"

type SongService interface {
  Create(*models.SongDto) (string, error)
  ReadAll() ([]models.Song, error)
  ReadOne(string) (*models.Song, error)
  Update(string, *models.SongDto) (error)
  Delete(string) (error)
}

func NewSongService(db *gorm.DB) SongService {
  return SongServiceImpl { db: db }
}
