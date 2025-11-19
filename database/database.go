package database

import (
	"fmt"
	"log"
	"os"
	
	"github.com/RomuloNd/Projeto-CruzeiroDoSul/config"
	"github.com/RomuloNd/Projeto-CruzeiroDoSul/models"
	// Importa o pacote 'seed' para ter acesso à função SeedDatabase
	"github.com/RomuloNd/Projeto-CruzeiroDoSul/database/seed" 
	
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB           *gorm.DB
	dbLoggerMode = logger.Silent
)

func SetupDatabase() {
	dbType := config.Get().DBType
	switch dbType {
	case "sqlite":
		setupSQLiteDatabase()
	case "postgres":
		setupPostgresDatabase()
	default:
		log.Panicf("Tipo de banco de dados não suportado: %s", dbType)
	}
}

func setupSQLiteDatabase() {
    // Verifica se o arquivo DB existe para decidir se logamos a criação.
    dbPath := config.Get().DBPath
    _, err := os.Stat(dbPath)
    isNewDB := os.IsNotExist(err)

    // Tenta abrir/criar o arquivo DB. O GORM o criará se não existir.
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	handleError("Falha ao conectar ao banco de dados SQLite", err)

	// 1. Migração das Tabelas
	handleError("Falha ao migrar os modelos", migrateModels(db))

    // 2. Seeding (População): Só é feito se a base for nova
    if isNewDB {
        log.Println("Base de dados criada. Executando seed inicial...")
	    handleError("Falha ao popular o banco de dados", seed.SeedDatabase(db))
        log.Println("Seed concluído com sucesso!")
    }

	db.Logger = logger.Default.LogMode(dbLoggerMode)
	DB = db
}

func setupPostgresDatabase() {
	cfg := config.Get()
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort, cfg.DBSSLMode, cfg.DBTimezone,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	handleError("Falha ao conectar ao banco de dados PostgreSQL", err)

	handleError("Falha ao migrar os modelos", migrateModels(db))
    // A lógica de seed na inicialização deve ser feita aqui também, se a base for nova.
	handleError("Falha ao popular o banco de dados", seed.SeedDatabase(db))

	db.Logger = logger.Default.LogMode(dbLoggerMode)
	DB = db
}

func handleError(msg string, err error) {
	if err != nil {
		log.Panicf("%s: %v", msg, err)
	}
}

func migrateModels(db *gorm.DB) error {
	modelsToMigrate := []interface{}{
		&models.User{},
		&models.Profile{},
		&models.Product{},
		&models.Stock{},
		&models.StoreConfig{},
		&models.Order{},
		&models.OrderDeliveryDetail{},
		&models.ShoppingCart{},
		&models.ShoppingCartItem{},
	}
	return db.AutoMigrate(modelsToMigrate...)
}

// **IMPORTANTE:** Você precisa garantir que esta função seja pública no seu arquivo 'database/seed.go'
// A função deve ser: func SeedDatabase(db *gorm.DB) error { ... }