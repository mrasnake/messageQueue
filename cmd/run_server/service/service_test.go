package service

import (
	"os"
	"testing"
)

// errString is used for comparison of error strings
// by returning the string "nil" when given a nil value
func errString(err error) string {
	if err == nil {
		return "nil"
	}
	return err.Error()
}

// AddItemTest is a struct to hold all testing parameters
// for each test performed on the AddItem.
type AddItemTest struct {
	name            string
	service         *ItemService
	logFileName		string
	inputRequest  	*AddItemRequest
	wantErr         string
}

// Test_AddItem builds a list of custom structs and loops through each of them
// performing the associated unit test on AddItem() with the specified parameters.
func Test_AddItem(t *testing.T) {

	// tests contains all the parameters, checks and expected results of each test.
	tests := []AddItemTest{
		// AddItem Test #1 checks that when given a bad 'AddItemRequest' value
		// the AddItem function fails correctly.
		{
			name: "AddItem Test #1 - Validation Failure",
			logFileName: "./output.txt",
			inputRequest: &AddItemRequest{},
			wantErr: "invalid request: Value: cannot be blank.",
		},
		// AddItem Test #1 checks when passing valid
		// parameters to AddItem, no errors are returned.
		{
			name: "AddItem Test #1 - Success",
			logFileName: "./output.txt",
			inputRequest: &AddItemRequest{
				Value: "foo",
			},
			wantErr: "nil",
		},
	}

	// This loops through each item in the tests list, uses the individual parameters
	// to prepare and perform the unit test and compares the results.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			s, err := NewService(tt.logFileName)
			if err != nil{
				t.Fatalf("could not initialize service: %v", err)
			}

			if err = s.AddItem(tt.inputRequest); tt.wantErr != errString(err) {
				t.Errorf("unexpected error, got=%v; want=%v", errString(err), tt.wantErr)
			}

			os.Remove(tt.logFileName)

		})
	}
}


