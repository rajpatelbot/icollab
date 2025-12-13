package utils

import (
	"fmt"

	"github.com/rajpatelbot/icollab/internal/config"
)

func GenerateDSN() string {
	env := config.EnvConfig

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		env.DB_HOST,
		env.DB_PORT,
		env.DB_USER,
		env.DB_PASSWORD,
		env.DB_NAME,
		env.DB_SSLMODE,
	)

	return dsn
}
