package models

type Song struct {
  Id string         `json:"id"`
  Title string      `json:"title"`
  Genre string      `json:"genre"`
  Duration int32    `json:"duration"`
  Year int16        `json:"year"`
}


