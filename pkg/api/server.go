package api

import (
	"fmt"
	"os"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	handler "github.com/donnyirianto/go-clean/pkg/api/handler"
	middleware "github.com/donnyirianto/go-clean/pkg/api/middleware"
	config "github.com/donnyirianto/go-clean/pkg/config"
	log "github.com/donnyirianto/go-clean/pkg/driver/log"
)

type Middlewares struct {
	ErrorHandler   *middleware.ErrorHandler
	Authentication *middleware.Authentication
}

type Handlers struct {
	UserHandler *handler.UserHandler
}

type ServerHTTP struct {
	app *fiber.App
}
type welcomeMessage struct {
	AppName string
	Message string
}

func NewServerHTTP(middlewares *Middlewares, handlers Handlers, log log.Logger, cfg config.Config) *ServerHTTP {
	app := fiber.New(
		fiber.Config{
			// NOTE: enable SO_REUSEPORT,
			// https://pkg.go.dev/github.com/valyala/fasthttp/reuseport, https://www.nginx.com/blog/socket-sharding-nginx-release-1-9-1/, https://github.com/gofiber/fiber/issues/180
			Prefork: cfg.Prefork,
			// NOTE: Override default JSON encoding, ref: https://docs.gofiber.io/guide/faster-fiber#custom-json-encoder-decoder
			JSONEncoder: sonic.Marshal,
			JSONDecoder: sonic.Unmarshal,
			// NOTE: Override default error handler
			ErrorHandler: middlewares.ErrorHandler.FiberErrorHandler(),
		})

	log.Info((fmt.Sprintf("Server is started with PID: %v and PPID: %v", os.Getpid(), os.Getppid())))

	// NOTE: Enable log tracing from Fiber, https://docs.gofiber.io/api/middleware/logger
	if cfg.Tracing {
		app.Use(logger.New())
	}

	if cfg.Recover {
		app.Use(recover.New())
	}

	// NOTE: Healthcheck
	Api := app.Group("/api")
	Api.Get("/", func(c *fiber.Ctx) error {
		m := welcomeMessage{"BE Clean Code", "Selamat Datang, ini adalah be go fiber clean code"}
		return c.JSON(m)
	})

	ApiV1 := Api.Group("/v1")
	// NOTE: Login route
	ApiV1.Get("/login", middleware.LoginHandler)

	// NOTE: User Api Route
	userAPI := ApiV1.Group("/users")
	userAPI.Get("/", handlers.UserHandler.FindAll)
	userAPI.Get("/:id<minLen(1)>", handlers.UserHandler.FindByID)
	userAPI.Post("/", handlers.UserHandler.Create)
	userAPI.Delete("/:id<minLen(1)>", handlers.UserHandler.Delete)
	userAPI.Put("/:id<minLen(1)>", handlers.UserHandler.Update)
	userAPI.Get("/name/:text<minLen(1)>", handlers.UserHandler.FindByMatchName)

	return &ServerHTTP{app}
}

func (sh *ServerHTTP) Start() {
	sh.app.Listen(":8080")
}
