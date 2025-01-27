package internal

import (
	"fmt"
	"log"

	"KFS_Backend/configs"
	"KFS_Backend/pkg/logger"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Veritabanı bağlantısı
var DB *gorm.DB

// Sunucuyu başlat
func StartServer() {
	// Config yükle
	config, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("Config yüklenirken hata: %v", err)
	}

	// Logger başlat
	logger.InitLogger()
	logger.Info("⚡ API başlatılıyor...")

	// Fiber başlat
	app := fiber.New()

	// Router'ı yükle
	SetupRouter(app)

	// ✅ Supabase için SSL bağlantısını ayarla
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require",
		config.Database.Host, config.Database.User, config.Database.Password,
		config.Database.Name, config.Database.Port,
	)

	var dbErr error
	DB, dbErr = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if dbErr != nil {
		logger.Error(fmt.Sprintf("⚠️  Supabase veritabanına bağlanılamadı: %v", dbErr))
	} else {
		logger.Info("✅ Supabase veritabanına başarıyla bağlandı!")
	}

	// Sunucuyu çalıştır
	port := ":" + config.Server.Port
	logger.Info("🚀 Sunucu " + port + " portunda çalışıyor...")
	log.Fatal(app.Listen(port))
}
