package service

import (
	"errors"
	myerrors "songdb/pkg/errors"
	"songdb/pkg/models"

	"github.com/aidarkhanov/nanoid"
	"gorm.io/gorm"
)

type SongRelServiceImpl struct {
  db *gorm.DB
}

func (sri SongRelServiceImpl) GetSongsInAlbum(id string) ([]models.Song, error) {

  // To get songs in album, we have to:
  // 1. Check if an album exists, return NotFound error if dont
  // 2. SELECT * from `songs` WHERE `albumId` = id
  var album models.Album
  result := sri.db.First(&album, "id = ?", id)
  if errors.Is(result.Error, gorm.ErrRecordNotFound) {
    return nil, &myerrors.NoData{ Message: "Cannot find requested id", What: id}
  }

  // Prepare array
  var songs []models.Song
  result = sri.db.Where("album_id = ?", id).Find(&songs)
  if result.Error != nil {
    return nil, result.Error
  }

  return songs, nil
}

func (sri SongRelServiceImpl) CreateOneSongInAlbum(albumId string, song *models.SongDto) (string, error) {

  // Check if albumId exists
  var album models.Album
  result := sri.db.First(&album, "id = ?", albumId)
  if errors.Is(result.Error, gorm.ErrRecordNotFound) {
    return "", &myerrors.NoData{ Message: "Cannot find requested id", What: albumId}
  }

  songId, err := nanoid.Generate(nanoidAlnum, nanoidSize)
  if err != nil {
    return "", err
  }

  song.AlbumId = album.Id;
  _song := song.ToEntity(); 
  _song.Id = songId;

  // Execute query to insert data to database
  result = sri.db.Create(&_song); 
  if result.Error != nil {
    return "", result.Error
  }

  return songId, nil
}
