package handler

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "go-gin-api/model"
    "go-gin-api/service"
)

type EmployeeHandler struct {
    EmployeeService *services.EmployeeService
}

func NewEmployeeHandler(service *services.EmployeeService) *EmployeeHandler {
    return &EmployeeHandler{EmployeeService: service}
}

func (h *EmployeeHandler) CreateEmployee(c *gin.Context) {
    var employee model.Employee
    if err := c.ShouldBindJSON(&employee); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := h.EmployeeService.CreateEmployee(&employee); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, employee)
}

func (h *EmployeeHandler) GetEmployeeByID(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.ParseUint(idParam, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    employee, err := h.EmployeeService.GetEmployeeByID(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
        return
    }
    c.JSON(http.StatusOK, employee)
}
