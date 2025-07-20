package repo

type table struct {
	Name    string
	Columns []string
}

var UserTable = table{
	"user",
	[]string{"name", "username", "password"},
}