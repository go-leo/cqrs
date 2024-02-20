package query

import (
	"context"
	"errors"
	"github.com/go-leo/cqrs"
)

type DemoDefaultAssembler[Req any, Resp any] interface {
	FromDemoDefaultReq(ctx context.Context, req Req) (*DemoDefaultQuery, error)
	ToDemoDefaultResp(ctx context.Context, res *DemoDefaultResult) (Resp, error)
}

type DemoDefaultQuery struct {
}

type DemoDefaultResult struct {
}

type DemoDefault cqrs.QueryHandler[*DemoDefaultQuery, *DemoDefaultResult]

func NewDemoDefault() DemoDefault {
	return &demoDefault{}
}

type demoDefault struct {
}

func (h *demoDefault) Handle(ctx context.Context, q *DemoDefaultQuery) (*DemoDefaultResult, error) {
	return nil, errors.New("this is error")
}
