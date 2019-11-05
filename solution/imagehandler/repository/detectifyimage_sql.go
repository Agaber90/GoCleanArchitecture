package repository

import (
	"context"
	"database/sql"
	"detectify/solution/imagehandler"
	"detectify/solution/models"
	"fmt"
	"time"
)

type sqlImageRepository struct {
	Conn *sql.DB
}

//NewImageRepository twill create an object that represent the image.Repository interface
func NewImageRepository(parConn *sql.DB) imagehandler.Repository {
	return &sqlImageRepository{parConn}
}

//AddImage will add a new snipped Image to the database
func (sql *sqlImageRepository) AddImage(parCtx context.Context, parImg *models.DetectifyImage) error {

	pvInsertquery := `INSERT detectify.imagedata SET ImagePath=? , Website=?, CreationDate=?`

	pvStmt, pvError := sql.Conn.PrepareContext(parCtx, pvInsertquery)
	if pvError != nil {
		return pvError
	}

	pvResult, pvError := pvStmt.ExecContext(parCtx, parImg.ImagePath, parImg.Website, time.Now())
	if pvError != nil {
		return pvError
	}

	pvLastID, pvError := pvResult.LastInsertId()
	if pvError != nil {
		return pvError
	}
	parImg.ID = pvLastID
	return nil
}

//UpdateImage will update an existing snippet
func (sql *sqlImageRepository) UpdateImage(parCtx context.Context, parImg *models.DetectifyImage) error {
	pvUpdateQuery := `UPDATE dbo.DetectifyTable set ImagePath=?, WebsitName=?, UpdateTime=? WHERE Id = ?`
	pvStmt, pvError := sql.Conn.PrepareContext(parCtx, pvUpdateQuery)
	if pvError != nil {
		return pvError
	}

	pvResult, pvError := pvStmt.ExecContext(parCtx, parImg.ImagePath, parImg.Website, parImg.UpdatedTime, parImg.ID)
	if pvError != nil {
		return pvError
	}

	pvRowsAffected, pvError := pvResult.RowsAffected()
	if pvError != nil {
		return pvError
	}

	if pvRowsAffected != 1 {
		pvError = fmt.Errorf("Total Affected: %d", pvRowsAffected)
	}

	return nil
}

func (sql *sqlImageRepository) GetImageByID(parCtx context.Context, parID int64) (img *models.DetectifyImage, err error) {
	PvFetchquery := `SELECT Id,ImagePath,WebsitName FROM article WHERE ID = ? ORDER BY ID`

	pvImageList, pvError := sql.fetch(parCtx, PvFetchquery, parID)
	if pvError != nil {
		return nil, pvError
	}

	if len(pvImageList) > 0 {
		img = pvImageList[0]
	} else {
		return nil, nil
	}
	return
}

func (sql *sqlImageRepository) fetch(parCtx context.Context, parQuery string, args ...interface{}) ([]*models.DetectifyImage, error) {
	pvRows, pvError := sql.Conn.QueryContext(parCtx, parQuery, args...)
	if pvError != nil {
		return nil, pvError
	}

	defer func() {
		pvError = pvRows.Close()
		if pvError != nil {
			panic(pvError)
		}
	}()

	pvResult := make([]*models.DetectifyImage, 0)
	for pvRows.Next() {
		pvTempResult := new(models.DetectifyImage)

		pvError = pvRows.Scan(&pvTempResult.ID, &pvTempResult.ImagePath, &pvTempResult.Website)
		if pvError != nil {
			return nil, pvError
		}
		pvResult = append(pvResult, pvTempResult)
	}
	return pvResult, nil
}
