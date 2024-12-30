package entity

// User represents a user entity
type User struct {
	ID       uint   `gorm:"primaryKey;autoIncrement" json:"id"` // ID is the primary key
	Name     string `gorm:"size:255;not null" json:"name"`      // Name is a required field
	Email    string `gorm:"size:100;unique" json:"email"`       // Email must be unique
	Password string `gorm:"size:255;not null" json:"password"`  // Password is a required field
}
