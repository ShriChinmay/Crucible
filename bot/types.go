package main

type Order struct {
    RunID         string
    BotID         string
    OrderID       string
    Type          string
    Side          string
    Price         float64
    Qty           int
    TimestampSent int64
}

type OrderResult struct {
    RunID         string
    BotID         string
    OrderID       string
    Type          string
    Status        string
    LatencyNs     int64
    TimestampSent int64
    ErrorMsg      string
}