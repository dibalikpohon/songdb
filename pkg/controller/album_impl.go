package controller

import (
	"errors"
	"net/http"
	myerror "songdb/pkg/errors"
	"songdb/pkg/models"
	"songdb/pkg/service"

	"github.com/labstack/echo/v4"
)

type AlbumControllerImpl struct {
  service service.AlbumService
}

func (ac AlbumControllerImpl) GetAll(c echo.Context) error {

  albums, err := ac.service.ReadAll(); 
  if err != nil {
    response := map[string]string {
      "message": "Something is wrong on the server side",
      "anotherMessage": err.Error(),
    }
    return c.JSON(http.StatusInternalServerError, response)
  }

  return c.JSON(http.StatusOK, AlbumListResponse{Data: albums})
}

func (ac AlbumControllerImpl) GetById(c echo.Context) error {

  id := c.Param("id")
  album, err := ac.service.ReadOne(id)

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

  return c.JSON(http.StatusOK, AlbumSingleResponse{Data: album})
}

func (ac AlbumControllerImpl) Post(c echo.Context) error {

  dto := new(models.AlbumDto)

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

  id, err := ac.service.Create(dto);
  if err != nil {
    response := map[string]string {
      "message": "Failed to POST /album",
      "anotherMessage": err.Error(),
    }
    return c.JSON(http.StatusInternalServerError, response)
  }

  response := AlbumIdResponse{}
  response.Data.Id = id

  return c.JSON(http.StatusCreated, response)
}

func (ac AlbumControllerImpl) Put(c echo.Context) error {
  
  id := c.Param("id")
  dto := new(models.AlbumDto)

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

  err := ac.service.Update(id, dto);
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

func (ac AlbumControllerImpl) Delete(c echo.Context) error {
  
  id := c.Param("id")
  err := ac.service.Delete(id)

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
