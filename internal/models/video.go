package models

type Video struct {
	ID          uint64 `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Title       string `json:"title"`
	Description string `json:"description" gorm:"type:text"`
	CreatedBy   uint64 `json:"created_by"`
	VideoID     string `json:"video_id"`
}
