package models
import "time"
type Post Struct{
	Title string `gorm:"not null" json:"title"`

	ContentMD string `gorm:"type:text;not null" json:"content_md"`

	ContentHTML string `gorm:"type:text" json:"content_html"`
	
