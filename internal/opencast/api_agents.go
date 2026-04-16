package opencast

import "fmt"

func (agent_api OpencastAgentAPI) Get(agentID string) (AgentStatus, error) {
	var agentData AgentStatus

	err := agent_api.requester.GetJSON(fmt.Sprintf("/api/agents/%s", agentID), &agentData)
	if err != nil {
		return AgentStatus{}, nil
	}

	return agentData, nil
}
