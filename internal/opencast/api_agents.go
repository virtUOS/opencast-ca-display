package opencast

import (
	"fmt"
)

func (agents) Get(agentID string) (AgentStatus, error) {
	var agentData AgentStatus

	err := opencastClient.GetJSON(fmt.Sprintf("/api/agents/%s", agentID), &agentData)
	if err != nil {
		return AgentStatus{}, nil
	}

	return agentData, nil
}
