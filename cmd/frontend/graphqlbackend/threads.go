package graphqlbackend

import (
	"context"
	"errors"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/sourcegraph/sourcegraph/cmd/frontend/graphqlbackend/graphqlutil"
)

// Threads is the implementation of the GraphQL API for threads queries and mutations. If it is not
// set at runtime, a "not implemented" error is returned to API clients who invoke it.
//
// This is contributed by enterprise.
var Threads ThreadsResolver

var errThreadsNotImplemented = errors.New("threads is not implemented")

// ThreadByID is called to look up a Thread given its GraphQL ID.
func ThreadByID(ctx context.Context, id graphql.ID) (Thread, error) {
	if Threads == nil {
		return nil, errors.New("threads is not implemented")
	}
	return Threads.ThreadByID(ctx, id)
}

// ThreadInRepository returns a specific thread in the specified repository.
func ThreadInRepository(ctx context.Context, repository graphql.ID, number string) (Thread, error) {
	if Threads == nil {
		return nil, errThreadsNotImplemented
	}
	return Threads.ThreadInRepository(ctx, repository, number)
}

// ThreadsForRepository returns an instance of the GraphQL ThreadConnection type with the list of
// threads defined in a repository.
func ThreadsForRepository(ctx context.Context, repository graphql.ID, arg *graphqlutil.ConnectionArgs) (ThreadConnection, error) {
	if Threads == nil {
		return nil, errThreadsNotImplemented
	}
	return Threads.ThreadsForRepository(ctx, repository, arg)
}

func (schemaResolver) Threads(ctx context.Context, arg *graphqlutil.ConnectionArgs) (ThreadConnection, error) {
	if Threads == nil {
		return nil, errThreadsNotImplemented
	}
	return Threads.Threads(ctx, arg)
}

func (r schemaResolver) CreateThread(ctx context.Context, arg *CreateThreadArgs) (Thread, error) {
	if Threads == nil {
		return nil, errThreadsNotImplemented
	}
	return Threads.CreateThread(ctx, arg)
}

func (r schemaResolver) UpdateThread(ctx context.Context, arg *UpdateThreadArgs) (Thread, error) {
	if Threads == nil {
		return nil, errThreadsNotImplemented
	}
	return Threads.UpdateThread(ctx, arg)
}

func (r schemaResolver) DeleteThread(ctx context.Context, arg *DeleteThreadArgs) (*EmptyResponse, error) {
	if Threads == nil {
		return nil, errThreadsNotImplemented
	}
	return Threads.DeleteThread(ctx, arg)
}

// ThreadsResolver is the interface for the GraphQL threads queries and mutations.
type ThreadsResolver interface {
	// Queries
	Threads(context.Context, *graphqlutil.ConnectionArgs) (ThreadConnection, error)

	// Mutations
	CreateThread(context.Context, *CreateThreadArgs) (Thread, error)
	UpdateThread(context.Context, *UpdateThreadArgs) (Thread, error)
	DeleteThread(context.Context, *DeleteThreadArgs) (*EmptyResponse, error)

	// ThreadByID is called by the ThreadByID func but is not in the GraphQL API.
	ThreadByID(context.Context, graphql.ID) (Thread, error)

	// ThreadInRepository is called by the ThreadInRepository func but is not in the GraphQL API.
	ThreadInRepository(ctx context.Context, repository graphql.ID, number string) (Thread, error)

	// ThreadsForRepository is called by the ThreadsForRepository func but is not in the GraphQL
	// API.
	ThreadsForRepository(ctx context.Context, repository graphql.ID, arg *graphqlutil.ConnectionArgs) (ThreadConnection, error)
}

type CreateThreadArgs struct {
	Input createThreadlikeInput
}

type UpdateThreadArgs struct {
	Input updateThreadlikeInput
}

type DeleteThreadArgs struct {
	Thread graphql.ID
}

type ThreadState string

const (
	ThreadStateOpen   ThreadState = "OPEN"
	ThreadStateClosed             = "CLOSED"
)

// Thread is the interface for the GraphQL type Thread.
type Thread interface {
	Threadlike
	State() ThreadState
}

// ThreadConnection is the interface for the GraphQL type ThreadConnection.
type ThreadConnection interface {
	Nodes(context.Context) ([]Thread, error)
	TotalCount(context.Context) (int32, error)
	PageInfo(context.Context) (*graphqlutil.PageInfo, error)
}
