package cash

import (
	"crypto/rand"
	"testing"

	"github.com/iov-one/weave"
	coin "github.com/iov-one/weave/coin"
	"github.com/iov-one/weave/errors"
)

func TestValidateSendMsg(t *testing.T) {
	addr1 := randomAddr(t)
	addr2 := randomAddr(t)

	cases := map[string]struct {
		msg     weave.Msg
		wantErr *errors.Error
	}{
		"success": {
			msg: &SendMsg{
				Amount: coin.NewCoinp(10, 0, "FOO"),
				Dest:   addr1,
				Src:    addr2,
				Memo:   "some memo message",
				Ref:    []byte("some reference"),
			},
			wantErr: nil,
		},
		"success with minimal amount of data": {
			msg: &SendMsg{
				Amount: coin.NewCoinp(10, 0, "FOO"),
				Dest:   addr1,
				Src:    addr2,
			},
			wantErr: nil,
		},
		"empty message": {
			msg:     &SendMsg{},
			wantErr: errors.ErrInvalidAmount,
		},
		"missing source": {
			msg: &SendMsg{
				Amount: coin.NewCoinp(10, 0, "FOO"),
				Dest:   addr1,
			},
			wantErr: errors.ErrEmpty,
		},
		"missing destination": {
			msg: &SendMsg{
				Amount: coin.NewCoinp(10, 0, "FOO"),
				Src:    addr2,
			},
			wantErr: errors.ErrEmpty,
		},
		"reference too long": {
			msg: &SendMsg{
				Amount: coin.NewCoinp(10, 0, "FOO"),
				Dest:   addr1,
				Src:    addr2,
				Ref:    []byte("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaB"),
			},
			wantErr: errors.ErrInvalidState,
		},
		"memo too long": {
			msg: &SendMsg{
				Amount: coin.NewCoinp(10, 0, "FOO"),
				Dest:   addr1,
				Src:    addr2,
				Ref:    []byte("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaB"),
			},
			wantErr: errors.ErrInvalidState,
		},
	}

	for testName, tc := range cases {
		t.Run(testName, func(t *testing.T) {
			if err := tc.msg.Validate(); !tc.wantErr.Is(err) {
				t.Fatalf("unexpected error: %+v", err)
			}
		})
	}
}

func TestValidateFeeTx(t *testing.T) {
	addr1 := randomAddr(t)

	cases := map[string]struct {
		info    *FeeInfo
		wantErr *errors.Error
	}{
		"success": {
			info: &FeeInfo{
				Fees:  coin.NewCoinp(1, 0, "IOV"),
				Payer: addr1,
			},
			wantErr: nil,
		},
		"empty": {
			info:    &FeeInfo{},
			wantErr: errors.ErrInvalidAmount,
		},
		"no fee": {
			info: &FeeInfo{
				Payer: addr1,
			},
			wantErr: errors.ErrInvalidAmount,
		},
		"no payer": {
			info: &FeeInfo{
				Fees: coin.NewCoinp(10, 0, "IOV"),
			},
			wantErr: errors.ErrEmpty,
		},
		"negative fee": {
			info: &FeeInfo{
				Fees:  coin.NewCoinp(-10, 0, "IOV"),
				Payer: addr1,
			},
			wantErr: errors.ErrInvalidAmount,
		},
		"invalid fee ticker": {
			info: &FeeInfo{
				Fees:  coin.NewCoinp(10, 0, "foobar"),
				Payer: addr1,
			},
			wantErr: errors.ErrCurrency,
		},
	}

	for testName, tc := range cases {
		t.Run(testName, func(t *testing.T) {
			if err := tc.info.Validate(); !tc.wantErr.Is(err) {
				t.Fatalf("unexpected error: %+v", err)
			}
		})
	}
}

func randomAddr(t testing.TB) weave.Address {
	a := make(weave.Address, weave.AddressLength)
	if _, err := rand.Read(a); err != nil {
		t.Fatalf("cannot read random data: %s", err)
	}
	if err := a.Validate(); err != nil {
		t.Fatalf("generated address is not valid: %s", err)
	}
	return a
}
