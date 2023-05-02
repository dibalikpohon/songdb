package controller

import (
	"errors"
	"net/http"
	"songdb/pkg/db"
	myerror "songdb/pkg/errors"
	"songdb/pkg/models"

	"github.com/labstack/echo/v4"
)

func GetAllAlbums(c echo.Context) error {

  albums, err := db.ReadAlbumsTable();
  if err != nil {
    response := map[string]string {
      "message": "Something is wrong on the server side",
      "anotherMessage": err.Error(),
    }
    return c.JSON(http.StatusInternalServerError, response)
  }

  return c.JSON(http.StatusOK, struct {
    Data []models.Album `json:"data"`
  }{albums})
}

func GetAlbumById(c echo.Context) error {

  id := c.Param("id")
  album, err := db.ReadOneAlbumById(id)

  if err != nil {
    var notFoundError *myerror.NoData
    if errors.As(err, &notFoundError) {
      return c.JSON(http.StatusNotFound, notFoundError)
    } else {
      response := map[string]string {
        "message": "Something is wrong on the server side",
        "anotherMessage": err.Error(),
      }
      return c.JSON(http.StatusInternalServerError, response)
    }
  }

  return c.JSON(http.StatusOK, struct {
    Data *models.Album `json:"data"`
  }{album})
}

func PostAlbum(c echo.Context) error {

  album := new(models.Album)

  if err := c.Bind(album); err != nil {
    response := map[string]string {
      "message": "Malformed payload",
      "anotherMessage": err.Error(),
    }
    return c.JSON(http.StatusBadRequest, response)
  }

  id, err := db.CreateNewAlbum(album);
  if err != nil {
    response := map[string]string {
      "message": "Failed to POST /album",
      "anotherMessage": err.Error(),
    }
    return c.JSON(http.StatusInternalServerError, response)
  }

  var data struct {
    Data struct {
      Id string `json:"id"`
    } `json:"data"`
  }

  data.Data.Id = id;

  return c.JSON(http.StatusCreated, data)
}

func PutAlbumUpdate(c echo.Context) error {

  id := c.Param("id")
  album := new(models.Album)

  if err := c.Bind(album); err != nil {
    response := map[string]string {
      "message": "Malformed payload",
      "anotherMessage": err.Error(),
    }
    return c.JSON(http.StatusBadRequest, response)
  }

  err := db.UpdateAlbumById(id, album);
  if err != nil {
    var notFoundError *myerror.NoData;
    if errors.As(err, &notFoundError) {
      return c.JSON(http.StatusNotFound, notFoundError);
    } else {
      response := map[string]string {
        "message": "Cannot PUT /albums/:id",
        "anotherMessage": err.Error(),
      }
      return c.JSON(http.StatusInternalServerError, response)
    }
  }

  return c.NoContent(http.StatusNoContent);
}

func DeleteAlbum(c echo.Context) error {

  id := c.Param("id")
  err := db.DeleteAlbumById(id)
  if err != nil {
    var notFoundError *myerror.NoData;
    if errors.As(err, &notFoundError) {
      return c.JSON(http.StatusNotFound, notFoundError)
    } else {
      response := map[string]string {
        "message": "Cannot DELETE /albums/:id",
        "anotherMessage": err.Error(),
      }
      return c.JSON(http.StatusInternalServerError, response)
    }
  }

  return c.NoContent(http.StatusNoContent);
}
