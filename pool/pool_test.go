package pool

import (
	"crawler/logger"
	"testing"
	"time"

	"gopkg.in/h2non/gock.v1"
)

func init() {
	logger.Mute()
}

func TestStartTask(t *testing.T) {
	defer gock.Off()
	for _, hostname := range []string{"a.com", "b.com", "c.com", "e.com"} {
		gock.New("http://" + hostname).Get("/").Reply(200).BodyString("")
	}
	ThrottleRate = time.Nanosecond
	for _, tc := range testCasesStartTask {
		task := StartTask(tc.Body)
		<-task.Done
		if task.Id != tc.Id {
			t.Fatalf("task.Id = %d. Want: %d",
				task.Id, tc.Id)
		}
		if Storage.Len() != tc.StorageSize {
			t.Fatalf("Storage.Len() = %d. Want: %d",
				Storage.Len(), tc.StorageSize)
		}
		if Throttle.Len() != tc.ThrottleSize {
			t.Fatalf("Throttle.Len() = %d. Want: %d",
				Throttle.Len(), tc.ThrottleSize)
		}
	}
}

func TestFindTask(t *testing.T) {
	for _, tc := range testCasesFindTask {
		tc.Setup()
		_, err := FindTask(tc.Id)
		if err != tc.Error {
			t.Fatalf("FindTask(%d) = _, %#v. Want: %#v",
				err, tc.Error)
		}
	}
}

func TestDeleteTask(t *testing.T) {
	for _, tc := range testCasesDeleteTask {
		tc.Setup()
		_, err := DeleteTask(tc.Id)
		if err != tc.Error {
			t.Fatalf("FindTask(%d) = _, %#v. Want: %#v",
				err, tc.Error)
		}
		if Storage.Len() != tc.StorageSize {
			t.Fatalf("Storage.Len() = %d. Want: %d",
				Storage.Len(), tc.StorageSize)
		}
	}
}
