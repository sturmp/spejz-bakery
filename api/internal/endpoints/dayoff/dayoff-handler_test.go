package dayoff

import (
	"api/internal/utility/assert"
	"api/internal/utility/test"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/mux"
)

func TestGetDayOffs_DBFetchIsSuccessful(t *testing.T) {
	// Arrange
	testDayOffs := []DayOff{
		{Id: 0, Day: time.Now()},
	}
	repository := &FakeRepository{}
	repository.FetchDayOffsMock = func() ([]DayOff, error) {
		return testDayOffs, nil
	}
	Repository = repository
	request := http.Request{}
	response := httptest.NewRecorder()

	// Act
	GetDayOffs(response, &request)
	var dayOffs []DayOff
	json.NewDecoder(response.Result().Body).Decode(&dayOffs)

	// Assert
	if response.Code != http.StatusOK {
		t.Fatalf("Response StatusCode should be: %d Actual: %d", http.StatusOK, response.Code)
	}
	if len(dayOffs) != len(testDayOffs) {
		t.Fatalf("Count of dayoffs in the response should be: %d Actual: %d",
			len(testDayOffs),
			len(dayOffs))
	}
	assert.Same(t, testDayOffs, dayOffs, isSameDayOff, handleNotFound)
}

func TestGetDayOffs_DBFetchIsFailing(t *testing.T) {
	// Arrange
	repository := &FakeRepository{}
	repository.FetchDayOffsMock = func() ([]DayOff, error) {
		return nil, test.DummyError{}
	}
	Repository = repository
	request := http.Request{}
	response := httptest.NewRecorder()

	// Act
	GetDayOffs(response, &request)
	var dayOffs []DayOff
	json.NewDecoder(response.Result().Body).Decode(&dayOffs)

	// Assert
	if response.Code != http.StatusInternalServerError {
		t.Fatalf("Response StatusCode should be: %d Actual: %d", http.StatusInternalServerError, response.Code)
	}
	if len(dayOffs) != 0 {
		t.Fatalf("Count of dayoffs in the response should be: %d Actual: %d",
			0,
			len(dayOffs))
	}
}

func TestCreateDayOff_InvalidRequestBody(t *testing.T) {
	body := strings.NewReader("")
	request, _ := http.NewRequest("POST", "/dayoff", body)
	response := httptest.NewRecorder()

	CreateDayOff(response, request)

	if response.Code != http.StatusBadRequest {
		t.Fatalf("Response StatusCode should be: %d Actual: %d", http.StatusBadRequest, response.Code)
	}
}

func TestCreateDayOff_InvalidDayString(t *testing.T) {
	body := strings.NewReader("\"2024081506:00:00Z\"")
	request, _ := http.NewRequest("POST", "/dayoff", body)
	response := httptest.NewRecorder()

	CreateDayOff(response, request)

	if response.Code != http.StatusBadRequest {
		t.Fatalf("Response StatusCode should be: %d Actual: %d", http.StatusBadRequest, response.Code)
	}
}

func TestCreateDayOff_DBOperationIsFailing(t *testing.T) {
	body := strings.NewReader("\"2024-08-15T06:00:00Z\"")
	request, _ := http.NewRequest("POST", "/dayoff", body)
	response := httptest.NewRecorder()

	repository := &FakeRepository{}
	repository.CreateDayOffMock = func(day time.Time) (int64, error) {
		return -1, test.DummyError{}
	}
	Repository = repository

	CreateDayOff(response, request)

	if response.Code != http.StatusInternalServerError {
		t.Fatalf("Response StatusCode should be: %d Actual: %d", http.StatusInternalServerError, response.Code)
	}
}

func TestCreateDayOff_DBOperationIsSuccessful(t *testing.T) {
	now := time.Now().Format(time.RFC3339)
	body := strings.NewReader(fmt.Sprintf("\"%s\"", now))
	request, _ := http.NewRequest("POST", "/dayoff", body)
	response := httptest.NewRecorder()

	repository := &FakeRepository{}
	id := 1
	repository.CreateDayOffMock = func(day time.Time) (int64, error) {
		return int64(id), nil
	}
	Repository = repository

	CreateDayOff(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("Response StatusCode should be: %d Actual: %d", http.StatusOK, response.Code)
	}

	var dayoff DayOff
	json.NewDecoder(response.Result().Body).Decode(&dayoff)
	if dayoff.Id != id {
		t.Fatalf("Id should be: %d Actual: %d", id, dayoff.Id)
	}

	if dayoff.Day.Format(time.RFC3339) != now {
		t.Fatalf("Id should be: %d Actual: %d", id, dayoff.Id)
	}
}

func TestDeleteDayOffs_IdIsMissing(t *testing.T) {
	// Arrange
	repository := &FakeRepository{}
	repository.DeleteDayOffMock = func(id int) error {
		return nil
	}
	Repository = repository

	request, _ := http.NewRequest("DELETE", "/dayoff", nil)
	response := httptest.NewRecorder()

	// Act
	DeleteDayOff(response, request)

	// Assert
	if response.Code != http.StatusBadRequest {
		t.Fatalf("Response StatusCode should be: %d Actual: %d", http.StatusBadRequest, response.Code)
	}
}

func TestDeleteDayOffs_IdIsInvalid(t *testing.T) {
	// Arrange
	repository := &FakeRepository{}
	repository.DeleteDayOffMock = func(id int) error {
		return nil
	}
	Repository = repository

	vars := map[string]string{
		"id": "abc",
	}
	request, _ := http.NewRequest("DELETE", "/dayoff/abc", nil)
	request = mux.SetURLVars(request, vars)
	response := httptest.NewRecorder()

	// Act
	DeleteDayOff(response, request)

	// Assert
	if response.Code != http.StatusBadRequest {
		t.Fatalf("Response StatusCode should be: %d Actual: %d", http.StatusBadRequest, response.Code)
	}
}

func TestDeleteDayOffs_DBOperationIsSuccessful(t *testing.T) {
	// Arrange
	repository := &FakeRepository{}
	repository.DeleteDayOffMock = func(id int) error {
		return nil
	}
	Repository = repository

	vars := map[string]string{
		"id": "1",
	}
	request, _ := http.NewRequest("DELETE", "/dayoff/1", nil)
	request = mux.SetURLVars(request, vars)
	response := httptest.NewRecorder()

	// Act
	DeleteDayOff(response, request)

	// Assert
	if response.Code != http.StatusOK {
		t.Fatalf("Response StatusCode should be: %d Actual: %d", http.StatusOK, response.Code)
	}
}

func TestDeleteDayOffs_DBOperationIsFailing(t *testing.T) {
	// Arrange
	repository := &FakeRepository{}
	repository.DeleteDayOffMock = func(id int) error {
		return test.DummyError{}
	}
	Repository = repository

	vars := map[string]string{
		"id": "1",
	}
	request, _ := http.NewRequest("DELETE", "/dayoff/1", nil)
	request = mux.SetURLVars(request, vars)
	response := httptest.NewRecorder()

	// Act
	DeleteDayOff(response, request)

	// Assert
	if response.Code != http.StatusInternalServerError {
		t.Fatalf("Response StatusCode should be: %d Actual: %d", http.StatusInternalServerError, response.Code)
	}
}

func isSameDayOff(a DayOff, b DayOff) bool {
	return a.Id == b.Id &&
		a.Day.Format(time.RFC3339) == b.Day.Format(time.RFC3339)
}

func handleNotFound(t *testing.T, dayOff DayOff) {
	t.Fatalf("Dayoff [Id: %d, Day: %s] was not found in the response.",
		dayOff.Id,
		dayOff.Day.Format(time.RFC3339))
}

type FakeRepository struct {
	FetchDayOffsMock func() ([]DayOff, error)
	DeleteDayOffMock func(id int) error
	CreateDayOffMock func(day time.Time) (int64, error)
}

func (repository *FakeRepository) FetchDayOffs() ([]DayOff, error) {
	return repository.FetchDayOffsMock()
}

func (repository *FakeRepository) DeleteDayOff(id int) error {
	return repository.DeleteDayOffMock(id)
}

func (repository *FakeRepository) CreateDayOff(day time.Time) (int64, error) {
	return repository.CreateDayOffMock(day)
}
