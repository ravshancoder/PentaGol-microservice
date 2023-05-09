package v1

import (
	"context"
	"net/http"
	"strconv"

	"github.com/PentaGol/api_getway/api/handlers/models"
	pu "github.com/PentaGol/api_getway/genproto/admin"
	"github.com/PentaGol/api_getway/pkg/etc"
	l "github.com/PentaGol/api_getway/pkg/logger"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
)

// @Summary create admin
// @Description This api creates a admin
// @Tags Admin
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param body body models.AdminRequest true "Create Admin"
// @Success 200 {object} models.Admin
// @Failure 400 {object} models.StandartErrorModel
// @Failure 500 {object} models.StandartErrorModel
// @Router /v1/admin [post]
func (h *handlerV1) CreateAdmin(c *gin.Context) {
	var (
		body        models.AdminRequest
		jspbMarshal protojson.MarshalOptions
	)

	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("Failed to bind json: ", l.Error(err))
		return
	}

	tokens, err := h.jwtHandler.GenerateAuthJWT()
	accessToken := tokens[0]
	refreshToken := tokens[1]
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("error while generating tokens", l.Error(err))
		return
	}

	password, err := etc.HashPassword(body.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to hash password", l.Error(err))
		return
	}

	response, err := h.serviceManager.AdminService().CreateAdmin(context.Background(), &pu.AdminRequest{
		Name:         body.Name,
		Email:        body.Email,
		Password:     password,
		AccesToken:   accessToken,
		RefreshToken: refreshToken,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create admin", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// @Summary get admin by id
// @Description This api gets a admin by id
// @Tags Admin
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path string true "Id"
// @Success 200 {object} models.Admin
// @Failure 400 {object} models.StandartErrorModel
// @Failure 500 {object} models.StandartErrorModel
// @Router /v1/admin/{id} [get]
func (h *handlerV1) GetAdminById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	response, err := h.serviceManager.AdminService().GetAdminById(context.Background(), &pu.IdRequest{Id: int64(id)})
	if err != nil {
		statusCode := http.StatusInternalServerError
		if status.Code(err) == codes.NotFound {
			statusCode = http.StatusNotFound
		}
		c.JSON(statusCode, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get admin by id: ", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}
