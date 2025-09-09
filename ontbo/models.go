package ontbo

import "time"

// Profile represents a user profile in the Ontbo system
type Profile struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// Fact represents a fact in the Ontbo system
type Fact struct {
	ID        string `json:"id,omitempty"`
	Data      string `json:"data,omitempty"`
	Source    string `json:"source,omitempty"`
	Timestamp int64  `json:"timestamp,omitempty"`
}

// SceneMessage represents a message in a scene conversation
type SceneMessage struct {
	Role      string  `json:"role"`      // "user" or "assistant"
	Content   string  `json:"content"`   // content of the message
	Timestamp float64 `json:"timestamp"` // epoch time in seconds
}

// UpdateStatus represent the status of an update operation
type UpdateStatus struct {
	Status   string  `json:"status"`   // "WORKING" ou "IDLE"
	Progress float64 `json:"progress"` // 0.0 → 100.0
}

// ContextResponse used in AddTextToScene response
type ContextResponse struct {
	Context string `json:"context"`
}

const (
	// FullData performs an extensive search on all profile facts.
	FullData = "FULL_DATA"
	// MultiHop makes multiple sub-queries to answer the query.
	MultiHop = "MULTI_HOP"
	// SingleHop simple call to the facts database.
	SingleHop = "SINGLE_HOP"
	// VectorSearch use embeddings to get the answer to the query.
	VectorSearch = "VECTOR_SEARCH"
)

// QueryResponse represent a query response from /scenes/query
type QueryResponse struct {
	Result string `json:"result"`
}

// ResponseWithID is a generic response containing an ID
type ResponseWithID struct {
	ID string `json:"id"`
}
