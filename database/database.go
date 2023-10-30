package database

import "fmt"

var connection string

func init() {
	fmt.Println("Init di jalankan") // to test initialization has run without function accessed.
	connection = "MySQL"
}

func GetDatabase() string {
	return connection
}
