package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func (bot *Bot) get(endpoint string, params any) (*APIResponse, error) {
	u, err := url.Parse(bot.baseURL + "/" + endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to parse url: %w", err)
	}

	q, err := structToQuery(params)
	if err != nil {
		return nil, fmt.Errorf("failed to parse params: %w", err)
	}
	u.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	res, err := bot.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer res.Body.Close()

	var apiRes APIResponse
	if err = json.NewDecoder(res.Body).Decode(&apiRes); err != nil {
		return nil, fmt.Errorf("failed to decode json response: %w", err)
	}

	if !apiRes.OK {
		return nil, fmt.Errorf("telegram BOT API error: %s (code %d)", apiRes.Description, apiRes.ErrorCode)
	}

	return &apiRes, nil
}

func (bot *Bot) post(endpoint string, payload any) (*APIResponse, error) {
	u, err := url.Parse(bot.baseURL + "/" + endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to parse url: %w", err)
	}

	jsonBody, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := http.NewRequest("POST", u.String(), bytes.NewReader(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := bot.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer res.Body.Close()

	var apiRes APIResponse
	if err = json.NewDecoder(res.Body).Decode(&apiRes); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if !apiRes.OK {
		return nil, fmt.Errorf("telegram BOT API error: %s (code %d)", apiRes.Description, apiRes.ErrorCode)
	}

	return &apiRes, nil
}
