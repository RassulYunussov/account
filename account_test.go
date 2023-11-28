package account

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestAccountCreated(t *testing.T) {
	account, err := CreateAccountNumber("KZ", "TP", 7399051)
	assert.NilError(t, err)
	assert.Equal(t, *account, "KZ98TPAC000007399051")
}

func TestAccountValid(t *testing.T) {
	assert.Equal(t, IsValid("KZ86TPAC000000000010"), true)
}

func TestAccountInvalid(t *testing.T) {
	assert.Equal(t, IsValid("KZ93TPAC000000000010"), false)
}

func TestAccountInvalidByLength(t *testing.T) {
	assert.Equal(t, IsValid("KZ93TPAC00010"), false)
}

func TestAccountNotCreatedBecauseOfCountry(t *testing.T) {
	_, err := CreateAccountNumber("KZT", "TP", 10)
	assert.ErrorIs(t, err, ErrInvalidArgument)
	_, err = CreateAccountNumber("K", "TP", 10)
	assert.ErrorIs(t, err, ErrInvalidArgument)
}

func TestAccountNotCreatedBecauseOfOrg(t *testing.T) {
	_, err := CreateAccountNumber("KZ", "TPS", 10)
	assert.ErrorIs(t, err, ErrInvalidArgument)
	_, err = CreateAccountNumber("KZ", "T", 10)
	assert.ErrorIs(t, err, ErrInvalidArgument)
}

func TestAccountNotCreatedBecauseOfId(t *testing.T) {
	_, err := CreateAccountNumber("KZ", "TP", 0)
	assert.ErrorIs(t, err, ErrIdOutOfRange)
	_, err = CreateAccountNumber("KZ", "TP", 1000000000000)
	assert.ErrorIs(t, err, ErrIdOutOfRange)
}
