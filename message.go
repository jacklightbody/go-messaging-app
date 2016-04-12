package main

import "time"

type Message struct {
    From        string      `json:"from"`
    To          string      `json:"to"`
    Body        string      `json:"message"`
    Timestamp   time.Time   `json:"timestamp"`
}

type Messages []Message