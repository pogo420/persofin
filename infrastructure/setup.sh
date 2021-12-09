echo "initializing database using: $SQLITE_DB"
sqlite3 $SQLITE_DB < infrastructure/setup.sql
echo "Done !"
