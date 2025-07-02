# Face Problems Analyzer API

A Go-based REST API for cosmetics and skincare analysis, featuring face condition analysis and product management capabilities.

## Endpoints

### Product Management

#### GET `/product/:barcode`
Retrieves product information by barcode from either internal database or external OpenBeautyFacts API.

**Parameters:**
- `barcode` (path parameter): Product barcode

**Response:**
```json
{
  "id": "string",
  "barcode": "string",
  "productName": "string",
  "brand": "string",
  "productQuantity": 100.0,
  "productQuantityUnit": "ml",
  "ingredients": ["ingredient1", "ingredient2"],
  "source": "internal|open_beauty_facts"
}
```

**Status Codes:**
- `200`: Product found successfully
- `400`: Missing barcode in path
- `404`: Product not found in any source
- `500`: Internal server error

### Face Analysis

#### POST `/face/analyze`
Analyzes uploaded face image for skin conditions using machine learning model and stores results.

**Content-Type:** `multipart/form-data`

**Parameters:**
- `uid` (form field): User identifier
- `image` (form field): Face image file

**Response:**
```json
{
  "message": "Face image processed successfully"
}
```

**Status Codes:**
- `200`: Image processed successfully
- `400`: Invalid request data or missing parameters
- `500`: Internal server error during processing

#### GET `/face/health_info`
Retrieves the latest face condition analysis for a user.

**Content-Type:** `application/json`

**Request Body:**
```json
{
  "uid": "string"
}
```

**Response:**
```json
{
  "faceCondition": {
    "probabilities": {
      "blackheads": 0.2,
      "dry_skin": 0.1,
      "oily_skin": 0.3
    },
    "overallScore": 75.5,
    "overallCondition": "Good"
  }
}
```

**Status Codes:**
- `200`: Face condition retrieved successfully
- `400`: Invalid request data or user not found

### Skincare Routine Management

#### POST `/routines/add_product`
Adds a product to user's skincare routine (morning, evening, or both).

**Content-Type:** `application/json`

**Request Body:**
```json
{
  "uid": "string",
  "productID": "string",
  "routineType": "morning|evening|both"
}
```

**Response:**
```json
{
  "message": "Product added to routine successfully"
}
```

**Status Codes:**
- `200`: Product added successfully
- `400`: Invalid request data
- `500`: Internal server error

#### GET `/routines/get_all`
Retrieves all products in user's skincare routines with ingredient impact analysis.

**Content-Type:** `application/json`

**Request Body:**
```json
{
  "uid": "string"
}
```

**Response:**
```json
{
  "routines": {
    "morningRoutineProducts": [
      {
        "id": "string",
        "barcode": "string",
        "productName": "string",
        "brand": "string",
        "productQuantity": 100.0,
        "productQuantityUnit": "ml",
        "ingredients": ["ingredient1", "ingredient2"],
        "ingredientsWithImpacts": [
          {
            "ingredient_name": "string",
            "blackheads_impact": "positive|negative|neutral",
            "blackheads_confidence": 0.85,
            "dry_skin_impact": "positive|negative|neutral",
            "dry_skin_confidence": 0.92
          }
        ],
        "source": "internal"
      }
    ],
    "eveningRoutineProducts": []
  }
}
```

**Status Codes:**
- `200`: Routines retrieved successfully
- `400`: Invalid request data
- `500`: Internal server error

## Features

- **Face Analysis**: Uses Python ML model to analyze skin conditions from uploaded images
- **Product Database**: Integrates with OpenBeautyFacts API for comprehensive product information
- **Ingredient Impact**: Provides detailed analysis of how ingredients affect various skin conditions
- **Routine Management**: Allows users to build and manage morning/evening skincare routines
- **Multi-source Data**: Combines internal database with external APIs for comprehensive product coverage

## Technology Stack

- **Backend**: Go with Hertz web framework
- **Database**: PostgreSQL
- **ML Processing**: Python with FastAI
- **External APIs**: OpenBeautyFacts
- **File Handling**: Multipart form uploads for image processing

## Configuration

The application uses environment variables for configuration:
- `SERVER_PORT`: Server port (default: 8080)
- `SUPABASE_HOST`, `SUPABASE_PORT`, `SUPABASE_USER`, `SUPABASE_PASSWORD`, `SUPABASE_DB_NAME`: Database connection
- `OPEN_BEAUTY_FACTS_BASE_URL`: External API base URL

## Running the Application

```bash
# Load environment variables
cp .env.example .env

# Build and run
./run.sh
