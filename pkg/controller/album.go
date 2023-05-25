package controller

import (
	"songdb/pkg/service"
	"songdb/pkg/models"

	"github.com/labstack/echo/v4"
)

type AlbumController interface {
  GetAll(echo.Context) error
  GetById(echo.Context) error
  Post(echo.Context) error
  Put(echo.Context) error
  Delete(echo.Context) error
}

func NewAlbumController(service service.AlbumService) AlbumController {
  return AlbumControllerImpl{ service: service }
}

type AlbumListResponse struct {
  Data []models.Album `json:"data"`
}

type AlbumSingleResponse struct {
  Data *models.Album `json:"data"`
}

type AlbumIdResponse struct {
  Data struct {
    Id string `json:"id"`
  } `json:"data"`
}
