package benchmark

import (
	"github.com/bytedance/sonic"
	"github.com/google/uuid"
)

type ArrayOfUUID struct {
	ID []uuid.UUID
}

type ArrayOfString struct {
	ID []string
}

func AppendIndexUUID(array ArrayOfUUID) []string {
	messages := make([]string, len(array.ID))

	for i, action := range array.ID {
		commandString, _ := sonic.Marshal(action)
		messages[i] = string(commandString)
	}
	return messages
}

func AppendLimitUUID(array ArrayOfUUID) []string {
	messages := make([]string, 0, len(array.ID))

	for _, action := range array.ID {
		commandString, _ := sonic.Marshal(action)
		messages = append(messages, string(commandString))
	}
	return messages
}

func AppendFullUUID(array ArrayOfUUID) []string {
	var messages []string

	for _, action := range array.ID {
		commandString, _ := sonic.Marshal(action)
		messages = append(messages, string(commandString))
	}
	return messages
}

func AppendIndexString(array ArrayOfString) []string {
	messages := make([]string, len(array.ID))

	for i, action := range array.ID {
		commandString, _ := sonic.Marshal(action)
		messages[i] = string(commandString)
	}
	return messages
}

func AppendLimitString(array ArrayOfString) []string {
	messages := make([]string, 0, len(array.ID))

	for _, action := range array.ID {
		commandString, _ := sonic.Marshal(action)
		messages = append(messages, string(commandString))
	}
	return messages
}

func AppendFullString(array ArrayOfString) []string {
	var messages []string

	for _, action := range array.ID {
		commandString, _ := sonic.Marshal(action)
		messages = append(messages, string(commandString))
	}
	return messages
}
