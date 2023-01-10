package main

import (
	//"time"
	//"database/sql"
	//"github.com/spf13/viper"
	//"github.com/rs/zerolog/log"
	"be/cmd/be/commands"
    //   "github.com/golang-migrate/migrate/v4"
	// _ "github.com/golang-migrate/migrate/v4/database/postgres"
	// _ "github.com/golang-migrate/migrate/v4/source/file"
)
	// type Config struct {
	// 	DBSource             string        `mapstructure:"DB_SOURCE"`
	// 	MigrationURL         string        `mapstructure:"MIGRATION_URL"`
	// 	}
func main() {
	// config, err := LoadConfig(".")
	// if err != nil {
	// 	log.Fatal().Err(err).Msg("cannot load config")
	// }
	// conn, err := sql.Open(config.DBDriver, config.DBSource)
	// if err != nil {
	// 	log.Fatal().Err(err).Msg("cannot connect to db")
	// }
	//runDBMigration(config.MigrationURL, config.DBSource)
	commands.Execute()
}
// func runDBMigration(migrationURL string, dbSource string) {
// 	migration, err := migrate.New(migrationURL, dbSource)
// 	if err != nil {
// 		log.Fatal().Err(err).Msg("cannot create new migrate instance")
// 	}

// 	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
// 		log.Fatal().Err(err).Msg("failed to run migrate up")
// 	}

// 	log.Info().Msg("db migrated successfully")
// }
// func LoadConfig(path string) (config Config, err error) {
// 	viper.AddConfigPath(path)
// 	viper.SetConfigName("app")
// 	viper.SetConfigType("env")

// 	viper.AutomaticEnv()

// 	err = viper.ReadInConfig()
// 	if err != nil {
// 		return
// 	}

// 	err = viper.Unmarshal(&config)
// 	return
// }

