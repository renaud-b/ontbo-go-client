package ontbo

import (
	"fmt"
	"net/http"
)

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

type HttpClient struct {
	baseURL    string
	httpClient *http.Client
	token      string
}

// NewClient crée un client Ontbo.
func NewClient(baseURL, token string) Client {
	return &HttpClient{
		baseURL:    baseURL,
		httpClient: http.DefaultClient,
		token:      token,
	}
}

func (c *HttpClient) newRequest(method, path string) (*http.Request, error) {
	url := fmt.Sprintf("%s%s", c.baseURL, path)
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.token)
	return req, nil
}
