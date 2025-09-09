package ontbo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// BuildContext → GET /profiles/{profile_id}/context?query=...
func (c *HttpClient) BuildContext(profileID, query string) (*QueryResponse, error) {
	u := fmt.Sprintf("/profiles/%s/context?query=%s",
		profileID, url.QueryEscape(query))

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
		return nil, fmt.Errorf("build context failed, status=%d", res.StatusCode)
	}

	var ctxResp QueryResponse
	if err := json.NewDecoder(res.Body).Decode(&ctxResp); err != nil {
		return nil, err
	}
	return &ctxResp, nil
}
