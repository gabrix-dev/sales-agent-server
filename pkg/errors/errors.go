package errors

// Type used to add context to errors without using reflection
type ErrorId string

// Type used to infer the http status code to return and handle
// specific errors without using errors.Is (expensive)
type ErrorType int

const (
	UnmarshalJSONError ErrorId = "UnmarshalJSONError"
	BytesReadingError  ErrorId = "BytesReadingError"

	EnumTypeConversionError ErrorId = "EnumTypeConversionError"

	EnvVariableNotFound     ErrorId = "EnvVariableNotFound"
	RepoInitializationError ErrorId = "RepoInitializationError"

	InvalidRequestData ErrorId = "InvalidRequestData"

	ParseMessagingProviderError ErrorId = "ParseMessagingProviderError"
	AddMessageError             ErrorId = "AddMessageErrro"

	GetStateError    ErrorId = "GetStateError"
	CreateStateError ErrorId = "CreateStateError"

	AgentAddMessageError         ErrorId = "AgentAddMessageError"
	AgentCreateChatError         ErrorId = "AgentCreateChatError"
	AgentSubmitActionOutputError ErrorId = "AgentSubmitActionOutputError"
	AgentCreateAgentError        ErrorId = "AgentCreateAgentError"

	ActionNotFound ErrorId = "ActionNotFound"
	RunActionError ErrorId = "RunActionError"

	StateNotFound ErrorId = "StateNotFound"

	SendMessageError ErrorId = "SendMessageError"

	OpenaiCreateThreadError     ErrorId = "OpenaiCreateThreadError"
	OpenaiCreateMessageError    ErrorId = "OpenaiCreateMessageError"
	OpenaiCreateRunError        ErrorId = "OpenaiCreateRunError"
	OpenaiRunExecutionError     ErrorId = "OpenaiRunExecutionError"
	OpenaiRetrieveRunError      ErrorId = "OpenaiRetrieveRunError"
	OpenaiListMessagesError     ErrorId = "OpenaiListMessagesError"
	OpenaiUnexpectedRunStatus   ErrorId = "OpenaiUnexpectedRunStatus"
	OpenaiSubmitToolOutputError ErrorId = "OpenaiSubmitToolOutputError"
	OpenaiCreateAssistantError  ErrorId = "OpenaiCreateAssistantError"

	ModelEngineNotFound ErrorId = "ModelEngineNotFound"
)

const (
	InternalError ErrorType = iota
	NotFoundError
	BadRequest
	SwitchDefaultCase
)
