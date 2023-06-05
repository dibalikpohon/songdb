package service

import "songdb/pkg/models"
import "gorm.io/gorm"

type SongRelService interface {
  GetSongsInAlbum(string) ([]models.Song, error)
  CreateOneSongInAlbum(string, *models.SongDto) (string, error)
}

func NewSongRelService(db *gorm.DB) SongRelService {
  return SongRelServiceImpl { db: db }
}
