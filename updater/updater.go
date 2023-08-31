package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"regexp"
)

func GetVersion() (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://cu8auck2lc.3z094n2681i06q8k14w31cu4q80d5p.com/71b8acf33b508c7543592acd9d9eb70d/updateApp", nil)
	if err != nil {
		return "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return "", err
	}

	latestVersion, ok := data["latestVersionCode"].(string)
	if !ok {
		return "", fmt.Errorf("latestVersionCode not found in response")
	}

	return latestVersion, nil
}

func Update(version string) error {
	fp := "request.go"
	content, err := os.ReadFile(fp)
	if err != nil {
		return err
	}

	re := regexp.MustCompile(`"appbuildcode",\s+"(\d+)"`)
	replacement := `"appbuildcode", "` + version + `"`
	modifiedContent := re.ReplaceAllString(string(content), replacement)
	fmt.Println(modifiedContent)

	err = os.WriteFile(fp, []byte(modifiedContent), 0644)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	version, err := GetVersion()
	if err != nil {
		fmt.Println("Error fetching latest version:", err)
		return
	}

	err = Update(version)
	if err != nil {
		fmt.Println("Error updating file content:", err)
		return
	}

	fmt.Println("File content updated successfully.")
}
