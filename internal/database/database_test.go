package database

import (
	"KFS_Backend/internal/modules/campaign"
	"KFS_Backend/internal/modules/investment"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() (*gorm.DB, error) {
	// **Test için hafızada çalışan bir SQLite veritabanı kullanıyoruz**
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// **Test için migration çalıştır**
	err = db.AutoMigrate(
		&campaign.Campaign{},
		&investment.Investment{},
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func TestDatabaseConnection(t *testing.T) {
	// 1️⃣ Test veritabanını oluştur
	testDB, err := setupTestDB()
	assert.NoError(t, err, "Veritabanı bağlantısı başarısız oldu!")

	// 2️⃣ Ping testi (Veritabanı gerçekten çalışıyor mu?)
	sqlDB, _ := testDB.DB()
	err = sqlDB.Ping()
	assert.NoError(t, err, "Veritabanı bağlantı testi başarısız oldu!")

	// 3️⃣ Migration kontrolü
	err = testDB.AutoMigrate(&campaign.Campaign{}, &investment.Investment{})
	assert.NoError(t, err, "Migration işlemi başarısız!")

	// 4️⃣ Test bitince bağlantıyı kapat
	sqlDB.Close()
}

func TestDatabaseFailure(t *testing.T) {
	// 1️⃣ Geçersiz bir DSN ile bağlantı açmayı deniyoruz
	invalidDB, err := gorm.Open(postgres.Open("host=invalid_host user=invalid_user dbname=invalid_db"), &gorm.Config{})

	// 2️⃣ Hata almalı çünkü yanlış bir veritabanı bağlantısı kurmaya çalışıyoruz
	assert.Error(t, err, "Geçersiz DSN ile bağlantı başarısız olmalı!")

	// 3️⃣ Eğer GORM bir `gorm.DB` nesnesi döndürdüyse, bu nesnenin gerçekten çalışmadığını test et
	if invalidDB != nil {
		sqlDB, dbErr := invalidDB.DB()

		// `sqlDB` bağlantısının olup olmadığını test edelim
		assert.NotNil(t, sqlDB, "sqlDB nesnesi nil olmamalı çünkü GORM nesne döndürüyor!")

		// `Ping()` çağrıldığında bağlantı gerçekten başarısız olmalı
		if sqlDB != nil {
			pingErr := sqlDB.Ping()
			assert.Error(t, pingErr, "Başarısız bağlantıda Ping başarısız olmalı!")
		}

		// Eğer bağlantı yanlışsa, `sqlDB.Close()` çağrılmamalı
		if dbErr == nil {
			sqlDB.Close()
		}
	}
}
