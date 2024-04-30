package picker

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

const NoValuesProvided = "no values provided"
const DefaultNoItems = "no items"

func TestEverything(t *testing.T) {
	// Test initialisation
	items := []Item{
		{Key: "Alpha", Value: "1"},
		{Key: "Bravo", Value: "2"},
		{Key: "Charlie", Value: "3"},
		{Key: "Delta", Value: "4"},
	}
	input := New(items)
	// We should have an array of 4 items
	if len(input.items) != len(items) {
		t.Errorf("Init list length = %d; expected %d", len(input.items), len(items))
	}
	// Check that View returns the first Key
	result := input.View()
	if result != "Alpha" {
		t.Errorf("View of cursor %d return '%s'; expected '%s'", input.cursor, result, items[input.cursor])
	}
	// Now test resetting the list of items
	input.SetItems([]Item{})
	// and check the View result again
	result = input.View()
	if result != DefaultNoItems {
		t.Errorf("Empty View returned '%s'; expected '%s'", result, DefaultNoItems)
	}
	// Test setting the empty prompt
	input.SetEmpty(NoValuesProvided)
	// and check the View result again
	result = input.View()
	if result != NoValuesProvided {
		t.Errorf("Empty View returned '%s'; expected '%s'", result, NoValuesProvided)
	}
	// Reset the items
	input.SetItems(items)
	if len(input.items) != len(items) {
		t.Errorf("SetItems list length = %d; expected %d", len(input.items), len(items))
	}
	// Check that we can fail to set the cursor
	err := input.SetCursor("non-existent")
	if err == nil {
		t.Errorf("Unexpected non-error from SetCursor")
	}
	// Check that we can set the cursor
	err = input.SetCursor(items[3].Value)
	if err != nil {
		t.Errorf("Unexpected error from SetCursor: %v", err)
	}
	if input.cursor != 3 {
		t.Errorf("SetCursor returned %d; expected 3", input.cursor)
	}
	// Check that the cursor function is valid
	if input.Cursor() != 3 {
		t.Errorf("Cursor returned %d; expected 3", input.Cursor())
	}
	// Check that we get the items back correctly
	return_items := input.Items()
	if len(return_items) != len(items) {
		t.Errorf("Items returned list of length %d; expected %d", len(return_items), len(items))
	}
	for i := 0; i < len(items); i++ {
		if items[i].Key != return_items[i].Key {
			t.Errorf("Items: index %d key is %s; expected %s", i, return_items[i].Key, items[i].Key)
		}
		if items[i].Value != return_items[i].Value {
			t.Errorf("Items: index %d value is %s; expected %s", i, return_items[i].Value, items[i].Value)
		}
	}
	// Set cursor to zero and try to move left. Nothing should happen.
	input.cursor = 0
	k := tea.KeyMsg{Type: tea.KeyLeft}
	model, cmd := input.Update(k)
	if cmd != nil {
		t.Errorf("Update returned non-nil command")
	}
	if model.cursor != 0 {
		t.Errorf("Update returned model with cursor of %d; expected 0", model.cursor)
	}
	// Set cursor to one and try to move left. Should go to zero.
	input.cursor = 1
	model, cmd = input.Update(k)
	if cmd != nil {
		t.Errorf("Update returned non-nil command")
	}
	if model.cursor != 0 {
		t.Errorf("Update returned model with cursor of %d; expected 0", model.cursor)
	}
	// Reset cursor to zero and try moving right. Should go to one.
	input.cursor = 0
	k = tea.KeyMsg{Type: tea.KeyRight}
	model, cmd = input.Update(k)
	if cmd != nil {
		t.Errorf("Update returned non-nil command")
	}
	if model.cursor != 1 {
		t.Errorf("Update returned model with cursor of %d; expected 1", model.cursor)
	}
	// Finally, go to the end of the list and try moving right.
	endpoint := len(input.items) - 1
	input.cursor = endpoint
	model, cmd = input.Update(k)
	if cmd != nil {
		t.Errorf("Update returned non-nil command")
	}
	if model.cursor != endpoint {
		t.Errorf("Update returned model with cursor of %d; expected %d", model.cursor, endpoint)
	}
}
