package adapters

import "testing"

func TestCSVUserRepositoryAdapter_GetUsers(t *testing.T) {
	adapter := NewCSVUserRepositoryAdapter()
	adapter.FilePath = "../fixtures/files/users.csv"
	users, _ := adapter.GetUsers()

	if len(users) != 3 {
		t.Fatalf("expected %d, actual %d", 3, len(users))
	}
}
