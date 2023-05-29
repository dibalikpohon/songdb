package main

import "github.com/labstack/echo/v4"

import "songdb/pkg/config"

func main() {
  e := echo.New()

  db, err := config.GetDb()
  if err != nil {
    panic(err.Error())
  }
  songRoutes := InitializeSongRoutes(db)
  albumRoutes := InitializeAlbumRoutes(db)
  songRelRoutes := InitializeSongRelRoutes(db)

  songRoutes.Register(e)
  albumRoutes.Register(e)
  songRelRoutes.Register(e)

  e.Logger.Fatal(e.Start(":9000"))
}
