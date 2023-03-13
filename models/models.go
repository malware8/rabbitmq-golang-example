package models

// AMQP

type LogItem struct {
	RequestID   string `json:"request_id"`
	ServiceName string `json:"service_name"`
	LogType     string `json:"log_type"`
	Info        string `json:"info"`
}
