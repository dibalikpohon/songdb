package models

type Song struct {
  Id string         `json:"id" gorm:"primaryKey"`
  Title string      `json:"title" gorm:"notNull"`
  Genre string      `json:"genre" gorm:"notNull"`
  Duration int32    `json:"duration"`
  Year int16        `json:"year" gorm:"notNull"`
}

type SongDto struct { 
  Title string      `json:"title"`
  Genre string      `json:"genre"`
  Duration int32    `json:"duration"`
  Year int16        `json:"year"`
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
}

func (s SongDto) ToEntity() Song {
  return Song{Title: s.Title, Genre: s.Genre,
              Duration: s.Duration, Year: s.Year}
}
