package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/madsilver/task-manager/internal/infra/env"
	"time"
)

type MysqlDB struct {
	Conn *sql.DB
}

func (m *MysqlDB) Query(query string, args any, fn func(scan func(dest ...any) error) error) error {
	var rows *sql.Rows
	var err error

	if args != nil && args != "" {
		rows, err = m.Conn.Query(query, args)
	} else {
		rows, err = m.Conn.Query(query)
	}
	if err != nil {
		return err
	}

	for rows.Next() {
		err = fn(rows.Scan)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *MysqlDB) QueryRow(query string, args any, fn func(scan func(dest ...any) error) error) error {
	row := m.Conn.QueryRow(query, args)
	if row.Err() != nil {
		return row.Err()
	}

	err := fn(row.Scan)
	if err != nil {
		return err
	}

	return nil
}

func (m *MysqlDB) Save(query string, args ...any) (any, error) {
	result, err := m.Conn.Exec(query, args...)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (m *MysqlDB) Update(query string, args ...any) error {
	_, err := m.Conn.Exec(query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (m *MysqlDB) Delete(query string, args any) error {
	_, err := m.Conn.Exec(query, args)
	if err != nil {
		return err
	}
	return nil
}

func NewMysqlDB() *MysqlDB {
	conn, err := sql.Open("mysql", getDSN())
	if err != nil {
		panic(err.Error())
	}

	if err := conn.Ping(); err != nil {
		panic(err.Error())
	}

	conn.SetConnMaxLifetime(time.Minute * 3)
	conn.SetMaxOpenConns(3)
	conn.SetMaxIdleConns(3)

	return &MysqlDB{
		Conn: conn,
	}
}

func getDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		env.GetString("MYSQL_USER", env.MysqlUser),
		env.GetString("MYSQL_PASSWORD", env.MysqlPassword),
		env.GetString("MYSQL_HOST", env.MysqlHost),
		env.GetString("MYSQL_PORT", env.MysqlPort),
		env.GetString("MYSQL_DATABASE", env.MysqlDatabase))
}
