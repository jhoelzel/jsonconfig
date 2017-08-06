package jsonconfig

// Configuration for all different Systems
type Configuration struct {
	Globals  GlobalSettings
	Emails   EmailSettings
	Login    LoginSettings
	Database DatabaseSettings
	Admin    AdminSettings
}

// GlobalSettings for general use
type GlobalSettings struct {
	URL              string `json:"url"`
	Theme            string `json:"theme"`
	DateFormat       string `json:"date_format"`
	TimeFormat       string `json:"time_format"`
	DecimalSeperator string `json:"decimal_seperator"`
	Language         string `json:"language"`
	Version          string `json:"version"`
}

// LoginSettings for authentication
type LoginSettings struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// AdminSettings for administrative contact
type AdminSettings struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// EmailSettings for smtp connections
type EmailSettings struct {
	Name      string `json:"name"`
	Host      string `json:"host"`
	Port      string `json:"port"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Mechanism string `json:"mechanism"`
}

// DatabaseSettings for db connection
type DatabaseSettings struct {
	Type     string `json:"type"`
	Username string `json:"username"`
	Password string `json:"password"`
	DBname   string `json:"dbname"`
	Host     string `json:"host"`
	Port     string `json:"port"`
}
