package logo

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func GetMerker(id string) (string, error) {

	reader, err := baseRequest("GETVARS:_local_=v0,M..1:" + id + "-1")

	if err != nil {
		return "body error", err
	}

	body, err := io.ReadAll(reader)

	defer close(reader)
	if err != nil {
		return "body error", err
	}

	parts := strings.Split(string(body), "'")

	return parts[5][1:2], nil
}

func SetMerker(id string, value string) error {

	_, err := baseRequest("SETVARS:_local_=v0,M..1:" + id + "-1,0" + value)

	return err
}

func baseRequest(body string) (io.ReadCloser, error) {

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(body)))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Cookie", "Security-Hint="+securityHint)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 401 {
		if err := generateSecurityHint(); err != nil {
			return nil, err
		}
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("not 200 but %v", resp.Status)
	}

	return resp.Body, nil
}
