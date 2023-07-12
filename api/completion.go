package api

import (
	"sparktogpt/types"
	"sparktogpt/utils"
)

func NewCompletionsResponse(data string, tokens int) types.CompletionsResponse {
	reTokens := utils.GetToken(data)
	return types.CompletionsResponse{
		ID:      "chatcmpl-" + utils.Make32ID(),
		Object:  "chat.completion",
		Created: 0,
		Choices: []types.CompletionsChoice{
			{
				Index: 0,
				Message: types.CompletionsMessage{
					Role:    "assistant",
					Content: data,
				},
				FinishReason: "stop",
			},
		},
		Usage: types.CompletionsUsage{
			CompletionTokens: reTokens,
			PromptTokens:     tokens,
			TotalTokens:      reTokens + tokens,
		},
	}
}
