package data

import (
	"database/sql"
	"fmt"
	upperDBconn "github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/mysql"
	"github.com/upper/db/v4/adapter/postgresql"
	"os"
)

// upperDB start upper db database connection session
var upperDBSession upperDBconn.Session

type Models struct {
	// any models inserted here (and in the New function)
	// are easily accessible throughout the entire application

}

// New upperDB takes the connection pool and creates adapter to access database
// return models type
func New(dbConnPool *sql.DB) Models {

	switch os.Getenv("DATABASE_TYPE") {
	case "mysql", "mariadb":
		// wrap the connection pool and create a new sqlbuilder session
		upperDBSession, _ = mysql.New(dbConnPool)
	case "postgresql", "postgres":
		// wrap the connection pool and create a new sqlbuilder session
		upperDBSession, _ = postgresql.New(dbConnPool)
	default:
		// no database specified
	}

	return Models{}
}

func getInsertId(i upperDBconn.ID) int {
	idType := fmt.Sprintf("%T", i)
	if idType == "int64" {
		return int(i.(int64))
	}
	return i.(int)
}
