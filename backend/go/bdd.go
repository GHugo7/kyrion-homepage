package main

import (
	"database/sql"

	_ "github.com/glebarez/go-sqlite"
	"github.com/metalblueberry/console"
)

func InitTable(db *sql.DB, containerNames []string) error {
	createSQL := `CREATE TABLE IF NOT EXISTS service_overrides (
		container_name TEXT PRIMARY KEY,
		category TEXT,
		url TEXT,
		enabled BOOLEAN DEFAULT 1
	);`

	if _, err := db.Exec(createSQL); err != nil {
		return err
	}

	for _, name := range containerNames {
		_, err := db.Exec(`
			INSERT OR IGNORE INTO service_overrides (container_name, category, url, enabled)
			VALUES (?, ?, ?, ?)`,
			name, "Autre", "", 0)
		if err != nil {
			return err
		}
	}

	return nil
}

func ConnectDB(containerNames []string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", "./home.db?_pragma=foreign_keys(1)")
	if err != nil {
		return nil, err
	}

	if err := InitTable(db, containerNames); err != nil {
		db.Close()
		return nil, err
	}

	console.Warn("DB Connected !")
	return db, nil
}

func loadOverride(db *sql.DB) (map[string]Override, error) {
	rows, err := db.Query("SELECT container_name, category, url, enabled FROM service_overrides")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	overrides := map[string]Override{}
	for rows.Next() {
		var name string
		var o Override
		if err := rows.Scan(&name, &o.Category, &o.URL, &o.Enabled); err != nil {
			return nil, err
		}
		overrides[name] = o
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return overrides, nil
}
