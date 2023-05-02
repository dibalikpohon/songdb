package controller

import (
	"net/http"
  "errors"
	"songdb/pkg/db"
	"songdb/pkg/models"
  myerror "songdb/pkg/errors"

	"github.com/labstack/echo/v4"
)

func GetAllSongs(c echo.Context) error {
  
  songs, err := db.ReadSongsTable();
  if err != nil {
    response := map[string]string {
      "message": "Something is wrong on the server side",
      "anotherMessage": err.Error(),
    }
    return c.JSON(http.StatusInternalServerError, response)
  }

  return c.JSON(http.StatusOK, struct {
    Data []models.Song `json:"data"`
  }{songs}) 
}


func GetSongById(c echo.Context) error {
  
  id := c.Param("id")
  song, err := db.ReadOneSongById(id);

  if err != nil {
    // How to compare errors in Go
    // https://stackoverflow.com/a/57613539
    var notFoundError *myerror.NoData;
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
    Data *models.Song `json:"data"`
  }{song})
}

func PostSong(c echo.Context) error {
  
  song := new(models.Song)

  if err := c.Bind(song); err != nil {
    response := map[string]string {
      "message": "Malformed payload",
      "anotherMessage": err.Error(),
    }
    return c.JSON(http.StatusBadRequest, response)
  }

  id, err := db.CreateNewSong(song);
  if err != nil {
    response := map[string]string {
      "message": "Failed to POST /song",
      "anotherMessage": err.Error(),
    }
    return c.JSON(http.StatusInternalServerError, response)
  }

  var data struct {
    Data struct {
      Id string `json:"id"`
    } `json:"data"`
  };

  data.Data.Id = id
  
  return c.JSON(http.StatusCreated, data)
}

func PutSongUpdate(c echo.Context) error {

  id := c.Param("id")
  song := new(models.Song)

  if err := c.Bind(song); err != nil {
    response := map[string]string {
      "message": "Malformed payload",
      "anotherMessage": err.Error(),
    }
    return c.JSON(http.StatusBadRequest, response)
  }

  err := db.UpdateSongById(id, song);
  if err != nil {
    var notFoundError *myerror.NoData;
    if errors.As(err, &notFoundError) {
      return c.JSON(http.StatusNotFound, notFoundError)
    } else {
      response := map[string]string {
        "message": "Cannot PUT /songs/:id",
        "anotherMessage": err.Error(),
      }
      return c.JSON(http.StatusInternalServerError, response)
    }
  }

  return c.NoContent(http.StatusNoContent);
}

func DeleteSong(c echo.Context) error {

  id := c.Param("id")
  err := db.DeleteSongById(id);
  if err != nil {
    var notFoundError *myerror.NoData;
    if errors.As(err, &notFoundError) {
      return c.JSON(http.StatusNotFound, notFoundError)
    } else {
      response := map[string]string {
        "message": "Cannot DELETE /songs/:id",
        "anotherMessage": err.Error(),
      }
      return c.JSON(http.StatusInternalServerError, response)
    }
  }

  return c.NoContent(http.StatusNoContent);
}
