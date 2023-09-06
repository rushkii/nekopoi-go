package nekopoi

import (
	"io"
	"net/http"
)

const API_URL = "https://cu8auck2lc.3z094n2681i06q8k14w31cu4q80d5p.com/330cceade91a6a9cd30fb8042222ed56/71b8acf33b508c7543592acd9d9eb70d"

func Request(endpoint string) (io.ReadCloser, error) {
	req, err := http.NewRequest("GET", API_URL+endpoint, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("appbuildcode", "25032")
	req.Header.Set("appsignature", "pOplm8IDEDGXN55IaYohQ8CzJFvWsfXyhGvwPRD9kWgzYSRuuvAOPfsE0AJbHVbAJyWGsGCNUIuQLJ7HbMbuFLMWwDgHNwxOrYMH")
	req.Header.Set("token", "XbGSFkQsJYbFC6pcUMCFL4oNHULvHU7WdDAXYgpmqYlh7p5ZCQ4QZ13GDgowiOGvAejz9X5H6DYvEQBMrc3A17SO3qwLwVkbn6YY")
	req.Header.Set("user-agent", "okhttp/4.9.0")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return res.Body, nil
}
