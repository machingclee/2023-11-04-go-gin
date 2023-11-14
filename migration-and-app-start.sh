
# cd /app/sql/schema
# export DB_SOURCE="postgresql://pguser:pguser@127.0.0.1:5432/pgdb?sslmode=disable"
# /go/bin/goose postgres $DB_SOURCE up
cd /app
/app/main