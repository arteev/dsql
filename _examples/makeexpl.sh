#!/bin/sh

echo "Clear repository in current directory"
rm -f repository.sqlite 

echo "Create examples databases"
isql-fb -i sql/firebird.db.sql  -q
echo "Add db into repository"
dsql -r=repository.sqlite db add --code=DB1  --engine=firebirdsql --uri=sysdba:masterkey@/tmp/dsql.exsample1.fdb
dsql -r=repository.sqlite db add --code=DB2  --engine=firebirdsql --uri=sysdba:masterkey@/tmp/dsql.exsample2.fdb
