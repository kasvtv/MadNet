package objs

import (
	"fmt"

	"github.com/MadBase/MadNet/application/objs/uint256"
	"github.com/MadBase/MadNet/errorz"
)

type PaginationToken struct {
	LastPaginatedType LastPaginatedType
	LastUtxoId        []byte
	TotalValue        *uint256.Uint256
}

type LastPaginatedType byte

const (
	LastPaginatedUtxo LastPaginatedType = iota
	LastPaginatedDeposit
)

const marshalledSize = 1 + 32 + 32

// UnmarshalBinary takes a byte slice and returns the corresponding
// PaginationToken object
func (pt *PaginationToken) UnmarshalBinary(data []byte) error {
	if pt == nil {
		return errorz.ErrInvalid{}.New("not initialized")
	}

	if data == nil || len(data) != marshalledSize || data[0] > 1 {
		return errorz.ErrInvalid{}.New("bytes invalid")
	}

	pt.LastPaginatedType = LastPaginatedType(data[0])

	pt.LastUtxoId = make([]byte, 32)
	copy(pt.LastUtxoId, data[1:33])

	TotalValue := &uint256.Uint256{}
	TotalValue.UnmarshalBinary(data[33:65])
	pt.TotalValue = TotalValue

	return nil
}

// MarshalBinary takes the PaginationToken object and returns the canonical
// byte slice
func (pt *PaginationToken) MarshalBinary() ([]byte, error) {
	if pt == nil {
		return nil, errorz.ErrInvalid{}.New("not initialized")
	}

	binaryMinValueLeft, err := pt.TotalValue.MarshalBinary()
	if err != nil {
		return nil, err
	}

	bytes := make([]byte, marshalledSize)
	bytes[0] = byte(pt.LastPaginatedType)
	copy(bytes[1:33], pt.LastUtxoId)
	copy(bytes[33:65], binaryMinValueLeft)

	return bytes, nil
}

func (b PaginationToken) String() string {
	return fmt.Sprintf("{LastPaginatedType: %d, LastUtxoId: 0x%x, TotalValue: %s}", b.LastPaginatedType, b.LastUtxoId, b.TotalValue)
}
