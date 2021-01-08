package config

const (
	// ServerPort -
	ServerPort string = "1324"
)

// Config -
type Config struct {
	Postgres   Postgres
	ServerPort int
}

// Postgres -
type Postgres struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}
