package logo

import (
	"bytes"
	"fmt"
	"hash/crc32"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"
)

func generateSecurityHint() error {

	var (
		a1 = rand.Uint32()
		a2 = rand.Uint32()
		b1 = rand.Uint32()
		b2 = rand.Uint32()
	)

	var (
		a1Str = strconv.FormatUint(uint64(a1), 10)
		a2Str = strconv.FormatUint(uint64(a2), 10)
		b1Str = strconv.FormatUint(uint64(b1), 10)
		b2Str = strconv.FormatUint(uint64(b2), 10)
	)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte("UAMCHAL:3,4,"+a1Str+","+a2Str+","+b1Str+","+b2Str)))
	if err != nil {
		return err
	}

	req.Header.Set("Cookie", "Security-Hint=p")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(resp.Body)

	defer close(resp.Body)

	if err != nil {
		return err
	}
	parts := strings.Split(string(body), ",")

	if len(parts) != 3 {
		return fmt.Errorf("Init Response did not return 3 parts: %v", parts)
	}

	if parts[0] != "700" {
		return fmt.Errorf("Init Response status not 700 but: %v", parts[0])
	}

	loginSecurityHint := parts[1]
	serverChallenge := parts[2]

	pwToken := truncate(password+"+"+serverChallenge, 32)

	serverChallengeInt, err := strconv.ParseUint(serverChallenge, 10, 32)

	pwTokenCRC := crc32.ChecksumIEEE([]byte(pwToken))

	loginPwToken := pwTokenCRC ^ uint32(serverChallengeInt)

	loginServerChallenge := a1 ^ a2
	loginServerChallenge = loginServerChallenge ^ b1
	loginServerChallenge = loginServerChallenge ^ b2
	loginServerChallenge = loginServerChallenge ^ uint32(serverChallengeInt)

	bodyStr := "UAMLOGIN:Web User," + strconv.FormatUint(uint64(loginPwToken), 10) + "," + strconv.FormatUint(uint64(loginServerChallenge), 10)

	req, err = http.NewRequest("POST", url, bytes.NewBuffer([]byte(bodyStr)))

	if err != nil {
		return err
	}

	req.Header.Set("Cookie", "Security-Header="+loginSecurityHint)

	resp, err = client.Do(req)
	if err != nil {
		return err
	}

	body, err = io.ReadAll(resp.Body)

	defer close(resp.Body)
	if err != nil {
		return err
	}

	parts = strings.Split(string(body), ",")

	if len(parts) != 2 {
		return fmt.Errorf("Login Response did not return 3 parts: %v", parts)
	}

	if parts[0] != "700" {
		return fmt.Errorf("Login Response status not 700 but: %v", parts[0])
	}

	securityHint = parts[1]

	return nil
}

func close(Body io.ReadCloser) {
	err := Body.Close()
	if err != nil {
		fmt.Print("Error while closing the body",
			err.Error())
	}
}

func truncate(s string, n int) string {
	if n >= len(s) {
		return s
	}
	for i := n; i >= n-3 && i >= 0; i-- {
		if utf8.RuneStart(s[i]) {
			return s[:i]
		}
	}
	return s[:n]
}
