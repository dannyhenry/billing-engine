package database

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseManager interface {
	GetMaster() *gorm.DB
	StartTransaction() *gorm.DB
	CommitTransaction(tx *gorm.DB) *gorm.DB
	RollbackTransaction(tx *gorm.DB) *gorm.DB

	Initialize(dsn string, maxIdleConns int, maxOpenConns int) error
}

func NewDatabaseManager() DatabaseManager {
	return &databaseManager{}
}

type databaseManager struct {
	Master *gorm.DB
}

func (dbManager *databaseManager) Initialize(dsn string, maxIdleConns int, maxOpenConns int) error {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(maxIdleConns)
	sqlDB.SetMaxOpenConns(maxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Hour * 2)

	dbManager.Master = db
	return nil
}

func (dbManager *databaseManager) GetMaster() *gorm.DB {
	if dbManager.Master == nil {
		return nil
	}

	return dbManager.Master
}

func (dbManager *databaseManager) StartTransaction() *gorm.DB {
	return dbManager.Master.Begin()
}

func (dbManager *databaseManager) CommitTransaction(tx *gorm.DB) *gorm.DB {
	return tx.Commit()
}

func (dbManager *databaseManager) RollbackTransaction(tx *gorm.DB) *gorm.DB {
	return tx.Rollback()
}

func PostgresURI(dbUserName, dbPassword, dbAddress, dbName, dbSchema string) string {
	return fmt.Sprintf(`postgres://%s:%s@%s/%s?sslmode=disable&search_path=%s`,
		dbUserName, dbPassword, dbAddress, dbName, dbSchema)
}
