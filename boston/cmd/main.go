// Author(s): Carl Saldanha

package cmd

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("application")
	viper.SetConfigType("toml")

	viper.BindEnv("go.path", "GOPATH")

	// Setup logging format
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
	})
}

var rootCmd = &cobra.Command{
		Use:   "api_server",
		Short: "Telemetry API Server",
}

var migrateCmd = &cobra.Command{
	Use:   "migrate <up|reset|to|version>",
	Short: "Run migration databases",
	RunE:  runMigrate,
}

var serveCmd = &cobra.Command{
	Use: "serve",
	Short: "Serve the server",
	RunE: startServer,

}

// Execute is called when main is executed. This is a required function for cobra to run.
func Execute() {
	// Declare flags for each command
	migrateCmd.
		Flags().
		String("conf", "conf/development", "Directory in which to find the application.toml file.")

	serveCmd.
		Flags().
		String("conf", "conf/development", "Directory in which to find the application.toml file.")

	rootCmd.AddCommand(migrateCmd)
	rootCmd.AddCommand(serveCmd)

	if err := rootCmd.Execute(); err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
}
