package db

import (
	"fmt"
	"log"
	"os"
	"time"
	"vaqua/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// Prefer DATABASE_URL if set (Render/Heroku style). Otherwise build from DB_* parts.
func dsnFromEnv() string {
	if url := os.Getenv("DATABASE_URL"); url != "" {
		// e.g. postgres://user:pass@host:5432/dbname?sslmode=require
		return url
	}
	host := getenv("DB_HOST", "localhost")
	user := getenv("DB_USER", "postgres")
	pass := getenv("DB_PASSWORD", "pass")
	name := getenv("DB_NAME", "vaqua")
	port := getenv("DB_PORT", "5432")
	ssl := getenv("DB_SSLMODE", "disable") // use "require" on cloud DBs

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		user, pass, host, port, name, ssl,
	)
}

func getenv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}

func InitDb() *gorm.DB {
	dsn := dsnFromEnv()

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("sql DB handle error: %v", err)
	}

	// Connection pool settings
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	// Auto-migrate schemas
	if err := DB.AutoMigrate(
		&models.User{},
		&models.Account{},
		&models.Transaction{},
		&models.Transfer{},
		&models.IncomeAndExpenses{},
		&models.BlacklistedToken{},
	); err != nil {
		log.Fatalf("failed to migrate schema: %v", err)
	}

	return DB
}

func Ping() error {
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Ping()
}

/*package db

import (
	"fmt"
	"log"
	"os"
	"time"
	"vaqua/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB
//changing this because of render

// Prefer DATABASE_URL if set (Render/Heroku style). 
func dsnFromEnv() string {
	if url := os.Getenv("DATABASE_URL"); url != "" {
		
		return url
	}
	host := getenv("DB_HOST", "localhost")
	user := getenv("DB_USER", "postgres")
	pass := getenv("DB_PASSWORD", "pass")
	name := getenv("DB_NAME", "vaqua")
	port := getenv("DB_PORT", "5432")
	ssl  := getenv("DB_SSLMODE", "disable") // use "require" for cloud DBs if needed

	// URL format works well with GORM's postgres driver
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		user, pass, host, port, name, ssl,
	)
}

func getenv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}

func InitDb() *gorm.DB {
	dsn := dsnFromEnv()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("sql DB handle error: %v", err)
	}
	// Reasonable pool defaults
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)


/*func InitDb() *gorm.DB {

	var err error

	err = godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	//open DB
	DB, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("connected to database successfully!")


	//migrate models to create db tables
	if err := DB.AutoMigrate(&models.User{},
		&models.Account{},
		&models.Transaction{},
		&models.Transfer{},
		&models.IncomeAndExpenses{},
		&models.BlacklistedToken{},
	);  err != nil {
		log.Fatal("failed to migrate schema", err)
	}
	return DB
}

func Ping() error {
	sqlDB, err := DB.DB() 
	if err != nil {
		return err
	}
	//ping the database.
	return sqlDB.Ping()
}/*
/*package db

import (
	"fmt"
	"log"
	"os"
	"time"
	"vaqua/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// Prefer DATABASE_URL if set (Render/Heroku style). Otherwise build from parts.
func dsnFromEnv() string {
	if url := os.Getenv("DATABASE_URL"); url != "" {
		// e.g. postgres://user:pass@host:5432/dbname?sslmode=require
		return url
	}
	host := getenv("DB_HOST", "localhost")
	user := getenv("DB_USER", "postgres")
	pass := getenv("DB_PASSWORD", "pass")
	name := getenv("DB_NAME", "vaqua")
	port := getenv("DB_PORT", "5432")
	ssl  := getenv("DB_SSLMODE", "disable") // use "require" for cloud DBs if needed

	// URL format works well with GORM's postgres driver
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		user, pass, host, port, name, ssl,
	)
}

func getenv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}

func InitDb() *gorm.DB {
	dsn := dsnFromEnv()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("sql DB handle error: %v", err)
	}
	// Reasonable pool defaults
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	// Auto-migrate schemas
	if err := db.AutoMigrate(
		&models.User{},
		&models.Account{},
		&models.Transaction{},
		&models.Transfer{},
		&models.IncomeAndExpenses{},
		&models.BlacklistedToken{},
	); err != nil {
		log.Fatalf("failed to migrate schema: %v", err)
	}

	DB = db
	return DB
}

func Ping() error {
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Ping()*/