package service

import "database/sql"
import "songdb/pkg/models"

type SongService interface {
  Create(*models.SongDto) (string, error)
  ReadAll() ([]models.Song, error)
  ReadOne(string) (*models.Song, error)
  Update(string, *models.SongDto) (error)
  Delete(string) (error)
}

func NewSongService(db *sql.DB) SongService {
  return SongServiceImpl { db: db }
}
