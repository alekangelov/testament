package database

import (
	"context"
	"fmt"
	"log"
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

type RanMigrationMap map[int]RanMigration

func getRanMigrations() RanMigrationMap {
	sql := "SELECT id, name FROM migrations"
	rows, err := Store.Conn().Query(context.Background(), sql)
	if err != nil {
		return make(RanMigrationMap)
	}
	defer rows.Close()
	var ran_migrations = make(RanMigrationMap)
	for rows.Next() {
		var ran_migration RanMigration
		err := rows.Scan(&ran_migration.Id, &ran_migration.Name)
		if err != nil {
			panic(err)
		}
		ran_migrations[ran_migration.Id] = ran_migration
	}

	return ran_migrations
}

func getPendingMigrations() []Migration {

	ran_migrations := getRanMigrations()

	var migrations []Migration

	files, err := filepath.Glob("./migrations/*.sql")

	if err != nil {
		log.Printf("Error getting migrations: %v", err)
		panic(err)
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

		if _, ok := ran_migrations[id_int]; ok {
			fmt.Printf("Migration with id: %d has been ran \n", id_int)
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

	if len(migrations) == 0 {
		log.Println("No migrations to run")
		return
	}

	for _, migration := range migrations {

		fmt.Printf("Running migration %s\n", migration.Name)

		file, err := os.ReadFile(migration.file)

		if err != nil {
			log.Println("1")
			panic(err)
		}

		if _, err := Store.Conn().Exec(context.Background(), string(file)); err != nil {
			log.Println("2")
			panic(err)
		}

		fmt.Printf("Migration %s ran successfully\n", migration.Name)
		if _, err := Store.Conn().Exec(context.Background(), "INSERT INTO migrations (id, name) VALUES ($1, $2) ON CONFLICT (id) DO UPDATE SET id = EXCLUDED.id, name = EXCLUDED.name, updated_at = now()", migration.Id, migration.Name); err != nil {
			log.Println("3")
			panic(err)
		}

		fmt.Printf("Migration %s inserted successfully\n", migration.Name)

	}
}
