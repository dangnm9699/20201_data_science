package model

type Stat struct {
	Payload Payload `json:"payload"`
}

type Payload struct {
	Last30Days Load `json:"Last_30_Days"`
	Last90Days Load `json:"Last_90_Days"`
}

type Load struct {
	High       float64 `json:"high"`
	Low        float64 `json:"low"`
	Average    float64 `json:"average"`
	Volatility float64 `json:"volatility"`
}
