package services

import (
    "go-gin-api/model"
    "go-gin-api/repository"
)

type EmployeeService struct {
    EmployeeRepo repository.EmployeeRepository
}

func NewEmployeeService(repo repository.EmployeeRepository) *EmployeeService {
    return &EmployeeService{EmployeeRepo: repo}
}

func (s *EmployeeService) CreateEmployee(employee *model.Employee) error {
    return s.EmployeeRepo.Create(employee)
}

func (s *EmployeeService) GetEmployeeByID(id uint) (*model.Employee, error) {
    return s.EmployeeRepo.FindByID(id)
}
