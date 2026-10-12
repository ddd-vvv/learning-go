package wakeup

import (
	"errors"
	"fmt"
	"math/rand"
)

func WakeUp(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("It's too late to wake up.")
	}
	// message := fmt.Sprintf(randomFormat(), name)
	message := fmt.Sprintf(randomFormat())
	return message, nil
}

func WakeUps(names []string) (map[string]string, error) {
	// A map to associate names with messages.
	messages := make(map[string]string)
	// Loop through the received slice of names, calling
	// the WakeUp function to get a message for each name.
	for _, name := range names {
		message, err := WakeUp(name)
		if err != nil {
			return nil, err
		}
		// In the map, associate the retrieved message with
		// the name.
		messages[name] = message
	}
	return messages, nil
}

func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Wake up, %[1]s...\nThe Matrix has you...\nFollow the white rabbit\n\n\nKnock, knock, %[1]s.",
		"Welcome to the real world, %[1]s",
		"This is your last chance, %[1]s. After this, there is no turning back.",
		"What are you waiting for? You're faster than this, %[1]s. Don't think you are, know you are.",
	}

	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}
