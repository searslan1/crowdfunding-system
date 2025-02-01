package controller

import (
    "net/http"
    "strconv"
    "entrepreneur-profiles/internal/model"
    "entrepreneur-profiles/internal/service"
    "github.com/gin-gonic/gin"
)

type EntrepreneurController struct {
    service service.EntrepreneurService
}

func NewEntrepreneurController(service service.EntrepreneurService) *EntrepreneurController {
    return &EntrepreneurController{service: service}
}

func (ec *EntrepreneurController) CreateEntrepreneurProfile(c *gin.Context) {
    var profile model.EntrepreneurProfile
    if err := c.ShouldBindJSON(&profile); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := ec.service.CreateProfile(&profile); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, profile)
}

func (ec *EntrepreneurController) GetEntrepreneurProfile(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    profile, err := ec.service.GetProfileByID(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Profile not found"})
        return
    }

    c.JSON(http.StatusOK, profile)
}