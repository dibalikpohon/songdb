package db

import "database/sql"
import "songdb/pkg/models"

type SongDb interface {
  Create(*models.SongDto) (string, error)
  ReadAll() ([]models.Song, error)
  ReadOne(string) (*models.Song, error)
  Update(string, *models.SongDto) (error)
  Delete(string) (error)
}

func NewSongDb(db *sql.DB) SongDb {
  return &SongDbImpl { db: db }
}
