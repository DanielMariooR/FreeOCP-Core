package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
	custom_validator "gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/databases/validator"
	"gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/models"
)

type UserController struct {
	userService UserService
}

func NewController(svc UserService) *UserController {
	return &UserController{
		userService: svc,
	}
}

//function only for testing authentication api will be removed after testing
func (ctl *UserController) HandleGetUserData(c echo.Context) error {
	ctx := c.Request().Context()
	user, err := ctl.userService.GetUserData(ctx, "danielmario@testmail.com")
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func (ctl *UserController) HandleUserSignIn(c echo.Context) error {
	ctx := c.Request().Context()

	input := new(models.SignInInput)
	if err := c.Bind(input); err != nil {
		return err
	}

	if err := c.Validate(input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, custom_validator.BuildCustomErrors(err))
	}

	resp, err := ctl.userService.SignIn(ctx, input)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

func (ctl *UserController) HandleUserSignUp(c echo.Context) error {
	ctx := c.Request().Context()

	input := new(models.SignUpInput)
	if err := c.Bind(input); err != nil {
		return err
	}

	if err := c.Validate(input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, custom_validator.BuildCustomErrors((err)))
	}

	resp, err := ctl.userService.SignUp(ctx, input)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}
