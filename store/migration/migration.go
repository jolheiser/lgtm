package migration

//go:generate go-bindata -pkg migration -o bindata.go sqlite3/ mysql/
//go:generate go fmt bindata.go
