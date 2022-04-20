package assignment

import (
	"net/http"

	"github.com/labstack/echo/v4"
	custom_validator "gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/databases/validator"
	"gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/models"
)

type AssignmentController struct {
	service AssignmentService
}

func NewController(svc AssignmentService) *AssignmentController {
	return &AssignmentController{
		service: svc,
	}
}

func (ctl *AssignmentController) HandleGetAssignment(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")

	resp, err := ctl.service.GetAssignment(ctx, id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

func (ctl *AssignmentController) HandleCreateAssignment(c echo.Context) error {
	ctx := c.Request().Context()

	input := new(models.AssignmentCreation)
	if err := c.Bind(input); err != nil {
		return err
	}

	if err := c.Validate(input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, custom_validator.BuildCustomErrors((err)))
	}

	userId := c.Get("userId").(string)
	input.Desc.Creator = userId
	resp, err := ctl.service.CreateAssignment(ctx, input)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

func (ctl *AssignmentController) HandleGetScore(c echo.Context) error {
	ctx := c.Request().Context()
	userId := c.Get("userId").(string)

	input := new(models.AssignmentSubmission)
	if err := c.Bind(input); err != nil {
		return err
	}

	if err := c.Validate(input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, custom_validator.BuildCustomErrors((err)))
	}

	resp, err := ctl.service.GetScore(ctx, userId, input)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, resp)
}