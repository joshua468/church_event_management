package model

import (
	"time"
)

type User struct {
ID 	    	  uint 	    `gorm:"primeryKey" json:"id"`
Username 	 string     `gorm:"unique" json:"username"`
Email    	 string     `gorm:"unique" json:"email"`
Password 	 string	   	`json:"_"`
Role      	 string	   	`json:"role"`
CreatedAt	 time.Time	`json:"createdat"`
}