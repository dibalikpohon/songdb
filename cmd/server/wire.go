// go:build wireinject
// +build wireinject

package main

import (
	"gorm.io/gorm"
	"songdb/pkg/controller"
	"songdb/pkg/routes"
	"songdb/pkg/service"

	"github.com/google/wire"
)

func InitializeSongRoutes(db *gorm.DB) routes.SongRoutes {
  wire.Build(service.NewSongService, 
             controller.NewSongController, 
             routes.NewSongRoutes)
  return routes.SongRoutes{}
}

func InitializeAlbumRoutes(db *gorm.DB) routes.AlbumRoutes {
  wire.Build(service.NewAlbumService, 
             controller.NewAlbumController, 
             routes.NewAlbumRoutes)
  return routes.AlbumRoutes{}
}

func InitializeSongRelRoutes(db *gorm.DB) routes.SongRelRoutes {
  wire.Build(service.NewSongRelService,
             controller.NewSongRelController,
             routes.NewSongRelRoutes)
  return routes.SongRelRoutes{}
}
