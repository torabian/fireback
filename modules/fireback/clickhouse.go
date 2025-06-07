package fireback

import (
	"context"
	"fmt"
	"log"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/urfave/cli"

	"github.com/ClickHouse/clickhouse-go/v2"
)

func connectToClickHouse() (driver.Conn, error) {
	// Create a connection to ClickHouse with HTTP protocol and basic authentication
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{
			config.ClickhouseDsn,
		},
		Auth: clickhouse.Auth{
			Database: "default",
			Username: "default",
			Password: "",
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to ClickHouse: %w", err)
	}
	return conn, nil
}

var ClickHouseTestConnectionCli = cli.Command{
	Name:        "clickhouse",
	ShortName:   "ch",
	Description: "Test the clickhouse connection in the project",
	Usage:       "Test the clickhouse connection in the project",
	Action: func(c *cli.Context) error {

		conn, err := connectToClickHouse()
		if err != nil {
			log.Fatalf("Query failed: %v", err)
		}

		ctx := context.Background()

		// Test query to verify the connection
		rows, err := conn.Query(ctx, "show tables")
		if err != nil {
			log.Fatalf("Query failed: %v", err)
		}
		defer rows.Close()

		var row string
		if rows.Next() {
			if err := rows.Scan(&row); err != nil {
				log.Fatalf("Failed to scan result: %v", err)
			}
			fmt.Println("Result:", row)
		}

		return nil
	},
}

var ALL_INT_TYPES = []string{
	"int",
	"int?",

	"int64",
	"int64?",

	"int32",
	"int32?",
}

func (x Module3Field) ClickhouseType() string {
	if x.Type == "string" || x.Type == "string?" || x.Type == "json" {
		return "String"
	}

	if x.Type == "int64" || x.Type == "int64?" {
		return "Int64"
	}

	if x.Type == "int32" || x.Type == "int32?" {
		return "Int32"
	}

	if x.Type == "float32" || x.Type == "float32?" {
		return "Float32"
	}

	if x.Type == "float64" || x.Type == "float64?" {
		return "Float64"
	}

	if x.Type == "int" || x.Type == "int?" {
		return "Int"
	}

	// It's important to return empty
	return ""
}
