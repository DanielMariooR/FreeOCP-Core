package problem

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
	custom_validator "gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/databases/validator"
	"gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/models"
)

type ProblemController struct {
	problemService ProblemService
}

func NewController(svc ProblemService) *ProblemController {
	return &ProblemController{
		problemService: svc,
	}
}

func (ctl *ProblemController) HandleGetProblemData(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")

	problem, err := ctl.problemService.GetProblemCandidate(ctx, id)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(problem.Detail.(string)), &(problem.Detail))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, problem)
}

func Bind(c echo.Context, input *models.ProblemCreationInput) error {
	var bodyBytes []byte
	if c.Request().Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request().Body)
	}

	c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	jsonBody := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&jsonBody)
	if err != nil {
		return err
	}

	detail, _ := json.Marshal(jsonBody["content"])

	if string(detail) != "null" {
		input.Detail = string(detail)
	}

	// Restore the io.ReadCloser to its original state
	c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	if err := c.Bind(input); err != nil {
		return err
	}

	return nil
}

func (ctl *ProblemController) HandleCreateProblem(c echo.Context) error {
	ctx := c.Request().Context()

	input := new(models.ProblemCreationInput)
	if err := Bind(c, input); err != nil {
		return err
	}

	input.Creator = c.Get("userId").(string)
	if err := c.Validate(input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, custom_validator.BuildCustomErrors((err)))
	}

	resp, err := ctl.problemService.CreateNewProblem(ctx, input)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

func (ctl *ProblemController) HandleGetProblemStatus(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Get("userId").(string)

	problemStatus, err := ctl.problemService.GetProblemStatus(ctx, id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, problemStatus)
}

func (ctl *ProblemController) HandleGetProblemCandidateTable(c echo.Context) error {
	ctx := c.Request().Context()

	filter := models.ProblemFilter{}
	filter.FromContext(c)

	resp, err := ctl.problemService.GetProblemCandidateList(ctx, filter)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

func (ctl *ProblemController) HandleGetProblemContent(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")

	problem, err := ctl.problemService.GetProblemDetail(ctx, id)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(problem.Detail.(string)), &(problem.Detail))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, problem)
}

func (ctl *ProblemController) HandleAcceptProblem(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")

	resp, err := ctl.problemService.AcceptProblem(ctx, id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

func (ctl *ProblemController) HandleRejectProblem(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")

	resp, err := ctl.problemService.RejectProblem(ctx, id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

func (ctl *ProblemController) HandleGetProblem(c echo.Context) error {
	ctx := c.Request().Context()

	filter := models.ProblemFilter{}
	filter.FromContext(c)

	resp, err := ctl.problemService.GetProblemList(ctx, filter)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}
