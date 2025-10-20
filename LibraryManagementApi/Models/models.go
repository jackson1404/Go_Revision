package models

import (
	"time"

	"gorm.io/gorm"
)

// User (Library member)
type User struct {
	gorm.Model
	Username     string        `gorm:"size:100;uniqueIndex;not null"`
	Email        string        `gorm:"size:150;uniqueIndex;not null"`
	PasswordHash string        `gorm:"size:255;not null"` // store hashed
	FullName     string        `gorm:"size:200"`
	Phone        string        `gorm:"size:30"`
	Loans        []Loan        `gorm:"foreignKey:UserID"`
	Reservations []Reservation `gorm:"foreignKey:UserID"`
}

// Category
type Category struct {
	gorm.Model
	Name  string `gorm:"size:100;uniqueIndex;not null"`
	Slug  string `gorm:"size:120;uniqueIndex;not null"`
	Books []Book `gorm:"foreignKey:CategoryID"`
}

// Author
type Author struct {
	gorm.Model
	Name  string `gorm:"size:150;not null;index"`
	Bio   string `gorm:"type:text"`
	Books []Book `gorm:"many2many:book_authors"`
}

// Book
type Book struct {
	gorm.Model
	Title        string        `gorm:"size:255;not null;index"`
	ISBN         string        `gorm:"size:20;uniqueIndex;not null"`
	Description  string        `gorm:"type:text"`
	CategoryID   uint          `gorm:"index;not null"`
	Category     Category      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Authors      []Author      `gorm:"many2many:book_authors"`
	TotalCopies  uint          `gorm:"default:1;not null"`
	Available    int           `gorm:"default:1;not null;index"` // current available count
	Loans        []Loan        `gorm:"foreignKey:BookID"`
	Reservations []Reservation `gorm:"foreignKey:BookID"`
}

// Loan (a borrow record)
// ReturnedAt = nil => currently on loan
type Loan struct {
	gorm.Model
	UserID     uint      `gorm:"index;not null"`
	User       User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	BookID     uint      `gorm:"index;not null"`
	Book       Book      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	BorrowedAt time.Time `gorm:"autoCreateTime"`
	DueAt      time.Time `gorm:"index"`
	ReturnedAt *time.Time
	FineCents  int64 `gorm:"default:0"` // store calculation at return time (in cents)
}

// Reservation (user reserves a book)
type Reservation struct {
	gorm.Model
	UserID     uint       `gorm:"index;not null"`
	User       User       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	BookID     uint       `gorm:"index;not null"`
	Book       Book       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ReservedAt time.Time  `gorm:"autoCreateTime"`
	Notified   bool       `gorm:"default:false"` // whether user was notified it is available
	ExpiresAt  *time.Time // optional expiry for reservation
}
