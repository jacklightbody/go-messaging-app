package main

import (
    "time"
    "encoding/json"
    "os"
    "database/sql"
    _ "github.com/ziutek/mymysql/godrv"
)
func db(){
	database, user, password = loadconfig()
	con, err := sql.Open("mymysql", database+"/"+user+"/"+password)
	defer con.Close()
}
func loadconfig(){
	file, _ := os.Open("conf.json")
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	decoder.Decode(&configuration)
	return configuration.Database, configuration.User, configuration.Password
}
func NewMessage(from, to, body){

}
type Config struct{
	Database string
	User string
	Password string
}
type Message struct {
    From        string      `json:"from"`
    To          string      `json:"to"`
    Body        string      `json:"message"`
    Timestamp   time.Time   `json:"timestamp"`
}

type Messages []Message