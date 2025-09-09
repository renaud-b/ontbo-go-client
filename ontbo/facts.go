package ontbo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// ListFacts → GET /profiles/{profile_id}/facts
func (c *HttpClient) ListFacts(profileID string, fields []string, skip, max int) ([]Fact, error) {
	u := fmt.Sprintf("/profiles/%s/facts", profileID)
	params := url.Values{}
	if len(fields) > 0 {
		for _, f := range fields {
			params.Add("fields", f)
		}
	}
	if skip > 0 {
		params.Add("skip_items", fmt.Sprintf("%d", skip))
	}
	if max > 0 {
		params.Add("max_items", fmt.Sprintf("%d", max))
	}
	if len(params) > 0 {
		u += "?" + params.Encode()
	}

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
		return nil, fmt.Errorf("list facts failed, status=%d, body=%s", res.StatusCode, res.Body)
	}

	var data []Fact
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}

// AddFact → POST /profiles/{profile_id}/facts?feedback=...&source_id=...
func (c *HttpClient) AddFact(profileID, feedback, sourceID string) error {
	u := fmt.Sprintf("/profiles/%s/facts?feedback=%s", profileID, url.QueryEscape(feedback))
	if sourceID != "" {
		u += "&source_id=" + url.QueryEscape(sourceID)
	}
	req, err := c.newRequest(http.MethodPost, u)
	if err != nil {
		return err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("add fact failed, status=%d, and read body failed: %w", res.StatusCode, err)
		}
		return fmt.Errorf("add fact failed, status=%d, body=%s", res.StatusCode, body)
	}

	return nil
}

// GetFact → GET /profiles/{profile_id}/facts/{fact_id}
func (c *HttpClient) GetFact(profileID, factID string) (*Fact, error) {
	u := fmt.Sprintf("/profiles/%s/facts/%s", profileID, factID)
	req, err := c.newRequest(http.MethodGet, u)
	if err != nil {
		return nil, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var fact Fact
	if err := json.NewDecoder(res.Body).Decode(&fact); err != nil {
		return nil, err
	}
	return &fact, nil
}

// DeleteFact → DELETE /profiles/{profile_id}/facts/{fact_id}
func (c *HttpClient) DeleteFact(profileID, factID string) error {
	u := fmt.Sprintf("/profiles/%s/facts/%s", profileID, factID)
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

// QueryFacts → GET /profiles/{profile_id}/facts/query
func (c *HttpClient) QueryFacts(profileID, query, queryType string) (*QueryResponse, error) {
	if queryType == "" {
		queryType = "FULL_DATA"
	}
	u := fmt.Sprintf("/profiles/%s/facts/query?query=%s&query_type=%s",
		profileID, url.QueryEscape(query), url.QueryEscape(queryType))

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
		return nil, fmt.Errorf("query facts failed, status=%d, body=%s", res.StatusCode, res.Body)
	}

	var result QueryResponse
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
