package main

import (
	"bytes"
	"log"
	"os"
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/stretchr/testify/assert"
)

func captureOutput(f func()) string {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	f()
	log.SetOutput(os.Stderr)
	return buf.String()
}

func TestMessage(t *testing.T) {
	s := &discordgo.Session{}
	m := &discordgo.MessageCreate{
		Message: &discordgo.Message{
			ID:        "messageid",
			Content:   "!ping",
			ChannelID: "channelid",
		},
	}
	output := captureOutput(func() {
		MessageCreate(s, m)
	})
	assert.Equal(t, "removed certificate www.example.com\n", output)
}
