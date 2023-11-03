package biz

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/snowflakedb/gosnowflake"
	"os"
	"talk2SQL/config"
)

const configFile = "config/snowflake.json"

func loadConfig(file string) (config.SnowflakeConfig, error) {
	var config config.SnowflakeConfig

	configFile, err := os.Open(file)
	if err != nil {
		return config, err
	}
	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)

	return config, nil
}

func Execute(query string) {

	config, _ := loadConfig(configFile)

	connectionString := fmt.Sprintf("%s:%s@%s/%s?warehouse=%s", config.Username, config.Password, config.Account,
		config.Database, config.Warehouse)

	db, err := sql.Open("snowflake", connectionString)

	if err != nil {
		fmt.Println(err)
	}

	rows, err := db.Query(query)

	if err != nil {
		fmt.Println(err)
	}

	var name string

	rows.Scan(&name)

	fmt.Println(name)

	defer db.Close()
}
