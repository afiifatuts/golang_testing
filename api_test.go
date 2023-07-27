package golangtest

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserDb struct {
	mock.Mock
}

func (m MockUserDb) FindById(ID string) (string, error) {
	args := m.Called(ID)
	return args.Get(0).(string), args.Error(1)
}

func TestGetUser(t *testing.T) {
	t.Run("can get user", func(t *testing.T) {
		userDbMock := new(MockUserDb)
		//tentukan behaviournya
		userDbMock.On("FindById", mock.Anything).Return("tsaani", nil)

		req, err := http.NewRequest(http.MethodGet, "/users", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := GetUser(userDbMock)

		// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
		// directly and pass in our Request and ResponseRecorder.
		handler.ServeHTTP(rr, req)

		// Check the status code is what we expect.
		assert.Equal(t, http.StatusOK, rr.Code)

		// Check the response body is what we expect.
		expected := `{"name": "tsaani"}`
		assert.JSONEq(t, expected, rr.Body.String())
	})

	t.Run("should return 5XX error", func(t *testing.T) {
		userDbMock := new(MockUserDb)
		//tentukan behaviournya
		userDbMock.On("FindById", "1").Return("", fmt.Errorf("db error"))

		req, err := http.NewRequest(http.MethodGet, "/users", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := GetUser(userDbMock)

		// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
		// directly and pass in our Request and ResponseRecorder.
		handler.ServeHTTP(rr, req)

		// Check the status code is what we expect.
		assert.Equal(t, http.StatusInternalServerError, rr.Code)

		// Check the response body is what we expect.
		userDbMock.AssertExpectations(t)

	})

}
