package service

import (
  "database/sql"
  "errors"
  "songdb/pkg/models"
  myerrors "songdb/pkg/errors"

  "github.com/aidarkhanov/nanoid"
)

type SongServiceImpl struct {
  db *sql.DB
}

func (si SongServiceImpl) Create(dto *models.SongDto) (string, error) {

  // Generate an ID
  newId, err := nanoid.Generate(nanoidAlnum, nanoidSize)
  if err != nil {
    return "", err
  }

  // Execute query to insert data to database
  _, err = si.db.Exec("INSERT INTO `songs` VALUES(?, ?, ?, ?, ?, ?)",
                   newId, dto.Title, dto.Genre, dto.Duration, dto.Year, nil)
  if err != nil {
    return "", err
  }

  return newId, nil
}

func  (si SongServiceImpl) ReadAll() ([]models.Song, error) {

  // execute query to read all rows in songs table
  rows, err := si.db.Query("SELECT `id`, `title`, `genre`, `duration`, `year` FROM songs")
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

func (si SongServiceImpl) ReadOne(id string) (*models.Song, error) {

  var song models.Song

  err := si.db.QueryRow("SELECT `id`, `title`, `genre`, `duration`, `year` fROM `songs` WHERE `id`=?", id).
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

func (si SongServiceImpl) Update(id string, dto *models.SongDto) (error) {

  // execute the query to update data
  result, err := si.db.Exec("UPDATE `songs` SET `title`=?, `genre`=?, `duration`=?, `year`=? WHERE `id`=?",
                   dto.Title, dto.Genre, dto.Duration, dto.Year, id)
  if err != nil {
    return err
  }
  
  if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
    return &myerrors.NoData{ Message: "Cannot find requested id", What: id };
  }

  return nil
}

func (si SongServiceImpl) Delete(id string) error {

  // Execute query to delete data
  result, err := si.db.Exec("DELETE FROM `songs` WHERE `id`=?", id)
  if err != nil {
    return err;
  }

  if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
    return &myerrors.NoData{ Message: "Cannot find requested id", What: id };
  }

  return nil
}
