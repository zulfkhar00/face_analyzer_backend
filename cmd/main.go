package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/zulfkhar00/cosmetics-backend/internal/config"
	handler "github.com/zulfkhar00/cosmetics-backend/internal/handler/http"
	repository "github.com/zulfkhar00/cosmetics-backend/internal/repository/postgres"
	"github.com/zulfkhar00/cosmetics-backend/internal/service"
	"github.com/zulfkhar00/cosmetics-backend/pkg/apiclient"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("No .env file found: %v\n", err)
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=require",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.DBName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	log.Println("Successfully connected to the database!")

	postgresProductRepository := repository.NewProductRepository(db)
	postgresFaceRepository := repository.NewFaceRepository(db)
	openBeautyFactsClient := apiclient.NewOpenBeautyFactsClient(cfg.ExternalAPI.OpenBeautyFactsBaseURL, nil)

	productService := service.NewProductService(postgresProductRepository, openBeautyFactsClient)
	faceAnalysisServer := service.NewFaceAnalysisService(postgresFaceRepository)

	productHandler := handler.NewProductHandler(productService)
	faceAnalysisHandler := handler.NewFaceAnalysisHandler(faceAnalysisServer)

	// create Hertz server
	h := server.New(server.WithHostPorts(":" + cfg.Server.Port))

	// Set up routes
	h.GET("/product/:barcode", productHandler.GetProductByBarcode)
	h.POST("/face/analyze", faceAnalysisHandler.SendUserFaceImage)
	h.GET("/face/health_info", faceAnalysisHandler.GetUserFaceCondition)
	h.POST("/routine/add_product", faceAnalysisHandler.AddProductToRoutine)

	// Start server
	h.Spin()
}
