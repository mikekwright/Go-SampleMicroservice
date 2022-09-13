package models

import (
	"testing"
)

func TestTaskCreation_IsSuccessful(t *testing.T) {
	expectedTaskName := "My Task"

	testModel := Task {TaskId: 1, TaskName: expectedTaskName, Description: "Testing task", IsComplete: false}

	if testModel.TaskName != expectedTaskName {
		t.Fatalf("Name should match, but go unexpected name %s", testModel.TaskName)
	}
}