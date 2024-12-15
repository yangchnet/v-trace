package verr

import "github.com/go-sql-driver/mysql"

func IsDuplicate(err error) bool {
	if errMySQL, ok := err.(*mysql.MySQLError); ok {
		return errMySQL.Number == 1062
	}

	return false
}

// func
