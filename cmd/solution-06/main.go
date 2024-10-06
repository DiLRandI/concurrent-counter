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
		tx, err := db.BeginTx(ctx, &sql.TxOptions{
			Isolation: sql.LevelSerializable,
		})
		checkError("failed to crete db transaction", err)

		rollbackFn := func(msg string, err error) {
			if err == nil {
				return
			}

			checkError("failed to rollback transaction", tx.Rollback())
			checkError(msg, err)
		}

		var value int

		row := tx.QueryRowContext(ctx, "SELECT `count_value` from `Solution06` WHERE id = 1")
		rollbackFn("error querying the Counter table", row.Err())

		err = row.Scan(&value)
		rollbackFn("error scanning the value", err)

		slog.Info("value of the counter", "value", value)

		// increment the value
		value++

		// write updated value
		_, err = tx.ExecContext(ctx, "UPDATE `Solution06` SET `count_value` = ? WHERE id = 1;", value)
		rollbackFn("error updating counter value", err)

		checkError("error when committing the transaction", tx.Commit())
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
