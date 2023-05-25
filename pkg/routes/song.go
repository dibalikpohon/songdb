package routes

import "github.com/labstack/echo/v4"
import "songdb/pkg/controller"

type SongRoutes struct {
  controller controller.SongController
}

func NewSongRoutes(controller controller.SongController) SongRoutes {
  return SongRoutes{controller}
}

func (r SongRoutes) Register(e *echo.Echo) {

  e.GET("/songs", r.controller.GetAll);
  e.GET("/songs/:id", r.controller.GetById);
  e.POST("/songs", r.controller.Post);
  e.PUT("/songs/:id", r.controller.Put);
  e.DELETE("/songs/:id", r.controller.Delete);
}
