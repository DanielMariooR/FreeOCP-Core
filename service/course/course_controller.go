package course

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	custom_validator "gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/databases/validator"
	"gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/models"
	"gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/models/pagination"
)

type CourseController struct {
	courseService CourseService
}

func NewController(serv CourseService) *CourseController {
	return &CourseController{
		courseService: serv,
	}
}

func (ctl *CourseController) HandleGetCourseData(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")

	course, err := ctl.courseService.GetCourseDetail(ctx, id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, course)
}

func (ctl *CourseController) HandleGetCompletedCoursePagination(c echo.Context) error {
	ctx := c.Request().Context()
	userId := c.Get("userId").(string)

	var resp pagination.PaginationResponse

	meta := pagination.Meta{}
	meta.FromContext(c)

	items, count, err := ctl.courseService.GetCompeletedCourse(ctx, &meta, userId)
	if err != nil {
		return err
	}

	resp = pagination.PaginationResponse{
		Meta: &pagination.Meta{
			Count: count,
			Page:  meta.Page,
			Limit: meta.Limit,
		},
		Items: items,
	}

	resp.SetTotalPage()

	return c.JSON(http.StatusOK, resp)
}

func (ctl *CourseController) HandleGetOnProgressCoursePagination(c echo.Context) error {
	ctx := c.Request().Context()
	userId := c.Get("userId").(string)

	var resp pagination.PaginationResponse

	meta := pagination.Meta{}
	meta.FromContext(c)

	items, count, err := ctl.courseService.GetOnProgressCourse(ctx, &meta, userId)
	if err != nil {
		return err
	}

	resp = pagination.PaginationResponse{
		Meta: &pagination.Meta{
			Count: count,
			Page:  meta.Page,
			Limit: meta.Limit,
		},
		Items: items,
	}

	resp.SetTotalPage()

	return c.JSON(http.StatusOK, resp)
}

func (ctl *CourseController) HandleGetCoursePagination(c echo.Context) error {
	ctx := c.Request().Context()

	var resp pagination.PaginationResponse

	meta := pagination.Meta{}
	meta.FromContext(c)

	items, count, err := ctl.courseService.GetCoursePagination(ctx, &meta)
	if err != nil {
		return err
	}

	resp = pagination.PaginationResponse{
		Meta: &pagination.Meta{
			Count: count,
			Page:  meta.Page,
			Limit: meta.Limit,
		},
		Items: items,
	}

	resp.SetTotalPage()

	return c.JSON(http.StatusOK, resp)
}

func (ctl *CourseController) HandleGetCourseSyllabus(c echo.Context) error {
	ctx := c.Request().Context()
	courseId := c.Param("id")

	resp, err := ctl.courseService.GetCourseSyllabus(ctx, courseId)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

func (ctl *CourseController) HandleGetMaterial(c echo.Context) error {
	ctx := c.Request().Context()
	courseId := c.Param("id")
	sectionId := c.Param("sectId")

	resp, err := ctl.courseService.GetCourseMaterial(ctx, courseId, sectionId)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

func (ctl *CourseController) HandleEnroll(c echo.Context) error {
	ctx := c.Request().Context()
	userId := c.Get("userId").(string)
	courseId := c.Param("id")

	resp, err := ctl.courseService.Enroll(ctx, userId, courseId)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, resp)
}

func (ctl *CourseController) HandleStoreProgress(c echo.Context) error {
	ctx := c.Request().Context()

	materialId := c.Param("id")
	userId := c.Get("userId").(string)

	resp, err := ctl.courseService.StoreUserProgress(ctx, userId, materialId)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

func (ctl *CourseController) HandleGetUserProgressPercentage(c echo.Context) error {
	ctx := c.Request().Context()

	courseId := c.Param("id")
	userId := c.Get("userId").(string)

	resp, err := ctl.courseService.ComputeUserProgress(ctx, userId, courseId)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

func (ctl *CourseController) HandleGetUserProgress(c echo.Context) error {
	ctx := c.Request().Context()

	courseId := c.Param("id")
	userId := c.Get("userId").(string)

	resp, err := ctl.courseService.GetUserProgress(ctx, userId, courseId)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

func (ctl *CourseController) HandleCourseCreation(c echo.Context) error {
	ctx := c.Request().Context()

	userId := c.Get("userId").(string)
	
	input := new(models.CourseCreation)
	if err := c.Bind(input); err != nil {
		return err
	}

	if err := c.Validate(input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, custom_validator.BuildCustomErrors((err)))
	}

	input.Course.Creator = userId
	resp, err := ctl.courseService.CreateNewCourse(ctx, input)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

func (ctl *CourseController) HandleUploadImage(c echo.Context) error {
	ctx := c.Request().Context()
	req := c.Request()
	url := "http://" + c.Request().Host + "/"

	fmt.Println("url from c: ", url)

	resp, err := ctl.courseService.UploadImage(ctx, req, url);
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

func (ctl *CourseController) HandleCreateDescription(c echo.Context) error {
	ctx := c.Request().Context()
	userId := c.Get("userId").(string)

	input := new(models.CourseDescriptionInput)
	if err := c.Bind(input); err != nil {
		return err
	}

	if err := c.Validate(input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, custom_validator.BuildCustomErrors((err)))
	}

	resp, err := ctl.courseService.CreateCourseDesc(ctx, input, userId)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)	
}

func (ctl *CourseController) HandleCreateSection(c echo.Context) error {
	ctx := c.Request().Context()
	userId := c.Get("userId").(string)

	input := new(models.CourseSectionInput)
	if err := c.Bind(input); err != nil {
		return err
	}

	if err := c.Validate(input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, custom_validator.BuildCustomErrors((err)))
	}

	resp, err := ctl.courseService.CreateCourseSection(ctx, input, userId)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)	
}

func (ctl *CourseController) HandleCreateMaterial(c echo.Context) error {
	ctx := c.Request().Context()
	userId := c.Get("userId").(string)

	input := new(models.CourseMaterialInput)
	if err := c.Bind(input); err != nil {
		return err
	}

	if err := c.Validate(input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, custom_validator.BuildCustomErrors((err)))
	}

	resp, err := ctl.courseService.CreateCourseMaterial(ctx, input, userId)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)	
}

func (ctl *CourseController) HandleCourseByCreatorId(c echo.Context) error {
	ctx := c.Request().Context()
	userId  := c.Param("userID")

	resp, err := ctl.courseService.GetCourseByCreatorID(ctx, userId)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

func (ctl *CourseController) HandleContributedCourse(c echo.Context) error {
	ctx := c.Request().Context()
	userId  := c.Get("userId").(string)

	resp, err := ctl.courseService.GetCourseByCreatorID(ctx, userId)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}