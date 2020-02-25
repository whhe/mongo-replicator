package operator

import "github.com/whhe/mongo-replicator/model"

// Operator interface can be used to provide the replicator.Replicator with custom
// implementations to deal with change event documents.
type Operator interface {
	Insert(model.ChangeEvent) error
	Delete(model.ChangeEvent) error
	Replace(model.ChangeEvent) error
	Update(model.ChangeEvent) error
	Drop(model.ChangeEvent) error
	Rename(model.ChangeEvent) error
	DropDatabase(model.ChangeEvent) error
	Invalidate(model.ChangeEvent) error
}

type noopOperator struct{}

func (n *noopOperator) Insert(model.ChangeEvent) error {
	return nil
}

func (n *noopOperator) Delete(model.ChangeEvent) error {
	return nil
}

func (n *noopOperator) Replace(model.ChangeEvent) error {
	return nil
}

func (n *noopOperator) Update(model.ChangeEvent) error {
	return nil
}

func (n *noopOperator) Drop(model.ChangeEvent) error {
	return nil
}

func (n *noopOperator) Rename(model.ChangeEvent) error {
	return nil
}

func (n *noopOperator) DropDatabase(model.ChangeEvent) error {
	return nil
}

func (n *noopOperator) Invalidate(model.ChangeEvent) error {
	return nil
}

// NewNoopOperator return a no-op Operator implementation.
func NewNoopOperator() Operator {
	return &noopOperator{}
}
