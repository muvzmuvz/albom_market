package sturct

type Album struct {
	ID        string `json:"id" gorm:"type:uuid;primaryKey"`
	Title     string `json:"title" gorm:"unique;not null"`
	Artist    string `json:"artist" gorm:"not null"`
	Desc      string `json:"desc" gorm:"not null; default:''"`
	Price     int    `json:"price" gorm:"not null"`
	ImagePath string `json:"image_path" gorm:"not null;default:''"`
}
