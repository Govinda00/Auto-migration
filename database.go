package internal

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/source/iofs"

	"embed"
)

var sqlFiles embed.FS

func MigrateSQLs(dsn string) {
	dsn = fmt.Sprintf("%s&multiStatements=true", dsn)

	source, err := iofs.New(sqlFiles, "migrations")
	if err != nil {
		log.Fatalf("cannot prepare migrations source:%v", err)
	}

	log.Println("opening mysql connection for migrations...")
	mig, err := migrate.NewWithSourceInstance("iofs", source, "mysql://"+dsn)
	if err != nil {
		log.Fatalf("cannot prepare migration instance: %v", err)
	}

	logMigrationVersions("[before migration]", mig)

	if err := mig.Up(); err != nil && err != migrate.ErrNoChange {
		log.Printf("up migrations failed: %v", err)

		if err := mig.Down(); err != nil {
			log.Printf("migrations rolled back failed: %v", err)
		} else {
			log.Println("rolled back to last version")
		}
		panic(err)
	}

	logMigrationVersions("[after migration]", mig)
}

func logMigrationVersions(logHead string, mig *migrate.Migrate) {
	if v, d, err := mig.Version(); err != nil && err != migrate.ErrNilVersion {
		log.Fatalf("cannot load migration version: %v", err)
	} else {
		log.Printf("%s: current database version is %d (dirty=%v)", logHead, v, d)
	}
}
