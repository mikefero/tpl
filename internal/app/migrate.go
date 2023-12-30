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
package app

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/mikefero/tpl/db/migrations"
	"github.com/mikefero/tpl/internal/config"
	"github.com/mikefero/tpl/internal/log"
	"github.com/pressly/goose/v3"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

type MigrateType int

const (
	MigrateTypeUp MigrateType = iota
	MigrateTypeDown
	MigrateTypeStatus
	MigrateTypeRedo
	MigrateTypeReset
)

func (m MigrateType) String() string {
	return [...]string{
		"up",
		"down",
		"status",
		"redo",
		"reset",
	}[m]
}

type gooseLogger struct {
	logger *zap.Logger
}

func (g *gooseLogger) Printf(format string, v ...interface{}) {
	msg := strings.TrimSpace(strings.Replace(fmt.Sprintf(format, v...), "goose:", "", 1))
	g.logger.Info(strings.TrimSpace(msg))
	fmt.Println(msg)
}

func (g *gooseLogger) Fatalf(format string, v ...interface{}) {
	msg := strings.TrimSpace(strings.Replace(fmt.Sprintf(format, v...), "goose:", "", 1))
	g.logger.Error(msg)
	fmt.Fprintln(os.Stderr, msg)
}

func NewMigrate(t MigrateType) *fx.App {
	return fx.New(
		fx.Provide(
			func() MigrateType { return t },
			func() log.LoggerCommandType { return log.LoggerCommandTypeMigrate },
			config.NewConfig,
			log.NewLogger,
		),
		fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: logger}
		}),
		fx.Invoke(registerMigrate),
	)
}

func registerMigrate(lc fx.Lifecycle, t MigrateType, config *config.Config, logger *zap.Logger) {
	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				db, err := sql.Open("pgx", config.DataSourceName())
				if err != nil {
					return fmt.Errorf("unable to connect to database: %w", err)
				}
				defer db.Close()

				goose.SetLogger(&gooseLogger{
					logger: logger,
				})
				goose.SetBaseFS(migrations.Migrations)

				switch t {
				case MigrateTypeUp:
					if err := goose.UpContext(ctx, db, "."); err != nil {
						return fmt.Errorf("unable to migrate up database: %w", err)
					}
				case MigrateTypeDown:
					if err := goose.DownContext(ctx, db, "."); err != nil {
						return fmt.Errorf("unable to migrate up database: %w", err)
					}
				case MigrateTypeStatus:
					if err := goose.StatusContext(ctx, db, "."); err != nil {
						return fmt.Errorf("unable to get the status of the migration database table: %w", err)
					}
				case MigrateTypeRedo:
					if err := goose.RedoContext(ctx, db, "."); err != nil {
						return fmt.Errorf("unable to redo the migration database table: %w", err)
					}
				case MigrateTypeReset:
					if err := goose.ResetContext(ctx, db, "."); err != nil {
						return fmt.Errorf("unable to reset the migration database table: %w", err)
					}
				}

				return nil
			},
			OnStop: func(ctx context.Context) error {
				if err := logger.Sync(); err != nil {
					return fmt.Errorf("unable to sync logger: %w", err)
				}
				return nil
			},
		},
	)
}
