package dfs

import (
	"fmt"
	"os"
	"testing"

	"github.com/fairdatasociety/fairOS-dfs/pkg/logging"
	"github.com/sirupsen/logrus"
)

func TestPodIntegration(t *testing.T) {
	username := "user"
	max := 10

	api, err := NewMockDfsAPI("mock", "", "", "", logging.New(os.Stdout, logrus.PanicLevel))
	if err != nil {
		t.Fatal(err)
	}
	_, _, ui, err := api.CreateUser(username, "123456789", "", "")
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < max; i++ {
		_, err := api.CreatePod(fmt.Sprintf("pod_%d", i), "123456789", ui.GetSessionId())
		if err != nil {
			t.Fatal(err)
		}
	}

	for i := 0; i < max; i++ {
		_, err := api.PodStat(fmt.Sprintf("pod_%d", i), ui.GetSessionId())
		if err != nil {
			t.Fatal(err)
		}
	}

	for i := 0; i < max; i++ {
		err = api.ClosePod(fmt.Sprintf("pod_%d", i), ui.GetSessionId())
		if err != nil {
			t.Fatal(err)
		}
	}

	for i := 0; i < max; i++ {
		_, _, err := api.ListPods(ui.GetSessionId())
		if err != nil {
			t.Fatal(err)
		}
	}

	for i := 0; i < max; i++ {
		exist := api.IsPodExist(fmt.Sprintf("pod_%d", i), ui.GetSessionId())
		if !exist {
			t.Fatal("pod should exist")
		}
	}

	for i := 0; i < max; i++ {
		_, err := api.OpenPod(fmt.Sprintf("pod_%d", i), "123456789", ui.GetSessionId())
		if err != nil {
			t.Fatal(err)
		}
	}

	for i := 0; i < max; i++ {
		err = api.DeletePod(fmt.Sprintf("pod_%d", i), "123456789", ui.GetSessionId())
		if err != nil {
			t.Fatal(err)
		}
	}

	for i := 0; i < max; i++ {
		exist := api.IsPodExist(fmt.Sprintf("pod_%d", i), ui.GetSessionId())
		if exist {
			t.Fatal("pod should nil exist")
		}
	}
}
