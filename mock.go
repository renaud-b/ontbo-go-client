package ontbo

import (
	"github.com/stretchr/testify/mock"
	"testing"
)

type Mock struct {
	mock.Mock
}

func NewMock(t *testing.T) *Mock {
	m := &Mock{}
	m.Test(t)
	return m
}

func (m *Mock) BuildContext(profileID, query string) (*QueryResponse, error) {
	args := m.Called(profileID, query)
	return args.Get(0).(*QueryResponse), args.Error(1)
}

func (m *Mock) ListFacts(profileID string, fields []string, skip, max int) ([]Fact, error) {
	args := m.Called(profileID, fields, skip, max)
	return args.Get(0).([]Fact), args.Error(1)
}

func (m *Mock) AddFact(profileID, feedback, sourceID string) error {
	args := m.Called(profileID, feedback, sourceID)
	return args.Error(0)
}

func (m *Mock) GetFact(profileID, factID string) (*Fact, error) {
	args := m.Called(profileID, factID)
	return args.Get(0).(*Fact), args.Error(1)
}

func (m *Mock) DeleteFact(profileID, factID string) error {
	args := m.Called(profileID, factID)
	return args.Error(0)
}

func (m *Mock) QueryFacts(profileID, query, queryType string) (*QueryResponse, error) {
	args := m.Called(profileID, query, queryType)
	return args.Get(0).(*QueryResponse), args.Error(1)
}

func (m *Mock) ListProfiles() ([]Profile, error) {
	args := m.Called()
	return args.Get(0).([]Profile), args.Error(1)
}

func (m *Mock) CreateProfile(requestedID string) (*Profile, error) {
	args := m.Called(requestedID)
	return args.Get(0).(*Profile), args.Error(1)
}

func (m *Mock) DeleteProfile(profileID string) error {
	args := m.Called(profileID)
	return args.Error(0)
}

func (m *Mock) RunProfileUpdate(profileID string) (*Profile, error) {
	args := m.Called(profileID)
	return args.Get(0).(*Profile), args.Error(1)
}

func (m *Mock) StopProfileUpdate(profileID string) (*Profile, error) {
	args := m.Called(profileID)
	return args.Get(0).(*Profile), args.Error(1)
}

func (m *Mock) GetProfileUpdateStatus(profileID string) (*UpdateStatus, error) {
	args := m.Called(profileID)
	return args.Get(0).(*UpdateStatus), args.Error(1)
}

func (m *Mock) ListScenes(profileID string) ([]string, error) {
	args := m.Called(profileID)
	return args.Get(0).([]string), args.Error(1)
}

func (m *Mock) CreateScene(profileID, requestedID string) (*ResponseWithID, error) {
	args := m.Called(profileID, requestedID)
	return args.Get(0).(*ResponseWithID), args.Error(1)
}

func (m *Mock) DeleteScene(profileID, sceneID string) error {
	args := m.Called(profileID, sceneID)
	return args.Error(0)
}

func (m *Mock) GetTextFromScene(profileID, sceneID string) ([]SceneMessage, error) {
	args := m.Called(profileID, sceneID)
	return args.Get(0).([]SceneMessage), args.Error(1)
}

func (m *Mock) AddTextToScene(profileID, sceneID string, messages []SceneMessage, updateNow, waitForResult bool) (*ResponseWithID, error) {
	args := m.Called(profileID, sceneID, messages, updateNow, waitForResult)
	return args.Get(0).(*ResponseWithID), args.Error(1)
}

func (m *Mock) QueryScenes(profileID, query string) ([]string, error) {
	args := m.Called(profileID, query)
	return args.Get(0).([]string), args.Error(1)
}
