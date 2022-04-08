package datastore

import (
	"sort"
	"testing"
)

// AddItemTest is a struct to hold all testing parameters
// for each test performed on the AddItem.
type AddItemTest struct {
	name            string
	input  			string
	wantErr         string
}

// Test_AddItem builds a list of custom structs and loops through each of them
// performing the associated unit test on AddItem() with the specified parameters.
func Test_AddItem(t *testing.T) {

	store := NewStorage()

	// tests contains all the parameters, checks and expected results of each test.
	tests := []AddItemTest{
		// AddItem Test #1 checks when passing valid
		// parameters to AddItem, no errors are returned.
		{
			name: "AddItem Test #1 - Success",
			input: "foo",
			wantErr: "nil",
		},
	}

	// This loops through each item in the tests list, uses the individual parameters
	// to prepare and perform the unit test and compares the results.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if err := store.AddItem(tt.input); tt.wantErr != errString(err) {
				t.Errorf("unexpected error, got=%v; want=%v", errString(err), tt.wantErr)
			}

			if !store.db[tt.input]{
				t.Errorf("Unable to find value: %v in datastore", tt.input)
			}

		})
	}
}

// errString is used for comparison of error strings
// by returning the string "nil" when given a nil value
func errString(err error) string {
	if err == nil {
		return "nil"
	}
	return err.Error()
}

// GetItemTest is a struct to hold all testing parameters
// for each test performed on the GetItem.
type GetItemTest struct {
	name            string
	input  			string
	want			string
	wantErr         string
}

// Test_GetItem builds a list of custom structs and loops through each of them
// performing the associated unit test on GetItem() with the specified parameters.
func Test_GetItem(t *testing.T) {

	store := NewStorage()
	store.db["foo"] = true

	// tests contains all the parameters, checks and expected results of each test.
	tests := []GetItemTest{
		// GetItem Test #1 checks when passing valid
		// parameters to GetItem, no errors are returned.
		{
			name:    "GetItem Test #1 - Success",
			input:   "foo",
			want: 	 "foo",
			wantErr: "nil",
		},
	}

	// This loops through each item in the tests list, uses the individual parameters
	// to prepare and perform the unit test and compares the results.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			ret, err := store.GetItem(tt.input)
			if tt.wantErr != errString(err) {
				t.Errorf("unexpected error, got=%v; want=%v", errString(err), tt.wantErr)
			}

			if ret != tt.want{
				t.Errorf("unexpected value, got=%v; want=%v", ret, tt.want)
			}

		})
	}
}

// RemoveItemTest is a struct to hold all testing parameters
// for each test performed on the RemoveItem.
type RemoveItemTest struct {
	name            string
	input  			string
	wantErr         string
}

// Test_RemoveItem builds a list of custom structs and loops through each of them
// performing the associated unit test on RemoveItem() with the specified parameters.
func Test_RemoveItem(t *testing.T) {

	store := NewStorage()
	store.db["foo"] = true

	// tests contains all the parameters, checks and expected results of each test.
	tests := []RemoveItemTest{
		// RemoveItem Test #1 checks that when given a bad input value
		// the RemoveItem function fails correctly.
		{
			name: "RemoveItem Test #1 - Does Not Exist Failure",
			input: "bar",
			wantErr: "cannot remove item that does not exist",
		},
		// RemoveItem Test #1 checks when passing valid
		// parameters to RemoveItem, no errors are returned.
		{
			name: "RemoveItem Test #2 - Success",
			input: "foo",
			wantErr: "nil",
		},
	}

	// This loops through each item in the tests list, uses the individual parameters
	// to prepare and perform the unit test and compares the results.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if err := store.RemoveItem(tt.input); tt.wantErr != errString(err) {
				t.Errorf("unexpected error, got=%v; want=%v", errString(err), tt.wantErr)
			}

		})
	}
}

// ListItemsTest is a struct to hold all testing parameters
// for each test performed on the ListItems.
type ListItemsTest struct {
	name            string
	want  			[]string
	wantErr         string
}

// Test_ListItems builds a list of custom structs and loops through each of them
// performing the associated unit test on ListItems() with the specified parameters.
func Test_ListItems(t *testing.T) {

	store := NewStorage()

	// tests contains all the parameters, checks and expected results of each test.
	tests := []ListItemsTest{
		// ListItems Test #1 checks that when given a bad input value
		// the ListItems function fails correctly.
		{
			name: "ListItems Test #1 - Does Not Exist Failure",
			want: []string{},
			wantErr: "no items to list",
		},
		// ListItems Test #1 checks when passing valid
		// parameters to ListItems, no errors are returned.
		{
			name: "ListItems Test #2 - Success",
			want: []string{"foo", "bar"},
			wantErr: "nil",
		},
	}

	// This loops through each item in the tests list, uses the individual parameters
	// to prepare and perform the unit test and compares the results.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			ret, err := store.ListItems()
			if tt.wantErr != errString(err) {
				t.Errorf("unexpected error, got=%v; want=%v", errString(err), tt.wantErr)
			}

			if !areEqual(ret, tt.want){
				t.Error("unexpeted values, want and ret are not equal")
			}

			store.db["foo"] = true
			store.db["bar"] = true

		})
	}
}

// areEqual
func areEqual(got, want []string) bool {

	if len(got) != len(want){
		return false
	}

	sort.Strings(got)
	sort.Strings(want)

	for i := range got{
		if got[i] != want[i]{
			return false
		}
	}

	return true

}