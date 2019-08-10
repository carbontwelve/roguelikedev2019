package main

import (
	"github.com/mitchellh/go-wordwrap"
	"raylibtinkering/position"
	"raylibtinkering/ui"
	"strings"
)

type SimpleMessage struct {
	Message string
	Colour  string
}

type Message struct {
	Turn    uint
	Message string
	Colour  string
}

type MessageLog struct {
	ui.Component
	Messages []Message
}

func (mL *MessageLog) AddMessage(m Message) {
	wrapped := strings.Split(wordwrap.WrapString(m.Message, mL.GetInnerWidth()), "\n")

	for _, line := range wrapped {
		// If the buffer is full, remove the first line to make room for the new one
		if uint(len(mL.Messages)) == mL.GetInnerHeight() {
			copy(mL.Messages[0:], mL.Messages[0+1:])
			mL.Messages[len(mL.Messages)-1] = Message{} // or the zero value of T
			mL.Messages = mL.Messages[:len(mL.Messages)-1]
		}

		mL.Messages = append(mL.Messages, Message{Turn: m.Turn, Message: line, Colour: m.Colour})
	}

	// Draw message buffer to component
	for y, msg := range mL.Messages {
		mL.ClearRow(uint(1 + y))
		mL.SetString(msg.Message, position.Position{X: 1, Y: 1 + y}, ui.GameColours[msg.Colour], ui.ColourNC)
	}
}

func (mL MessageLog) Count() int {
	return len(mL.Messages)
}

func NewMessageLog(name string, w, h uint, offX, offY int) *MessageLog {
	mL := &MessageLog{
		Component: *ui.NewComponent(name, w, h, offX, offY, true),
		Messages:  make([]Message, 0),
	}

	mL.SetAutoClear(false)
	return mL
}
