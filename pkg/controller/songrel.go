package controller

import (
	"songdb/pkg/models"
	"songdb/pkg/service"

	"github.com/labstack/echo/v4"
)

type SongRelController interface {
  GetAllSongsInAlbum(echo.Context) error
  PostSongInAlbum(echo.Context) error
}

func NewSongRelController(service service.SongRelService) SongRelController {
  return SongRelControllerImpl{service}
}

type AllSongsInAlbumResponse struct {
  Data struct {
    AlbumId string `json:"albumId"`
    Songs []models.Song `json:"songs"`
  } `json:"data"`
}

type SongIdInAlbumResponse struct {
  Data struct {
    AlbumId string `json:"albumId"`
    SongId string `json:"songId"`
  } `json:"data"`
}
