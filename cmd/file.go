// Author(s): Carl Saldanha
package cmd

import (
	"fmt"

	"github.com/spf13/viper"
)

// GetFilePath takes in a sub path relative to the package root and returns the absolute path.
// Example:
// - input: "storage/file.go"
// - output: "$GOROOT/src/fetchcore-services-platform/storage/file.go"
func GetFilePath(subPath string) string {
	goPath := viper.GetString("go.path")
	pkgRoot := viper.GetString("pkg_root")
	return fmt.Sprintf("%s/%s/%s", goPath, pkgRoot, subPath)
}
