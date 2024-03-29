package config

import (
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

var _ = Describe("Config", func() {
	AfterEach(func() {
		defer viper.Reset()
	})

	Context("when configuring", func() {
		It("defaults will be provided when no overrides are present", func() {
			config, err := NewConfig()
			Expect(err).To(BeNil())
			Expect(config).To(Equal(&Config{
				Database: Database{
					Host:     "localhost",
					Port:     5432,
					Username: "tpl",
					Password: "tpl",
					DBName:   "tpl",
				},
				Logger: Logger{
					Level:     "info",
					Filename:  "tpl-log.json",
					Retention: 60,
				},
			}))
		})

		It("defaults will be overridden when environment variables are used", func() {
			t := GinkgoT()
			t.Setenv("TPL_DATABASE_HOST", "tpl")
			t.Setenv("TPL_DATABASE_PORT", "1234")
			t.Setenv("TPL_DATABASE_USERNAME", "pink")
			t.Setenv("TPL_DATABASE_PASSWORD", "purple")
			t.Setenv("TPL_DATABASE_DBNAME", "yellow")
			t.Setenv("TPL_LOGGER_LEVEL", "debug")
			t.Setenv("TPL_LOGGER_FILENAME", "brown-log.json")
			t.Setenv("TPL_LOGGER_RETENTION", "1")
			config, err := NewConfig()
			Expect(err).To(BeNil())
			Expect(config).To(Equal(&Config{
				Database: Database{
					Host:     "tpl",
					Port:     1234,
					Username: "pink",
					Password: "purple",
					DBName:   "yellow",
				},
				Logger: Logger{
					Level:     "debug",
					Filename:  "brown-log.json",
					Retention: 1,
				},
			}))
		})

		It("defaults will be overridden when configuration file is used", func() {
			tDir := GinkgoT().TempDir()
			expected := &Config{
				Database: Database{
					Host:     "lpt",
					Port:     4321,
					Username: "kinp",
					Password: "elprup",
					DBName:   "wolley",
				},
				Logger: Logger{
					Level:     "debug",
					Filename:  "gol-nworb.nosj",
					Retention: 2,
				},
			}
			yData, err := yaml.Marshal(expected)
			if err != nil {
				Fail(err.Error())
			}
			tFile, err := os.Create(filepath.Join(tDir, "tpl.yaml"))
			if err != nil {
				Fail(err.Error())
			}
			_, err = tFile.Write(yData)
			if err != nil {
				Fail(err.Error())
			}
			defer tFile.Close()

			viper.AddConfigPath(tDir)
			config, err := NewConfig()
			Expect(err).To(BeNil())
			Expect(config).To(Equal(expected))
		})

		It("defaults will be assigned when override is not present", func() {
			t := GinkgoT()
			t.Setenv("TPL_DATABASE_PORT", "1234")
			t.Setenv("TPL_DATABASE_USERNAME", "pink")
			t.Setenv("TPL_DATABASE_PASSWORD", "purple")
			t.Setenv("TPL_DATABASE_DBNAME", "yellow")
			config, err := NewConfig()
			Expect(err).To(BeNil())
			Expect(config).To(Equal(&Config{
				Database: Database{
					Host:     "localhost",
					Port:     1234,
					Username: "pink",
					Password: "purple",
					DBName:   "yellow",
				},
				Logger: Logger{
					Level:     "info",
					Filename:  "tpl-log.json",
					Retention: 60,
				},
			}))
		})

		It("error will occur when failing to marshal configuration value", func() {
			t := GinkgoT()
			t.Setenv("TPL_DATABASE_PORT", "meh")
			config, err := NewConfig()
			Expect(err).To(Not(BeNil()))
			Expect(config).To(BeNil())
		})

		It("DSN can be generated from default configuration values", func() {
			config, err := NewConfig()
			Expect(err).To(BeNil())
			Expect(config.DataSourceName()).To(Equal("postgres://tpl:tpl@localhost:5432/tpl?sslmode=disable"))
		})

		It("DSN can be generated from overridden configuration values", func() {
			t := GinkgoT()
			t.Setenv("TPL_DATABASE_PORT", "1234")
			t.Setenv("TPL_DATABASE_USERNAME", "pink")
			t.Setenv("TPL_DATABASE_PASSWORD", "purple")
			t.Setenv("TPL_DATABASE_DBNAME", "yellow")
			config, err := NewConfig()
			Expect(err).To(BeNil())
			Expect(config.DataSourceName()).To(Equal("postgres://pink:purple@localhost:1234/yellow?sslmode=disable"))
		})
	})
})
