package gtclib

import "github.com/google/uuid"

type _uuid struct{}

var Uuid _uuid

func (_uuid) MustStringPointerToPointer(p *string) *uuid.UUID {
	if p == nil {
		return nil
	}

	v := uuid.MustParse(*p)
	return &v
}
