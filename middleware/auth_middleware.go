package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/config"
	er "gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/error"
)

func DecodeJWTToken() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			header := c.Request().Header
			authv := header.Get("Authorization")

			if !strings.HasPrefix(strings.ToLower(authv), "bearer") {
				return er.NewError(fmt.Errorf("Forbidden Access"), http.StatusForbidden, nil)
			}

			values := strings.Split(authv, " ")
			if len(values) < 2 {
				return er.NewError(fmt.Errorf("Forbidden Access"), http.StatusForbidden, nil)
			}

			tokenString := values[1]
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, er.NewError(fmt.Errorf("Unexpected Signing Method"), http.StatusForbidden, nil)
				}

				return []byte(config.GetConfig().JWTSecret), nil
			})
			if err != nil {
				return er.NewError(fmt.Errorf("Forbidden Access"), http.StatusForbidden, nil)
			}

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				c.Set("userId", claims["id"])
				c.Set("userEmail", claims["email"])
				c.Set("isAdmin", claims["admin"])
			} else {
				return er.NewError(fmt.Errorf("Forbidden Access"), http.StatusForbidden, nil)
			}

			return next(c)
		}
	}
}

func VerifyAdmin() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			isAdmin := c.Get("isAdmin")
			if isAdmin != true {
				return er.NewError(fmt.Errorf("Forbidden Access"), http.StatusForbidden, nil)
			}

			return next(c)
		}
	}
}
