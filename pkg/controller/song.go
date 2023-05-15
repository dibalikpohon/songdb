package controller

import (
	"songdb/pkg/models"
	"songdb/pkg/service"

	"github.com/labstack/echo/v4"
)

type SongController interface {
  GetAll(echo.Context) error
  GetById(echo.Context) error
  Post(echo.Context) error
  Put(echo.Context) error
  Delete(echo.Context) error
}

func NewSongController(service service.SongService) SongController {
  return SongControllerImpl{service}
}

type SongListResponse struct {
  Data []models.Song `json:"data"`
}

type SongSingleResponse struct {
  Data *models.Song `json:"data"`
}

type SongIdResponse struct {
  Data struct {
    Id string `json:"id"`
  } `json:"data"`
}
