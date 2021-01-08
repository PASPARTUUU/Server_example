package postgres

import (
	"fmt"

	"github.com/go-pg/pg"
	"go.uber.org/atomic"

	"github.com/PASPARTUUU/Server_example/service/config"
	"github.com/PASPARTUUU/Server_example/tools/errpath"
)

// Pg -
type Pg struct {
	DB *pg.DB
}

// NewPg -
func NewPg(db *pg.DB) (*Pg, error) {

	return &Pg{DB: db}, nil
}

// -------------------------------------------------
// -------------------------------------------------
// -------------------------------------------------

// Connect -
func Connect(cfg config.Postgres) (*Pg, error) {

	// Connect to the db and remember to close it
	db := pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%v", cfg.Host, cfg.Port),
		User:     cfg.User,
		Password: cfg.Password,
		Database: cfg.DBName,
	})

	// Test connection
	var ping int
	_, err := db.QueryOne(pg.Scan(&ping), "SELECT 1")
	if err != nil {
		return nil, errpath.Err(err, "failed to connect to the db")
	}

	defer db.Close()

	return &Pg{
		DB: db,
	}, nil
}

// -------------------------------------------------
// -------------------------------------------------
// -------------------------------------------------

// DBQueryTraceHook -
type DBQueryTraceHook struct {
	enableCounter *atomic.Int32
}

// BeforeQuery -
func (db *DBQueryTraceHook) BeforeQuery(q *pg.QueryEvent) {
	if db.enableCounterVal() > 0 {
		query, err := q.FormattedQuery()
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		}
		fmt.Printf(errpath.InfofWithFuncCaller(12, "\x1b[35m %s \x1b[0m\n\n", query))

		db.enableCounterDec()
	}
}

// AfterQuery -
func (db *DBQueryTraceHook) AfterQuery(q *pg.QueryEvent) {}

// DebugQueryHook - активирует логирование SQL запроса
type DebugQueryHook interface {
	// StartTrace - начать логировать SQL
	StartTrace()
	enableCounterInc()
	enableCounterDec()
	enableCounterVal() int
}

// StartTrace - начать логировать SQL
func (db *DBQueryTraceHook) StartTrace() { db.enableCounterInc() }

func (db *DBQueryTraceHook) enableCounterInc()     { db.enableCounter.Inc() }
func (db *DBQueryTraceHook) enableCounterDec()     { db.enableCounter.Dec() }
func (db *DBQueryTraceHook) enableCounterVal() int { return int(db.enableCounter.Load()) }

// InitDebugSQLQueryHook -
func InitDebugSQLQueryHook(conn *pg.DB) *DBQueryTraceHook {
	hook := DBQueryTraceHook{
		enableCounter: atomic.NewInt32(0),
	}
	conn.AddQueryHook(&hook)
	return &hook
}
