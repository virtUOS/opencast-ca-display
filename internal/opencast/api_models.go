package opencast

// Model for /api/agents/{agentID}
type AgentStatus struct {
	Status  string   `json:"status"`
	AgentID string   `json:"agent_id"`
	Update  string   `json:"update"`
	URL     string   `json:"url"`
	Inputs  []string `json:"inputs"`
}

type Event struct {
	ArchiveVersion    int             `json:"archive_version"`
	Created           string          `json:"created"`
	Creator           string          `json:"creator"`
	Contributor       []string        `json:"contributor"`
	Description       string          `json:"description"`
	HasPreviews       bool            `json:"has_previews"`
	EventID           string          `json:"identifier"`
	Location          string          `json:"location"`
	Presenter         []string        `json:"presenter"`
	Language          string          `json:"language"`
	RightsHolder      string          `json:"rightsholder"`
	License           string          `json:"license"`
	SeriesID          string          `json:"is_part_of"`
	SeriesTitle       string          `json:"series"`
	Source            string          `json:"source"`
	Status            string          `json:"status"`
	PublicationStatus []string        `json:"publication_status"`
	ProcessingState   string          `json:"processing_state"`
	Start             string          `json:"start"`
	Duration          int             `json:"duration"`
	Subjects          []string        `json:"subjects"`
	Title             string          `json:"title"`
	Scheduling        EventScheduling `json:"scheduling"`
	Acls              []EventACL      `json:"acl"`
}

type EventScheduling struct {
	Start   string   `json:"start"`
	End     string   `json:"end"`
	AgentID string   `json:"agent_id"`
	Inputs  []string `json:"inputs"`
}

type EventACL struct {
	Allow  bool   `json:"allow"`
	Action string `json:"action"`
	Role   string `json:"role"`
}
