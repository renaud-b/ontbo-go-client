package ontbo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// ListScenes → GET /profiles/{profile_id}/scenes
func (c *HttpClient) ListScenes(profileID string) ([]string, error) {
	u := fmt.Sprintf("/profiles/%s/scenes", profileID)
	req, err := c.newRequest(http.MethodGet, u)
	if err != nil {
		return nil, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("list scenes failed, status=%d", res.StatusCode)
	}

	var data []string
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}

// CreateScene → POST /profiles/{profile_id}/scenes?requested_id=...
func (c *HttpClient) CreateScene(profileID, requestedID string) (*ResponseWithID, error) {
	u := fmt.Sprintf("/profiles/%s/scenes", profileID)
	if requestedID != "" {
		u += "?requested_id=" + url.QueryEscape(requestedID)
	}

	req, err := c.newRequest(http.MethodPost, u)
	if err != nil {
		return nil, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("create scene failed, status=%d", res.StatusCode)
	}

	var response ResponseWithID
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}
	return &response, nil
}

// DeleteScene → DELETE /profiles/{profile_id}/scenes/{scene_id}
func (c *HttpClient) DeleteScene(profileID, sceneID string) error {
	u := fmt.Sprintf("/profiles/%s/scenes/%s", profileID, sceneID)
	req, err := c.newRequest(http.MethodDelete, u)
	if err != nil {
		return err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("delete failed, status=%d", res.StatusCode)
	}
	return nil
}

// GetTextFromScene → GET /profiles/{profile_id}/scenes/{scene_id}/text
func (c *HttpClient) GetTextFromScene(profileID, sceneID string) ([]SceneMessage, error) {
	u := fmt.Sprintf("/profiles/%s/scenes/%s/text", profileID, sceneID)
	req, err := c.newRequest(http.MethodGet, u)
	if err != nil {
		return nil, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("get text from scene failed, status=%d", res.StatusCode)
	}

	var messages []SceneMessage
	if err := json.NewDecoder(res.Body).Decode(&messages); err != nil {
		return nil, err
	}
	return messages, nil
}

// AddTextToScene → POST /profiles/{profile_id}/scenes/{scene_id}/text
// Envoie une liste de messages au format attendu
func (c *HttpClient) AddTextToScene(profileID, sceneID string, messages []SceneMessage, updateNow, waitForResult bool) (*ResponseWithID, error) {
	u := fmt.Sprintf("/profiles/%s/scenes/%s/text", profileID, sceneID)

	params := url.Values{}
	if updateNow {
		params.Add("update_now", "true")
	}
	if waitForResult {
		params.Add("wait_for_result", "true")
	}
	if len(params) > 0 {
		u += "?" + params.Encode()
	}

	// encode en JSON
	body, err := json.Marshal(messages)
	if err != nil {
		return nil, err
	}

	req, err := c.newRequest(http.MethodPost, u)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Body = io.NopCloser(strings.NewReader(string(body)))

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("add text to scene failed, status=%d", res.StatusCode)
	}

	var response ResponseWithID
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}
	return &response, nil
}

// QueryScenes → GET /profiles/{profile_id}/scenes/query?query=...
func (c *HttpClient) QueryScenes(profileID, query string) ([]string, error) {
	u := fmt.Sprintf("/profiles/%s/scenes/query?query=%s", profileID, url.QueryEscape(query))

	req, err := c.newRequest(http.MethodGet, u)
	if err != nil {
		return nil, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("query scenes failed, status=%d", res.StatusCode)
	}

	var result []string
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}
