package http

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rnair86/godemo-BkStore_oauth-api/src/domain/access_token"
)

type AccessTokenHandler interface {
	GetById(*gin.Context)
	//GetById(string) (*access_token.AccessToken, *errors.RestErr)
}

type accesstokenHandler struct {
	service access_token.Service
}

func NewHandler(service access_token.Service) AccessTokenHandler {
	return &accesstokenHandler{service: service}
}

func (h *accesstokenHandler) GetById(g *gin.Context) {
	accessTokenId := strings.TrimSpace( g.Param("access_token_id"))
	accesstoken, err := h.service.GetById(accessTokenId)
	if err != nil {
		g.JSON(err.Status, err)
		return
	}
	g.JSON(http.StatusOK, accesstoken)

}
