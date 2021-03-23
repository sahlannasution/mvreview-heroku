package middlewares

import (
	"fmt"
	"log"
	"os"
	"time"

	logger "github.com/sahlannasution/mvreview-heroku/log"
	"github.com/sahlannasution/mvreview-heroku/models"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

type login struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// type users struct {
// 	ID       string
// 	FullName string
// 	Role     string
// }

var identityKey = "id"

var (
	loginVals login
	user      models.Users
)

func (StrDB *StrDB) MiddleWare() (mw *jwt.GinJWTMiddleware) {
	// dbPG := config.Connection()
	if err := godotenv.Load(".env"); err != nil {
		log.Println("ENV File Not Found!")
	}
	// the jwt middleware
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte(os.Getenv("SECRET_KEY")),
		Timeout:     time.Minute * 15, // SET TIME EXPIRED TO 15 MINUTE
		MaxRefresh:  time.Minute * 15, // SET TIME EXPIRED TO 15 MINUTE
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.Users); ok {
				return jwt.MapClaims{
					identityKey: v.Email,
					"fullName":  v.FullName,
					"role":      v.Role,
				}
			}
			return jwt.MapClaims{}
		},
		// IdentityHandler: func(c *gin.Context) interface{} {
		// 	claims := jwt.ExtractClaims(c)
		// 	return &users{
		// 		ID:       claims[identityKey].(string),
		// 		FullName: claims["fullName"].(string),
		// 		Role:     claims["role"].(string),
		// 	}
		// },
		Authenticator: func(c *gin.Context) (interface{}, error) {

			// baca dari json di raw data
			if err := c.ShouldBind(&loginVals); err != nil {
				logger.Sentry(err)
				return "", jwt.ErrMissingLoginValues
			}

			// check di database, apakah ada data yang dimaksud -> (inputan dari raw data)
			StrDB.DB.Where("email = ? ", loginVals.Email).First(&user)

			// compare password yang dimasukan di raw data apakah sama dengan yang ada di database
			if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginVals.Password)); err != nil {
				// log.Println(user)
				// log.Println(user.Password)
				// log.Println(loginVals.Password)
				logger.Sentry(err)
				log.Println("Password does not match!")
			} else { // ketika passwordnya match maka hit code dibawahnya.
				return &user, nil
			}
			// logger.Sentry(err)
			return nil, jwt.ErrFailedAuthentication
		},

		// menentukan role nya
		Authorizator: func(data interface{}, c *gin.Context) bool {
			fmt.Println()
			// method := c.Request.Method
			claims := jwt.ExtractClaims(c)
			// fmt.Println("Masuk authorizator", claims["role"])
			var result bool
			if claims["role"] == "admin" || claims["role"] == "guest" {
				result = true
			} else {
				result = false
			}
			// fmt.Println("Ini result nya", result)
			return result
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			logger.SentryStr(message)
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	}) // close authMidlleware

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
		logger.Sentry(err)
	}

	errInit := authMiddleware.MiddlewareInit()

	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
		logger.Sentry(errInit)
	}

	// fmt.Println("ini param valuenya : ", mw)
	// fmt.Println("ini returnya : ", authMiddleware)
	return authMiddleware
}
