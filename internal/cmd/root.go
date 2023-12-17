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
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var license string

type Options struct {
	License string
}

var rootCmd = &cobra.Command{
	Use:   "tpl",
	Short: "The Pinball League",
	Long: `The Pinball League for hosting pinball leagues and tournaments with
standings.`,
}

func Execute(opts Options) {
	license = opts.License
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	viper.SetEnvPrefix("TPL")
	viper.AutomaticEnv()
}