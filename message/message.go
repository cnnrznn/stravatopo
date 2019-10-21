package message

type Request struct {
    PingURL string `json:"PingURL"`
    GetURL string `json:"GetURL"`
}

type Response struct {
    PingLats []int `json:"PingLats"`
    GetLats []int `json:"GetLats"`
}

