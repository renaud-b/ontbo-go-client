package ontbo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// ListProfiles → GET /profiles
func (c *HttpClient) ListProfiles() ([]Profile, error) {
	req, err := c.newRequest(http.MethodGet, "/profiles")
	if err != nil {
		return nil, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var data []Profile
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}

// CreateProfile → POST /profiles?requested_id=xxx
func (c *HttpClient) CreateProfile(requestedID string) (*Profile, error) {
	u := "/profiles"
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

	var profile Profile
	if err := json.NewDecoder(res.Body).Decode(&profile); err != nil {
		return nil, err
	}
	return &profile, nil
}

// DeleteProfile → DELETE /profiles/{profile_id}
func (c *HttpClient) DeleteProfile(profileID string) error {
	req, err := c.newRequest(http.MethodDelete, "/profiles/"+profileID)
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

// RunProfileUpdate → PUT /profiles/{profile_id}/update/run
func (c *HttpClient) RunProfileUpdate(profileID string) (*Profile, error) {
	u := fmt.Sprintf("/profiles/%s/update/run", profileID)
	req, err := c.newRequest(http.MethodPut, u)
	if err != nil {
		return nil, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var profile Profile
	if err := json.NewDecoder(res.Body).Decode(&profile); err != nil {
		return nil, err
	}
	return &profile, nil
}

// StopProfileUpdate → PUT /profiles/{profile_id}/update/stop
func (c *HttpClient) StopProfileUpdate(profileID string) (*Profile, error) {
	u := fmt.Sprintf("/profiles/%s/update/stop", profileID)
	req, err := c.newRequest(http.MethodPut, u)
	if err != nil {
		return nil, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var profile Profile
	if err := json.NewDecoder(res.Body).Decode(&profile); err != nil {
		return nil, err
	}
	return &profile, nil
}

// GetProfileUpdateStatus → GET /profiles/{profile_id}/update/status
func (c *HttpClient) GetProfileUpdateStatus(profileID string) (*UpdateStatus, error) {
	u := fmt.Sprintf("/profiles/%s/update/status", profileID)
	req, err := c.newRequest(http.MethodGet, u)
	if err != nil {
		return nil, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var status UpdateStatus
	if err := json.NewDecoder(res.Body).Decode(&status); err != nil {
		return nil, err
	}
	return &status, nil
}
