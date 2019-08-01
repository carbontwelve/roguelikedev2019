package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/mitchellh/go-wordwrap"
	"strings"
)

type SimpleMessage struct {
	Message string
	Colour  rl.Color
}

type Message struct {
	Turn    uint
	Message string
	Colour  rl.Color
}

type MessageLog struct {
	Messages []Message
	X        uint
	Width    uint
	Height   uint
}

func (mL *MessageLog) AddMessage(m Message) {
	wrapped := strings.Split(wordwrap.WrapString(m.Message, mL.Width), "\n")

	for _, line := range wrapped {
		// If the buffer is full, remove the first line to make room for the new one
		if uint(len(mL.Messages)) == mL.Height {
			copy(mL.Messages[0:], mL.Messages[0+1:])
			mL.Messages[len(mL.Messages)-1] = Message{} // or the zero value of T
			mL.Messages = mL.Messages[:len(mL.Messages)-1]
		}

		mL.Messages = append(mL.Messages, Message{Turn: m.Turn, Message: line, Colour: m.Colour})
	}
}

func (mL MessageLog) Count() int {
	return len(mL.Messages)
}

func NewMessageLog(x, width, height uint) *MessageLog {
	return &MessageLog{
		Messages: make([]Message, 0),
		X:        x,
		Width:    width,
		Height:   height,
	}
}
