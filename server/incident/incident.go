package incident

import "github.com/pkg/errors"

// Incident holds the detailed information of an incident.
type Incident struct {
	Header
	ChannelIDs []string `json:"channel_ids"`
	PostID     string   `json:"post_id"`
}

// Header holds the summary information of an incident.
type Header struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	IsActive        bool   `json:"is_active"`
	CommanderUserID string `json:"commander_user_id"`
	TeamID          string `json:"team_id"`
	CreatedAt       int64  `json:"created_at"`
}

// HeaderFilterOptions specifies the optional parameters when getting headers.
type HeaderFilterOptions struct {
	// Gets all the headers with this TeamID.
	TeamID string
}

// ErrNotFound used to indicate entity not found.
var ErrNotFound = errors.New("not found")

// Service is the incident/service interface.
type Service interface {
	// GetHeaders returns filtered headers.
	GetHeaders(options HeaderFilterOptions) ([]Header, error)

	// CreateIncident Creates a new incident.
	CreateIncident(incident *Incident) (*Incident, error)

	// CreateIncidentDialog Opens an interactive dialog to start a new incident.
	CreateIncidentDialog(commanderID string, triggerID string, postID string) error

	// EndIncident Completes the incident associated to the given channelID.
	EndIncident(channelID string) (*Incident, error)

	// GetIncident Gets an incident by ID.
	GetIncident(id string) (*Incident, error)

	// NukeDB Removes all incident related data.
	NukeDB() error
}

// Store defines the methods the ServiceImpl needs from the interfaceStore.
type Store interface {
	// GetHeaders returns filtered headers.
	GetHeaders(options HeaderFilterOptions) ([]Header, error)

	// CreateIncident Creates a new incident.
	CreateIncident(incident *Incident) (*Incident, error)

	// UpdateIncident updates an incident.
	UpdateIncident(incident *Incident) error

	// GetIncident Gets an incident by ID.
	GetIncident(id string) (*Incident, error)

	// GetIncidentByChannel Gets an incident associated with the given channel id.
	GetIncidentByChannel(channelID string, active bool) (*Incident, error)

	// NukeDB Removes all incident related data.
	NukeDB() error
}