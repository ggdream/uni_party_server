package uuid

import "github.com/google/uuid"

// NewV1 version1的uuid
func NewV1() (string, error) {
	res, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	return res.String(), nil
}

// NewV4 version4的uuid
func NewV4() string {
	return uuid.NewString()
}

// New equals function NewV4
var New = NewV4
