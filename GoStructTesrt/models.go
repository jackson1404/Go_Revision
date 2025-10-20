package main

type User struct {
	Id        int
	Email     string
	Name      string
	ContactId int
	Contact   Contact `gorm:"foreignKey:ContactId;references:Id"`
}

type Contact struct {
	Id        int
	Handphone string
	Whatsapp  string
}
