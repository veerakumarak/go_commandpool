package commandbus

import (
	"encoding/json"
	"errors"
)

type Message json.RawMessage

func (m Message) valid() error {
	if !json.Valid(m) {
		return errors.New("message is not a valid json")
	}
	return nil
}

func UnMarshal[v any](message json.RawMessage, variable v) (v, error) {
	err := json.Unmarshal(message, variable)
	return variable, err
}

func Marshal[v any](variable v) (Message, error) {
	return json.Marshal(variable)
}

//func (m Message) to(s T) error {
//
//}
