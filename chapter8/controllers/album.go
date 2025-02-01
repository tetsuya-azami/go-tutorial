package controllers

import (
	"go-tutorial/chapter8/api"
	"go-tutorial/chapter8/app/models"
	"go-tutorial/chapter8/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AlbumHandler struct{}

func (a *AlbumHandler) CreateAlbum(c *gin.Context) {
	var requestBody api.CreateAlbumJSONRequestBody
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		logger.Warn(err.Error())
		c.JSON(http.StatusBadRequest, api.ErrorResponse{Message: err.Error()})
		return
	}

	createdAlbum, err := models.CreateAlbum(
		*requestBody.Title,
		requestBody.ReleaseDate.Time,
		string(requestBody.Category.Name))
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, api.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, createdAlbum)
}

func (a *AlbumHandler) GetAlbumById(c *gin.Context, ID int) {
	album, err := models.GetAlbum(ID)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, api.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, album)
}

// func (a *AlbumHandler) UpdateAlbumById(c *gin.Context, ID int) {
// 	var requestBody api.UpdateAlbumByIdJSONRequestBody
// }
