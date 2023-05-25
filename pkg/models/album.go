package models

type Album struct {
  Id string `json:"id"`
  Name string `json:"name"`
  Year int16 `json:"year"`
}

type AlbumDto struct {
  Name string `json:"name"`
  Year int16 `json:"year"` 
}
