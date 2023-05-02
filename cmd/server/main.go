package main

import (
  "github.com/labstack/echo/v4"
  "songdb/pkg/routes"
)

func main() {
  e := echo.New()
  routes.RegisterSongRoutes(e)
  routes.RegisterAlbumRoutes(e)
  e.Logger.Fatal(e.Start(":9000"))
}
