package biz

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/jedib0t/go-pretty/table"
	_ "github.com/snowflakedb/gosnowflake"
	"os"
	"talk2SQL/config"
)

const configFile = "config/snowflake.json"

const (
	InfoColor    = "\033[1;34m%s\033[0m"
	NoticeColor  = "\033[1;36m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
	DebugColor   = "\033[0;36m%s\033[0m"
)

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
		fmt.Printf(ErrorColor, err)
		return
	}

	rows, err := db.Query(query)

	if err != nil {
		fmt.Printf(ErrorColor, err)
		return
	}

	defer rows.Close()

	columns, err := rows.Columns()

	if err != nil {
		fmt.Printf(ErrorColor, err)
		return
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	header := table.Row{}
	for _, col := range columns {
		header = append(header, col)
	}
	t.AppendHeader(header)

	values := make([]interface{}, len(columns))
	valuePtrs := make([]interface{}, len(columns))

	for rows.Next() {
		for i := range columns {
			valuePtrs[i] = &values[i]
		}

		rows.Scan(valuePtrs...)

		record := make(table.Row, len(columns))
		for i, val := range values {
			b, ok := val.([]byte)

			if ok {
				record[i] = string(b)
			} else {
				record[i] = fmt.Sprint(val)
			}
		}

		t.AppendRow(record)
	}

	t.Render()

	defer db.Close()
}
