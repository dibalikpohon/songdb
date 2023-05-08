package routes

import "github.com/labstack/echo/v4"
import "songdb/pkg/controller"

func RegisterSongrelRoutes (e *echo.Echo) {
  e.GET("/albums/:id/songs", controller.GetAllSongsInAlbum)
  e.POST("/albums/:id/songs", controller.PostSongInAlbum)
}
