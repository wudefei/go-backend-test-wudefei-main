package model

import "time"

type User struct {
	ID          int64     `gorm:"column:id"`          // user ID
	Name        string    `gorm:"column:name"`        // user name
	Birth       time.Time `gorm:"column:dob"`         // date of birth
	Address     string    `gorm:"column:address"`     // user address
	Description string    `gorm:"column:description"` // user description
	Phone       string    `json:"phone"`              // user phone
	CreatedAt   time.Time `gorm:"column:created_at"`  // user created date
	UpdatedAt   time.Time `gorm:"column:updated_at"`  // update time
	IsDeleted   int32     `gorm:"column:is_deleted"`  // soft delete
}

func (User) TableName() string {
	return "t_user"
}
