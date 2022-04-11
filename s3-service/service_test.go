package s3_service

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	mocks "unit-test-tut/mocks/s3-service"
)

func TestCopyObject(t *testing.T) {
	// given
	mockClient := &mocks.S3Client{}
	SetS3Client(mockClient)
	mockClient.On("CopyObject", mock.Anything, mock.Anything).Return(nil, nil)

	// when
	err := CopyS3Object()

	// then
	assert.Equal(t, err, nil)
}
