package issues

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/cli/go-gh"
	"github.com/cli/go-gh/pkg/api"
)

type Author struct {
	Login string `json: "login"`
}

type Issue struct {
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	State     string    `json:"state"`
	Author    Author    `json:"author,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

func CreateIssue(issue *Issue, user, repo string) error {
	client, err := gh.RESTClient(&api.ClientOptions{
		Headers: map[string]string{"Content-Type": "application/json"},
	})
	if err != nil {
		return err
	}

	issueJson, err := json.Marshal(issue)
	if err != nil {
		return err
	}
	reqBody := bytes.NewBuffer(issueJson)
	return client.Post(
		fmt.Sprintf("repos/%s/%s/issues", user, repo), reqBody, nil)
}
