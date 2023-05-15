package main

import "github.com/labstack/echo/v4"

import (
	"songdb/pkg/config"
	"songdb/pkg/controller"
	"songdb/pkg/routes"
	"songdb/pkg/service"
)


func main() {
  e := echo.New()
  db, err := config.GetDb()
  if err != nil {
    panic(err.Error())
  }
  defer db.Close()

  // Create all services
  songService := service.NewSongService(db)
  albumService := service.NewAlbumService(db)
  songRelService := service.NewSongRelService(db)

  // Create all controllers
  songController := controller.NewSongController(songService)
  albumController := controller.NewAlbumController(albumService)
  songRelController := controller.NewSongRelController(songRelService)

  // Create all routes
  songRoutes := routes.NewSongRoutes(songController)
  albumRoutes := routes.NewAlbumRoutes(albumController)
  songRelRoutes := routes.NewSongRelRoutes(songRelController)

  // Register all routes to Echo
  songRoutes.Register(e)
  albumRoutes.Register(e)
  songRelRoutes.Register(e)

  e.Logger.Fatal(e.Start(":9000"))
}
