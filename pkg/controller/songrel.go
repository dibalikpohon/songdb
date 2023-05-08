package controller

import (
  "net/http"
  "errors"
  "github.com/labstack/echo/v4"
)

import (
  "songdb/pkg/db"
  "songdb/pkg/models"
  myerror "songdb/pkg/errors"
)

func GetAllSongsInAlbum(c echo.Context) error {
  
  id := c.Param("id")
  songs, err := db.GetSongsInAlbum(id)
  if err != nil {
    var noDataError *myerror.NoData
    if errors.As(err, &noDataError) {
      return c.JSON(http.StatusNotFound, noDataError)
    }
    response := map[string]string {
      "message": "Something is wrong on the server side",
      "anotherMessage": err.Error(),
    }
    return c.JSON(http.StatusInternalServerError, response)
  }

  var response struct {
    Data struct {
      AlbumId string `json:"albumId"`
      Songs []models.Song `json:"songs"` 
    } `json:"data"`
  }
  response.Data.AlbumId = id
  response.Data.Songs = songs

  return c.JSON(http.StatusOK, response)
}

func PostSongInAlbum(c echo.Context) error {

  id := c.Param("id")
  song := new(models.Song)

  if err := c.Bind(song); err != nil {
    response := map[string]string {
      "message": "Malformed payload",
      "anotherMessage": err.Error(),
    }
    return c.JSON(http.StatusBadRequest, response)
  }

  songId, err := db.CreateOneSongInAlbum(id, song);
  if err != nil {
    var notFoundError *myerror.NoData
    if errors.As(err, &notFoundError) {
      return c.JSON(http.StatusNotFound, notFoundError)
    }
    response := map[string]string {
      "message": "Cannot POST /albums/:id/songs",
      "anotherMessage": err.Error(),
    }
    return c.JSON(http.StatusInternalServerError, response)
  }

  var response struct {
    Data struct {
      AlbumId string `json:"albumId"`
      SongId string `json:"songId"`
    } `json:"data"`
  }

  response.Data.AlbumId = id
  response.Data.SongId = songId

  return c.JSON(http.StatusCreated, response)
}
