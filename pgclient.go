package pgclient

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Client struct {
	pool *pgxpool.Pool
}

func (c *Client) Connect(connectionString string) error {
	var err error

	c.pool, err = pgxpool.Connect(context.Background(), connectionString)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) Close() {
	c.pool.Close()
}

func (c *Client) Ping() error {
	return c.pool.Ping(context.Background())
}

func (c *Client) Exec(sql string, arguments ...interface{}) error {
	if _, err := c.pool.Exec(context.Background(), sql, arguments...); err != nil {
		return err
	}

	return nil
}

func (c *Client) QueryRow(sql string, args ...interface{}) pgx.Row {
	return c.pool.QueryRow(context.Background(), sql, args...)
}

func (c *Client) CopyFromRows(table pgx.Identifier, columns []string, rows [][]interface{}) (int64, error) {
	var count int64 = 0

	conn, err := c.pool.Acquire(context.Background())
	if err != nil {
		return count, err
	}
	defer conn.Release()

	count, err = conn.CopyFrom(context.Background(), table, columns, pgx.CopyFromRows(rows))
	if err != nil {
		return count, err
	}

	return count, nil
}
