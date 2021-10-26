package tools

import "log"

// FatalError return faltal error
func FatalError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

// DomainError Print Domain error
func DomainError(err error, msg string) {
	if err != nil {
		log.Printf("%s: %s", msg, err)
	}
}
