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
	"github.com/mikefero/tpl/internal/app"
	"github.com/spf13/cobra"
)

var migrateDownCmd = &cobra.Command{
	Use:   "down",
	Short: "Execute migration down operations against The Pinball League database",
	Long: `Execute migration down operations against The Pinball League database
reverting the database structure to its previous state.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return migrate(app.MigrateTypeDown)
	},
}

func init() {
	migrateCmd.AddCommand(migrateDownCmd)
}
