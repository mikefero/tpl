/*
Copyright © 2023 Michael Fero

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	version   string
	commit    string
	osArch    string
	goVersion string
	buildDate string
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the The Pinball League version",
	Long: `The version command prints the version of The Pinball League along
with a git commit hash of the source tree, OS, architecture, go version, and
build date.`,
	Run: func(cmd *cobra.Command, args []string) {
		formatVersion := version
		if len(formatVersion) == 0 {
			formatVersion = "dev"
		}
		if len(commit) > 0 {
			formatVersion = fmt.Sprintf("%s (%s)", formatVersion, commit)
		}
		if len(osArch) > 0 {
			formatVersion = fmt.Sprintf("%s %s", formatVersion, osArch)
		}
		fmt.Printf("The Pinball League version %s\n", formatVersion) //nolint:forbidigo
		if len(goVersion) > 0 {
			fmt.Printf("go version %s\n", goVersion) //nolint:forbidigo
		}
		if len(buildDate) > 0 {
			fmt.Printf("Built on %s\n", buildDate) //nolint:forbidigo
		}
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
