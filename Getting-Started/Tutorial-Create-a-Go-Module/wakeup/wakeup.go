package wakeup

import (
	"errors"
	"fmt"
)

func WakeUp(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("It's too late to wake up.")
	}
	message := fmt.Sprintf("Wake up, %v...\nThe Matrix has you...\nFollow the white rabit\n\n\nKnock, knock, %v.", name, name)
	return message, nil
}
