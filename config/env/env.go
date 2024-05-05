package env

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
}

func DBName() string {
	return os.Getenv("DB_NAME")
}

func DB_Port() string {
	return os.Getenv("DB_PORT")
}

func Host_Name() string {
	return os.Getenv("HOST_NAME")
}

func UserName() string {
	return os.Getenv("USERNAME")
}

func Password() string {
	return os.Getenv("PASSWORD")
}

func MigrationDir() string {
	return os.Getenv("MIGRATION_DIR")
}

func GetJWTKey() string {
	return os.Getenv("JWT_KEY")
}
