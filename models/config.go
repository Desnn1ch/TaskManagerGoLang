package models

type DataBase struct {
	Host     string
	Port     string
	User     string
	Name     string
	Password string
	SSLMode  string
}

type JWT struct {
	Secret_key string
}

type Redis struct {
	Port string
	Host string
	Ttl  int
}
