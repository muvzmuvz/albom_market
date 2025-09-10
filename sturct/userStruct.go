package sturct

type User struct {
	ID       string `json:"id" gorm:"type:uuid;primaryKey"`
	Role     string `json:"role" gorm:"not null;default:'user'"`
	Username string `json:"username" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
}
