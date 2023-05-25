package routes

import "github.com/labstack/echo/v4"
import "songdb/pkg/controller"

type SongRelRoutes struct {
  controller controller.SongRelController
}

func NewSongRelRoutes(controller controller.SongRelController) SongRelRoutes {
  return SongRelRoutes{controller}
}

func (r SongRelRoutes) Register (e *echo.Echo) {
  e.GET("/albums/:id/songs", r.controller.GetAllSongsInAlbum)
  e.POST("/albums/:id/songs", r.controller.PostSongInAlbum)
}
