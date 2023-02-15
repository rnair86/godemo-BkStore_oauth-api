package app

import (
	"github.com/gin-gonic/gin"
	"github.com/rnair86/godemo-BkStore_oauth-api/src/domain/access_token"
	"github.com/rnair86/godemo-BkStore_oauth-api/src/http"
	"github.com/rnair86/godemo-BkStore_oauth-api/src/repository/db"
)
var(
	router = gin.Default()
)

func StartApplication() {
	dbRepository:=db.NewRepository()
	atservice := access_token.NewService(dbRepository)
	atHandler:=http.NewHandler(atservice)

	//router.GET

	router.GET("/oauth/access_token/:access_token_id",atHandler.GetById)
	router.Run(":8080")
}
