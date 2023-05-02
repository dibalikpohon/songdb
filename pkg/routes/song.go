package routes

import "github.com/labstack/echo/v4"
import "songdb/pkg/controller"

func RegisterSongRoutes(e *echo.Echo) {
  e.GET("/songs", controller.GetAllSongs);
  e.GET("/songs/:id", controller.GetSongById);
  e.POST("/songs", controller.PostSong);
  e.PUT("/songs/:id", controller.PutSongUpdate);
  e.DELETE("/songs/:id", controller.DeleteSong)
}
