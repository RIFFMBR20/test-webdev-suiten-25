package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"test-webdev-suiten-25/internal/service"
	"test-webdev-suiten-25/internal/util"
)

type DivisionController interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type divisionControllerImpl struct {
	service service.DivisionService
	log     *log.Logger
}

func ProvideDivisionController(
	service service.DivisionService,
	logger *log.Logger,
) DivisionController {
	return &divisionControllerImpl{
		service: service,
		log:     logger,
	}
}

func (c *divisionControllerImpl) GetAll(w http.ResponseWriter, r *http.Request) {
	data, err := c.service.GetAll()
	if err != nil {
		c.log.Println("error get all divisions:", err)
		util.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.JSON(w, http.StatusOK, data)
}

func (c *divisionControllerImpl) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/divisions/")
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

func (c *divisionControllerImpl) Create(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.Error(w, http.StatusBadRequest, "invalid body")
		return
	}

	data, err := c.service.Create(req.Name)
	if err != nil {
		util.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	util.JSON(w, http.StatusCreated, data)
}

func (c *divisionControllerImpl) Update(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/divisions/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		util.Error(w, http.StatusBadRequest, "invalid id")
		return
	}

	var req struct {
		Name string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.Error(w, http.StatusBadRequest, "invalid body")
		return
	}

	data, err := c.service.Update(id, req.Name)
	if err != nil {
		util.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	util.JSON(w, http.StatusOK, data)
}

func (c *divisionControllerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/divisions/")
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
