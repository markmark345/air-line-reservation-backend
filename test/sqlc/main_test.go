package sqlc_test

import (
	"context"
	"log"
	"os"
	"testing"

	db "air-line-reservation-backend/internal/infrastucture/postgres"

	"github.com/jackc/pgx/v5/pgxpool"

	_ "github.com/lib/pq"
)

const DBSource = "postgresql://root:password@localhost:5050/db_air_line_reservation?sslmode=disable"

var testQueries *db.Queries

func TestMain(m *testing.M) {
	connPool, err := pgxpool.New(context.Background(), DBSource)

	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}
	
	testQueries = db.New(connPool)

	os.Exit(m.Run())
}
