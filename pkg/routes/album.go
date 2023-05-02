package routes

import "github.com/labstack/echo/v4"
import "songdb/pkg/controller"

func RegisterAlbumRoutes(e *echo.Echo) {
  e.GET("/albums", controller.GetAllAlbums);
  e.GET("/albums/:id", controller.GetAlbumById);
  e.POST("/albums", controller.PostAlbum);
  e.PUT("/albums/:id", controller.PutAlbumUpdate);
  e.DELETE("/albums/:id", controller.DeleteAlbum);
}





