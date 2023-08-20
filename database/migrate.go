package database

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Migration struct {
	Id   int
	Name string
	file string
}

type RanMigration struct {
	Id   int
	Name string
}

func getRanMigrations() []RanMigration {
	sql := "SELECT id, name FROM migrations"
	rows, err := Store.Conn().Query(context.Background(), sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var ran_migrations []RanMigration
	for rows.Next() {
		var ran_migration RanMigration
		err := rows.Scan(&ran_migration.Id, &ran_migration.Name)
		if err != nil {
			panic(err)
		}
		fmt.Printf(
			"Migration with id: %d has been ran \n", ran_migration.Id,
		)
		ran_migrations = append(ran_migrations, ran_migration)
	}

	return ran_migrations
}

func getPendingMigrations() []Migration {
	ran_migrations := getRanMigrations()
	var migrations []Migration
	files, err := filepath.Glob("./migrations/*.sql")

	if err != nil {
		panic(`Error getting migrations`)
	}

	for _, file := range files {
		name := filepath.Base(file)
		name_without_extension := strings.Replace(name, ".sql", "", -1)
		id_name := strings.Split(name_without_extension, "_")
		id := id_name[0]
		id_int, err := strconv.Atoi(id)
		if err != nil {
			id_int = 0
		}
		migration_exists := false
		for _, ran_migration := range ran_migrations {
			if ran_migration.Id == id_int {
				migration_exists = true
				break
			}
		}
		if migration_exists {
			continue
		}

		migration_name := ""
		if len(id_name) > 1 {
			migration_name = id_name[1]
		}

		migrations = append(migrations, Migration{
			Id:   id_int,
			Name: migration_name,
			file: file,
		})

	}

	return migrations

}

func Migrate() {
	migrations := getPendingMigrations()

	for _, migration := range migrations {
		fmt.Printf("Running migration %s\n", migration.Name)
		file, err := os.ReadFile(migration.file)
		if err != nil {
			panic(err)
		}
		if _, err := Store.Conn().Exec(context.Background(), string(file)); err != nil {
			panic(err)
		}

		fmt.Printf("Migration %s ran successfully\n", migration.Name)
		if _, err := Store.Conn().Exec(context.Background(), "INSERT INTO migrations (id, name) VALUES ($1, $2) ON CONFLICT (id) DO UPDATE SET id = EXCLUDED.id, name = EXCLUDED.name, updated_at = now()", migration.Id, migration.Name); err != nil {
			panic(err)
		}
		fmt.Printf("Migration %s inserted successfully\n", migration.Name)

	}
}
