package configs

import (
	"testing"

	"os"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	// 1️⃣ .env dosyasını yükleyelim
	err := godotenv.Load("../.env")
	assert.NoError(t, err, "⚠️ .env dosyası yüklenemedi!")

	// 2️⃣ Beklenen değerleri `.env` dosyasından okuyarak test et
	expectedEnv := map[string]string{
		"PORT":        os.Getenv("PORT"),
		"DB_HOST":     os.Getenv("DB_HOST"),
		"DB_PORT":     os.Getenv("DB_PORT"),
		"DB_USER":     os.Getenv("DB_USER"),
		"DB_PASSWORD": os.Getenv("DB_PASSWORD"),
		"DB_NAME":     os.Getenv("DB_NAME"),
		"DB_SSLMODE":  os.Getenv("DB_SSLMODE"),
	}

	// 3️⃣ Config yükleme fonksiyonunu çağır
	config, err := LoadConfig()
	assert.NoError(t, err, "LoadConfig() çağrılırken hata oluştu!")

	// 4️⃣ `.env` ile `config` içindeki değerleri kıyasla
	assert.Equal(t, expectedEnv["PORT"], config.Server.Port, "PORT değişkeni yanlış!")
	assert.Equal(t, expectedEnv["DB_HOST"], config.Database.Host, "DB_HOST değişkeni yanlış!")
	assert.Equal(t, expectedEnv["DB_PORT"], config.Database.Port, "DB_PORT değişkeni yanlış!")
	assert.Equal(t, expectedEnv["DB_USER"], config.Database.User, "DB_USER değişkeni yanlış!")
	assert.Equal(t, expectedEnv["DB_PASSWORD"], config.Database.Password, "DB_PASSWORD değişkeni yanlış!")
	assert.Equal(t, expectedEnv["DB_NAME"], config.Database.Name, "DB_NAME değişkeni yanlış!")
	assert.Equal(t, expectedEnv["DB_SSLMODE"], config.Database.SSLMode, "DB_SSLMODE değişkeni yanlış!")
}
