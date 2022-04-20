package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/config"
	"gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/databases"
	custom_validator "gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/databases/validator"
	er "gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/error"
	"gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/helper"
	mid "gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/middleware"
	"gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/models"
	"gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/service/assignment"
	"gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/service/assignment/assignment_repository"
	"gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/service/course"
	"gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/service/course/course_repository"
	"gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/service/problem"
	"gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/service/problem/problem_repository"
	"gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/service/user"
	"gitlab.informatika.org/andrc1613/if3250_2022_08_freeocp/service/user/user_repository"
)

type App struct {
	DBManager *databases.Manager
	config    *config.Config
	E         *echo.Echo
}

func New(config *config.Config) *App {
	app := &App{
		config:    config,
		E:         echo.New(),
		DBManager: &databases.Manager{},
	}

	app.initErrorHandler()
	app.initValidator()
	app.initDatabase()
	app.initMiddleware()

	app.initRoutes()

	return app
}

func (app *App) initDatabase() {
	addr := fmt.Sprintf("%s:%s", app.config.DBHost, app.config.DBPort)
	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s", app.config.DBUser, app.config.DBPassword, addr, app.config.DBName)
	fmt.Println(connectionString)
	db, err := sqlx.Connect("mysql", connectionString)
	if err != nil {
		fmt.Println(err.Error())
	}

	app.DBManager.DB = db
}

// Fill new endpoint here
// use the middleware.DecodeJWTToken() to limit api for authenticated use only
func (app *App) initRoutes() {
	userRepository := user_repository.NewRepository()
	courseRepository := course_repository.NewRepository()
	problemRepository := problem_repository.NewRepository()
	assignmentRepository := assignment_repository.NewRepository()

	userService := user.NewService(app.DBManager.DB)
	_ = userService.InjectUserRepository(userRepository)

	courseService := course.NewService(app.DBManager.DB)
	_ = courseService.InjectCourseRepository(courseRepository)
	_ = courseService.InjectUserRepository(userRepository)

	problemService := problem.NewService(app.DBManager.DB)
	_ = problemService.InjectRepository(problemRepository)

	assignmentService := assignment.NewService(app.DBManager.DB)
	_ = assignmentService.InjectAssignmentRepository(assignmentRepository)

	userController := user.NewController(userService)
	app.E.GET("/", userController.HandleGetUserData, mid.DecodeJWTToken())

	user := app.E.Group("/v1/user")
	user.POST("/signin", userController.HandleUserSignIn)
	user.POST("/signup", userController.HandleUserSignUp)

	courseController := course.NewController(courseService)
	course := app.E.Group("/v1/course")
	course.GET("/", courseController.HandleGetCoursePagination)
	course.GET("/:id", courseController.HandleGetCourseData, mid.DecodeJWTToken())
	course.GET("/completed", courseController.HandleGetCompletedCoursePagination, mid.DecodeJWTToken())
	course.GET("/on-progress", courseController.HandleGetOnProgressCoursePagination, mid.DecodeJWTToken())
	course.GET("/syllabus/:id", courseController.HandleGetCourseSyllabus, mid.DecodeJWTToken())
	course.GET("/syllabus/:id/:sectId", courseController.HandleGetMaterial, mid.DecodeJWTToken())
	course.POST("/enroll/:id", courseController.HandleEnroll, mid.DecodeJWTToken())
	course.GET("/progress/:id", courseController.HandleGetUserProgress, mid.DecodeJWTToken())
	course.GET("/progress/percentage/:id", courseController.HandleGetUserProgressPercentage, mid.DecodeJWTToken())
	course.POST("/progress/:id", courseController.HandleStoreProgress, mid.DecodeJWTToken())
	// course.POST("/create", courseController.HandleCourseCreation, mid.DecodeJWTToken())
	course.POST("/create/description", courseController.HandleCreateDescription, mid.DecodeJWTToken())
	course.POST("/create/section", courseController.HandleCreateSection, mid.DecodeJWTToken())
	course.POST("/create/material", courseController.HandleCreateMaterial, mid.DecodeJWTToken())
	course.POST("/upload-image", courseController.HandleUploadImage)
	course.GET("/contributed", courseController.HandleContributedCourse, mid.DecodeJWTToken())
	course.GET("/contributed/:userID", courseController.HandleCourseByCreatorId)

	app.E.Static("/static", "static")

	problemController := problem.NewController(problemService)
	problem := app.E.Group("/v1/problem")
	problem.GET("/", problemController.HandleGetProblem, mid.DecodeJWTToken())
	problem.GET("/status", problemController.HandleGetProblemStatus, mid.DecodeJWTToken())
	problem.GET("/content/:id", problemController.HandleGetProblemContent, mid.DecodeJWTToken())
	problem.GET("/candidate/:id", problemController.HandleGetProblemData, mid.DecodeJWTToken())
	problem.GET("/candidate", problemController.HandleGetProblemCandidateTable, mid.DecodeJWTToken(), mid.VerifyAdmin())
	problem.POST("/create", problemController.HandleCreateProblem, mid.DecodeJWTToken())
	problem.PUT("/accept/:id", problemController.HandleAcceptProblem, mid.DecodeJWTToken(), mid.VerifyAdmin())
	problem.PUT("/reject/:id", problemController.HandleRejectProblem, mid.DecodeJWTToken(), mid.VerifyAdmin())

	assignmentController := assignment.NewController(assignmentService)
	assignment := app.E.Group("v1/assignment")
	assignment.GET("/:id", assignmentController.HandleGetAssignment, mid.DecodeJWTToken())
	assignment.POST("/create", assignmentController.HandleCreateAssignment, mid.DecodeJWTToken())
	assignment.POST("/:id", assignmentController.HandleGetScore, mid.DecodeJWTToken())
}

func (app *App) initValidator() {
	db := &databases.Manager{
		DB: app.DBManager.DB,
	}

	custom_validator.Init(app.E, db)
}

// Start the server and handle graceful shutdown
func (app *App) Start() {

	// Start server
	go func() {
		if err := app.E.Start(":" + app.config.AppPort); err != nil {
			app.E.Logger.Info("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	// Graceful Shutdown see: https://echo.labstack.com/cookbook/graceful-shutdown
	// Make sure no more in-flight request within 10seconds timeout
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := app.E.Shutdown(ctx); err != nil {
		app.E.Logger.Fatal(err)
	}
}

func (app *App) PreStop() {
	app.DBManager.DB.Close()
}

func main() {
	conf := config.GetConfig()
	app := New(&conf)

	defer app.PreStop()
	app.Start()
}

func (app *App) initErrorHandler() {
	app.E.HTTPErrorHandler = customHTTPErrorHandler
}

func (app *App) initMiddleware() {
	app.E.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))
}

func customHTTPErrorHandler(err error, c echo.Context) {
	var (
		response helper.RequestError
		errors   []custom_validator.CustomError
		errs     []helper.ErrorStruct
	)

	code := http.StatusInternalServerError
	message := "Server Error"

	switch t := err.(type) {
	case *models.CustomError:
		message = t.Message
		code = t.Code
	case *echo.HTTPError:
		code = t.Code
		switch m := t.Message.(type) {
		case []custom_validator.CustomError:
			errors = m
			message = "Validation Errors"
		default:
			message = fmt.Sprintf("%s", m)
		}
	case er.Error:
		message = t.Error()
		code = t.HTTPStatusCode()
		ers := t.GetErrors()
		if ers != nil {
			for _, e := range *ers {
				customEr := custom_validator.CustomError{
					Field:  e.Field,
					Reason: e.Reason,
				}
				errors = append(errors, customEr)
			}
		}
	default:
		message = t.Error()
	}

	if len(errors) > 0 {
		for _, v := range errors {
			errs = append(errs, helper.ErrorStruct{Field: v.Field, Reason: v.Reason})
		}

		response.Error = helper.RequestErrors{
			StatusCode: code,
			Message:    message,
			Errors:     errs,
		}
	} else {
		response.Error = helper.RequestSingleError{
			StatusCode: code,
			Message:    message,
		}
	}

	c.JSON(code, response)
}
