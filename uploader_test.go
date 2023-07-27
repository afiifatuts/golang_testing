package golangtest

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockS3 struct {
	mock.Mock
}

func (m MockS3) PutObject(input *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
	args := m.Called(input)
	return args.Get(0).(*s3.PutObjectOutput), args.Error(1)
}

func TestUploader_WithMock(t *testing.T) {
	m := new(MockS3)
	u := Uploader{
		svc: m,
	}
	m.On("PutObject", mock.Anything).Return(&s3.PutObjectOutput{},nil)

	err := u.Upload()
	assert.NoError(t, err)

}
