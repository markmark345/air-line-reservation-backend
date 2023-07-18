cp ./db/sqlc/db.go ./internal/infrastucture/postgres/db.go
cp ./db/sqlc/models.go ./internal/infrastucture/postgres/model/models.go
cp ./db/sqlc/querier.go ./internal/infrastucture/postgres/querier_repository.go
cp ./db/sqlc/*.sql.go ./internal/infrastucture/postgres/

# rm -rf ./db/sqlc

