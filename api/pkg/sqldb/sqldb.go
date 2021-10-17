package sqldb

import "database/sql"

// globalDB => global variable to store database connection
var globalDB *sql.DB
