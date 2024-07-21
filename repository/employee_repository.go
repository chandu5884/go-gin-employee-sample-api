package repository

import (
    "github.com/jinzhu/gorm"
    "go-gin-api/model"
)

type EmployeeRepository interface {
    Create(employee *model.Employee) error
    FindByID(id uint) (*model.Employee, error)
}

type GormEmployeeRepository struct {
    DB *gorm.DB
}

func NewGormEmployeeRepository(db *gorm.DB) EmployeeRepository {
    return &GormEmployeeRepository{DB: db}
}

func (r *GormEmployeeRepository) Create(employee *model.Employee) error {
    return r.DB.Create(employee).Error
}

func (r *GormEmployeeRepository) FindByID(id uint) (*model.Employee, error) {
    var employee model.Employee
    if err := r.DB.First(&employee, id).Error; err != nil {
        return nil, err
    }
    return &employee, nil
}
