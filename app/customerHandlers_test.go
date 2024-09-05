package app

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ShaizanKhan/go-banking-lib/errs"
	"github.com/ShaizanKhan/go-banking/dto"
	"github.com/ShaizanKhan/go-banking/mocks/service"
	"github.com/gorilla/mux"
	"go.uber.org/mock/gomock"
)

var router *mux.Router
var ch CustomerHandlers
var mockService *service.MockCustomerService

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockService = service.NewMockCustomerService(ctrl)
	ch = CustomerHandlers{service: mockService}
	router = mux.NewRouter()
	router.HandleFunc("/customers", ch.getAllCustomers)
	return func() {
		router = nil
		defer ctrl.Finish()
	}
}

func Test_should_return_customers_with_status_code_200(t *testing.T) {
	//Arrange
	tearDown := setup(t)
	defer tearDown()

	dummyCustomers := []dto.CustomerResponse{
		{Id: "1001", Name: "John Doe", City: "New York", Zipcode: "10001", DateofBirth: "1985-05-20", Status: "Active"},
		{Id: "1002", Name: "Jane Smith", City: "Los Angeles", Zipcode: "90001", DateofBirth: "1990-08-15", Status: "Inactive"},
	}
	mockService.EXPECT().GetAllCustomer("").Return(dummyCustomers, nil)

	request, _ := http.NewRequest(http.MethodGet, "/customers", nil)

	//Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusOK {
		t.Error("Failed while testing the status code")
	}
}

func Test_should_return_status_code_500(t *testing.T) {
	//Arrange
	tearDown := setup(t)
	defer tearDown()
	mockService.EXPECT().GetAllCustomer("").Return(nil, errs.NewUnExpectedError("some database error"))
	request, _ := http.NewRequest(http.MethodGet, "/customers", nil)

	//Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusInternalServerError {
		t.Error("Failed while testing the status code")
	}
}
