package apiclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type OpenBeautyFactsResponse struct {
	Code    string                 `json:"code"`
	Errors  []interface{}          `json:"errors"`
	Product OpenBeautyFactsProduct `json:"product"`
	Result  struct {
		ID     string `json:"id"`
		LcName string `json:"lc_name"`
		Name   string `json:"name"`
	} `json:"result"`
	Status   int           `json:"status"`
	Warnings []interface{} `json:"warnings"`
}

type OpenBeautyFactsProduct struct {
	ProductID                string        `json:"_id"`
	Keywords                 []string      `json:"_keywords"`
	AddedCountriesTags       []interface{} `json:"added_countries_tags"`
	AdditivesN               int           `json:"additives_n"`
	AdditivesOriginalTags    []string      `json:"additives_original_tags"`
	AdditivesTags            []string      `json:"additives_tags"`
	Allergens                string        `json:"allergens"`
	AllergensFromIngredients string        `json:"allergens_from_ingredients"`
	AllergensFromUser        string        `json:"allergens_from_user"`
	AllergensHierarchy       []interface{} `json:"allergens_hierarchy"`
	AllergensLc              string        `json:"allergens_lc"`
	AllergensTags            []interface{} `json:"allergens_tags"`
	AminoAcidsTags           []interface{} `json:"amino_acids_tags"`
	Brands                   string        `json:"brands"`
	BrandsOld                string        `json:"brands_old"`
	BrandsTags               []string      `json:"brands_tags"`
	Categories               string        `json:"categories"`
	CategoriesHierarchy      []string      `json:"categories_hierarchy"`
	CategoriesLc             string        `json:"categories_lc"`
	CategoriesProperties     struct {
	} `json:"categories_properties"`
	CategoriesPropertiesTags []string      `json:"categories_properties_tags"`
	CategoriesTags           []string      `json:"categories_tags"`
	CheckersTags             []interface{} `json:"checkers_tags"`
	CitiesTags               []interface{} `json:"cities_tags"`
	Code                     string        `json:"code"`
	CodesTags                []string      `json:"codes_tags"`
	Complete                 int           `json:"complete"`
	Completeness             float64       `json:"completeness"`
	CorrectorsTags           []string      `json:"correctors_tags"`
	Countries                string        `json:"countries"`
	CountriesHierarchy       []string      `json:"countries_hierarchy"`
	CountriesLc              string        `json:"countries_lc"`
	CountriesTags            []string      `json:"countries_tags"`
	CreatedT                 int           `json:"created_t"`
	Creator                  string        `json:"creator"`
	DataQualityBugsTags      []interface{} `json:"data_quality_bugs_tags"`
	DataQualityErrorsTags    []interface{} `json:"data_quality_errors_tags"`
	DataQualityInfoTags      []string      `json:"data_quality_info_tags"`
	DataQualityTags          []string      `json:"data_quality_tags"`
	DataQualityWarningsTags  []interface{} `json:"data_quality_warnings_tags"`
	DataSources              string        `json:"data_sources"`
	DataSourcesTags          []string      `json:"data_sources_tags"`
	DebugParamSortedLangs    []string      `json:"debug_param_sorted_langs"`
	EcoscoreData             struct {
		Adjustments struct {
			OriginsOfIngredients struct {
				AggregatedOrigins []struct {
					Origin  string `json:"origin"`
					Percent int    `json:"percent"`
				} `json:"aggregated_origins"`
				EpiScore                int      `json:"epi_score"`
				EpiValue                int      `json:"epi_value"`
				OriginsFromCategories   []string `json:"origins_from_categories"`
				OriginsFromOriginsField []string `json:"origins_from_origins_field"`
				TransportationScores    struct {
					Ad    int `json:"ad"`
					Al    int `json:"al"`
					At    int `json:"at"`
					Ax    int `json:"ax"`
					Ba    int `json:"ba"`
					Be    int `json:"be"`
					Bg    int `json:"bg"`
					Ch    int `json:"ch"`
					Cy    int `json:"cy"`
					Cz    int `json:"cz"`
					De    int `json:"de"`
					Dk    int `json:"dk"`
					Dz    int `json:"dz"`
					Ee    int `json:"ee"`
					Eg    int `json:"eg"`
					Es    int `json:"es"`
					Fi    int `json:"fi"`
					Fo    int `json:"fo"`
					Fr    int `json:"fr"`
					Gg    int `json:"gg"`
					Gi    int `json:"gi"`
					Gr    int `json:"gr"`
					Hr    int `json:"hr"`
					Hu    int `json:"hu"`
					Ie    int `json:"ie"`
					Il    int `json:"il"`
					Im    int `json:"im"`
					Is    int `json:"is"`
					It    int `json:"it"`
					Je    int `json:"je"`
					Lb    int `json:"lb"`
					Li    int `json:"li"`
					Lt    int `json:"lt"`
					Lu    int `json:"lu"`
					Lv    int `json:"lv"`
					Ly    int `json:"ly"`
					Ma    int `json:"ma"`
					Mc    int `json:"mc"`
					Md    int `json:"md"`
					Me    int `json:"me"`
					Mk    int `json:"mk"`
					Mt    int `json:"mt"`
					Nl    int `json:"nl"`
					No    int `json:"no"`
					Pl    int `json:"pl"`
					Ps    int `json:"ps"`
					Pt    int `json:"pt"`
					Ro    int `json:"ro"`
					Rs    int `json:"rs"`
					Se    int `json:"se"`
					Si    int `json:"si"`
					Sj    int `json:"sj"`
					Sk    int `json:"sk"`
					Sm    int `json:"sm"`
					Sy    int `json:"sy"`
					Tn    int `json:"tn"`
					Tr    int `json:"tr"`
					Ua    int `json:"ua"`
					Uk    int `json:"uk"`
					Us    int `json:"us"`
					Va    int `json:"va"`
					World int `json:"world"`
					Xk    int `json:"xk"`
				} `json:"transportation_scores"`
				TransportationValues struct {
					Ad    int `json:"ad"`
					Al    int `json:"al"`
					At    int `json:"at"`
					Ax    int `json:"ax"`
					Ba    int `json:"ba"`
					Be    int `json:"be"`
					Bg    int `json:"bg"`
					Ch    int `json:"ch"`
					Cy    int `json:"cy"`
					Cz    int `json:"cz"`
					De    int `json:"de"`
					Dk    int `json:"dk"`
					Dz    int `json:"dz"`
					Ee    int `json:"ee"`
					Eg    int `json:"eg"`
					Es    int `json:"es"`
					Fi    int `json:"fi"`
					Fo    int `json:"fo"`
					Fr    int `json:"fr"`
					Gg    int `json:"gg"`
					Gi    int `json:"gi"`
					Gr    int `json:"gr"`
					Hr    int `json:"hr"`
					Hu    int `json:"hu"`
					Ie    int `json:"ie"`
					Il    int `json:"il"`
					Im    int `json:"im"`
					Is    int `json:"is"`
					It    int `json:"it"`
					Je    int `json:"je"`
					Lb    int `json:"lb"`
					Li    int `json:"li"`
					Lt    int `json:"lt"`
					Lu    int `json:"lu"`
					Lv    int `json:"lv"`
					Ly    int `json:"ly"`
					Ma    int `json:"ma"`
					Mc    int `json:"mc"`
					Md    int `json:"md"`
					Me    int `json:"me"`
					Mk    int `json:"mk"`
					Mt    int `json:"mt"`
					Nl    int `json:"nl"`
					No    int `json:"no"`
					Pl    int `json:"pl"`
					Ps    int `json:"ps"`
					Pt    int `json:"pt"`
					Ro    int `json:"ro"`
					Rs    int `json:"rs"`
					Se    int `json:"se"`
					Si    int `json:"si"`
					Sj    int `json:"sj"`
					Sk    int `json:"sk"`
					Sm    int `json:"sm"`
					Sy    int `json:"sy"`
					Tn    int `json:"tn"`
					Tr    int `json:"tr"`
					Ua    int `json:"ua"`
					Uk    int `json:"uk"`
					Us    int `json:"us"`
					Va    int `json:"va"`
					World int `json:"world"`
					Xk    int `json:"xk"`
				} `json:"transportation_values"`
				Values struct {
					Ad    int `json:"ad"`
					Al    int `json:"al"`
					At    int `json:"at"`
					Ax    int `json:"ax"`
					Ba    int `json:"ba"`
					Be    int `json:"be"`
					Bg    int `json:"bg"`
					Ch    int `json:"ch"`
					Cy    int `json:"cy"`
					Cz    int `json:"cz"`
					De    int `json:"de"`
					Dk    int `json:"dk"`
					Dz    int `json:"dz"`
					Ee    int `json:"ee"`
					Eg    int `json:"eg"`
					Es    int `json:"es"`
					Fi    int `json:"fi"`
					Fo    int `json:"fo"`
					Fr    int `json:"fr"`
					Gg    int `json:"gg"`
					Gi    int `json:"gi"`
					Gr    int `json:"gr"`
					Hr    int `json:"hr"`
					Hu    int `json:"hu"`
					Ie    int `json:"ie"`
					Il    int `json:"il"`
					Im    int `json:"im"`
					Is    int `json:"is"`
					It    int `json:"it"`
					Je    int `json:"je"`
					Lb    int `json:"lb"`
					Li    int `json:"li"`
					Lt    int `json:"lt"`
					Lu    int `json:"lu"`
					Lv    int `json:"lv"`
					Ly    int `json:"ly"`
					Ma    int `json:"ma"`
					Mc    int `json:"mc"`
					Md    int `json:"md"`
					Me    int `json:"me"`
					Mk    int `json:"mk"`
					Mt    int `json:"mt"`
					Nl    int `json:"nl"`
					No    int `json:"no"`
					Pl    int `json:"pl"`
					Ps    int `json:"ps"`
					Pt    int `json:"pt"`
					Ro    int `json:"ro"`
					Rs    int `json:"rs"`
					Se    int `json:"se"`
					Si    int `json:"si"`
					Sj    int `json:"sj"`
					Sk    int `json:"sk"`
					Sm    int `json:"sm"`
					Sy    int `json:"sy"`
					Tn    int `json:"tn"`
					Tr    int `json:"tr"`
					Ua    int `json:"ua"`
					Uk    int `json:"uk"`
					Us    int `json:"us"`
					Va    int `json:"va"`
					World int `json:"world"`
					Xk    int `json:"xk"`
				} `json:"values"`
				Warning string `json:"warning"`
			} `json:"origins_of_ingredients"`
			Packaging struct {
				Value   int    `json:"value"`
				Warning string `json:"warning"`
			} `json:"packaging"`
			ProductionSystem struct {
				Labels  []interface{} `json:"labels"`
				Value   int           `json:"value"`
				Warning string        `json:"warning"`
			} `json:"production_system"`
			ThreatenedSpecies struct {
			} `json:"threatened_species"`
		} `json:"adjustments"`
		Agribalyse struct {
			Warning string `json:"warning"`
		} `json:"agribalyse"`
		Grade   string `json:"grade"`
		Missing struct {
			AgbCategory int `json:"agb_category"`
			Labels      int `json:"labels"`
			Origins     int `json:"origins"`
			Packagings  int `json:"packagings"`
		} `json:"missing"`
		MissingAgribalyseMatchWarning int    `json:"missing_agribalyse_match_warning"`
		MissingKeyData                int    `json:"missing_key_data"`
		Status                        string `json:"status"`
	} `json:"ecoscore_data"`
	EcoscoreGrade            string        `json:"ecoscore_grade"`
	EcoscoreTags             []string      `json:"ecoscore_tags"`
	EditorsTags              []string      `json:"editors_tags"`
	EmbCodes                 string        `json:"emb_codes"`
	EmbCodesTags             []interface{} `json:"emb_codes_tags"`
	EntryDatesTags           []string      `json:"entry_dates_tags"`
	ExpirationDate           string        `json:"expiration_date"`
	FoodGroupsTags           []interface{} `json:"food_groups_tags"`
	GenericName              string        `json:"generic_name"`
	GenericNameEn            string        `json:"generic_name_en"`
	ID                       string        `json:"id"`
	ImageFrontSmallURL       string        `json:"image_front_small_url"`
	ImageFrontThumbURL       string        `json:"image_front_thumb_url"`
	ImageFrontURL            string        `json:"image_front_url"`
	ImageIngredientsSmallURL string        `json:"image_ingredients_small_url"`
	ImageIngredientsThumbURL string        `json:"image_ingredients_thumb_url"`
	ImageIngredientsURL      string        `json:"image_ingredients_url"`
	ImageNutritionSmallURL   string        `json:"image_nutrition_small_url"`
	ImageNutritionThumbURL   string        `json:"image_nutrition_thumb_url"`
	ImageNutritionURL        string        `json:"image_nutrition_url"`
	ImageSmallURL            string        `json:"image_small_url"`
	ImageThumbURL            string        `json:"image_thumb_url"`
	ImageURL                 string        `json:"image_url"`
	Images                   struct {
		Num1 struct {
			Sizes struct {
				Num100 struct {
					H int `json:"h"`
					W int `json:"w"`
				} `json:"100"`
				Num400 struct {
					H int `json:"h"`
					W int `json:"w"`
				} `json:"400"`
				Full struct {
					H int `json:"h"`
					W int `json:"w"`
				} `json:"full"`
			} `json:"sizes"`
			UploadedT int    `json:"uploaded_t"`
			Uploader  string `json:"uploader"`
		} `json:"1"`
		Num2 struct {
			Sizes struct {
				Num100 struct {
					H int `json:"h"`
					W int `json:"w"`
				} `json:"100"`
				Num400 struct {
					H int `json:"h"`
					W int `json:"w"`
				} `json:"400"`
				Full struct {
					H int `json:"h"`
					W int `json:"w"`
				} `json:"full"`
			} `json:"sizes"`
			UploadedT int    `json:"uploaded_t"`
			Uploader  string `json:"uploader"`
		} `json:"2"`
		Num3 struct {
			Sizes struct {
				Num100 struct {
					H int `json:"h"`
					W int `json:"w"`
				} `json:"100"`
				Num400 struct {
					H int `json:"h"`
					W int `json:"w"`
				} `json:"400"`
				Full struct {
					H int `json:"h"`
					W int `json:"w"`
				} `json:"full"`
			} `json:"sizes"`
			UploadedT int    `json:"uploaded_t"`
			Uploader  string `json:"uploader"`
		} `json:"3"`
		FrontEn struct {
			Angle                int         `json:"angle"`
			CoordinatesImageSize string      `json:"coordinates_image_size"`
			Geometry             string      `json:"geometry"`
			Imgid                string      `json:"imgid"`
			Normalize            interface{} `json:"normalize"`
			Rev                  string      `json:"rev"`
			Sizes                struct {
				Num100 struct {
					H int `json:"h"`
					W int `json:"w"`
				} `json:"100"`
				Num200 struct {
					H int `json:"h"`
					W int `json:"w"`
				} `json:"200"`
				Num400 struct {
					H int `json:"h"`
					W int `json:"w"`
				} `json:"400"`
				Full struct {
					H int `json:"h"`
					W int `json:"w"`
				} `json:"full"`
			} `json:"sizes"`
			WhiteMagic interface{} `json:"white_magic"`
			X1         string      `json:"x1"`
			X2         string      `json:"x2"`
			Y1         string      `json:"y1"`
			Y2         string      `json:"y2"`
		} `json:"front_en"`
		IngredientsEn struct {
			Angle                int         `json:"angle"`
			CoordinatesImageSize string      `json:"coordinates_image_size"`
			Geometry             string      `json:"geometry"`
			Imgid                string      `json:"imgid"`
			Normalize            interface{} `json:"normalize"`
			Rev                  string      `json:"rev"`
			Sizes                struct {
				Num100 struct {
					H int `json:"h"`
					W int `json:"w"`
				} `json:"100"`
				Num200 struct {
					H int `json:"h"`
					W int `json:"w"`
				} `json:"200"`
				Num400 struct {
					H int `json:"h"`
					W int `json:"w"`
				} `json:"400"`
				Full struct {
					H int `json:"h"`
					W int `json:"w"`
				} `json:"full"`
			} `json:"sizes"`
			WhiteMagic interface{} `json:"white_magic"`
			X1         string      `json:"x1"`
			X2         string      `json:"x2"`
			Y1         string      `json:"y1"`
			Y2         string      `json:"y2"`
		} `json:"ingredients_en"`
		NutritionEn struct {
			Angle                int         `json:"angle"`
			CoordinatesImageSize string      `json:"coordinates_image_size"`
			Geometry             string      `json:"geometry"`
			Imgid                string      `json:"imgid"`
			Normalize            interface{} `json:"normalize"`
			Rev                  string      `json:"rev"`
			Sizes                struct {
				Num100 struct {
					H int `json:"h"`
					W int `json:"w"`
				} `json:"100"`
				Num200 struct {
					H int `json:"h"`
					W int `json:"w"`
				} `json:"200"`
				Num400 struct {
					H int `json:"h"`
					W int `json:"w"`
				} `json:"400"`
				Full struct {
					H int `json:"h"`
					W int `json:"w"`
				} `json:"full"`
			} `json:"sizes"`
			WhiteMagic interface{} `json:"white_magic"`
			X1         string      `json:"x1"`
			X2         string      `json:"x2"`
			Y1         string      `json:"y1"`
			Y2         string      `json:"y2"`
		} `json:"nutrition_en"`
	} `json:"images"`
	InformersTags []string `json:"informers_tags"`
	Ingredients   []struct {
		ID              string  `json:"id"`
		IsInTaxonomy    int     `json:"is_in_taxonomy"`
		PercentEstimate float64 `json:"percent_estimate"`
		PercentMax      float64 `json:"percent_max"`
		PercentMin      float64 `json:"percent_min"`
		Text            string  `json:"text"`
	} `json:"ingredients"`
	IngredientsAnalysis struct {
		EnPalmOilContentUnknown   []string `json:"en:palm-oil-content-unknown"`
		EnVeganStatusUnknown      []string `json:"en:vegan-status-unknown"`
		EnVegetarianStatusUnknown []string `json:"en:vegetarian-status-unknown"`
	} `json:"ingredients_analysis"`
	IngredientsAnalysisTags              []string      `json:"ingredients_analysis_tags"`
	IngredientsHierarchy                 []string      `json:"ingredients_hierarchy"`
	IngredientsLc                        string        `json:"ingredients_lc"`
	IngredientsN                         int           `json:"ingredients_n"`
	IngredientsNTags                     []string      `json:"ingredients_n_tags"`
	IngredientsNonNutritiveSweetenersN   int           `json:"ingredients_non_nutritive_sweeteners_n"`
	IngredientsOriginalTags              []string      `json:"ingredients_original_tags"`
	IngredientsPercentAnalysis           int           `json:"ingredients_percent_analysis"`
	IngredientsSweetenersN               int           `json:"ingredients_sweeteners_n"`
	IngredientsTags                      []string      `json:"ingredients_tags"`
	IngredientsText                      string        `json:"ingredients_text"`
	IngredientsTextEn                    string        `json:"ingredients_text_en"`
	IngredientsTextWithAllergens         string        `json:"ingredients_text_with_allergens"`
	IngredientsTextWithAllergensEn       string        `json:"ingredients_text_with_allergens_en"`
	IngredientsWithSpecifiedPercentN     int           `json:"ingredients_with_specified_percent_n"`
	IngredientsWithSpecifiedPercentSum   int           `json:"ingredients_with_specified_percent_sum"`
	IngredientsWithUnspecifiedPercentN   int           `json:"ingredients_with_unspecified_percent_n"`
	IngredientsWithUnspecifiedPercentSum int           `json:"ingredients_with_unspecified_percent_sum"`
	IngredientsWithoutCiqualCodes        []string      `json:"ingredients_without_ciqual_codes"`
	IngredientsWithoutCiqualCodesN       int           `json:"ingredients_without_ciqual_codes_n"`
	IngredientsWithoutEcobalyseIds       []string      `json:"ingredients_without_ecobalyse_ids"`
	IngredientsWithoutEcobalyseIdsN      int           `json:"ingredients_without_ecobalyse_ids_n"`
	InterfaceVersionCreated              string        `json:"interface_version_created"`
	InterfaceVersionModified             string        `json:"interface_version_modified"`
	KnownIngredientsN                    int           `json:"known_ingredients_n"`
	Labels                               string        `json:"labels"`
	LabelsHierarchy                      []interface{} `json:"labels_hierarchy"`
	LabelsLc                             string        `json:"labels_lc"`
	LabelsTags                           []interface{} `json:"labels_tags"`
	Lang                                 string        `json:"lang"`
	Languages                            struct {
		EnEnglish int `json:"en:english"`
	} `json:"languages"`
	LanguagesCodes struct {
		En int `json:"en"`
	} `json:"languages_codes"`
	LanguagesHierarchy      []string      `json:"languages_hierarchy"`
	LanguagesTags           []string      `json:"languages_tags"`
	LastEditDatesTags       []string      `json:"last_edit_dates_tags"`
	LastEditor              string        `json:"last_editor"`
	LastImageDatesTags      []string      `json:"last_image_dates_tags"`
	LastImageT              int           `json:"last_image_t"`
	LastModifiedBy          string        `json:"last_modified_by"`
	LastModifiedT           int           `json:"last_modified_t"`
	LastUpdatedT            int           `json:"last_updated_t"`
	Lc                      string        `json:"lc"`
	Link                    string        `json:"link"`
	MainCountriesTags       []interface{} `json:"main_countries_tags"`
	ManufacturingPlaces     string        `json:"manufacturing_places"`
	ManufacturingPlacesTags []interface{} `json:"manufacturing_places_tags"`
	MaxImgid                string        `json:"max_imgid"`
	MineralsTags            []string      `json:"minerals_tags"`
	MiscTags                []string      `json:"misc_tags"`
	NoNutritionData         string        `json:"no_nutrition_data"`
	NovaGroup               int           `json:"nova_group"`
	NovaGroupDebug          string        `json:"nova_group_debug"`
	NovaGroups              string        `json:"nova_groups"`
	NovaGroupsMarkers       struct {
		Num4 [][]string `json:"4"`
	} `json:"nova_groups_markers"`
	NovaGroupsTags  []string      `json:"nova_groups_tags"`
	NucleotidesTags []interface{} `json:"nucleotides_tags"`
	NutrientLevels  struct {
	} `json:"nutrient_levels"`
	NutrientLevelsTags []interface{} `json:"nutrient_levels_tags"`
	Nutriments         struct {
		FruitsVegetablesLegumesEstimateFromIngredients100G    int `json:"fruits-vegetables-legumes-estimate-from-ingredients_100g"`
		FruitsVegetablesLegumesEstimateFromIngredientsServing int `json:"fruits-vegetables-legumes-estimate-from-ingredients_serving"`
		FruitsVegetablesNutsEstimateFromIngredients100G       int `json:"fruits-vegetables-nuts-estimate-from-ingredients_100g"`
		FruitsVegetablesNutsEstimateFromIngredientsServing    int `json:"fruits-vegetables-nuts-estimate-from-ingredients_serving"`
		NovaGroup                                             int `json:"nova-group"`
		NovaGroup100G                                         int `json:"nova-group_100g"`
		NovaGroupServing                                      int `json:"nova-group_serving"`
	} `json:"nutriments"`
	Nutriscore struct {
		Num2021 struct {
			CategoryAvailable int `json:"category_available"`
			Data              struct {
				Energy                                   interface{} `json:"energy"`
				Fiber                                    int         `json:"fiber"`
				FruitsVegetablesNutsColzaWalnutOliveOils int         `json:"fruits_vegetables_nuts_colza_walnut_olive_oils"`
				IsBeverage                               int         `json:"is_beverage"`
				IsCheese                                 int         `json:"is_cheese"`
				IsFat                                    int         `json:"is_fat"`
				IsWater                                  int         `json:"is_water"`
				Proteins                                 interface{} `json:"proteins"`
				SaturatedFat                             interface{} `json:"saturated_fat"`
				Sodium                                   interface{} `json:"sodium"`
				Sugars                                   interface{} `json:"sugars"`
			} `json:"data"`
			Grade                string `json:"grade"`
			NutrientsAvailable   int    `json:"nutrients_available"`
			NutriscoreApplicable int    `json:"nutriscore_applicable"`
			NutriscoreComputed   int    `json:"nutriscore_computed"`
		} `json:"2021"`
		Num2023 struct {
			CategoryAvailable int `json:"category_available"`
			Data              struct {
				Energy                  interface{} `json:"energy"`
				Fiber                   interface{} `json:"fiber"`
				FruitsVegetablesLegumes int         `json:"fruits_vegetables_legumes"`
				IsBeverage              int         `json:"is_beverage"`
				IsCheese                int         `json:"is_cheese"`
				IsFatOilNutsSeeds       int         `json:"is_fat_oil_nuts_seeds"`
				IsRedMeatProduct        int         `json:"is_red_meat_product"`
				IsWater                 int         `json:"is_water"`
				Proteins                interface{} `json:"proteins"`
				Salt                    interface{} `json:"salt"`
				SaturatedFat            interface{} `json:"saturated_fat"`
				Sugars                  interface{} `json:"sugars"`
			} `json:"data"`
			Grade                string `json:"grade"`
			NutrientsAvailable   int    `json:"nutrients_available"`
			NutriscoreApplicable int    `json:"nutriscore_applicable"`
			NutriscoreComputed   int    `json:"nutriscore_computed"`
		} `json:"2023"`
	} `json:"nutriscore"`
	Nutriscore2021Tags                                                       []string      `json:"nutriscore_2021_tags"`
	Nutriscore2023Tags                                                       []string      `json:"nutriscore_2023_tags"`
	NutriscoreGrade                                                          string        `json:"nutriscore_grade"`
	NutriscoreTags                                                           []string      `json:"nutriscore_tags"`
	NutriscoreVersion                                                        string        `json:"nutriscore_version"`
	NutritionData                                                            string        `json:"nutrition_data"`
	NutritionDataPer                                                         string        `json:"nutrition_data_per"`
	NutritionDataPrepared                                                    string        `json:"nutrition_data_prepared"`
	NutritionDataPreparedPer                                                 string        `json:"nutrition_data_prepared_per"`
	NutritionGradeFr                                                         string        `json:"nutrition_grade_fr"`
	NutritionGrades                                                          string        `json:"nutrition_grades"`
	NutritionGradesTags                                                      []string      `json:"nutrition_grades_tags"`
	NutritionScoreBeverage                                                   int           `json:"nutrition_score_beverage"`
	NutritionScoreDebug                                                      string        `json:"nutrition_score_debug"`
	NutritionScoreWarningFruitsVegetablesLegumesEstimateFromIngredients      int           `json:"nutrition_score_warning_fruits_vegetables_legumes_estimate_from_ingredients"`
	NutritionScoreWarningFruitsVegetablesLegumesEstimateFromIngredientsValue int           `json:"nutrition_score_warning_fruits_vegetables_legumes_estimate_from_ingredients_value"`
	NutritionScoreWarningFruitsVegetablesNutsEstimateFromIngredients         int           `json:"nutrition_score_warning_fruits_vegetables_nuts_estimate_from_ingredients"`
	NutritionScoreWarningFruitsVegetablesNutsEstimateFromIngredientsValue    int           `json:"nutrition_score_warning_fruits_vegetables_nuts_estimate_from_ingredients_value"`
	NutritionScoreWarningNoFiber                                             int           `json:"nutrition_score_warning_no_fiber"`
	Origin                                                                   string        `json:"origin"`
	OriginEn                                                                 string        `json:"origin_en"`
	Origins                                                                  string        `json:"origins"`
	OriginsHierarchy                                                         []interface{} `json:"origins_hierarchy"`
	OriginsLc                                                                string        `json:"origins_lc"`
	OriginsTags                                                              []interface{} `json:"origins_tags"`
	OtherNutritionalSubstancesTags                                           []interface{} `json:"other_nutritional_substances_tags"`
	PackagingMaterialsTags                                                   []interface{} `json:"packaging_materials_tags"`
	PackagingRecyclingTags                                                   []interface{} `json:"packaging_recycling_tags"`
	PackagingShapesTags                                                      []interface{} `json:"packaging_shapes_tags"`
	PackagingText                                                            string        `json:"packaging_text"`
	PackagingTextEn                                                          string        `json:"packaging_text_en"`
	Packagings                                                               []interface{} `json:"packagings"`
	PackagingsComplete                                                       int           `json:"packagings_complete"`
	PackagingsMaterials                                                      struct {
	} `json:"packagings_materials"`
	PhotographersTags    []string      `json:"photographers_tags"`
	PnnsGroups1          string        `json:"pnns_groups_1"`
	PnnsGroups1Tags      []string      `json:"pnns_groups_1_tags"`
	PnnsGroups2          string        `json:"pnns_groups_2"`
	PnnsGroups2Tags      []string      `json:"pnns_groups_2_tags"`
	PopularityKey        int           `json:"popularity_key"`
	PopularityTags       []string      `json:"popularity_tags"`
	ProductName          string        `json:"product_name"`
	ProductNameEn        string        `json:"product_name_en"`
	ProductQuantity      float64       `json:"product_quantity"`
	ProductQuantityUnit  string        `json:"product_quantity_unit"`
	ProductType          string        `json:"product_type"`
	PurchasePlaces       string        `json:"purchase_places"`
	PurchasePlacesTags   []interface{} `json:"purchase_places_tags"`
	Quantity             string        `json:"quantity"`
	RemovedCountriesTags []interface{} `json:"removed_countries_tags"`
	Rev                  int           `json:"rev"`
	ScansN               int           `json:"scans_n"`
	SchemaVersion        int           `json:"schema_version"`
	SelectedImages       struct {
		Front struct {
			Display struct {
				En string `json:"en"`
			} `json:"display"`
			Small struct {
				En string `json:"en"`
			} `json:"small"`
			Thumb struct {
				En string `json:"en"`
			} `json:"thumb"`
		} `json:"front"`
		Ingredients struct {
			Display struct {
				En string `json:"en"`
			} `json:"display"`
			Small struct {
				En string `json:"en"`
			} `json:"small"`
			Thumb struct {
				En string `json:"en"`
			} `json:"thumb"`
		} `json:"ingredients"`
		Nutrition struct {
			Display struct {
				En string `json:"en"`
			} `json:"display"`
			Small struct {
				En string `json:"en"`
			} `json:"small"`
			Thumb struct {
				En string `json:"en"`
			} `json:"thumb"`
		} `json:"nutrition"`
	} `json:"selected_images"`
	States                string        `json:"states"`
	StatesHierarchy       []string      `json:"states_hierarchy"`
	StatesTags            []string      `json:"states_tags"`
	Stores                string        `json:"stores"`
	StoresTags            []interface{} `json:"stores_tags"`
	Teams                 string        `json:"teams"`
	TeamsTags             []string      `json:"teams_tags"`
	Traces                string        `json:"traces"`
	TracesFromIngredients string        `json:"traces_from_ingredients"`
	TracesFromUser        string        `json:"traces_from_user"`
	TracesHierarchy       []interface{} `json:"traces_hierarchy"`
	TracesLc              string        `json:"traces_lc"`
	TracesTags            []interface{} `json:"traces_tags"`
	UniqueScansN          int           `json:"unique_scans_n"`
	UnknownIngredientsN   int           `json:"unknown_ingredients_n"`
	UnknownNutrientsTags  []interface{} `json:"unknown_nutrients_tags"`
	UpdateKey             string        `json:"update_key"`
	VitaminsTags          []interface{} `json:"vitamins_tags"`
	WeighersTags          []interface{} `json:"weighers_tags"`
}

type OpenBeautyFactsClient interface {
	GetProductByBarcode(ctx context.Context, barcode string) (*OpenBeautyFactsResponse, error)
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type openBeautyFactsHTTPClient struct {
	baseURL    string
	httpClient HTTPClient
}

var ErrOpenBeautyFactsNotFound = fmt.Errorf("open beauty facts product not found")

func NewOpenBeautyFactsClient(baseURL string, client HTTPClient) OpenBeautyFactsClient {
	if client == nil {
		client = &http.Client{Timeout: 10 * time.Second}
	}
	return &openBeautyFactsHTTPClient{
		baseURL:    baseURL,
		httpClient: client,
	}
}

func (c *openBeautyFactsHTTPClient) GetProductByBarcode(ctx context.Context, barcode string) (*OpenBeautyFactsResponse, error) {
	url := fmt.Sprintf("%s/api/v2/product/%s.json", c.baseURL, barcode)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("apiclient: failed to create HTTP request for barcode %s: %w", barcode, err)
	}

	req.Header.Set("User-Agent", "PersonalCosmeticsApp/1.0 (olxzulfar@gmail.com)")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("apiclient: failed to make HTTP request to OpenBeautyFacts for barcode %s: %w", barcode, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			return nil, fmt.Errorf("apiclient: product not found in OpenBeautyFacts for barcode %s: %w", barcode, ErrOpenBeautyFactsNotFound)
		}
		return nil, fmt.Errorf("apiclient: OpenBeautyFacts API returned non-OK status %d for barcode %s", resp.StatusCode, barcode)
	}

	var apiResp OpenBeautyFactsResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, fmt.Errorf("apiclient: failed to decode OpenBeautyFacts response for barcode %s: %w", barcode, err)
	}

	if apiResp.Status != 1 {
		return nil, fmt.Errorf("apiclient: OpenBeautyFacts API product status %d (not found) for barcode %s: %w", apiResp.Status, barcode, ErrOpenBeautyFactsNotFound)
	}

	return &apiResp, nil
}
