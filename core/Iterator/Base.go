package Iterator

import "ApricusFuzz/core/payload"

type BaseIterator interface {
	IfEnd() bool
	Scan() bool
	Value() []string
	Exec(Payloads []payload.Payload)
	Channel() chan []string
}
