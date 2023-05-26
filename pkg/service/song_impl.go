package service

import (
  "errors"
  "songdb/pkg/models"
  myerror "songdb/pkg/errors"

  "github.com/aidarkhanov/nanoid"
)

import "gorm.io/gorm"

type SongServiceImpl struct {
  db *gorm.DB
}

func (si SongServiceImpl) Create(dto *models.SongDto) (string, error) {

  // Generate an ID
  newId, err := nanoid.Generate(nanoidAlnum, nanoidSize)
  if err != nil {
    return "", err
  }

  // Execute query to insert data to database
  // _, err = si.db.Exec("INSERT INTO `songs` VALUES(?, ?, ?, ?, ?, ?)",
  //                  newId, dto.Title, dto.Genre, dto.Duration, dto.Year, nil)
  song := dto.ToEntity()
  song.Id = newId
  result := si.db.Create(&song)
  if result.Error != nil {
    return "", result.Error
  }

  return newId, nil
}

func  (si SongServiceImpl) ReadAll() ([]models.Song, error) {

  // execute query to read all rows in songs table
  // rows, err := si.db.Query("SELECT `id`, `title`, `genre`, `duration`, `year` FROM songs")
  var songs []models.Song

  // Find all data with no conditions
  // SELECT * FROM songs;
  result := si.db.Find(&songs)
  if result.Error != nil {
    return nil, result.Error
  }

  return songs, nil
}

func (si SongServiceImpl) ReadOne(id string) (*models.Song, error) {

  var song models.Song

  result := si.db.First(&song, "id = ?", id)

  // === I don't think checking err != nil is required
  //     because errors.Is already did the null checking ===
  // if err != nil {
  //   // How to compare errors in Go
  //   // https://stackoverflow.com/a/57613539
  //   if errors.Is(err, sql.ErrNoRows) {
  //     return nil, &myerrors.NoData{ Message: "Cannot find requested id", What: id } 
  //   }
  //   return nil, err
  // }
  if errors.Is(result.Error, gorm.ErrRecordNotFound) { 
      return nil, &myerror.NoData{ Message: "Cannot find requested id", What: id }
  }
 
  return &song, nil 
}

func (si SongServiceImpl) Update(id string, dto *models.SongDto) (error) {

  // execute the query to update data
  // result, err := si.db.Exec("UPDATE `songs` SET `title`=?, `genre`=?, `duration`=?, `year`=? WHERE `id`=?",
  //                  dto.Title, dto.Genre, dto.Duration, dto.Year, id)
  var song models.Song

  // Grab first song that matches the primary key: id
  result := si.db.First(&song, "id = ?", id)
  if errors.Is(result.Error, gorm.ErrRecordNotFound) {
      return &myerror.NoData{ Message: "Cannot find requested id", What: id }
  }

  // Modify them
  song.UpdateFromDto(*dto)

  si.db.Save(&song)

  return nil
}

func (si SongServiceImpl) Delete(id string) error {

  // Execute query to delete data
  // result, err := si.db.Exec("DELETE FROM `songs` WHERE `id`=?", id)
  result := si.db.Delete(&models.Song{}, "id = ?", id)
  if errors.Is(result.Error, gorm.ErrRecordNotFound) {
    return &myerror.NoData{ Message: "Cannot find requested id", What: id }
  }

  return nil
}
