package main

import "songdb/pkg/models"
import "songdb/pkg/config"

func main() {
  db, err := config.GetDb();
  if err != nil {
    panic(err);
  }

  db.AutoMigrate(&models.Album{}, &models.Song{})
}
