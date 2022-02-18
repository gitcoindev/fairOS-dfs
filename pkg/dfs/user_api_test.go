package dfs

import (
	"fmt"
	"os"
	"testing"

	"github.com/fairdatasociety/fairOS-dfs/pkg/logging"
	"github.com/sirupsen/logrus"
)

func TestUserIntegration(t *testing.T) {
	username := "user_1"
	max := 10

	api, err := NewMockDfsAPI("mock", "", "", "", logging.New(os.Stdout, logrus.DebugLevel))
	if err != nil {
		t.Fatal(err)
	}

	sessionIds := make([]string, max)
	for i := 0; i < max; i++ {
		_, _, ui, err := api.CreateUser(fmt.Sprintf("%s_%d", username, i), "123456789", "", "")
		if err != nil {
			t.Fatal(err)
		}
		sessionIds[i] = ui.GetSessionId()
	}

	for i := 0; i < max; i++ {
		available := api.IsUserNameAvailable(fmt.Sprintf("%s_%d", username, i))
		if !available {
			t.Fatal("some user is not available in server")
		}
	}

	users, err := api.Users()
	if err != nil {
		t.Fatal(err)
	}
	if len(users) != max {
		t.Fatal("user map count mismatch")
	}

	for i := 0; i < max; i++ {
		err := api.LogoutUser(sessionIds[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	for i := 0; i < max; i++ {
		loggedIn := api.IsUserLoggedIn(fmt.Sprintf("%s_%d", username, i))
		if loggedIn {
			t.Fatal(fmt.Sprintf("%s_%d user is still in loggedIn state", username, i))
		}
	}

	for i := 0; i < max; i++ {
		ui, err := api.LoginUser(fmt.Sprintf("%s_%d", username, i), "123456789", "")
		if err != nil {
			t.Fatal(err)
		}
		sessionIds[i] = ui.GetSessionId()
	}

	for i := 0; i < max; i++ {
		err = api.DeleteUser("123456789", sessionIds[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	users, err = api.Users()
	if err != nil {
		t.Fatal(err)
	}
	if len(users) != 0 {
		t.Fatal("user map should be zero")
	}
}
