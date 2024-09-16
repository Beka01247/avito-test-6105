package db

import (
	"fmt"
	"log"
	"os"
	"zadanie-6105/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func Connect() {
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
        os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_USERNAME"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DATABASE"), os.Getenv("POSTGRES_PORT"))

    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
        NamingStrategy: schema.NamingStrategy{
            SingularTable: true,
        },
    })

    if err != nil {
        log.Fatal("Failed to connect to the database:", err)
    }

    log.Println("Connected to the database")

    
    err = DB.Exec(`DO $$ BEGIN
                    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'organization_type') THEN
                        CREATE TYPE organization_type AS ENUM ('IE', 'LLC', 'JSC');
                    END IF;
                   END $$;`).Error
    if err != nil {
        log.Fatal("Failed to create enum type:", err)
    }

    err = DB.Migrator().DropTable(&models.Tender{})
    if err != nil {
        log.Fatal("Failed to drop Tender table:", err)
    }
    log.Println("Tender table dropped successfully")

    err = DB.AutoMigrate(&models.Organization{}, &models.Tender{}, &models.Bid{}, &models.Employee{}, &models.OrganizationResponsible{}, &models.BidFeedback{}, &models.TenderVersion{}, &models.BidVersion{})
    if err != nil {
        log.Fatal("Failed to migrate database schema:", err)
    }

    log.Println("Database schema migrated successfully")

}

