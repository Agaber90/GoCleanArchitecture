package models

import "time"

//DetectifyImage that would hold the data model
type DetectifyImage struct {
	ID          int64
	ImagePath   string
	Website     string
	CreatedTime time.Time
	UpdatedTime time.Time
}
