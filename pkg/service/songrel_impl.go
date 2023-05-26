package service

import (
  "errors"
  "songdb/pkg/models"
  myerrors "songdb/pkg/errors"

  "github.com/aidarkhanov/nanoid"
)

import "gorm.io/gorm"

type SongRelServiceImpl struct {
  db *gorm.DB
}

func (sri SongRelServiceImpl) GetSongsInAlbum(id string) ([]models.Song, error) {

  // To get songs in album, we have to:
  // 1. Check if an album exists, return NotFound error if dont
  // 2. SELECT * from `songs` WHERE `albumId` = id
  var _id string
  result := sri.db.Limit(1).Where("id = ?", id).Select("id").Scan(&_id)
  if errors.Is(result.Error, gorm.ErrRecordNotFound) {
    return nil, &myerrors.NoData{ Message: "Cannot find requested id", What: id}
  }

  // Prepare array
  var songs []models.Song
  result = sri.db.Select("id", "title", "genre", "duration").Find(&songs)
  if result.Error != nil {
    return nil, result.Error
  }

  return songs, nil
}

func (sri SongRelServiceImpl) CreateOneSongInAlbum(albumId string, song *models.SongDto) (string, error) {

  // Check if albumId exists
  var _albumId string
  // err := sri.db.QueryRow("SELECT `id` FROM `albums` WHERE `id`=?", albumId).Scan(&_albumId)
  result := sri.db.Limit(1).Select("id").Where("id = ?", albumId).Scan(&_albumId)
  if errors.Is(result.Error, gorm.ErrRecordNotFound) {
    return "", &myerrors.NoData{ Message: "Cannot find requested id", What: albumId}
  }

  songId, err := nanoid.Generate(nanoidAlnum, nanoidSize)
  if err != nil {
    return "", err
  }

  // Execute query to insert data to database
  result = sri.db.Exec("INSERT INTO `songs` VALUES(?, ?, ?, ?, ?, ?)",
                   songId, song.Title, song.Genre, song.Duration, song.Year, albumId)
  if result.Error != nil {
    return "", result.Error
  }

  return songId, nil
}
