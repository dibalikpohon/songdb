package db

import (
	"database/sql"
  "errors"
	"songdb/pkg/config"
	"songdb/pkg/models"
  myerrors "songdb/pkg/errors"

	"github.com/aidarkhanov/nanoid"
)

func CreateNewSong(song *models.Song) (string, error) {
  // Create a DB connection
  db, err := config.GetDb();
  if err != nil {
    return "", err;
  }
  defer db.Close();

  // Generate an ID
  newId, err := nanoid.Generate(nanoidAlnum, nanoidSize)
  if err != nil {
    return "", err
  }

  // Execute query to insert data to database
  _, err = db.Exec("INSERT INTO `songs` VALUES(?, ?, ?, ?, ?)",
                   newId, song.Title, song.Genre, song.Duration, song.Year)
  if err != nil {
    return "", err
  }

  return newId, nil
}

func ReadSongsTable() ([]models.Song, error) {
  // Create a DB connection
  db, err := config.GetDb();
  if err != nil {
    return nil, err
  }
  defer db.Close()

  // execute query to read all rows in songs table
  rows, err := db.Query("SELECT `id`, `title`, `genre`, `duration`, `year` FROM songs")
  if err != nil {
    return nil, err
  }
  defer rows.Close()
  
  // Prepare the array
  var songs []models.Song

  // Iterate each rows by .Next()
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

func ReadOneSongById(id string) (*models.Song, error) {
  // Create a DB connection
  db, err := config.GetDb()
  if err != nil {
    return nil, err
  }
  // Let's defer db.Close()!!!
  defer db.Close();

  var song models.Song

  err = db.QueryRow("SELECT `id`, `title`, `genre`, `duration`, `year` fROM `songs` WHERE `id`=?", id).
             Scan(&song.Id, &song.Title, &song.Genre, &song.Duration, &song.Year)

  if err != nil {
    // How to compare errors in Go
    // https://stackoverflow.com/a/57613539
    if errors.Is(err, sql.ErrNoRows) {
      return nil, &myerrors.NoData{ Message: "Cannot find requested id", What: id } 
    }
    return nil, err
  }
 
  return &song, nil 
}

func UpdateSongById(id string, newSong *models.Song) (error) {
  // Create a DB connection
  db, err := config.GetDb()
  if err != nil {
    return err
  }
  defer db.Close()

  // execute the query to update data
  result, err := db.Exec("UPDATE `songs` SET `title`=?, `genre`=?, `duration`=?, `year`=? WHERE `id`=?",
                   newSong.Title, newSong.Genre, newSong.Duration, newSong.Year, id)
  if err != nil {
    return err
  }
  
  if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
    return &myerrors.NoData{ Message: "Cannot find requested id", What: id };
  }

  return nil
}

func DeleteSongById(id string) (error) {
  // Create a DB connection
  db, err := config.GetDb()
  if err != nil {
    return err
  }
  defer db.Close()

  // Execute query to delete data
  result, err := db.Exec("DELETE FROM `songs` WHERE `id`=?", id)
  if err != nil {
    return err;
  }

  if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
    return &myerrors.NoData{ Message: "Cannot find requested id", What: id };
  }

  return nil
}
