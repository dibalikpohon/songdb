package service

import (
  "errors"
  "github.com/aidarkhanov/nanoid"

  "songdb/pkg/models"
  myerror "songdb/pkg/errors"
)

import "gorm.io/gorm"

type AlbumServiceImpl struct {
  db *gorm.DB
}

func (ai AlbumServiceImpl) Create(dto *models.AlbumDto) (string, error) {
  
  newId, err := nanoid.Generate(nanoidAlnum, nanoidSize)
  if err != nil {
    return  "", err
  }

  // Execute query to insert data to database
  album := dto.ToEntity()
  album.Id = newId
  result := ai.db.Create(&album)
  if result.Error != nil {
    return "", result.Error
  }

  return newId, nil
}

func (ai AlbumServiceImpl) ReadAll() ([]models.Album, error) {

  // execute query to read all rows in albums table
  var albums []models.Album

  // Find all data with no conditions
  result := ai.db.Find(&albums)
  if result.Error != nil {
    return nil, result.Error
  }

  return albums, nil
}

func (ai AlbumServiceImpl) ReadOne(id string) (*models.Album, error) {

  // Prepare the object
  var album models.Album

  // Grab the first data that matches the primary key: id
  result := ai.db.First(&album, "id = ?", id)
  if errors.Is(result.Error, gorm.ErrRecordNotFound) {
      return nil, &myerror.NoData{ Message: "Cannot find requested id", What: id }
  }

  return &album, nil;
}

func (ai AlbumServiceImpl) Update(id string, dto *models.AlbumDto) error {

  // execute the query to update data
  var album models.Album

  // Grab first album that matches primary key: id
  result := ai.db.First(&album, "id = ?", id)
  if errors.Is(result.Error, gorm.ErrRecordNotFound) {
      return &myerror.NoData{ Message: "Cannot find requested id", What: id }
  }

  // Modify them
  album.UpdateFromDto(*dto)
  
  // Update data
  ai.db.Save(&album)

  return nil
}

func (ai AlbumServiceImpl) Delete(id string) error {

  // execute query to delete data
  var album models.Album

  result := ai.db.First(&album, "id = ?", id)
  
  if errors.Is(result.Error, gorm.ErrRecordNotFound) {
    return &myerror.NoData{ Message: "Cannot find requested id", What: id }
  }

  ai.db.Delete(&album)
  return nil
}
