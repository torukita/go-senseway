package model

type Gateway struct {
	Date      string  `json:"date"`
	Rssi      int32   `json:"rssi"`
	Snr       float32 `json:"snr"`
	GatewayId string  `json:"gwid"`
}

type Module struct {
	Datarate    string  `json:"dr"`
	FCntUp      uint32  `json:"cnt"`
	Frequency   float32 `json:"fq"`
	FPort       uint8   `json:"port"`
	MessageType string  `json:"mt"`
	DevEUI      string  `json:"devEUI"`
	Data        string  `json:"data"` // hexadecimal string
}

type UplinkMessage struct {
	Gateways []Gateway `json:"gw"`
	Module   Module    `json:"mod"`
}

type DownlinkMessage struct {
	Confirmed bool   `json:"cnf"`
	Reference string `json:"ref"`
	FPort     uint8  `json:"port"`
	Data      string `json:"data"` // hexadecimal string
}
