package routes

import "github.com/labstack/echo/v4"
import "songdb/pkg/controller"

type AlbumRoutes struct {
  controller controller.AlbumController
}

func NewAlbumRoutes(controller controller.AlbumController) AlbumRoutes {
  return AlbumRoutes{ controller }
}

func (r AlbumRoutes) Register(e *echo.Echo) {
  
  e.GET("/albums", r.controller.GetAll);
  e.GET("/albums/:id", r.controller.GetById);
  e.POST("/albums", r.controller.Post);
  e.PUT("/albums/:id", r.controller.Put);
  e.DELETE("/albums/:id", r.controller.Delete);
}
