package myservice

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"v3/mock_test/mocks"
)

func TestDoSomething(t *testing.T) {
	// Create an instance of our test object
	apiMock := new(mocks.ExternalAPIMock)

	// Setup expectations
	apiMock.On("FetchData", "foo").Return("bar", nil)

	// Create an instance of MyService using the mock
	myService := New(apiMock)

	// Call the method we want to test
	result, err := myService.DoSomething("foo")

	// Assert expectations
	assert.NoError(t, err)
	assert.Equal(t, "bar", result)
	apiMock.AssertExpectations(t)
}
