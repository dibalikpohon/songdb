package models

type Album struct {
  Id string `json:"id" gorm:"primaryKey"`
  Name string `json:"name" gorm:"notNull"`
  Year int16 `json:"year" gorm:"notNull"`
}

type AlbumDto struct {
  Name string `json:"name"`
  Year int16 `json:"year"` 
}

func (a *Album) UpdateFromDto(dto AlbumDto) {
  if len(dto.Name) > 0 {
    a.Name = dto.Name
  }
  if dto.Year > 0 {
    a.Year = dto.Year
  }
}

func (a AlbumDto) ToEntity() Album {
  return Album{Name: a.Name, Year: a.Year}
}
