package replicator

import (
	"errors"

	"github.com/whhe/mongo-replicator/model"
	"github.com/whhe/mongo-replicator/operator"
)

// Replicator is a wrapper for operator.Operator.
type Replicator struct {
	op operator.Operator
}

// NewReplicator creates a Replicator instance.
func NewReplicator(operator operator.Operator) *Replicator {
	return &Replicator{operator}
}

// Replicate calls the corresponding method of Operator according to the type of change event.
func (r *Replicator) Replicate(e model.ChangeEvent) error {
	switch e.OperationType {
	case "insert":
		return r.op.Insert(e)
	case "delete":
		return r.op.Delete(e)
	case "replace":
		return r.op.Replace(e)
	case "update":
		return r.op.Update(e)
	case "drop":
		return r.op.Drop(e)
	case "rename":
		return r.op.Rename(e)
	case "dropDatabase":
		return r.op.DropDatabase(e)
	case "invalidate":
		return r.op.Invalidate(e)
	default:
		return errors.New("invalid operation type: " + e.OperationType)
	}
}
