package initializers

import (
	"os"
)

func LoadEnvVariables() {
	os.Setenv("DB_URL", "host=silly.db.elephantsql.com user=qfjrxaof password=9rF7AHfJSA9Fi--wK88u7ibVc9WthoJY dbname=qfjrxaof   port=5432 sslmode=disable")
	os.Setenv("PORT", "8080")
	os.Setenv("GIN_MODE", "release")
	os.Setenv("JWT_SECRET", "secret")
}
