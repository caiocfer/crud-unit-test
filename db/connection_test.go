package db

import "testing"

func TestConnectToDatabase(t *testing.T) {
	_, err := ConnectToDatabase()

	if err != nil {
		t.Errorf("Error in database")
	}

}
