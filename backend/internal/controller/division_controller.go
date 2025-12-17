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

// GetAll godoc
// @Summary      List divisions
// @Description  Get all divisions
// @Tags         divisions
// @Produce      json
// @Success      200  {array}   dto.DivisionDTO
// @Failure      500  {object}  util.ErrorResponse
// @Router       /api/divisions [get]
func (c *divisionControllerImpl) GetAll(w http.ResponseWriter, r *http.Request) {
	data, err := c.service.GetAll()
	if err != nil {
		c.log.Println("error get all divisions:", err)
		util.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.JSON(w, http.StatusOK, data)
}

// GetByID godoc
// @Summary      Get division by ID
// @Description  Get a division by its ID
// @Tags         divisions
// @Produce      json
// @Param        id   path      int  true  "Division ID"
// @Success      200  {object}  dto.DivisionDTO
// @Failure      400  {object}  util.ErrorResponse
// @Failure      404  {object}  util.ErrorResponse
// @Router       /api/divisions/{id} [get]
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

// Create godoc
// @Summary      Create division
// @Description  Create a new division
// @Tags         divisions
// @Accept       json
// @Produce      json
// @Param        request  body      dto.DivisionDTO  true  "Create division payload"
// @Success      201      {object}  dto.DivisionDTO
// @Failure      400      {object}  util.ErrorResponse
// @Router       /api/divisions [post]
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

// Update godoc
// @Summary      Update division
// @Description  Update a division by ID
// @Tags         divisions
// @Accept       json
// @Produce      json
// @Param        id       path      int             true  "Division ID"
// @Param        request  body      dto.DivisionDTO  true  "Update division payload"
// @Success      200      {object}  dto.DivisionDTO
// @Failure      400      {object}  util.ErrorResponse
// @Failure      404      {object}  util.ErrorResponse
// @Router       /api/divisions/{id} [put]
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

// Delete godoc
// @Summary      Delete division
// @Description  Delete a division by ID
// @Tags         divisions
// @Param        id   path  int  true  "Division ID"
// @Success      204  "No Content"
// @Failure      400  {object}  util.ErrorResponse
// @Failure      404  {object}  util.ErrorResponse
// @Router       /api/divisions/{id} [delete]
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
