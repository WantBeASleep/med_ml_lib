package gtc

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type _timestamp struct{}

var Timestamp _timestamp

func (_timestamp) TimePointerToPointer(p *time.Time) *timestamppb.Timestamp {
	if p == nil {
		return nil
	}
	return timestamppb.New(*p)
}

func (_timestamp) PointerToTimePointer(p *timestamppb.Timestamp) *time.Time {
	if p == nil {
		return nil
	}

	t := p.AsTime()
	return &t
}
