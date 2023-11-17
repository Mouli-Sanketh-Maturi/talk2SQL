package helper

import (
	"context"
	"github.com/sashabaranov/go-openai"
	"os"
	"strings"
)

const promptFile = "config/prompt.txt"

func QueryToSql(query string) (string, error) {
	promptString, err := readPrompt(promptFile)
	if err != nil {
		return "", err
	}

	promptString = strings.Replace(promptString, "{question}", query, -1)

	return queryGptApi(promptString)
}

func readPrompt(file string) (string, error) {
	fileContent, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}
	return string(fileContent), nil
}

func queryGptApi(query string) (string, error) {
	client := openai.NewClient("sk-xxx")
	resp, err := client.CreateChatCompletion(context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4TurboPreview,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: query,
				},
			},
		})
	if err != nil {
		return "", err
	}

	sqlQuery := resp.Choices[0].Message.Content
	sqlQuery = strings.Replace(sqlQuery, "```", "", -1)
	sqlQuery = strings.Replace(sqlQuery, "sql", "", -1)
	sqlQuery = strings.Replace(sqlQuery, "\n", " ", -1)
	sqlQuery = strings.Replace(sqlQuery, "\"", "", -1)
	return sqlQuery, nil
}
