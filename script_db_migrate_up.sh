DB_URL=postgres://tqsmxzpn:bkluSc8c64S14JRc_GZbVNAkmyyGw8Id@rain.db.elephantsql.com/tqsmxzpn

cd sql/schema
goose postgres $DB_URL up
read -p "Press any key to leave ..."