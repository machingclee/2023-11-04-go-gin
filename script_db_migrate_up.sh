DB_URL=postgresql://pguser:pguser@127.0.0.1:5432/pgdb

cd sql/schema
goose postgres $DB_URL up
read -p "Press any key to leave ..."