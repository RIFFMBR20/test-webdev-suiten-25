package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"test-webdev-suiten-25/internal/models/dto"
	"test-webdev-suiten-25/internal/service"
	"test-webdev-suiten-25/internal/util"
)

type EmployeeController interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type employeeControllerImpl struct {
	service service.EmployeeService
	log     *log.Logger
}

func ProvideEmployeeController(
	service service.EmployeeService,
	logger *log.Logger,
) EmployeeController {
	return &employeeControllerImpl{
		service: service,
		log:     logger,
	}
}

func (c *employeeControllerImpl) GetAll(w http.ResponseWriter, r *http.Request) {
	data, err := c.service.GetAll()
	if err != nil {
		c.log.Println("error get all employees:", err)
		util.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.JSON(w, http.StatusOK, data)
}

func (c *employeeControllerImpl) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/employees/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		util.Error(w, http.StatusBadRequest, "invalid id")
		return
	}

	data, err := c.service.GetByID(id)
	if err != nil {
		util.Error(w, http.StatusNotFound, err.Error())
		return
	}

	util.JSON(w, http.StatusOK, data)
}

func (c *employeeControllerImpl) Create(w http.ResponseWriter, r *http.Request) {
	var employeeDTO dto.EmployeeInputModifyDTO

	if err := json.NewDecoder(r.Body).Decode(&employeeDTO); err != nil {
		c.log.Printf("Error decoding employee body for create: %v\n", err)
		util.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	data, err := c.service.Upsert(0, employeeDTO)
	if err != nil {
		status := http.StatusBadRequest
		if strings.Contains(err.Error(), "not found") {
			status = http.StatusNotFound
		}

		c.log.Printf("Error creating employee: %v\n", err)
		util.Error(w, status, err.Error())
		return
	}

	util.JSON(w, http.StatusCreated, data)
}

func (c *employeeControllerImpl) Update(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/employees/")
	id, err := strconv.Atoi(idStr)
	if err != nil || id == 0 {
		util.Error(w, http.StatusBadRequest, "invalid or missing employee id for update")
		return
	}

	var employeeDTO dto.EmployeeInputModifyDTO

	if err := json.NewDecoder(r.Body).Decode(&employeeDTO); err != nil {
		c.log.Printf("Error decoding employee body for update: %v\n", err)
		util.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	data, err := c.service.Upsert(id, employeeDTO)
	if err != nil {
		status := http.StatusBadRequest
		if strings.Contains(err.Error(), "employee not found") {
			status = http.StatusNotFound
		} else if strings.Contains(err.Error(), "division not found") {
			status = http.StatusNotFound
		}

		c.log.Printf("Error updating employee (id=%d): %v\n", id, err)
		util.Error(w, status, err.Error())
		return
	}

	util.JSON(w, http.StatusOK, data)
}

func (c *employeeControllerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/employees/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		util.Error(w, http.StatusBadRequest, "invalid id")
		return
	}

	if err := c.service.Delete(id); err != nil {
		util.Error(w, http.StatusNotFound, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
