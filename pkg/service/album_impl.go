package service

import (
  "database/sql"
  "errors"
  "github.com/aidarkhanov/nanoid"

  "songdb/pkg/models"
  myerror "songdb/pkg/errors"
)

type AlbumServiceImpl struct {
  db *sql.DB
}

func (ai AlbumServiceImpl) Create(dto *models.AlbumDto) (string, error) {
  
  newId, err := nanoid.Generate(nanoidAlnum, nanoidSize)
  if err != nil {
    return  "", err
  }

  // Execute query to insert data to database
  _, err = ai.db.Exec("INSERT INTO `albums` VALUES (?, ?, ?)", newId, dto.Name, dto.Year)
  if err != nil {
    return "", err
  }

  return newId, nil
}

func (ai AlbumServiceImpl) ReadAll() ([]models.Album, error) {

  // execute query to read all rows in albums table
  rows, err := ai.db.Query("SELECT `id`, `name`, `year` FROM `albums`")
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

func (ai AlbumServiceImpl) ReadOne(id string) (*models.Album, error) {

  // Prepare the object
  var album models.Album

  // Create and execute the query
  err := ai.db.QueryRow("SELECT `id`, `name`, `year` FROM `albums` WHERE `id`=?", id).Scan(&album.Id, &album.Name, &album.Year)
  if err != nil {
    if errors.Is(err, sql.ErrNoRows) {
      return nil, &myerror.NoData{ Message: "Cannot find requested id", What: id }
    }
    return nil, err
  }

  return &album, nil;
}

func (ai AlbumServiceImpl) Update(id string, dto *models.AlbumDto) error {

  // execute the query to update data
  result, err := ai.db.Exec("UPDATE `albums` SET `name`=?, `year`=? WHERE `id`=?", dto.Name, dto.Year, id)
  if err != nil {
    return err;
  }

  if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
    return &myerror.NoData{ Message: "Cannot find requested id", What: id }
  }

  return nil
}

func (ai AlbumServiceImpl) Delete(id string) error {

  // execute query to delete data
  result, err := ai.db.Exec("DELETE FROM `albums` WHERE `id`=?", id)
  if err != nil {
    return err
  }

  if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
    return &myerror.NoData{ Message: "Cannot find requested id", What: id }
  }

  return nil
}
