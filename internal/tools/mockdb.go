package tools

import "time"

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"kofil": {
		AuthToken: "123",
		Username:  "koffil",
	},
}

var mockCardDetails = map[string]CardDetails{
	"kofil": {
		Cards:    109,
		Username: "koffil",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserCards(username string) *CardDetails {
	time.Sleep(time.Second * 1)

	var clientData = CardDetails{}
	clientData, ok := mockCardDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}