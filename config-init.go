package main

import "github.com/spf13/viper"

// Config Struct for the API
type ApiConfig struct {
	port string
	enable bool
}

// Config Struct which will hold all configuration for the programm operation.
type Config struct {
	dbType string
	dbPath string
	sqlUser string
	sqlPassword string
	sqlDatabase string
	sqlServer string
	sqlPort string
	validityPrefix string
	api ApiConfig
}

func initConfig() *Config{
	// Set configuration Defaults
	viper.SetDefault("useFlatDB", true)
	viper.SetDefault("useKeyValueDB", false)
	viper.SetDefault("useMysqlDB", false)
	viper.SetDefault("dbPath", "data/badger.bdb")
	viper.SetDefault("flatPath", "data/database.csv")
	viper.SetDefault("apiEndpointMode", false)
	viper.SetDefault("apiEndpointPort", "8080")
	viper.SetDefault("mysqlUser", "barcodeapp")
	viper.SetDefault("mysqlPassword", "barcodeapp")
	viper.SetDefault("mysqlDatabaseName", "barcodes")
	viper.SetDefault("mysqlServerAddress", "localhost")
	viper.SetDefault("mysqlServerPort", "3306")
	viper.SetDefault("validityPrefix", "H24")

	// If it doesnt exist, create a new config file with the default values.
	viper.SafeWriteConfigAs(".config.yaml")

	// Read the config File
	viper.SetConfigName(".config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	check(err)



	// Create new Config object
	var config = Config{}

	// Set correct values depending on the selected Database type
	if viper.GetBool("useFlatDB") {
		config.dbType = "flat"
		config.dbPath = viper.GetString("dbPath")
	} else if viper.GetBool("useKeyValueDB") {
		config.dbType = "kv"
		config.dbPath = viper.GetString("flatPath")
	} else if viper.GetBool("useMysqlDB") {
		config.dbType = "sql"
		config.sqlUser = viper.GetString("mysqlUser")
		config.sqlPassword = viper.GetString("mysqlPassword")
		config.sqlDatabase = viper.GetString("mysqlDatabaseName")
		config.sqlServer = viper.GetString("mysqlServerAddress")
		config.sqlPort = viper.GetString("mysqlPort")
	}

	// Set Api Configuration if API Mode is on
	if viper.GetBool("apiEndpointMode") {
		config.api.enable = true
		config.api.port = viper.GetString("apiEndpointPort")
	}

	// Set the Prefix, to check if a barcode is valid.
	config.validityPrefix = viper.GetString("validityPrefix")

	return &config

}