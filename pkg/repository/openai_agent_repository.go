package repository

import (
	"context"
	"os"
	"time"

	"github.com/sashabaranov/go-openai"
	"notjustadeveloper.com/sales-agent-server/pkg/errors"
	"notjustadeveloper.com/sales-agent-server/pkg/models"
)

const (
	salesAssistantId   = "asst_b350IO8RIzPA5kIaDfKtNGB9"
	gpt4oModelID       = "gpt4o"
	assistantsBaseName = "Twinfluence-"
)

type OpenaiAgentRepository struct {
	client *openai.Client
}

func NewOpenaiAgentRepository() (OpenaiAgentRepository, error) {
	openaiApiKey := os.Getenv("OPENAI_API_KEY")
	if openaiApiKey == "" {
		return OpenaiAgentRepository{}, errors.New(errors.EnvVariableNotFound, errors.InternalError)
	}
	return OpenaiAgentRepository{
		client: openai.NewClient(openaiApiKey),
	}, nil
}

func (a *OpenaiAgentRepository) CreateAgent(ctx context.Context, userId string, agentSettings *models.AgentSettings) (string, error) {
	agentName := assistantsBaseName + userId
	assistant, err := a.client.CreateAssistant(ctx, openai.AssistantRequest{
		Model:        gpt4oModelID,
		Name:         &agentName,
		Instructions: &agentSettings.SystemPrompt,
		//TODO: add schedule meeting tool
	})
	if err != nil {
		return "", errors.NewWrap(errors.OpenaiCreateAssistantError, err)
	}
	return assistant.ID, nil
}

// func (a *OpenaiAgentRepository) UpdateAgent(ctx context.Context, agentId string, agentSettings *models.AgentSettings) (string, error) {
// 	assistant, err := a.client.ModifyAssistant(ctx, ...)
// 	if err != nil {
// 		return "", errors.NewWrap(errors.OpenaiCreateAssistantError, err)
// 	}
// 	return assistant.ID, nil
// }

func (a *OpenaiAgentRepository) AddUserMessage(ctx context.Context, message string, state *models.State) (*models.AgentOutput, error) {
	messageReq := openai.MessageRequest{
		Role:    string(openai.ThreadMessageRoleUser),
		Content: message,
	}
	_, err := a.client.CreateMessage(ctx, state.ChatId, messageReq)
	if err != nil {
		return nil, errors.NewWrap(errors.OpenaiCreateMessageError, err)
	}
	run, err := a.client.CreateRun(ctx, state.ChatId, openai.RunRequest{AssistantID: salesAssistantId})
	if err != nil {
		return nil, errors.NewWrap(errors.OpenaiCreateRunError, err)
	}
	return a.getRunOutput(ctx, &run, state.ChatId, state.UserId, state.Provider)
}

func (a *OpenaiAgentRepository) SubmitActionOutput(ctx context.Context, actionOutput string, metadata models.ActionMetadata) (*models.AgentOutput, error) {
	toolOutput := openai.ToolOutput{
		ToolCallID: metadata.OpenaiToolCallId,
		Output:     actionOutput,
	}
	submitToolOutputRequest := openai.SubmitToolOutputsRequest{ToolOutputs: []openai.ToolOutput{toolOutput}}
	run, err := a.client.SubmitToolOutputs(ctx, metadata.ChatId, metadata.OpenaiRunId, submitToolOutputRequest)
	if err != nil {
		return nil, errors.NewWrap(errors.OpenaiSubmitToolOutputError, err)
	}
	return a.getRunOutput(ctx, &run, metadata.ChatId, metadata.ToId, metadata.Provider)

}

func (a *OpenaiAgentRepository) getRunOutput(ctx context.Context, run *openai.Run, chatId string, toId string, provider models.MessagingProvider) (*models.AgentOutput, error) {
	for run.Status == openai.RunStatusQueued || run.Status == openai.RunStatusInProgress {
		time.Sleep(time.Millisecond * 250)
		runCp, err := a.client.RetrieveRun(ctx, chatId, run.ID)
		if err != nil {
			return nil, errors.NewWrap(errors.OpenaiRetrieveRunError, err)
		}
		*run = runCp
	}
	if run.Status == openai.RunStatusRequiresAction {
		actionReq, err := models.NewActionRequestFromOpenai(run.RequiredAction.SubmitToolOutputs.ToolCalls[0], run.ID, chatId, toId, provider)
		if err != nil {
			return nil, err
		}
		return &models.AgentOutput{ActionRequest: actionReq}, nil
	}
	if run.Status != openai.RunStatusCompleted {
		return nil, errors.New(errors.OpenaiUnexpectedRunStatus, errors.InternalError)
	}
	messageList, err := a.client.ListMessage(ctx, chatId, a.getLastMessageLimit(), nil, nil, nil)
	if err != nil {
		return nil, errors.NewWrap(errors.OpenaiListMessagesError, err)
	}
	return &models.AgentOutput{Message: messageList.Messages[0].Content[0].Text.Value}, nil
}

func (a *OpenaiAgentRepository) getLastMessageLimit() *int {
	limit := new(int)
	*limit = 1
	return limit
}

func (a *OpenaiAgentRepository) CreateChat(ctx context.Context) (string, error) {
	response, err := a.client.CreateThread(ctx, openai.ThreadRequest{})
	if err != nil {
		return "", errors.NewWrapWithType(err, errors.InternalError, errors.OpenaiCreateThreadError)
	}
	return response.ID, nil
}
