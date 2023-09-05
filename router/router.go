package router

import (
	"merchandise-review-list-backend/controller"
	"net/http"

	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(
	uc controller.IUserController,
) *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", os.Getenv("FE_URL")},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowCredentials: true,
	}))
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		CookiePath:     "/",
		CookieDomain:   os.Getenv("API_DOMAIN"),
		CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteNoneMode,
		// CookieSameSite: http.SameSiteDefaultMode, //PostMan使用する時に使用
		// CookieMaxAge: 60,
	}))

	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.LogIn)
	e.POST("/logout", uc.LogOut)
	e.GET("/csrf", uc.CsrfToken)

	u := e.Group("/user")
	u.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))

	u.Use((middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:    []byte(os.Getenv("SECRET")), // JWTの署名鍵を指定
		SigningMethod: "HS256",                     // 使用する署名アルゴリズムを指定
		ContextKey:    "token",                     // ユーザー情報を格納するコンテキストキーを指定
		TokenLookup:   "cookie:token",              // JWTを探す場所を指定
	})))

	// JWTが必須なエンドポイント
	u.GET("", uc.GetLoggedInUser)
	u.PUT("", uc.UpdateUser)
	u.DELETE("/:userId", uc.DeleteUser)

	p := e.Group("/posts")
	// JWTが必須なエンドポイント
	p.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))

	return e
}
