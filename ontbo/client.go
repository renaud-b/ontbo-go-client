package ontbo

import (
	"fmt"
	"net/http"
)

// Client is the interface for interacting with the Ontbo service
type Client interface {
	BuildContext(profileID, query string) (*QueryResponse, error)
	ListFacts(profileID string, fields []string, skip, max int) ([]Fact, error)
	AddFact(profileID, feedback, sourceID string) error
	GetFact(profileID, factID string) (*Fact, error)
	DeleteFact(profileID, factID string) error
	QueryFacts(profileID, query, queryType string) (*QueryResponse, error)
	ListProfiles() ([]Profile, error)
	CreateProfile(requestedID string) (*Profile, error)
	DeleteProfile(profileID string) error
	RunProfileUpdate(profileID string) (*Profile, error)
	StopProfileUpdate(profileID string) (*Profile, error)
	GetProfileUpdateStatus(profileID string) (*UpdateStatus, error)
	ListScenes(profileID string) ([]string, error)
	CreateScene(profileID, requestedID string) (*ResponseWithID, error)
	DeleteScene(profileID, sceneID string) error
	GetTextFromScene(profileID, sceneID string) ([]SceneMessage, error)
	AddTextToScene(profileID, sceneID string, messages []SceneMessage, updateNow, waitForResult bool) (*ResponseWithID, error)
	QueryScenes(profileID, query string) ([]string, error)
}

// HttpClient is an implementation of the Client interface using HTTP
type HttpClient struct {
	baseURL    string
	httpClient *http.Client
	token      string
}

// NewClient create a new Ontbo client
func NewClient(baseURL, token string) Client {
	return &HttpClient{
		baseURL:    baseURL,
		httpClient: http.DefaultClient,
		token:      token,
	}
}

// newRequest creates a new HTTP request with the appropriate headers
func (c *HttpClient) newRequest(method, path string) (*http.Request, error) {
	url := fmt.Sprintf("%s%s", c.baseURL, path)
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.token)
	return req, nil
}
