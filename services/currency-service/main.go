package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"ariefsn/game-currency/global/helper"
	_ "ariefsn/game-currency/global/models/swagger"
	"ariefsn/game-currency/services/currency-service/controllers"
	"ariefsn/game-currency/services/currency-service/models"

	. "ariefsn/game-currency/services/currency-service/docs"

	logger "github.com/chi-middleware/logrus-logger"
	"github.com/common-nighthawk/go-figure"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @version 3.0

// @contact.name API Support
// @contact.url https://www.ariefsn.dev
// @contact.email ayiexz22@gmail.com

// @BasePath /

func main() {
	err := helper.InitEnv()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	err = helper.InitDB()

	if err != nil {
		log.Fatal("Init database failed. Error:", err)
	}

	helper.InitRender()

	env := helper.GetEnv()

	SwaggerInfo.Title = env.ServiceName
	SwaggerInfo.Description = fmt.Sprintf("%s API Documentation.", env.ServiceName)
	SwaggerInfo.BasePath = env.Swagger.BasePath

	// Migrate
	helper.GetDB().AutoMigrate(&models.CurrencyModel{}, &models.RateModel{})

	r := chi.NewRouter()

	r.Use(middleware.Heartbeat("/ping"))
	r.Use(logger.Logger("router", logrus.New()))
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(helper.SecureSetup().Handler)

	// Register Routes
	controllers.NewRootController().Register(r)
	controllers.NewCurrencyController().Register(r)
	controllers.NewRateController().Register(r)
	controllers.NewConvertController().Register(r)

	// Register Swagger
	r.Group(func(r chi.Router) {
		r.Use(helper.CorsSetup(true).Handler)

		docsUrl := fmt.Sprintf("%s%s:%s/docs/doc.json", env.Swagger.Protocol, env.Swagger.Host, env.Swagger.Port)

		if strings.ToLower(env.Swagger.Port) == "none" {
			docsUrl = fmt.Sprintf("%s%s/docs/doc.json", env.Swagger.Protocol, env.Swagger.Host)
		}

		r.Get("/docs/*", httpSwagger.Handler(
			httpSwagger.URL(docsUrl),
		))
	})

	address := fmt.Sprintf("%s:%s", env.Host, env.Port)

	addressLog := fmt.Sprintf("Start on %s\n", address)

	figure.NewColorFigure(env.ServiceName, "", "green", true).Print()

	fmt.Println("\n", addressLog)

	helper.RouterConfig(r)

	err = http.ListenAndServe(address, r)

	if err != nil {
		log.Fatal("Serve failed. Error", err)
		return
	}
}
