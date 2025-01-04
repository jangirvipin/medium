package models

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"size:100"`
	Email    string `gorm:"uniqueIndex"`
	Password string `gorm:"size:100"`
	Posts    []Post `gorm:"foreignKey:AuthorID"` // Explicitly define the foreign key
}

type Post struct {
	ID        uint   `gorm:"primaryKey"`
	Title     string `gorm:"size:100"`
	Content   string `gorm:"size:100"`
	Published bool   `gorm:"default:false"`
	AuthorID  uint   // Foreign key to reference User
	Author    User   `gorm:"foreignKey:AuthorID"`
}
