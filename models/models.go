package models

type User struct {
	ID       uint      `gorm:"primaryKey"`
	Segments []Segment `gorm:"many2many:user_segments"`
}

type Segment struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"uniqueIndex"`
	Users []User `gorm:"many2many:user_segments"`
}
