package imagehandler

import (
	"context"
	"detectify/solution/models"
)

//Usecase represent the image's handler usecases
type UseCase interface {
	AddImage(parCtx context.Context, parURLReq *models.SitesRequest) error
	UpdateImage(parCtx context.Context, parImg *models.DetectifyImage) error
	GetImageByID(parCtx context.Context, parID int64) (*models.DetectifyImage, error)
}
