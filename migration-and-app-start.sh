
cd /app/sql/schema
/go/bin/goose postgres $DB_SOURCE up
cd /app
/app/main