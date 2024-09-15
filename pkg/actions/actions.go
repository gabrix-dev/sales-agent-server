package actions

import (
	"context"

	"notjustadeveloper.com/sales-agent-server/pkg/models"
)

const (
	sendScheduleCallLink string = "send_call_schedule_link"
)

func (am *ActionsManager) sendScheduleCallLink(params map[string]interface{}, metadata models.ActionMetadata) (string, error) {
	am.messageRepo.SendMessage(context.Background(), "Usa este link para reservar una llamada con mi equipo: https://calendly.com/A4rtg3H", "", metadata.ToId, metadata.Provider)
	return "Link sent succesfully", nil
}
