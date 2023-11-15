#!/bin/sh
cd sql/schema
goose postgres $DB_URL up
# read -p "Press enter to continue"