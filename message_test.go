package main

import (
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/stretchr/testify/assert"
)

func Session() discordgo.Session {
	s := discordgo.Session{StateEnabled: true, State: discordgo.NewState()}
	s.State.User = &discordgo.User{ID: "1234"}
	return s
}

func TestMessageCreate(t *testing.T) {
	s := Session()
	m := &discordgo.MessageCreate{
		Message: &discordgo.Message{
			ID:      "messageid",
			Content: "!ping",
			Author: &discordgo.User{
				ID: "123",
			},
			ChannelID: "channelid",
		},
	}
	output := MessageCreate(&s, m)
	assert.Equal(t, `{"id":"messageid","channel_id":"channelid","content":"!ping","author":"123"}`, output)
}

func TestMessageDelete(t *testing.T) {
	s := Session()
	m := &discordgo.MessageDelete{
		Message: &discordgo.Message{
			ID:      "messageid",
			Content: "!ping",
			Author: &discordgo.User{
				ID: "123",
			},
			ChannelID: "channelid",
		},
	}
	output := MessageDelete(&s, m)
	assert.Equal(t, `{"id":"messageid","content":"!ping"}`, output)
}

func TestMessageEdit(t *testing.T) {
	s := Session()
	content := "!ping"
	m := &discordgo.MessageEdit{
		ID:      "messageid",
		Content: &content,
	}
	output := MessageEdit(&s, m)
	assert.Equal(t, `{"id":"messageid","content":"!ping"}`, output)
}
