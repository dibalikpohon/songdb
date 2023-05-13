package db

import (
  "database/sql"
  "errors"
  "github.com/aidarkhanov/nanoid"

  "songdb/pkg/models"
  myerror "songdb/pkg/errors"
)

type AlbumDbImpl struct {
  db *sql.DB
}

func (c *AlbumDbImpl) Create(dto *models.AlbumDto) (string, error) {
  
  newId, err := nanoid.Generate(nanoidAlnum, nanoidSize)
  if err != nil {
    return  "", err
  }

  // Execute query to insert data to database
  _, err = c.db.Exec("INSERT INTO `albums` VALUES (?, ?, ?)", newId, dto.Name, dto.Year)
  if err != nil {
    return "", err
  }

  return newId, nil
}

func (c *AlbumDbImpl) ReadAll() ([]models.Album, error) {

  // execute query to read all rows in albums table
  rows, err := c.db.Query("SELECT `id`, `name`, `year` FROM `albums`")
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

func (c *AlbumDbImpl) ReadOne(id string) (*models.Album, error) {

  // Prepare the object
  var album models.Album

  // Create and execute the query
  err := c.db.QueryRow("SELECT `id`, `name`, `year` FROM `albums` WHERE `id`=?", id).Scan(&album.Id, &album.Name, &album.Year)
  if err != nil {
    if errors.Is(err, sql.ErrNoRows) {
      return nil, &myerror.NoData{ Message: "Cannot find requested id", What: id }
    }
    return nil, err
  }

  return &album, nil;
}

func (c *AlbumDbImpl) Update(id string, dto *models.AlbumDto) error {

  // execute the query to update data
  result, err := c.db.Exec("UPDATE `albums` SET `name`=?, `year`=? WHERE `id`=?", dto.Name, dto.Year, id)
  if err != nil {
    return err;
  }

  if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
    return &myerror.NoData{ Message: "Cannot find requested id", What: id }
  }

  return nil
}

func (c *AlbumDbImpl) Delete(id string) error {

  // execute query to delete data
  result, err := c.db.Exec("DELETE FROM `albums` WHERE `id`=?", id)
  if err != nil {
    return err
  }

  if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
    return &myerror.NoData{ Message: "Cannot find requested id", What: id }
  }

  return nil
}
