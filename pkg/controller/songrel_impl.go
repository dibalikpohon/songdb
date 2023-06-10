package controller

import (
	"errors"
	"net/http"
	"songdb/pkg/models"
	"songdb/pkg/service"

	myerror "songdb/pkg/errors"

	"github.com/labstack/echo/v4"
)

type SongRelControllerImpl struct {
  service service.SongRelService
}

func (sri SongRelControllerImpl) GetAllSongsInAlbum(c echo.Context) error {
  
  id := c.Param("id")
  songs, err := sri.service.GetSongsInAlbum(id)
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

  response := AllSongsInAlbumResponse{}

  response.Data.AlbumId = id
  response.Data.Songs = songs

  return c.JSON(http.StatusOK, response)
}

func (sri SongRelControllerImpl) PostSongInAlbum(c echo.Context) error {

  id := c.Param("id")
  dto := new(models.SongDto)

  if err := c.Bind(dto); err != nil {
    response := map[string]string {
      "message": "Malformed payload",
      "anotherMessage": err.Error(),
    }
    return c.JSON(http.StatusBadRequest, response)
  }

  if err := c.Validate(dto); err != nil {
    return err;
  }

  songId, err := sri.service.CreateOneSongInAlbum(id, dto);
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

  response := SongIdInAlbumResponse{}
  response.Data.AlbumId = id
  response.Data.SongId = songId

  return c.JSON(http.StatusCreated, response)
}
