// services/database.go
package services

type MySQLDatabase struct{}

func (db *MySQLDatabase) Query() string {
	return "Executing MySQL query"
}
