package db

import (
  "database/sql"
  "songdb/pkg/models"
)

type SongRelDb interface {
  GetSongsInAlbum(string) ([]models.Song, error)
  CreateOneSongInAlbum(string, *models.SongDto) (string, error)
}

func NewSongRelDb(db *sql.DB) SongRelDb {
  return &SongRelDbImpl { db: db }
}
