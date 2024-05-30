package todo

type (
	// User представляет данные о пользователе.
	User struct {
		ID          int         `json:"id" db:"id"`
		Username    string      `json:"username" binding:"required"`
		Email       string      `json:"email" binding:"required"`
		Password    string      `json:"password" binding:"required"`
		Nationality string      `json:"nationality" binding:"required"`
		Height      float64     `json:"height" binding:"required"`
		Weight      float64     `json:"weight" binding:"required"`
		BodyType    string      `json:"body_type" db:"body_type" binding:"required"`
		Allergies   []string    `json:"allergies" binding:"required"`
		Goal        string      `json:"goal" binding:"required"`
		TrackerData TrackerData `json:"tracker_data"`
	}
)
type (
	// TrackerData представляет данные трекера пользователя.
	TrackerData struct {
		TrackerData float64 `json:"tracker_data"`
	}
)

type (
	// DrinkRecord представляет запись о выпитом напитке пользователем.
	DrinkRecord struct {
		UserId            int     `json:"user_id" binding:"required"`
		Date              string  `json:"date" binding:"required"`
		Time              string  `json:"time" binding:"required"`
		Quantity          float64 `json:"quantity" binding:"required"`
		DrinkName         string  `json:"drink_name" binding:"required"`
		AlcoholPercentage float64 `json:"alcohol_percentage" binding:"required"`
	}
)

type (
	// UserStatistics представляет статистику пользователя.
	UserStatistics struct {
		UserID              int64   `db:"user_id"`
		LightDrinks         float64 `db:"light_drinks"`
		MediumDrinks        float64 `db:"medium_drinks"`
		HeavyDrinks         float64 `db:"heavy_drinks"`
		LastConsumptionDate string  `db:"last_consumption_date"`
	}
)

type (
	UserStatsRequest struct {
		UserID              int64  `json:"user_id" db:"user_id"`
		MoreDrinks          string `json:"more_drinks"`
		LastConsumptionDate string `json:"last_consumption_date" db:"last_consumption_date"`
	}
)
type (
	// Drink представляет данные о напитке.
	Drink struct {
		DrinkID     int64   `json:"drink_id"`
		DrinkName   string  `json:"drink_name"`
		DrinkType   string  `json:"drink_type"`
		MinStrength float64 `json:"min_strength"`
		MaxStrength float64 `json:"max_strength"`
	}
)
type (
	// Formula представляет данные о формуле.
	Formula struct {
		ID      int64  `json:"id"`
		Formula string `json:"formula"`
	}
)

type (
	UserData struct {
		ID          int      `json:"id"`
		Username    string   `json:"username"`
		Email       string   `json:"email"`
		Nationality string   `json:"nationality"`
		Allergy     string   `json:"allergy" db:"allergies"`
		Height      float64  `json:"height"`
		Weight      float64  `json:"weight"`
		BodyType    string   `json:"body_type" db:"body_type"`
		Allergies   []string `json:"allergies"`
		Goal        string   `json:"goal"`
		FormulaCode string   `json:"formula_code"`
	}
)

type (
	Response struct {
		User  UserResponse `json:"user"`
		Token string       `json:"token"`
	}
)

type (
	UserResponse struct {
		ID          int      `json:"id"`
		Username    string   `json:"username"`
		Email       string   `json:"email"`
		Nationality string   `json:"nationality"`
		Height      float64  `json:"height"`
		Weight      float64  `json:"weight"`
		BodyType    string   `json:"body_type"`
		Allergies   []string `json:"allergies"`
		Goal        string   `json:"goal"`
	}
)

type (
	Request struct {
		ID int `json:"id" binding:"required"`
	}
)

type (
	Norm struct {
		AlcoholPercentage float64 `json:"alcohol_percentage" binding:"required"`
	}
)
