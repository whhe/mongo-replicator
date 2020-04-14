package token

import (
	"go.mongodb.org/mongo-driver/bson"
)

// Manager interface can be used to customize how to get and update latest resume token.
type Manager interface {
	// Get gets the latest resume token.
	Get() (bson.Raw, error)
	// Set saves the latest resume token.
	Set(bson.Raw) error
	// Close closes the Manager.
	Close()
}
