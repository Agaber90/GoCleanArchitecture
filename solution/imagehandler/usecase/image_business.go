package usecase

import (
	"context"

	"detectify/solution/imagehandler"
	"detectify/solution/models"
	"detectify/solution/pkg/helper"
	"log"
	"time"
)

type imageBusiness struct {
	imageRepo      imagehandler.Repository
	contextTimeout time.Duration
	imgeHelper     helper.ImageHelper
}

//NewImageBusiness will create new an articleUsecase object representation of imagehandler.UseCase interface
func NewImageBusiness(parImgeRepo imagehandler.Repository, parImgeHelper helper.ImageHelper, parTimeout time.Duration) imagehandler.UseCase {
	return &imageBusiness{
		imageRepo:      parImgeRepo,
		contextTimeout: parTimeout,
		imgeHelper:     parImgeHelper,
	}
}

func (imgBL *imageBusiness) AddImage(parCtx context.Context, parURLReq *models.SitesRequest) error {
	pvContext, pvCancel := context.WithTimeout(parCtx, imgBL.contextTimeout)
	var pvImageDate models.DetectifyImage
	//var PVImgHelper helper.ImageHelper
	defer pvCancel()
	go func() {
		select {
		case <-time.After(1000 * time.Second):
			log.Fatal("overslept")
		case <-parCtx.Done():
			log.Fatal(parCtx.Err())
		}
	}()

	imgBL.imgeHelper.StartTakeScreenShot(&imgBL.imgeHelper, parURLReq.URLS)

	for _, pvURLS := range parURLReq.URLS {
		pvImageDate.CreatedTime = time.Now()
		pvImageDate.Website = pvURLS
		pvImageDate.ImagePath = imgBL.imgeHelper.Path
		pvError := imgBL.imageRepo.AddImage(pvContext, &pvImageDate)
		if pvError != nil {
			return pvError
		}

	}
	return nil
}

func (imgBL *imageBusiness) UpdateImage(parCtx context.Context, parImg *models.DetectifyImage) error {
	pvContext, pvCancel := context.WithTimeout(parCtx, imgBL.contextTimeout)
	defer pvCancel()
	parImg.UpdatedTime = time.Now()
	return imgBL.imageRepo.UpdateImage(pvContext, parImg)
}

func (imgBL *imageBusiness) GetImageByID(parCtx context.Context, parID int64) (*models.DetectifyImage, error) {
	pvContext, pvCancel := context.WithTimeout(parCtx, imgBL.contextTimeout)
	defer pvCancel()
	pvResult, pvError := imgBL.imageRepo.GetImageByID(pvContext, parID)
	if pvError != nil {
		return nil, pvError
	}

	return pvResult, nil
}
