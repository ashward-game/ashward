package eventerror

import "encoding/json"

type eventError struct {
	Module  string `json:"module"`
	TxHash  string `json:"txhash"`
	Message string `json:"message"`
}

func New(module, tx, message string) *eventError {
	return &eventError{module, tx, message}
}

func (e *eventError) Error() string {
	err, _ := json.Marshal(e)
	return string(err)
}
