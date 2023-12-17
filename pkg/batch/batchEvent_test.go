package batch

import "testing"

func TestNewBatchEvent(t *testing.T) {
	batchID := uint(123)
	eventType := AddIngredientEvent

	// Test valid event type
	event, err := NewBatchEvent(batchID, eventType)
	if err != nil {
		t.Errorf("NewBatchEvent(%d, %q) returned an unexpected error: %v", batchID, eventType, err)
	}
	if event.BatchID != batchID {
		t.Errorf("NewBatchEvent(%d, %q) returned an event with incorrect BatchID. Expected %d, got %d", batchID, eventType, batchID, event.BatchID)
	}
	if event.Type != CleanInputString(eventType) {
		t.Errorf("NewBatchEvent(%d, %q) returned an event with incorrect Type. Expected %q, got %q", batchID, eventType, CleanInputString(eventType), event.Type)
	}

	// Test invalid event type
	invalidEventType := "invalidEventType"
	_, err = NewBatchEvent(batchID, invalidEventType)
	if err != ErrInvalidEventType {
		t.Errorf("NewBatchEvent(%d, %q) did not return the expected error. Expected %v, got %v", batchID, invalidEventType, ErrInvalidEventType, err)
	}
}
