package models

import (
	"encoding/json"
	"fmt"

	myerrors "notjustadeveloper.com/sales-agent-server/pkg/errors"

	"github.com/sashabaranov/go-openai"
)

type AgentOutput struct {
	Message       string
	ActionRequest *ActionRequest
}

type ActionRequest struct {
	ActionId         string
	ActionParameters map[string]interface{}
	Metadata         ActionMetadata
}

type ActionMetadata struct {
	OpenaiRunId      string
	OpenaiToolCallId string
	ChatId           string
	ToId             string
	Provider         MessagingProvider
}

func NewActionRequestFromOpenai(toolCall openai.ToolCall, runId string, chatId string, toId string, provider MessagingProvider) (*ActionRequest, error) {
	var actionParams map[string]interface{}
	err := json.Unmarshal([]byte(toolCall.Function.Arguments), &actionParams)
	if err != nil {
		return nil, myerrors.NewWrap(myerrors.UnmarshalJSONError, err)
	}
	return &ActionRequest{
		ActionId:         toolCall.Function.Name,
		ActionParameters: actionParams,
		Metadata: ActionMetadata{
			OpenaiRunId:      runId,
			ChatId:           chatId,
			OpenaiToolCallId: toolCall.ID,
			ToId:             toId,
			Provider:         provider,
		},
	}, nil
}

type AgentEngine string

const OpenaiAgentEngine AgentEngine = "openai"

type AgentSettings struct {
	AgentEngine    AgentEngine
	AnswerExamples []AnswerExample
	SystemPrompt   string
}

type AnswerExample struct {
	userMessage string
	agentAnswer string
}

func (q AnswerExample) ToString(agentName string) string {
	return fmt.Sprintf("USER: %s\n%s:%s", q.userMessage, agentName, q.agentAnswer)
}
