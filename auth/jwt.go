package auth

import (
	"time"

	jwtapple2 "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/isnakolah/todoAPI/config"
	"github.com/isnakolah/todoAPI/model"
)

func SetupAuth() (*jwtapple2.GinJWTMiddleware, error) {
	authMiddleware, err := jwtapple2.New(&jwtapple2.GinJWTMiddleware{
		Realm:           "apitodogo",
		Key:             []byte(config.Key),
		Timeout:         time.Hour * 24,
		MaxRefresh:      time.Hour,
		IdentityKey:     config.IdentityKey,
		PayloadFunc:     payload,
		IdentityHandler: identityHandler,
		Authenticator:   authenticator,
		Authorizator:    authorizator,
		Unauthorized:    unauthorized,
		LoginResponse:   loginResponse,
		TokenLookup:     "header: Authorization, query: token, cookie: jwtapple2",
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,
	})

	return authMiddleware, err
}

func payload(data interface{}) jwtapple2.MapClaims {
	if v, ok := data.(*model.User); ok {
		return jwtapple2.MapClaims{
			config.IdentityKey: v.ID,
		}
	}

	return jwtapple2.MapClaims{}
}

func identityHandler(c *gin.Context) interface{} {
	claims := jwtapple2.ExtractClaims(c)
	var user model.User
	config.GetDB().Where("id = ?", claims[config.IdentityKey]).First(&user)

	return user
}

func authenticator(c *gin.Context) (interface{}, error) {
	var loginVals model.User
	if err := c.ShouldBind(&loginVals); err != nil {
		return "", jwtapple2.ErrMissingLoginValues
	}

	var result model.User
	config.GetDB().Where("username = ? AND password = ?",
		loginVals.Username, loginVals.Password).First(&result)

	if result.ID == 0 {
		return nil, jwtapple2.ErrFailedAuthentication
	}

	return &result, nil
}

func authorizator(data interface{}, c *gin.Context) bool {
	if v, ok := data.(model.User); ok && v.ID != 0 {
		return true
	}

	return false
}

func unauthorized(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"message": message,
	})
}

func loginResponse(c *gin.Context, code int, token string, expire time.Time) {
	c.JSON(code, gin.H{
		"expire": expire,
		"token":  token,
	})
}
