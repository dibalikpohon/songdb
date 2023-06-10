package controller

import (
	"errors"
	"net/http"
	myerror "songdb/pkg/errors"
	"songdb/pkg/models"
	"songdb/pkg/service"

	"github.com/labstack/echo/v4"
)

type SongControllerImpl struct {
  service service.SongService
}

func (si SongControllerImpl) GetAll(c echo.Context) error {

  songs, err := si.service.ReadAll();
  if err != nil {
    response := map[string]string {
      "message": "Something is wrong on the server side",
      "anotherMessage": err.Error(),
    }
    return c.JSON(http.StatusInternalServerError, response)
  }

  return c.JSON(http.StatusOK, SongListResponse{songs})
}

func (si SongControllerImpl) GetById(c echo.Context) error {

  id := c.Param("id")
  song, err := si.service.ReadOne(id)

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

  return c.JSON(http.StatusOK, SongSingleResponse{song})
}

func (si SongControllerImpl) Post(c echo.Context) error {

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

  id, err := si.service.Create(dto);
  if err != nil {
    response := map[string]string {
      "message": "Failed to POST /song",
      "anotherMessage": err.Error(),
    }
    return c.JSON(http.StatusInternalServerError, response)
  }

  response := SongIdResponse{}
  response.Data.Id = id

  return c.JSON(http.StatusCreated, response)
}

func (si SongControllerImpl) Put(c echo.Context) error {

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

  err := si.service.Update(id, dto);
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

  return c.NoContent(http.StatusNoContent)
}

func (si SongControllerImpl) Delete(c echo.Context) error {

  id := c.Param("id")
  err := si.service.Delete(id);
  
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
