package db

import (
  "database/sql"
  "errors"
  "github.com/aidarkhanov/nanoid"

	"songdb/pkg/config"
	"songdb/pkg/models"
  myerror "songdb/pkg/errors"
)

func CreateNewAlbum(album *models.Album) (string, error) {
  // Create a DB connection
  db, err := config.GetDb();
  if err != nil {
    return "", err;
  }
  defer db.Close()

  // Generate an ID
  newId, err := nanoid.Generate(nanoidAlnum, nanoidSize);
  if err != nil {
    return "", err
  }

  // Execute query to insert data to database
  _, err = db.Exec("INSERT INTO `albums` VALUES (?, ?, ?)", newId, album.Name, album.Year)
  if err != nil {
    return "", err
  }

  return newId, nil
}

func ReadAlbumsTable() ([]models.Album, error) {
  // Create a DB connection
  db, err := config.GetDb()
  if err != nil {
    return nil, err
  }
  defer db.Close()

  // execute query to read all rows in albums table
  rows, err := db.Query("SELECT `id`, `name`, `year` FROM `albums`")
  if err != nil {
    return nil, err
  }
  defer rows.Close()

  // Prepare the array
  var result []models.Album

  // Iterate each rows by .Next()
  for rows.Next() {
    var each = models.Album{}
    var err = rows.Scan(&each.Id, &each.Name, &each.Year)

    if err != nil {
      return nil, err
    }

    result = append(result, each)
  }

  if err = rows.Err(); err != nil {
    return nil, err
  }

  return result, nil
}

func ReadOneAlbumById(id string) (*models.Album, error) {
  // Create a DB connection
  db, err := config.GetDb()
  if err != nil {
    return nil, err
  }
  defer db.Close()

  // Prepare the object
  var album models.Album

  // Create and execute the query
  err = db.QueryRow("SELECT `id`, `name`, `year` FROM `albums` WHERE `id`=?", id).Scan(&album.Id, &album.Name, &album.Year)
  if err != nil {
    if errors.Is(err, sql.ErrNoRows) {
      return nil, &myerror.NoData{ Message: "Cannot find requested id", What: id }
    }
    return nil, err
  }

  return &album, nil;
}

func UpdateAlbumById(id string, newAlbum *models.Album) (error) {
  // Create a DB connection
  db, err := config.GetDb()
  if err != nil {
    return err
  }
  defer db.Close()

  // execute the query to update data
  result, err := db.Exec("UPDATE `albums` SET `name`=?, `year`=? WHERE `id`=?", newAlbum.Name, newAlbum.Year, id)
  if err != nil {
    return err;
  }

  if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
    return &myerror.NoData{ Message: "Cannot find requested id", What: id }
  }

  return nil
}

func DeleteAlbumById(id string) (error) {
  // Create a DB connection
  db, err := config.GetDb()
  if err != nil {
    return err
  }
  defer db.Close()

  // execute query to delete data
  result, err := db.Exec("DELETE FROM `albums` WHERE `id`=?", id)
  if err != nil {
    return err
  }

  if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
    return &myerror.NoData{ Message: "Cannot find requested id", What: id }
  }

  return nil
}
