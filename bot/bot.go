package bot

import (
	"context"
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"

	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
)

var BotID string
var goBot *discordgo.Session

var openaiClient *openai.Client

func Start() {

	err := godotenv.Load()
    if err != nil {
    fmt.Println("err",(err.Error()))
		return
}
	goBot, err := discordgo.New("Bot " +os.Getenv("TOKEN") )

	fmt.Println("token",os.Getenv("TOKEN"))
	if err != nil {
		fmt.Println("err1",(err.Error()))
		return
	}

	u, err := goBot.User("@me")
	if err != nil {
		fmt.Println("err2",(err.Error()))
		return
	}

	BotID = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()
	if err != nil {
		fmt.Println((err.Error()))
		return
	}

	openaiClient = openai.NewClient(os.Getenv("API_KEY"))

	fmt.Println("Bot is running")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotID {
		return
	}

	if m.Content == "hello" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Hello")
	} else {
		response, err := generateResponse(m.Content)
		if err != nil {
			fmt.Println("Failed to generate OpenAI response:", err)
			return
		}

		_, _ = s.ChannelMessageSend(m.ChannelID, response)
	}
}

func generateResponse(input string) (string, error) {
	req := openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: input,
			},
		},
	}

	resp, err := openaiClient.CreateChatCompletion(context.Background(), req)
	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}