package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/supertokens/supertokens-golang/recipe/session"
)

func CurrentUserID(c *gin.Context) (userID string) {
	sessionContainer := session.GetSessionFromRequestContext(c.Request.Context())
	userID = sessionContainer.GetUserID()

	return
}
