package service

import (
  "database/sql"
  "errors"
  "songdb/pkg/models"
  myerrors "songdb/pkg/errors"

  "github.com/aidarkhanov/nanoid"
)

type SongRelServiceImpl struct {
  db *sql.DB
}

func (sri SongRelServiceImpl) GetSongsInAlbum(id string) ([]models.Song, error) {

  // To get songs in album, we have to:
  // 1. Check if an album exists, return NotFound error if dont
  // 2. SELECT * from `songs` WHERE `albumId` = id
  var _id string
  err := sri.db.QueryRow("SELECT `id` from `albums` WHERE `id`=?", id).Scan(&_id);
  if err != nil {
    if errors.Is(err, sql.ErrNoRows) {
      // FIXME: Should return NoSuchId
      //        as NoData is considered too generic.
      return nil, &myerrors.NoData{ Message: "Cannot find requested id", What: id }
    }
    return nil, err
  }

  rows, err := sri.db.Query("SELECT `id`, `title`, `genre`, `duration`, `year` FROM `songs` WHERE `albumId`=?", id)
  if err != nil {
    return nil, err
  }
  defer rows.Close()

  // Prepare array
  var songs []models.Song

  for rows.Next() {
    var song = models.Song{}
    var err = rows.Scan(&song.Id, &song.Title, &song.Genre, &song.Duration, &song.Year)

    if err != nil {
      return nil, err
    }

    songs = append(songs, song)
  }

  if err = rows.Err(); err != nil {
    return nil, err
  }

  return songs, nil
}

func (sri SongRelServiceImpl) CreateOneSongInAlbum(albumId string, song *models.SongDto) (string, error) {

  // Check if albumId exists
  var _albumId string
  err := sri.db.QueryRow("SELECT `id` FROM `albums` WHERE `id`=?", albumId).Scan(&_albumId)
  if err != nil {
    if errors.Is(err, sql.ErrNoRows) {
      return "", &myerrors.NoData{ Message: "Cannot find requested id", What: albumId }
    }
  }

  songId, err := nanoid.Generate(nanoidAlnum, nanoidSize)
  if err != nil {
    return "", err
  }

  // Execute query to insert data to database
  _, err = sri.db.Exec("INSERT INTO `songs` VALUES(?, ?, ?, ?, ?, ?)",
                   songId, song.Title, song.Genre, song.Duration, song.Year, albumId)
  if err != nil {
    return "", err
  }

  return songId, nil
}
