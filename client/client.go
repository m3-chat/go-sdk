package client

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/m3-chat/go-sdk/models"
	"github.com/m3-chat/go-sdk/types"
)

type M3ChatClient struct {
	stream bool
}

func NewClient(opts *types.ClientOptions) *M3ChatClient {
	stream := false
	if opts != nil {
		stream = opts.Stream
	}
	return &M3ChatClient{stream: stream}
}

func (c *M3ChatClient) GetResponse(params types.RequestParams) error {
	validModels := models.GetAvailableModels()
	isValid := false
	for _, m := range validModels {
		if m == params.Model {
			isValid = true
			break
		}
	}

	if !isValid {
		fmt.Printf("%s is not an available model.\nSupported models: %v\n", params.Model, validModels)
		os.Exit(1)
	}

	endpoint, _ := url.Parse("https://m3-chat.vercel.app/api/gen")
	query := endpoint.Query()
	query.Set("model", params.Model)
	query.Set("content", params.Content)
	endpoint.RawQuery = query.Encode()

	req, err := http.NewRequest("GET", endpoint.String(), nil)
	if err != nil {
		return err
	}
	if c.stream {
		req.Header.Set("Accept", "text/event-stream")
	} else {
		req.Header.Set("Accept", "application/json")
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("Request failed: %d %s", res.StatusCode, res.Status)
	}

	if c.stream {
		reader := bufio.NewReader(res.Body)
		for {
			line, err := reader.ReadBytes('\n')
			if err == io.EOF {
				break
			}
			if err != nil {
				return err
			}
			fmt.Print(string(line))
		}
	} else {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		fmt.Println(string(body))
	}

	return nil
}

func (c *M3ChatClient) BatchRequests(messages []string, opts *types.BatchRequestOptions) ([]string, error) {
	if c.stream {
		return nil, errors.New("Streaming batch requests not implemented")
	}

	results := []string{}

	for _, msg := range messages {
		endpoint, _ := url.Parse("https://m3-chat.vercel.app/api/gen")
		query := endpoint.Query()
		query.Set("content", msg)
		if opts != nil && opts.Model != "" {
			query.Set("model", opts.Model)
		}
		endpoint.RawQuery = query.Encode()

		res, err := http.Get(endpoint.String())
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("Request failed with status %d", res.StatusCode)
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		results = append(results, string(body))
	}

	return results, nil
}
