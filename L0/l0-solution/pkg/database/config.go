package database

type ConfigDB struct {
	databaseURL string `toml:"database_url"`
}

func NewConfig() *ConfigDB {
	return &ConfigDB{
		databaseURL: "host=localhost dbname=streaming_api sslmode=disable",
	}
}
