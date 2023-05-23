// go:build wireinject
// +build wireinject


package main

import (
	"database/sql"
	"songdb/pkg/controller"
	"songdb/pkg/routes"
	"songdb/pkg/service"

	"github.com/google/wire"
)

func InitializeSongRoutes(db *sql.DB) routes.SongRoutes {
  wire.Build(service.NewSongService, 
             controller.NewSongController, 
             routes.NewSongRoutes)
  return routes.SongRoutes{}
}

func InitializeAlbumRoutes(db *sql.DB) routes.AlbumRoutes {
  wire.Build(service.NewAlbumService, 
             controller.NewAlbumController, 
             routes.NewAlbumRoutes)
  return routes.AlbumRoutes{}
}

func InitializeSongRelRoutes(db *sql.DB) routes.SongRelRoutes {
  wire.Build(service.NewSongRelService,
             controller.NewSongRelController,
             routes.NewSongRelRoutes)
  return routes.SongRelRoutes{}
}
