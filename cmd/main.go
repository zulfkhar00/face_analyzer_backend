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
	openBeautyFactsClient := apiclient.NewOpenBeautyFactsClient(cfg.ExternalAPI.OpenBeautyFactsBaseURL, nil)
	productService := service.NewProductService(postgresProductRepository, openBeautyFactsClient)
	productHandler := handler.NewProductHandler(productService)

	// create Hertz server
	h := server.New(server.WithHostPorts(":" + cfg.Server.Port))

	// Set up routes
	h.GET("/product/:barcode", productHandler.GetProductByBarcode)

	// Start server
	h.Spin()
}

// func getProductByBarcode(ctx context.Context, c *app.RequestContext) {
// 	barcode := c.Param("barcode")
//
// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=require",
// 		host, port, user, password, dbname)
//
// 	db, err := sql.Open("postgres", psqlInfo)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()
//
// 	productRepo := repository.NewProductRepository(db)
// 	product, err := productRepo.GetByBarcode(barcode)
//
// 	if err != nil {
// 		// try to fetch from OpenBeautyFacts
// 		log.Println(err)
// 		openBeautyFactsResp, err := getProductByBarcodeFromOpenBeautyFacts(barcode)
// 		if err != nil {
// 			var apiErr *ExternalAPIError
// 			if errors.As(err, &apiErr) {
// 				c.JSON(apiErr.StatusCode, map[string]string{"error": apiErr.Message})
// 			} else {
// 				c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal server error"})
// 			}
// 			return
// 		}
//
// 		ingredients := []string{}
// 		for _, ingredient := range openBeautyFactsResp.Product.Ingredients {
// 			ingredients = append(ingredients, ingredient.Text)
// 		}
//
// 		barcodeInt, err := strconv.Atoi(openBeautyFactsResp.Product.Code)
// 		if err != nil {
// 			c.JSON(consts.StatusInternalServerError, map[string]string{"err": "cannot convert barcode to int"})
// 			return
// 		}
//
// 		product := data.Product{
// 			ProductID:           -1,
// 			Barcode:             barcodeInt,
// 			ProductName:         openBeautyFactsResp.Product.ProductNameEn,
// 			Brand:               openBeautyFactsResp.Product.Brands,
// 			ProductQuantityUnit: openBeautyFactsResp.Product.ProductQuantityUnit,
// 			ProductQuantity:     openBeautyFactsResp.Product.ProductQuantity,
// 			Ingredients:         ingredients,
// 			Source:              "open_beauty_facts",
// 		}
//
// 		c.JSON(consts.StatusOK, product)
// 		return
// 	}
//
// 	c.JSON(consts.StatusOK, product)
// }

// func getProductByBarcodeFromOpenBeautyFacts(barcode string) (*data.GetProductAPIResponse, error) {
// 	client := http.DefaultClient
// 	url := fmt.Sprintf("https://world.openbeautyfacts.org/api/v2/product/%s.json", barcode)
// 	req, err := http.NewRequest(http.MethodGet, url, nil)
// 	if err != nil {
// 		log.Fatalf("Failed to create request: %v", err)
// 		return nil, &ExternalAPIError{
// 			StatusCode: http.StatusInternalServerError,
// 			Message:    "Couldn't create request to OpenBeautyFacts API",
// 			Err:        err,
// 		}
// 	}
// 	req.Header.Set("User-Agent", "PersonalCosmeticsApp/1.0 (olxzulfar@gmail.com)")
//
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		log.Fatalf("Request failed: %v", err)
// 		return nil, &ExternalAPIError{
// 			StatusCode: http.StatusInternalServerError,
// 			Message:    "API request failed",
// 			Err:        err,
// 		}
// 	}
// 	defer resp.Body.Close()
//
// 	var apiResp data.GetProductAPIResponse
// 	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
// 		log.Fatalf("Error decoding JSON: %v", err)
// 		return nil, &ExternalAPIError{
// 			StatusCode: http.StatusInternalServerError,
// 			Message:    "Error decoding JSON",
// 			Err:        err,
// 		}
// 	}
//
// 	if apiResp.Status != 1 {
// 		return nil, &ExternalAPIError{
// 			StatusCode: http.StatusNotFound,
// 			Message:    "Product not found",
// 			Err:        errors.New("product does not exist"),
// 		}
// 	}
//
// 	return &apiResp, nil
// }
