package tests

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// получаем URL для тестирования
func getURL(path string) string {
	port := 8080
	envPort := os.Getenv("SQNStest_PORT")
	if len(envPort) > 0 {
		if eport, err := strconv.ParseInt(envPort, 10, 32); err == nil {
			port = int(eport)
		}
	}
	path = strings.TrimPrefix(strings.ReplaceAll(path, `\`, `/`), `../web/`)
	return fmt.Sprintf("http://localhost:%d/%s", port, path)
}

// получаем тело ответа от сервера
func getBody(path string) ([]byte, error) {
	resp, err := http.Get(getURL(path))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// проверяем статус ответа
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("неверный статус ответа: %d", resp.StatusCode)
	}

	return io.ReadAll(resp.Body)
}

// рекурсивно обходим директорию
func walkDir(path string, f func(fname string) error) error {
	dirs, err := os.ReadDir(path)
	if err != nil {
		return err
	}
	for _, v := range dirs {
		fname := filepath.Join(path, v.Name())
		if v.IsDir() {
			if err = walkDir(fname, f); err != nil {
				return err
			}
			continue
		}
		if err = f(fname); err != nil {
			return err
		}
	}
	return nil
}

// основной тест сервера
func TestServer(t *testing.T) {
	// функция сравнения файлов
	cmp := func(fname string) error {
		// читаем файл с диска
		fbody, err := os.ReadFile(fname)
		if err != nil {
			return fmt.Errorf("ошибка чтения файла %s: %w", fname, err)
		}

		// получаем данные с сервера
		body, err := getBody(fname)
		if err != nil {
			return fmt.Errorf("ошибка получения данных с сервера для %s: %w", fname, err)
		}

		// проверяем только размер файла
		assert.Equal(t, len(fbody), len(body),
			fmt.Sprintf("сервер возвращает для %s данные другого размера", fname))

		return nil
	}

	// запускаем проверку всех файлов
	assert.NoError(t, walkDir("../web", cmp), "ошибка при проверке файлов")

	// проверяем корневую директорию
	_, err := getBody("/")
	assert.NoError(t, err, "ошибка при запросе корневой директории")
}
