package commandbus

import "encoding/json"

type HandlerFunc func(payload json.RawMessage) (json.RawMessage, error)

func (h HandlerFunc) valid() error {
	return nil
}
