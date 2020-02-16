package replicator

import (
	"errors"

	"github.com/whhe/mongo-replicator/pkg/model"
	"github.com/whhe/mongo-replicator/pkg/operator"
)

// Replicator is a wrapper for operator.Operator.
type Replicator struct {
	operator.Operator
}

// New creates a Replicator instance.
func New(operator operator.Operator) *Replicator {
	return &Replicator{operator}
}

// Replicate calls the corresponding method of Operator according to the type of change event.
func (r *Replicator) Replicate(e model.ChangeEvent) error {
	switch e.OperationType {
	case "insert":
		return r.Insert(e)
	case "delete":
		return r.Delete(e)
	case "replace":
		return r.Replace(e)
	case "update":
		return r.Update(e)
	case "drop":
		return r.Drop(e)
	case "rename":
		return r.Rename(e)
	case "dropDatabase":
		return r.DropDatabase(e)
	case "invalidate":
		return r.Invalidate(e)
	default:
		return errors.New("invalid operation type: " + e.OperationType)
	}
}
