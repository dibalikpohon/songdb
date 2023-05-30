package models

import "database/sql"

type Song struct {
  Id string         `json:"id" gorm:"primaryKey;size:10"`
  Title string      `json:"title" gorm:"notNull;size:30"`
  Genre string      `json:"genre" gorm:"notNull;size:30"`
  Duration int32    `json:"duration"`
  Year int16        `json:"year" gorm:"notNull"`
  AlbumId sql.NullString `json:"-"`
}

type SongDto struct { 
  Title string      `json:"title"`
  Genre string      `json:"genre"`
  Duration int32    `json:"duration"`
  Year int16        `json:"year"`
  AlbumId string    `json:"albumId"`
}

func (s *Song) UpdateFromDto(dto SongDto) {
  if len(dto.Title) > 0 {
    s.Title = dto.Title
  }
  if len(dto.Genre) > 0 {
    s.Genre = dto.Genre
  }
  if dto.Duration > 0 {
    s.Duration = dto.Duration
  }
  if dto.Year > 0 {
    s.Year = dto.Year
  }
  if len(dto.AlbumId) > 0 {
    s.AlbumId.String = dto.AlbumId
    s.AlbumId.Valid = true
  }
}

func (s SongDto) ToEntity() Song {
  if len(s.AlbumId) > 0 {
    return Song{Title: s.Title, Genre: s.Genre,
                Duration: s.Duration, Year: s.Year, 
                AlbumId: sql.NullString{Valid: true, String: s.AlbumId}}
  } else {
    return Song{Title: s.Title, Genre: s.Genre,
                Duration: s.Duration, Year: s.Year, 
                AlbumId: sql.NullString{Valid: false}}
  }
}
