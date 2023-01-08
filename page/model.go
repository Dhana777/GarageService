package page

import (
	"time"

	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type UserReq struct {
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Phonenumber string `json:"phonenumber"`
	Password    string `json:"password"`
}

type User struct {
	Id          string
	UserName    string
	Phonenumber string
	Password    string
}

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}

type CarRequest struct {
	CarNumber   string
	PhoneNumber string
}

type Car struct {
	Model             string    `json:"model"`
	DateOfManufacture time.Time `json:"date_of_manufacture"`
	LastServicedDate  time.Time `json:"last_serviced_date"`
	Id                string    `json:"id"`
	LastUsedDate      time.Time `json:"last_used_date"`
}

type UserLoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
