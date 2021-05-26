package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

type CordMessageCreate struct {
	ID string `json:"id"`

	// The ID of the channel in which the message was sent.
	ChannelID string `json:"channel_id"`

	// The content of the message.
	Content string `json:"content"`

	// The author of the message. This is not guaranteed to be a
	// valid user (webhook-sent messages do not possess a full author).
	AuthorID string `json:"author"`
}

type CordMessageDelete struct {
	ID string `json:"id"`

	// The content of the message.
	Content string `json:"content"`
}

type CordMessageEdit struct {
	ID string `json:"id"`

	// The content of the message.
	Content *string `json:"content"`
}

func SerializeCreate(m *discordgo.MessageCreate) ([]byte, error) {
	return json.Marshal(CordMessageCreate{
		ID:        m.ID,
		ChannelID: m.ChannelID,
		AuthorID:  m.Author.ID,
		Content:   m.Content,
	})
}

func SerializeEdit(m *discordgo.MessageEdit) ([]byte, error) {
	return json.Marshal(CordMessageEdit{
		ID:      m.ID,
		Content: m.Content,
	})
}

func SerializeDelete(m *discordgo.MessageDelete) ([]byte, error) {
	return json.Marshal(CordMessageDelete{
		ID:      m.ID,
		Content: m.Content,
	})
}

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	b, err := SerializeCreate(m)
	if err != nil {
		fmt.Println("error:", err)
	}

	os.Stdout.Write(b)
}

func MessageDelete(s *discordgo.Session, m *discordgo.MessageDelete) {
	b, err := SerializeDelete(m)
	if err != nil {
		fmt.Println("error:", err)
	}

	os.Stdout.Write(b)
}

func MessageEdit(s *discordgo.Session, m *discordgo.MessageEdit) {
	b, err := SerializeEdit(m)
	if err != nil {
		fmt.Println("error:", err)
	}

	os.Stdout.Write(b)
}
