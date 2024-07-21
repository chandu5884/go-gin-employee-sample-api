package model

import "github.com/jinzhu/gorm"

type Employee struct {
    gorm.Model
    Name  string `json:"name"`
    Email string `json:"email"`
}
