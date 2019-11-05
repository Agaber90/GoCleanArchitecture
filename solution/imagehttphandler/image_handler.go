package imagehttphandler

import (
	"context"
	"detectify/solution/imagehandler"
	"detectify/solution/models"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	validator "gopkg.in/go-playground/validator.v9"
)

var (
	// ErrInternalServerError will throw if any the Internal Server Error happen
	ErrInternalServerError = errors.New("Internal Server Error")
	// ErrNotFound will throw if the requested item is not exists
	ErrNotFound = errors.New("Your requested Item is not found")
	// ErrConflict will throw if the current action already exists
	ErrConflict = errors.New("Your Item already exist")
	// ErrBadParamInput will throw if the given request-body or params is not valid
	ErrBadParamInput = errors.New("Given Param is not valid")
)

//ResponseError represent the reseponse error struct
type ResponseError struct {
	Message string `json:"message"`
}

//ImageHandler  represent the httphandler for image hndler snippet
type ImageHandler struct {
	ImageBL imagehandler.UseCase
}

// NewImageHandler will initialize the NewImageHandler/ resources endpoint
func NewImageHandler(parEcho *echo.Echo, parImgBL imagehandler.UseCase) {
	pvImaghandler := &ImageHandler{
		ImageBL: parImgBL,
	}

	parEcho.POST("/images", pvImaghandler.AddImage)

}

// AddImage will store the snippetImage
func (img *ImageHandler) AddImage(parEco echo.Context) error {
	var pvError error
	var PvUrlsRequest models.SitesRequest
	pvBodyRequest, pvError := ioutil.ReadAll(parEco.Request().Body)
	if pvError != nil {
		return parEco.JSON(http.StatusBadRequest, pvError.Error())
	}

	if pvError = json.Unmarshal(pvBodyRequest, &PvUrlsRequest); pvError != nil {
		return parEco.JSON(http.StatusUnprocessableEntity, "Cannot serialize the JSON"+fmt.Sprint(pvError.Error()))
	}

	if len(PvUrlsRequest.URLS) == 0 {
		return parEco.JSON(http.StatusNoContent, "URLs cannot be empty")
	}

	if pvOK, PvError := isValidRequest(&PvUrlsRequest); !pvOK {
		return parEco.JSON(http.StatusBadRequest, PvError.Error())
	}
	pvCxt := parEco.Request().Context()
	if pvCxt != nil {
		pvCxt = context.Background()
	}

	if pvError := img.ImageBL.AddImage(pvCxt, &PvUrlsRequest); pvError != nil {
		return parEco.JSON(getStatusCode(pvError), ResponseError{Message: pvError.Error()})
	}

	return parEco.JSON(getStatusCode(pvError), "URLS have been snapped successfully and the images data has been added")
}

func isValidRequest(parModel *models.SitesRequest) (bool, error) {
	pvValidate := validator.New()
	pvError := pvValidate.Struct(parModel)
	if pvError != nil {
		return false, pvError
	}
	return true, nil
}

func getStatusCode(parError error) int {
	if parError == nil {
		return http.StatusOK
	}
	logrus.Error(parError)
	switch parError {
	case ErrInternalServerError:
		return http.StatusInternalServerError
	case ErrNotFound:
		return http.StatusNotFound
	case ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
