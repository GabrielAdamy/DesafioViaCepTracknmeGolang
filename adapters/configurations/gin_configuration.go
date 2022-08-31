package configurations

import (
	"os"

	"github.com/gin-gonic/gin"
)

var ACCESS_SECRET string = os.Getenv("ACCESS_SECRET")

func NewHttpGinServer() *gin.Engine {
	ginServer := gin.Default()
	ginServer.Use(gin.Recovery())
	ginServer.Use(gin.Logger())

	return ginServer
}

// func validateToken(signedToken string) (jwt.MapClaims, string) {
// 	signedToken = strings.Split(signedToken, " ")[1]
// 	token, err := jwt.Parse(
// 		signedToken,
// 		func(token *jwt.Token) (interface{}, error) {
// 			return []byte(ACCESS_SECRET), nil
// 		},
// 	)

// 	if err != nil && err.Error() != "signature is invalid" {
// 		return nil, err.Error()
// 	}

// 	claims, ok := token.Claims.(jwt.MapClaims)
// 	if !ok {
// 		return nil, "the token is invalid"
// 	}

// 	return claims, ""
// }

// func Authenticate() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		if PROFILE != "" {
// 			clientToken := c.Request.Header.Get("Authorization")
// 			if clientToken == "" {
// 				c.JSON(http.StatusUnauthorized, gin.H{
// 					"message": "No Authorization header provided",
// 					"code":    http.StatusUnauthorized,
// 				})
// 				c.Abort()
// 				return
// 			}

// 			claims, err := validateToken(clientToken)
// 			if err != "" {
// 				c.JSON(http.StatusUnauthorized, gin.H{
// 					"message": err,
// 					"code":    http.StatusUnauthorized,
// 				})
// 				c.Abort()
// 				return
// 			}

// 			sub := claims["sub"].(string)
// 			c.Set("profile", strings.Split(sub, ":")[0])
// 			c.Set("user", strings.Split(sub, ":")[1])
// 		} else {
// 			c.Set("profile", "ADMINISTRATOR")
// 			c.Set("user", "0")
// 		}

// 		c.Next()
// 	}
//}
