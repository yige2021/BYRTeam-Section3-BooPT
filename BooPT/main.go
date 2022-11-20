package main

import (
	"BooPT/config"
	"BooPT/database"
	r "BooPT/router"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/sirupsen/logrus"
)

func main() {

	/* read config file
	 * 读取根目录下的配置文件config.yaml
	 * config.Read函数返回读取到的文件的正确性
	 * 如果文件读入有问题，就输出错误代码（方式：logrus输出日志）并且退出程序
	 */
	if config.Read("./config.yaml") != nil {
		logrus.Println("Error reading config file")
		os.Exit(1)
	}

	/* get database connection
	 * 连接数据库，若连接失败，则输出错误日志；若成功，则创建相应的表
	 */
	if database.Connect() != nil {
		logrus.Println("Error connecting to database")
		os.Exit(1)
	}

	// init app
	app := fiber.New()

	//调用日志中间件
	app.Use(logger.New())

	//实现跨域访问
	app.Use((cors.New(cors.Config{
		AllowOrigins: "http://localhost:8080",
	})))

	/* public routes
	 * 设定公共路由，也就是登录注册
	 */
	setPublicRoutes(app)

	/* jwt middleware
	 * jwt中间件
	 */
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: config.JWTSALT,
		ContextKey: "jwt",
	}))

	// private routes
	setPrivateRoutes(app)

	err := app.Listen(":3000")
	if err != nil {
		logrus.Errorf("Error: %v ", err)
		os.Exit(1)
	}

}

func setPublicRoutes(app *fiber.App) {
	//创建群组路由api
	api := app.Group("/api")
	r.SetupAccountRouterPub(api)
}

func setPrivateRoutes(app *fiber.App) {
	api := app.Group("/api")
	r.SetupBookRouter(api)
	r.SetupLinkRouter(api)
	r.SetupTypeRouter(api)
	r.SetupTagRouter(api)
	r.SetupS3Router(api)
}

/*
debug

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World 👋!")
    })
*/
