// Author(s): Carl Saldaha

package cmd

import (
	"franklin/boston/server"
	"net/http"
        "franklin/boston/model"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/jinzhu/gorm"
	  _ "github.com/jinzhu/gorm/dialects/sqlite"
)

func startServer(cmd *cobra.Command, args []string) error {
	// Read in configuration files
	configDir, err := cmd.Flags().GetString("conf")
	if err != nil {
		return err
	}

	viper.AddConfigPath(configDir)
	if err := viper.ReadInConfig(); err != nil {
		logrus.Errorf("config file not found: %s\n", err)
		return err
	}

	// Run migration
	// TODO (Carl): Remove and replace with GORM Auto Migrate
	if err := migrateSqlite([]string{migrateUp}); err != nil {
		logrus.Warnf("db migration: %s", err)
	}

	//
	db, err := gorm.Open("sqlite3", "test.db")

	db.AutoMigrate(&model.Task{})

	//dbOptions, err := getDatabaseOptions()
	server := &http.Server{
		Addr:    viper.GetString("server.addr"),
		Handler: server.LoadRoutes(db),
	}

	logrus.Infof("listening and serving on %v", viper.GetString("server.addr"))
	if err := server.ListenAndServe(); err != nil {
		logrus.Fatalf("failed to start HTTP server: %v", err)
	}
	return nil
}
