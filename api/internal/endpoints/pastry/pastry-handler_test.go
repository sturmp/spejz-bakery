package pastry

import (
	"api/internal/utility/test"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetPastries_DBFetchIsFailing(t *testing.T) {
	// Arrange
	repository := FakeRepository{}
	repository.FetchAllPastriesMock = func(languageCode string) ([]Pastry, error) {
		return nil, test.DummyError{}
	}
	Repository = &repository
	request := http.Request{}
	response := httptest.NewRecorder()

	// Act
	GetPastries(response, &request)
	var pastries []Pastry
	json.NewDecoder(response.Result().Body).Decode(&pastries)

	// Assert
	if response.Code != http.StatusInternalServerError {
		t.Fatalf("Response StatusCode should be: %d Actual: %d", http.StatusInternalServerError, response.Code)
	}
	if len(pastries) != 0 {
		t.Fatalf("Count of pastries in the response should be: %d Actual: %d",
			0,
			len(pastries))
	}
}

func TestGetPastries_DBFetchIsSuccessful(t *testing.T) {
	// Arrange
	testPastries := []Pastry{
		{
			Id:               0,
			Name:             "test",
			Description:      "test",
			Price:            "123",
			UnitOfMeasure:    "db",
			QuantityPerPiece: "123",
			Enabled:          true,
		},
	}
	repository := FakeRepository{}
	repository.FetchAllPastriesMock = func(languageCode string) ([]Pastry, error) {
		return testPastries, nil
	}
	Repository = &repository
	request := http.Request{}
	response := httptest.NewRecorder()

	// Act
	GetPastries(response, &request)
	var pastries []Pastry
	json.NewDecoder(response.Result().Body).Decode(&pastries)

	// Assert
	if response.Code != http.StatusOK {
		t.Fatalf("Response StatusCode should be: %d Actual: %d", http.StatusOK, response.Code)
	}
	if len(pastries) != len(testPastries) {
		t.Fatalf("Count of pastries in the response should be: %d Actual: %d",
			len(testPastries),
			len(pastries))
	}
}

func TestGetPastries_ReturnsOnlyEnabledPastries(t *testing.T) {
	// Arrange
	testPastries := []Pastry{
		{
			Id:               0,
			Name:             "test",
			Description:      "test",
			Price:            "123",
			UnitOfMeasure:    "db",
			QuantityPerPiece: "123",
			Enabled:          true,
		},
		{
			Id:               1,
			Name:             "test2",
			Description:      "test2",
			Price:            "123",
			UnitOfMeasure:    "db",
			QuantityPerPiece: "123",
			Enabled:          false,
		},
	}
	repository := FakeRepository{}
	repository.FetchAllPastriesMock = func(languageCode string) ([]Pastry, error) {
		return testPastries, nil
	}
	Repository = &repository
	request := http.Request{}
	response := httptest.NewRecorder()

	// Act
	GetPastries(response, &request)
	var pastries []Pastry
	json.NewDecoder(response.Result().Body).Decode(&pastries)

	// Assert
	if response.Code != http.StatusOK {
		t.Fatalf("Response StatusCode should be: %d Actual: %d", http.StatusOK, response.Code)
	}
	if len(pastries) != 1 {
		t.Fatalf("Count of pastries in the response should be: %d Actual: %d",
			1,
			len(pastries))
	}
	if !pastries[0].Enabled {
		t.Fatal("Only enabled pastries should be returned!")
	}
}

func TestGetAllPastries_DBFetchIsFailing(t *testing.T) {
	// Arrange
	repository := FakeRepository{}
	repository.FetchAllPastriesMock = func(languageCode string) ([]Pastry, error) {
		return nil, test.DummyError{}
	}
	Repository = &repository
	request := http.Request{}
	response := httptest.NewRecorder()

	// Act
	GetAllPastries(response, &request)
	var pastries []Pastry
	json.NewDecoder(response.Result().Body).Decode(&pastries)

	// Assert
	if response.Code != http.StatusInternalServerError {
		t.Fatalf("Response StatusCode should be: %d Actual: %d", http.StatusInternalServerError, response.Code)
	}
	if len(pastries) != 0 {
		t.Fatalf("Count of pastries in the response should be: %d Actual: %d",
			0,
			len(pastries))
	}
}

func TestGetAllPastries_DBFetchIsSuccessful(t *testing.T) {
	// Arrange
	testPastries := []Pastry{
		{
			Id:               0,
			Name:             "test",
			Description:      "test",
			Price:            "123",
			UnitOfMeasure:    "db",
			QuantityPerPiece: "123",
			Enabled:          true,
		},
	}
	repository := FakeRepository{}
	repository.FetchAllPastriesMock = func(languageCode string) ([]Pastry, error) {
		return testPastries, nil
	}
	Repository = &repository
	request := http.Request{}
	response := httptest.NewRecorder()

	// Act
	GetAllPastries(response, &request)
	var pastries []Pastry
	json.NewDecoder(response.Result().Body).Decode(&pastries)

	// Assert
	if response.Code != http.StatusOK {
		t.Fatalf("Response StatusCode should be: %d Actual: %d", http.StatusOK, response.Code)
	}
	if len(pastries) != len(testPastries) {
		t.Fatalf("Count of pastries in the response should be: %d Actual: %d",
			len(testPastries),
			len(pastries))
	}
}

func TestGetAllPastries_ReturnsEnabledAndDisabledPastries(t *testing.T) {
	// Arrange
	testPastries := []Pastry{
		{
			Id:               0,
			Name:             "test",
			Description:      "test",
			Price:            "123",
			UnitOfMeasure:    "db",
			QuantityPerPiece: "123",
			Enabled:          true,
		},
		{
			Id:               1,
			Name:             "test2",
			Description:      "test2",
			Price:            "123",
			UnitOfMeasure:    "db",
			QuantityPerPiece: "123",
			Enabled:          false,
		},
	}
	repository := FakeRepository{}
	repository.FetchAllPastriesMock = func(languageCode string) ([]Pastry, error) {
		return testPastries, nil
	}
	Repository = &repository
	request := http.Request{}
	response := httptest.NewRecorder()

	// Act
	GetAllPastries(response, &request)
	var pastries []Pastry
	json.NewDecoder(response.Result().Body).Decode(&pastries)

	// Assert
	if response.Code != http.StatusOK {
		t.Fatalf("Response StatusCode should be: %d Actual: %d", http.StatusOK, response.Code)
	}
	if len(pastries) != len(testPastries) {
		t.Fatalf("Count of pastries in the response should be: %d Actual: %d",
			len(testPastries),
			len(pastries))
	}
}

func TestUpdatePastry_InvalidRequestBody(t *testing.T) {
	// Arrange
	request, _ := http.NewRequest("PUT", "/pastry", strings.NewReader(""))
	response := httptest.NewRecorder()

	// Act
	UpdatePastry(response, request)

	// Assert
	if response.Code != http.StatusBadRequest {
		t.Fatalf("Response StatusCode should be: %d Actual: %d", http.StatusBadRequest, response.Code)
	}
}

func TestUpdatePastry_DBOperationIsFailing(t *testing.T) {
	// Arrange
	repository := FakeRepository{}
	repository.UpdatePastryMock = func(pastry Pastry, languageCode string) error {
		return test.DummyError{}
	}
	Repository = &repository

	pastry := Pastry{
		Id:               0,
		Name:             "new name",
		Description:      "test",
		Price:            "123",
		UnitOfMeasure:    "db",
		QuantityPerPiece: "123",
		Enabled:          true,
	}
	body, _ := json.Marshal(pastry)
	request, _ := http.NewRequest("PUT", "/pastry", strings.NewReader(string(body)))
	response := httptest.NewRecorder()

	// Act
	UpdatePastry(response, request)

	// Assert
	if response.Code != http.StatusInternalServerError {
		t.Fatalf("Response StatusCode should be: %d Actual: %d", http.StatusInternalServerError, response.Code)
	}
}

func TestUpdatePastry_DBOperationIsSuccessful(t *testing.T) {
	// Arrange
	repository := FakeRepository{}
	repository.UpdatePastryMock = func(pastry Pastry, languageCode string) error {
		return nil
	}
	Repository = &repository

	pastry := Pastry{
		Id:               0,
		Name:             "new name",
		Description:      "test",
		Price:            "123",
		UnitOfMeasure:    "db",
		QuantityPerPiece: "123",
		Enabled:          true,
	}
	body, _ := json.Marshal(pastry)
	request, _ := http.NewRequest("PUT", "/pastry", strings.NewReader(string(body)))
	response := httptest.NewRecorder()

	// Act
	UpdatePastry(response, request)
	var pastryResult Pastry
	json.NewDecoder(response.Result().Body).Decode(&pastryResult)

	// Assert
	if response.Code != http.StatusOK {
		t.Fatalf("Response StatusCode should be: %d Actual: %d", http.StatusOK, response.Code)
	}
	if pastryResult != pastry {
		t.Fatalf("Returned Pastry is not equal to the one in the request")
	}
}

func TestCreatePastry_InvalidRequestBody(t *testing.T) {
	// Arrange
	request, _ := http.NewRequest("POST", "/pastry", strings.NewReader(string("")))
	response := httptest.NewRecorder()

	// Act
	CreatePastry(response, request)

	// Assert
	if response.Code != http.StatusBadRequest {
		t.Fatalf("Response StatusCode should be: %d Actual: %d", http.StatusBadRequest, response.Code)
	}
}

func TestCreatePastry_DBOperationIsFailing(t *testing.T) {
	// Arrange
	repository := FakeRepository{}
	repository.CreatePastryMock = func(createRequest CreatePastryRequest, languageCode string) (Pastry, error) {
		return Pastry{}, test.DummyError{}
	}
	Repository = &repository

	pastry := CreatePastryRequest{
		Name:             "new name",
		Description:      "test",
		Price:            "123",
		UnitOfMeasure:    1,
		QuantityPerPiece: "123",
	}
	body, _ := json.Marshal(pastry)
	request, _ := http.NewRequest("POST", "/pastry", strings.NewReader(string(body)))
	response := httptest.NewRecorder()

	// Act
	CreatePastry(response, request)

	// Assert
	if response.Code != http.StatusInternalServerError {
		t.Fatalf("Response StatusCode should be: %d Actual: %d", http.StatusInternalServerError, response.Code)
	}
}

func TestCreatePastry_DBOperationIsSuccessful(t *testing.T) {
	// Arrange
	repository := FakeRepository{}
	pastry := Pastry{
		Id:               0,
		Name:             "new name",
		Description:      "test",
		Price:            "123",
		UnitOfMeasure:    "db",
		QuantityPerPiece: "123",
		Enabled:          true,
	}
	repository.CreatePastryMock = func(createRequest CreatePastryRequest, languageCode string) (Pastry, error) {
		return pastry, nil
	}
	Repository = &repository

	createPastryRequest := CreatePastryRequest{
		Name:             "new name",
		Description:      "test",
		Price:            "123",
		UnitOfMeasure:    1,
		QuantityPerPiece: "123",
	}
	body, _ := json.Marshal(createPastryRequest)
	request, _ := http.NewRequest("POST", "/pastry", strings.NewReader(string(body)))
	response := httptest.NewRecorder()

	// Act
	CreatePastry(response, request)
	var pastryResult Pastry
	json.NewDecoder(response.Result().Body).Decode(&pastryResult)

	// Assert
	if response.Code != http.StatusOK {
		t.Fatalf("Response StatusCode should be: %d Actual: %d", http.StatusOK, response.Code)
	}
	if pastryResult != pastry {
		t.Fatalf("Returned Pastry is not equal to the one in the request")
	}
}

type FakeRepository struct {
	FetchAllPastriesMock func(languageCode string) ([]Pastry, error)
	UpdatePastryMock     func(pastry Pastry, languageCode string) error
	CreatePastryMock     func(createRequest CreatePastryRequest, languageCode string) (Pastry, error)
}

func (repository *FakeRepository) FetchAllPastries(languageCode string) (pastries []Pastry, err error) {
	return repository.FetchAllPastriesMock(languageCode)
}

func (repository *FakeRepository) UpdatePastry(pastry Pastry, languageCode string) error {
	return repository.UpdatePastryMock(pastry, languageCode)
}

func (repository *FakeRepository) CreatePastry(createRequest CreatePastryRequest, languageCode string) (pastry Pastry, err error) {
	return repository.CreatePastryMock(createRequest, languageCode)
}
