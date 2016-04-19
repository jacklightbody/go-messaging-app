package main

import (
    "time"
    "log"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)
func loadconfig() (string, string, string, string){
	// because working with a sensible config file is a pain
	// just do this. Good for a quick hack
	Database := "SimpleMessaging"
	User := "root"
    Password := "password"
    Host := "localhost"
	return Database, User, Password, Host
}
func NewMessage(from string, to string, body string){
	// controller validated already
	now := time.Now()
    t := now.Unix()
    // get current timestamp
	database, user, password, host := loadconfig()
	con, err := sql.Open("mysql", user+":"+password+"@tcp("+host+")/"+database)
	defer con.Close()
	if err != nil { log.Fatal(err)}
	_, err = con.Exec("INSERT into Messages (sender, recipient, body, sent_at) values (?, ?, ?, ?)", from, to, body, t)
	if err != nil { log.Fatal(err)}
}
func GetMessagesAfter(from string, to string, timestamp int) Messages{
    database, user, password, host := loadconfig()
	con, err := sql.Open("mysql", user+":"+password+"@tcp("+host+")/"+database)
	defer con.Close()
	rows, err := con.Query("SELECT * FROM Messages WHERE sent_at>? AND ((sender=? AND recipient=?) OR (sender=? AND recipient=?))", timestamp, from, to, to, from)
	// Get messages sent either way
	defer rows.Close()
	if err != nil { log.Fatal(err)}
	items := Messages{}
	var (
		body, sender, recipient string
		sent_at, id 				int
	)
	// we need empty vars that scan can read into
	for rows.Next() {
	    err = rows.Scan(&id, &sender, &recipient, &body, &sent_at)// read 'em in
	    if err != nil { log.Fatal(err)}
	    // 
	    items = append(items, Message{From: sender, To: recipient, Body: body, Timestamp: sent_at})
	    // add to our list of messages
	}
	err = rows.Err()
	if err != nil { log.Fatal(err)}
	return items
}
type Message struct {
    From        string      `json:"from"`
    To          string      `json:"to"`
    Body        string      `json:"message"`
    Timestamp   int   		`json:"at"`
}
type Messages []Message