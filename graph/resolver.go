package graph

import (
	"context"

	gmodel "github.com/pashest/task-manager-api/graph/model"
)

type Resolver struct{}

func (r *Resolver) Tasks(ctx context.Context) ([]gmodel.Task, error) {
	return nil, nil
}

func (r *Resolver) CreateTasks(ctx context.Context) (*gmodel.Task, error) {
	return nil, nil
}
