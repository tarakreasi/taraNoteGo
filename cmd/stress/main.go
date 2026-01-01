package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

const (
	baseURL     = "http://localhost:3000"
	concurrency = 436
)

type LoginResponse struct {
	Token string `json:"token"`
}

func main() {
	// 1. Login to get token
	token, err := login("ajarsinau@gmail.com", "password")
	if err != nil {
		log.Fatalf("Failed to login: %v", err)
	}
	fmt.Println("Successfully logged in, starting stress test...")

	// 2. Run stress test
	var wg sync.WaitGroup
	var successCount int32
	var failCount int32

	start := time.Now()

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			if err := createNote(token, id); err != nil {
				atomic.AddInt32(&failCount, 1)
				// fmt.Printf("Request %d failed: %v\n", id, err) // Optional: verbose logging
			} else {
				atomic.AddInt32(&successCount, 1)
			}
		}(i)
	}

	wg.Wait()
	duration := time.Since(start)

	fmt.Printf("\nStress Test Results:\n")
	fmt.Printf("Total Requests: %d\n", concurrency)
	fmt.Printf("Success: %d\n", successCount)
	fmt.Printf("Failed: %d\n", failCount)
	fmt.Printf("Duration: %v\n", duration)
	fmt.Printf("Success Rate: %.2f%%\n", float64(successCount)/float64(concurrency)*100)
}

func login(email, password string) (string, error) {
	payload := map[string]string{
		"email":    email,
		"password": password,
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", baseURL+"/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	// Prevent redirect to capture cookie
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != 302 {
		return "", fmt.Errorf("bad status: %s", resp.Status)
	}

	cookies := resp.Header.Get("Set-Cookie")
	if cookies == "" {
		return "", fmt.Errorf("no cookies received")
	}
	return cookies, nil
}

func createNote(cookie string, id int) error {
	payload := map[string]interface{}{
		"title":   fmt.Sprintf("Stress Note %d", id),
		"content": "<p>Stress test content</p>",
		"status":  "DRAFT",
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", baseURL+"/api/v1/admin/notes", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", cookie)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 && resp.StatusCode != 201 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("status %d: %s", resp.StatusCode, string(bodyBytes))
	}

	return nil
}
