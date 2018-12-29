// Author(s): Carl Saldanha

package cmd

import (
	"fmt"
	"strconv"
	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/sqlite3" // Driver
	_ "github.com/golang-migrate/migrate/source/file"       // Driver
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const migrateUsage = `
Commands:
    up                   Migrate the DB to the most recent version available
    reset                Resets the database
    to VERSION         	 Migrates up or down to the given direction
    version              Print the current version of the database
Usage:
	api_server migrate <command>
`

const driver = "sqlite"
const migrateUp = "up"
const migrateReset = "reset"
const migrateTo = "to"
const migrateVersion = "version"

func runMigrate(cmd *cobra.Command, args []string) error {
	configDir, err := cmd.Flags().GetString("conf")
	if err != nil {
		return err
	}

	// Read in configuration files
	viper.AddConfigPath(configDir)
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("Config file not found: %s", err)
	}

	return migrateSqlite(args)
}

// migrateSqlite runs sqlite migrations
func migrateSqlite(args []string) error {
	if len(args) == 0 {
		fmt.Println(migrateUsage)
		return fmt.Errorf("no commands provided")
	}

	command := args[0]

	if command == "-h" || command == "--help" {
		fmt.Println(migrateUsage)
		return nil
	}

	arguments := []string{}
	if len(args) > 1 {
		arguments = append(arguments, args[1:]...)
	}

	migrationFiles := viper.GetString("migration.file_path")
	migrationDir := fmt.Sprintf("file://%s", GetFilePath(migrationFiles))
	address := fmt.Sprintf("sqlite3://%s", viper.GetString("database.file"))

	logrus.Infof("Running Sqlite  migrations from files: %s", migrationDir)
	logrus.Infof("Connecting to Sqlite3: %s", address)

	migration, err := migrate.New(migrationDir, address)
	if err != nil {
		return err
	}

	switch command {
	case migrateUp:
		oldVersion := getDBVersion(migration)
		if err = migration.Up(); err != nil {
			return err
		}
		newVersion := getDBVersion(migration)
		logrus.Infof("Migrated from version %v to version %v", oldVersion, newVersion)
	case migrateReset:
		if err = migration.Drop(); err != nil {
			return err
		}
		logrus.Info("Finished resetting database")
	case migrateTo:
		if len(arguments) == 0 {
			return fmt.Errorf("No version provided")
		}
		version, err := strconv.ParseUint(arguments[0], 10, 64)
		if err != nil {
			return fmt.Errorf("Expected a version number, but got: %s", arguments[0])
		}
		oldVersion := getDBVersion(migration)
		if err = migration.Migrate(uint(version)); err != nil {
			return err
		}
		newVersion := getDBVersion(migration)
		logrus.Infof("Migrated from version %v to version %v", oldVersion, newVersion)
	case migrateVersion:
		version := getDBVersion(migration)
		logrus.Infof("Version: %v", version)
	default:
		return fmt.Errorf("Not one of the commands %s, %s, %s, %s", migrateUp, migrateReset, migrateTo, migrateVersion)

	}

	return nil
}

func getDBVersion(migration *migrate.Migrate) uint {
	version, dirty, err := migration.Version()
	if err != nil {
		logrus.Warn(err)
	}
	if dirty {
		logrus.Info("Version is dirty, a previous migration failed and user interaction is required.")
	}
	return version
}

