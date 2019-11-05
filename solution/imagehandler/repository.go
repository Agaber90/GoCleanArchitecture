package imagehandler

import (
	"context"
	"detectify/solution/models"
)

//Repository Would be called on the Repofile to implement
//the AddImageOperation, and get Image
type Repository interface {
	AddImage(parCtx context.Context, parImg *models.DetectifyImage) error
	UpdateImage(parCtx context.Context, parImg *models.DetectifyImage) error
	GetImageByID(parCtx context.Context, parID int64) (*models.DetectifyImage, error)
}
