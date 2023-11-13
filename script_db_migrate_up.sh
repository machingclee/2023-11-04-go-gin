DB_URL=postgresql://root:ogCOdacTPw3suLPrhuHF@simple-bank.cin7oq1qemkd.ap-northeast-2.rds.amazonaws.com/simple_bank

cd sql/schema
goose postgres $DB_URL up
read -p "Press any key to leave ..."