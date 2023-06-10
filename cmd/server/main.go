package main

import (
  "strings"
  "reflect"
	"net/http"
	"songdb/pkg/config"

	myerror "songdb/pkg/errors"
	myvalidator "songdb/pkg/validator"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
  validator := validator.New();
  // FieldError .Field function returns struct name instead of JSON tag.
  // https://github.com/go-playground/validator/issues/337#issuecomment-357731246
  validator.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	e.Validator = myvalidator.NewTheValidator(validator)
	e.HTTPErrorHandler = echoHttpErrorHandler

	db, err := config.GetDb()
	if err != nil {
		panic(err.Error())
	}
	songRoutes := InitializeSongRoutes(db)
	albumRoutes := InitializeAlbumRoutes(db)
	songRelRoutes := InitializeSongRelRoutes(db)

	songRoutes.Register(e)
	albumRoutes.Register(e)
	songRelRoutes.Register(e)

	e.Logger.Fatal(e.Start(":9000"))
}

func echoHttpErrorHandler(err error, c echo.Context) {

  if httpError, ok := err.(*echo.HTTPError); ok {
    handleHttpError(httpError, c)
  } else if validationErrors, ok := err.(validator.ValidationErrors); ok {
    handleValidationErrors(validationErrors, c)
  } else {
    c.Error(echo.NewHTTPError(http.StatusInternalServerError, "Unknown error"))
  }
}

func handleHttpError(httpError *echo.HTTPError, c echo.Context) {

  c.JSON(httpError.Code, httpError.Error())
}

func handleValidationErrors(validationErrors validator.ValidationErrors, c echo.Context) {

  fieldValidationError := myerror.NewFieldValidationError();
  for _, err := range validationErrors {
    switch err.Tag() {
    case "printascii":
      fieldValidationError.AppendError(err.Field(), "Contains non-printable ASCII character(s)")
    case "number":
      fieldValidationError.AppendError(err.Field(), "Contains non-number character(s)")
    case "gte":
      fieldValidationError.AppendError(err.Field(), "Subceeds the limit")
    case "lte":
      fieldValidationError.AppendError(err.Field(), "Exceeds the limit")
    case "max":
      fieldValidationError.AppendError(err.Field(), "Character limit exceeded")
    }
  }
  c.JSON(http.StatusBadRequest, fieldValidationError)
}
