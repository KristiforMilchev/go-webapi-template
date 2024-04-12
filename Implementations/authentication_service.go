package implementations

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/google/uuid"

	"leanmeal/api/dtos"
)

type AuthenticationService struct {
	AuthRequests map[uuid.UUID]dtos.StoredAuthRequest
}

func (authService *AuthenticationService) GetMessage(email string) dtos.InitAuthReponse {

	if authService.AuthRequests == nil {
		authService.AuthRequests = make(map[uuid.UUID]dtos.StoredAuthRequest)
	}

	id := uuid.New()
	uuid := uuid.New()
	code := generateRandomString(32)

	authResponse := dtos.StoredAuthRequest{
		Id:   id,
		Name: email,
		Code: code,
		Uuid: uuid.String(),
		Time: time.Time{}.Add(time.Duration(time.Second * 30)),
	}

	authService.AuthRequests[id] = authResponse

	fmt.Println(&authService.AuthRequests)

	response := dtos.InitAuthReponse{
		Code: authResponse.Code,
		Uuid: authResponse.Uuid,
	}

	return response
}

func (authService *AuthenticationService) VerifySignature(response dtos.FinishAuthResponse) bool {

	return true
}

func (authService *AuthenticationService) Start() {
	// Create a channel to receive signals
	signal := make(chan struct{})

	// Start a goroutine to send signals at regular intervals
	go func() {
		for {
			time.Sleep(30 * time.Second) // Wait for 30 seconds
			signal <- struct{}{}         // Send a signal to the channel
		}
	}()

	for {
		select {
		case <-signal:
			fmt.Println("Something has happened at", time.Now())
			for i, d := range authService.AuthRequests {
				fmt.Println(i)
				fmt.Println("Checking if request has expired")
				expiresAt := d.Time.UTC()
				fmt.Fprintln(os.Stdout, "Expires at ", expiresAt, "and the time now is", time.Now().UTC())
				expired := d.Time.UTC().Before(time.Now().UTC())
				fmt.Println(expired)
				if expired {
					delete(authService.AuthRequests, i)
					fmt.Println("Timer has expired")
					fmt.Println(d.Uuid)
					fmt.Println(d.Time)
				}
			}
			fmt.Println(authService.AuthRequests)
		}
	}
}

// generateRandomString generates a random string of specified length
func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var result string
	for i := 0; i < length; i++ {
		randomIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		result += string(charset[randomIndex.Int64()])
	}
	return result
}
