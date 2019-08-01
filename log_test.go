package main

import (
	"fmt"
	"testing"
)

func TestLineWrapping(t *testing.T) {
	messageLog := NewMessageLog(0, 3, 5)

	if len(messageLog.Messages) != 0 {
		t.FailNow()
	}

	messageLog.AddMessage(Message{Message: "Big Bad Wolf"})

	if len(messageLog.Messages) != 3 {
		t.FailNow()
	}

	messageLog.AddMessage(Message{Message: "Som and"})

	if len(messageLog.Messages) != 5 {
		t.FailNow()
	}

	messageLog.AddMessage(Message{Message: "Hello World"})

	if len(messageLog.Messages) != 5 {
		t.FailNow()
	}

	fmt.Println(messageLog.Messages)
}
