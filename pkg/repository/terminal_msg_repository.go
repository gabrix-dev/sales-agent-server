package repository

import "fmt"

type TerminalMsgRepository struct {
}

func NewTerminalMsgRepository() *TerminalMsgRepository {
	return &TerminalMsgRepository{}
}

func (t *TerminalMsgRepository) DisplayUserMessage(message string) {
	fmt.Println("USER: " + message)
}

func (t *TerminalMsgRepository) DisplayResponse(message string) {
	fmt.Println("ELPATRON: " + message)
}
