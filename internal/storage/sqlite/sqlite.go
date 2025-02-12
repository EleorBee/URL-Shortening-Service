package sqlite

import (
	"URLShort/internal/model"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

type Storage struct {
	db *sql.DB
}

func New(storagePath string) (*Storage, error) {

	db, err := sql.Open("sqlite3", storagePath)

	if err != nil {
		return nil, fmt.Errorf("Error opening SQLite database: %v", err)
	}

	stmt, err := db.Prepare(
		"CREATE TABLE IF NOT EXISTS url(id INTEGER PRIMARY KEY AUTOINCREMENT," +
			"url TEXT NOT NULL," +
			"shortCode TEXT NOT NULL UNIQUE," +
			"createAt DATETIME NOT NULL," +
			"updateAt DATETIME NOT NULL," +
			"accessCount INTEGER DEFAULT 0);" +
			"INDEX url_code IF NOT EXISTS ON url(shortCode);")

	if err != nil {
		return nil, fmt.Errorf("Error preparing SQLite database: %v", err)
	}
	_, err = stmt.Exec()

	if err != nil {
		return nil, fmt.Errorf("Error executing SQLite database: %v", err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) GetUrl(code string) (*model.ShortUrl, error) {
	stmt, err := s.db.Prepare("SELECT * FROM url WHERE shortCode=?")

	if err != nil {
		return nil, fmt.Errorf("Error preparing SQLite database: %v", err)
	}

	model := new(model.ShortUrl)

	err = stmt.QueryRow(code).Scan(&model.Id, &model.Url, &model.ShortCode, &model.CreatedAt, &model.UpdatedAt, &model.AccessCount)

	if err != nil {
		return nil, fmt.Errorf("Error executing SQLite database: %v code: %d", err, code)
	}

	return model, nil
}

func (s *Storage) VisitUrl(count *int64, code string) error {
	stmt, err := s.db.Prepare("UPDATE url SET accessCount=? WHERE shortCode=?")

	if err != nil {
		return fmt.Errorf("Error preparing SQLite database: %v", err)
	}
	*count = *count + 1
	_, err = stmt.Exec(count, code)

	if err != nil {
		return fmt.Errorf("Error executing SQLite database: %v", err)
	}

	return nil
}

func (s *Storage) GetUrlStats(code string) (*model.ShortUrl, error) {
	stmt, err := s.db.Prepare("SELECT * FROM url WHERE shortCode=?")
	if err != nil {
		return nil, fmt.Errorf("Error preparing SQLite database: %v", err)
	}

	um := new(model.ShortUrl)

	err = stmt.QueryRow(code).Scan(&um.Id, &um.Url, &um.ShortCode, &um.CreatedAt, &um.UpdatedAt, &um.AccessCount)

	if err != nil {
		return nil, fmt.Errorf("Error executing SQLite database: %v", err)
	}

	return um, nil
}

func (s *Storage) SaveUrl(url string, code string) error {
	stmt, err := s.db.Prepare("INSERT INTO url(shortCode,url,createAt,updateAt) VALUES(?,?,?,?)")

	if err != nil {
		return fmt.Errorf("Error preparing SQLite database: %v", err)
	}

	_, err = stmt.Exec(code, url, time.Now(), time.Now())

	if err != nil {
		return fmt.Errorf("Error executing SQLite database: %v", err)
	}

	return nil
}

func (s *Storage) DeleteUrl(code string) error {
	stmt, err := s.db.Prepare("DELETE FROM url WHERE shortCode=?")

	if err != nil {
		return fmt.Errorf("Error preparing SQLite database: %v", err)
	}

	_, err = stmt.Exec(code)

	if err != nil {
		return fmt.Errorf("Error executing SQLite database: %v", err)
	}

	return nil
}

func (s *Storage) UpdateUrl(code string, url string) error {
	stmt, err := s.db.Prepare("UPDATE url SET url=?,updateAt=? WHERE shortCode=?")

	if err != nil {
		return fmt.Errorf("Error preparing SQLite database: %v", err)
	}

	_, err = stmt.Exec(url, time.Now(), code)

	if err != nil {
		return fmt.Errorf("Error executing SQLite database: %v", err)
	}

	return nil
}
