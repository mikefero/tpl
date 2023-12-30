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

var migrateRedoCmd = &cobra.Command{
	Use:   "redo",
	Short: "Execute migration redo operations against The Pinball League database",
	Long: `Execute migration redo operations against The Pinball League database
reverting the most recently applied migration and executing it again.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return migrate(app.MigrateTypeRedo)
	},
}

func init() {
	migrateCmd.AddCommand(migrateRedoCmd)
}
