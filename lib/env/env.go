package env

import "os"

// Env has environment stored
type Env struct {
	ServerPort      string
	Environment     string
	LogOutput       string
	DBUsername      string
	DBPassword      string
	DBHost          string
	DBPort          string
	DBName          string
	JWTSecret       string
	DBEUsername     string
	DBEPassword     string
	DBEHost         string
	Endpoint        string
	AcessKeyID      string
	SecretAccessKey string
	UseSsl          string
}

// NewEnv creates a new environment
func NewEnv() Env {
	env := Env{}
	env.LoadEnv()
	return env
}

// LoadEnv loads environment
func (env *Env) LoadEnv() {
	env.ServerPort = os.Getenv("ServerPort")
	env.Environment = os.Getenv("Environment")
	env.LogOutput = os.Getenv("LogOutput")

	env.DBUsername = os.Getenv("DBUsername")
	env.DBPassword = os.Getenv("DBPassword")
	env.DBHost = os.Getenv("DBHost")
	env.DBPort = os.Getenv("DBPort")
	env.DBName = os.Getenv("DBName")

	env.JWTSecret = os.Getenv("JWTSecret")

	env.DBEUsername = os.Getenv("DBEUsername")
	env.DBEPassword = os.Getenv("DBEPassword")
	env.DBEHost = os.Getenv("DBEHost")

	env.Endpoint = os.Getenv("ENDPOINT")
	env.AcessKeyID = os.Getenv("ACCESS_KEY_ID")
	env.SecretAccessKey = os.Getenv("SECRET_ACCESS_KEY")
	env.UseSsl = os.Getenv("USE_SSL")
}
