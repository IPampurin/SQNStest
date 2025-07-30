package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// таблицы БД
// logoTextUsers - информация о пользователях
const schemaUsers = `
CREATE TABLE logo_Text_users (
user_id INT PRIMARY KEY AUTO_INCREMENT,
registration_date DATE NOT NULL DEFAULT (CURRENT_DATE),
username VARCHAR(32) NOT NULL DEFAULT '',
login VARCHAR(32) NOT NULL DEFAULT '',
eMail VARCHAR(32) NOT NULL DEFAULT '',
pass VARCHAR(32) NOT NULL DEFAULT ''
);`

// logoTextComments - информация о комментариях
const schemaComments = `
CREATE TABLE logo_Text_Comments (
 comment_id INT PRIMARY KEY AUTO_INCREMENT,
 user_id INT NOT NULL,
 comment_date DATE NOT NULL DEFAULT (CURRENT_DATE),
 title VARCHAR(200) NOT NULL DEFAULT '',
 comment TEXT NOT NULL,
 recommendation VARCHAR(6) NOT NULL DEFAULT '',
 FOREIGN KEY (user_id) REFERENCES logo_Text_users (user_id)
);`

// baseTables список таблиц в базе данных
var baseTables = map[string]string{
	"logo_Text_users":    schemaUsers,
	"logo_Text_Comments": schemaComments,
}

// createDateIndexes и createUserIndexes для добавления индексов
// сверх ТЗ, но теоретически должно пригодится
const createDateIndexes = `
CREATE INDEX logoTextComments_date ON logo_Text_Comments (comment_date);`
const createUserIndexes = `
CREATE INDEX logoTextComments_user ON logo_Text_Comments (user_id);`

// InitDB обеспечивает подключается к базе данных, проверяет наличие необходимых
// таблиц и при необходимости их создаёт
func InitDB() error {

	// создаём подключение к базе данных с учётом параметров из .env
	dbHost, ok := os.LookupEnv("SQNStest_DB_Host_PORT")
	if !ok {
		dbHost = "localhost:3306"
	}
	dbUser, ok := os.LookupEnv("SQNStest_MYSQL_USER")
	if !ok {
		dbUser = "SQNS_User"
	}
	dbPassword, ok := os.LookupEnv("SQNStest_MYSQL_PASSWORD")
	if !ok {
		dbPassword = "123"
	}
	dbName, ok := os.LookupEnv("SQNStest_MYSQL_DATABASE")
	if !ok {
		dbName = "SQNS_DB"
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPassword, dbHost, dbName)
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("ошибка открытия базы данных: %v", err)
	}

	// проверяем подключение
	err = DB.Ping()
	if err != nil {
		return fmt.Errorf("ошибка подключения к базе данных: %v", err)
	}
	log.Println("База данных подключена")

	// проверяем наличие необходимых таблиц в базе
	for table, schema := range baseTables {
		if !CheckTableExists(DB, table) {
			_, err := DB.Exec(schema)
			if err != nil {
				return fmt.Errorf("ошибка создания таблицы %s: %v", table, err)
			}
			log.Printf("Таблица %s создана", table)
		}
	}

	// добавляем индексы если их нет
	if !CheckIndexExists(DB, "logo_Text_Comments", "logoTextComments_date") {
		_, err = DB.Exec(createDateIndexes)
		if err != nil {
			return fmt.Errorf("ошибка добавления индекса 'logoTextComments_date' %v", err)
		}
	}
	if !CheckIndexExists(DB, "logo_Text_Comments", "logoTextComments_user") {
		_, err = DB.Exec(createUserIndexes)
		if err != nil {
			return fmt.Errorf("ошибка добавления индекса 'logoTextComments_user' %v", err)
		}
	}

	return nil
}

// CheckTableExists проверяет наличие таблицы в базе данных
func CheckTableExists(db *sql.DB, tableName string) bool {
	query := `
    SELECT COUNT(*) 
    FROM INFORMATION_SCHEMA.TABLES 
    WHERE TABLE_SCHEMA = DATABASE() 
    AND TABLE_NAME = ?
    `

	var count int
	err := db.QueryRow(query, tableName).Scan(&count)
	if err != nil {
		return false
	}
	return count > 0
}

// CheckIndexExists проверяет наличие индекса в базе данных
func CheckIndexExists(db *sql.DB, tableName, indexName string) bool {
	query := `
    SELECT COUNT(*) 
    FROM INFORMATION_SCHEMA.STATISTICS 
    WHERE TABLE_SCHEMA = DATABASE() 
    AND TABLE_NAME = ?
    AND INDEX_NAME = ?
    `

	var count int
	err := db.QueryRow(query, tableName, indexName).Scan(&count)
	if err != nil {
		return false
	}
	return count > 0
}
