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

type TransactionReq struct {
	Amount    float32   `json:"amount"`
	CreatedAt time.Time `json:"timestamp"`
	UserId    string    `json:"userId"`
}

type Transaction struct {
	Id        string    `json:"id"`
	Amount    float32   `json:"amount"`
	CreatedAt time.Time `json:"createdAt"`
	IsDeleted bool      `json:"isDeleted"`
	UserId    string    `json:"userId"`
}

type LocationReq struct {
	City string `json:"city"`
}

type UserDetails struct {
	Id   string `json:"id"`
	City string `json:"city"`
}

type TransactionDetails struct {
	Max   float32 `json:"max"`
	Min   float32 `json:"min"`
	Avg   float32 `json:"avg"`
	Sum   float32 `json:"sum"`
	Count int     `json:"count"`
}
