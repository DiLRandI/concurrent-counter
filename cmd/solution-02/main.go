package main

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"os"
	"os/signal"

	"github.com/go-sql-driver/mysql"
)

func main() {
	slog.Info("app started")
	defer slog.Info("app ended")

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	db := getMysqlConnection(ctx)

	checkError("error when doing the work", doWork(ctx, db))
}

func doWork(ctx context.Context, db *sql.DB) error {
	for range 10_000 {
		select {
		case <-ctx.Done():
			return fmt.Errorf("error context done, %w", ctx.Err())
		default:
		}

		// reading the value
		var value int

		row := db.QueryRowContext(ctx, "SELECT `count_value` from `Solution02` WHERE id = 1")
		checkError("error querying the Counter table", row.Err())
		checkError("error scanning the value", row.Scan(&value))

		slog.Info("value of the counter", "value", value)

		// increment the value
		value++

		// write updated value
		_, err := db.ExecContext(ctx, "UPDATE `Solution02` SET `count_value` = ? WHERE id = 1;", value)
		checkError("error updating counter value", err)
	}

	return nil
}

func getMysqlConnection(ctx context.Context) *sql.DB {
	config := mysql.NewConfig()
	config.Addr = "db:3306"
	config.User = "root"
	config.Passwd = "password"
	config.DBName = "concurrent_counter"

	connector, err := mysql.NewConnector(config)
	checkError("error creating mysql connector", err)

	db := sql.OpenDB(connector)
	checkError("error when ping", db.PingContext(ctx))

	return db
}

func checkError(msg string, err error) {
	if err == nil {
		return
	}

	slog.Error(msg, "error", err)
	os.Exit(1)
}
