package tests

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	"SQNStest/pkg/db"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

// Единая функция тестирования для проверки подключения, таблиц и индексов
func TestDatabase(t *testing.T) {
	// Настройки подключения
	dbPass, ok := os.LookupEnv("SQNStest_MYSQL_ROOT_PASSWORD")
	if !ok {
		dbPass = "rootpassword"
	}
	dbHost, ok := os.LookupEnv("SQNStest_DB_Host_PORT")
	if !ok {
		dbHost = "localhost:3306"
	}

	// Создаем тестовую базу данных
	testDBName := "test_SQNS_db"
	setupTestDB(t, testDBName, dbPass, dbHost)
	defer cleanupTestDB(t, testDBName, dbPass, dbHost)

	// Инициализируем базу данных
	err := db.InitDB()
	assert.NoError(t, err, "Ошибка инициализации БД")

	// Проверяем подключение
	err = db.DB.Ping()
	assert.NoError(t, err, "Ошибка проверки подключения к БД")

	// Проверяем существование таблиц
	tables := []string{"logo_Text_users", "logo_Text_Comments"}
	for _, table := range tables {
		assert.True(t, db.CheckTableExists(db.DB, table),
			fmt.Sprintf("Таблица %s не существует", table))
	}

	// Проверяем индексы в таблицах
	indexes := map[string][]string{
		"logo_Text_Comments": {
			"logoTextComments_date",
			"logoTextComments_user",
		},
	}

	for table, indexList := range indexes {
		for _, index := range indexList {
			assert.True(t, checkIndexExists(db.DB, table, index),
				fmt.Sprintf("Индекс %s в таблице %s не существует", index, table))
		}
	}
}

// Функция проверки существования индекса
func checkIndexExists(db *sql.DB, table, indexName string) bool {
	query := `
    SELECT COUNT(*) 
    FROM INFORMATION_SCHEMA.STATISTICS 
    WHERE TABLE_NAME = ? 
    AND INDEX_NAME = ?
    `

	var count int
	err := db.QueryRow(query, table, indexName).Scan(&count)
	if err != nil {
		return false
	}
	return count > 0
}

// Создание тестовой БД
func setupTestDB(t *testing.T, dbName, dbPass, dbHost string) {
	testDSN := fmt.Sprintf("root:%s@tcp(%s)/", dbPass, dbHost)
	testDB, err := sql.Open("mysql", testDSN)
	assert.NoError(t, err, "Ошибка подключения к MySQL")
	defer testDB.Close()

	_, err = testDB.Exec(fmt.Sprintf("CREATE DATABASE %s", dbName))
	assert.NoError(t, err, "Ошибка создания тестовой БД")

	dbName = fmt.Sprintf("root:%s@tcp(%s)/%s", dbPass, dbHost, dbName)
	db.DB, err = sql.Open("mysql", dbName)
	assert.NoError(t, err, "Ошибка подключения к тестовой БД")
}

// Удаление тестовой БД
func cleanupTestDB(t *testing.T, dbName, dbPass, dbHost string) {
	testDSN := fmt.Sprintf("root:%s@tcp(%s)/", dbPass, dbHost)
	testDB, err := sql.Open("mysql", testDSN)
	assert.NoError(t, err, "Ошибка подключения к MySQL")
	defer testDB.Close()

	_, err = testDB.Exec(fmt.Sprintf("DROP DATABASE %s", dbName))
	assert.NoError(t, err, "Ошибка удаления тестовой БД")
}
