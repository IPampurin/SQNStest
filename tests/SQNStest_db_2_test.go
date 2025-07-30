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

// TestInitDB проверяет корректность инициализации базы данных
func TestInitDB(t *testing.T) {

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

	// Инициализируем подключение
	err := db.InitDB() // Используем функцию из пакета db
	assert.NoError(t, err, "Ошибка инициализации БД")

	// Проверяем существование таблиц
	tables := []string{"logo_Text_users", "logo_Text_Comments"}
	for _, table := range tables {
		assert.True(t, db.CheckTableExists(db.DB, table),
			fmt.Sprintf("Таблица %s не существует", table))
	}
}

// TestCheckTableExists проверяет работу функции проверки существования таблицы
func TestCheckTableExists(t *testing.T) {
	// Проверяем существующую таблицу
	assert.True(t, db.CheckTableExists(db.DB, "logo_Text_users"),
		"Неверно определена существующая таблица")

	// Проверяем несуществующую таблицу
	assert.False(t, db.CheckTableExists(db.DB, "non_existent_table"),
		"Неверно определена несуществующая таблица")
}

func setupTestDB(t *testing.T, dbName, dbPass, dbHost string) {
	// Создаем подключение к MySQL
	testDSN := fmt.Sprintf("root:%s@tcp(%s)/", dbPass, dbHost)
	testDB, err := sql.Open("mysql", testDSN)
	assert.NoError(t, err, "Ошибка подключения к MySQL")
	defer testDB.Close()

	// Создаем тестовую БД
	_, err = testDB.Exec(fmt.Sprintf("CREATE DATABASE %s", dbName))
	assert.NoError(t, err, "Ошибка создания тестовой БД")

	// Обновляем DSN для использования тестовой БД
	dbName = fmt.Sprintf("root:%s@tcp(%s)/%s", dbPass, dbHost, dbName)
	db.DB, err = sql.Open("mysql", dbName)
	assert.NoError(t, err, "Ошибка подключения к тестовой БД")
}

func cleanupTestDB(t *testing.T, dbName, dbPass, dbHost string) {
	// Подключаемся к MySQL
	testDSN := fmt.Sprintf("root:%s@tcp(%s)/", dbPass, dbHost)
	testDB, err := sql.Open("mysql", testDSN)
	assert.NoError(t, err, "Ошибка подключения к MySQL")
	defer testDB.Close()

	// Удаляем тестовую БД
	_, err = testDB.Exec(fmt.Sprintf("DROP DATABASE %s", dbName))
	assert.NoError(t, err, "Ошибка удаления тестовой БД")
}

// TestIndexes проверяет наличие индексов
func TestIndexes(t *testing.T) {
	indexes := []string{
		"logoTextComments_date",
		"logoTextComments_user",
	}

	for _, index := range indexes {
		assert.True(t, checkIndexExists(db.DB, index),
			fmt.Sprintf("Индекс %s не существует", index))
	}
}

func checkIndexExists(db *sql.DB, indexName string) bool {
	query := `
    SELECT COUNT(*) 
    FROM INFORMATION_SCHEMA.STATISTICS 
    WHERE TABLE_SCHEMA = DATABASE() 
    AND INDEX_NAME = ?
    `

	var count int
	err := db.QueryRow(query, indexName).Scan(&count)
	if err != nil {
		return false
	}
	return count > 0
}
