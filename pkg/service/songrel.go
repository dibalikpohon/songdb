package service

import (
  "database/sql"
  "songdb/pkg/models"
)

type SongRelService interface {
  GetSongsInAlbum(string) ([]models.Song, error)
  CreateOneSongInAlbum(string, *models.SongDto) (string, error)
}

func NewSongRelDb(db *sql.DB) SongRelService {
  return SongRelServiceImpl { db: db }
}
