package database

var connection string

// function yang otomatis dijalankan ketika method dalam package ini dipanggil
func init() {
	connection = "MySQL"
}

func GetDatabase() string {
	return connection
}