package validations

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	exception "backend_community_grpc/exceptions"
)


// Github Username Validation
func ValidateGitHubUsername(username string) (bool, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s", username)
	resp, err := http.Get(url)
	if err != nil {
		exception.TryCatchError(err)
		return false, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return false, nil
	} else if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("GitHub API request failed with status code: %d", resp.StatusCode)
	}

	// Read The Response Body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		exception.TryCatchError(err)
		return false, err
	}

	// Parse The JSON Response
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		exception.TryCatchError(err)
		return false, err
	}

	// Check If The Username Exists
	if _, ok := data["login"]; ok {
		return true, nil
	}

	return false, nil
}