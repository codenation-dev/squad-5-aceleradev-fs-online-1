package controller

import (
	"app/domain/errors"
	"app/domain/model"
	"app/domain/validator"
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/textproto"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type mockCustomer struct {
	customer       *model.Customer
	customerInsert *model.CustomerInsert
	customerList   *model.CustomerList
	err            error
}

func (mk mockCustomer) Parse(file multipart.File) (*model.CustomerInsert, error) {

	return mk.customerInsert, mk.err
}

func (mk mockCustomer) CreateCustomer(customer *model.Customer) (*model.Customer, error) {

	return mk.customer, mk.err
}

func (mk mockCustomer) UpdateCustomer(id string, customer *model.Customer) (*model.Customer, error) {
	mk.customer.ID = id
	return mk.customer, mk.err
}

func (mk mockCustomer) ListCustomer(q *validator.CustomerListRequest) (*model.CustomerList, error) {

	return mk.customerList, mk.err
}

func TestCustomerController_UploadCustomer(t *testing.T) {

	mock := mockCustomer{
		customerInsert: &model.CustomerInsert{
			Success:      1,
			AlreadyExist: 0,
		},
		err: nil,
	}

	cc := CustomerController{mock}

	router := gin.Default()

	router.POST("/customers", cc.UploadCustomer)

	file, err := os.Create("clintes.csv")
	assert.Nil(t, err)

	defer file.Close()
	defer os.Remove("clintes.csv")

	file.WriteString("test\n")

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition", `form-data; name=file; filename="clintes.csv"`)
	h.Set("Content-Type", "text/csv")

	part, err := writer.CreatePart(h)
	assert.Nil(t, err)

	_, err = io.Copy(part, file)
	assert.Nil(t, err)

	err = writer.Close()
	assert.Nil(t, err)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/customers", body)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"success\":1,\"alreadyExist\":0}", w.Body.String())
}

func TestCustomerController_CreateCustomer(t *testing.T) {

	mock := mockCustomer{
		customer: &model.Customer{
			ID:     "1111",
			Name:   "test",
			Salary: 1111.11,
		},
		err: nil,
	}

	cc := CustomerController{mock}

	router := gin.Default()

	router.POST("/customer", cc.CreateCustomer)

	b := bytes.NewReader([]byte(`{
		"name": "test",
		"salary": 1111.11
	  }`))

	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/customer", b)
	assert.Nil(t, err)
	router.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
	assert.Equal(t, w.Body.String(), "{\"id\":\"1111\",\"name\":\"test\",\"salary\":1111.11}")

}

func TestCustomerController_CreateCustomer_DuplicatedCustomerError(t *testing.T) {

	mock := mockCustomer{
		customer: nil,
		err:      errors.DuplicatedCustomerError,
	}

	cc := CustomerController{mock}

	router := gin.Default()

	router.POST("/customer", cc.CreateCustomer)

	b := bytes.NewReader([]byte(`{
		"name": "test",
		"salary": 1111.11
	  }`))

	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/customer", b)
	assert.Nil(t, err)
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, w.Body.String(), "[{\"message\":\"Customer já existe\"}]")

}

func TestCustomerController_CreateCustomer_ValidationErrorCustomerName(t *testing.T) {

	mock := mockCustomer{
		customer: nil,
		err:      nil,
	}

	cc := CustomerController{mock}

	router := gin.Default()

	router.POST("/customer", cc.CreateCustomer)

	b := bytes.NewReader([]byte(`{
		"name": "testtesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttest",
		"salary": 1111.11
	  }`))

	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/customer", b)
	assert.Nil(t, err)
	router.ServeHTTP(w, req)

	assert.Equal(t, 422, w.Code)
	assert.Equal(t, w.Body.String(), "[{\"field\":\"Name\",\"message\":\"Field validation for 'Name' failed on the 'max' tag\"}]")

}

func TestCustomerController_CreateCustomer_ValidationErrorCustomerSalary(t *testing.T) {

	mock := mockCustomer{
		customer: nil,
		err:      nil,
	}

	cc := CustomerController{mock}

	router := gin.Default()

	router.POST("/customer", cc.CreateCustomer)

	b := bytes.NewReader([]byte(`{
		"name": "test",
		"salary": 0
	  }`))

	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/customer", b)
	assert.Nil(t, err)
	router.ServeHTTP(w, req)

	assert.Equal(t, 422, w.Code)
	assert.Equal(t, w.Body.String(), `[{"field":"Salary","message":"Field validation for 'Salary' failed on the 'required' tag"}]`)

}

func TestCustomerController_UpdateCustomer(t *testing.T) {

	mock := mockCustomer{
		customer: &model.Customer{
			ID:     "1111",
			Name:   "test",
			Salary: 1111.11,
		},
		err: nil,
	}

	cc := CustomerController{mock}

	router := gin.Default()

	router.PUT("/customer/:customerId", cc.UpdateCustomer)

	uri := "/customer/11111111111111111111111111"

	b := bytes.NewReader([]byte(`{
		"name": "test",
		"salary": 100.10
	}`))

	w := httptest.NewRecorder()
	req, err := http.NewRequest("PUT", uri, b)
	assert.Nil(t, err)
	router.ServeHTTP(w, req)

	assert.Equal(t, 204, w.Code)
	assert.Equal(t, "", w.Body.String())

}

func TestCustomerController_UpdateCustomer_ValidationErrorCustomerName(t *testing.T) {

	mock := mockCustomer{
		customer: &model.Customer{
			ID:     "1111",
			Name:   "test",
			Salary: 1111.11,
		},
		err: nil,
	}

	cc := CustomerController{mock}

	router := gin.Default()

	router.PUT("/customer/:customerId", cc.UpdateCustomer)

	uri := "/customer/11111111111111111111111111"

	b := bytes.NewReader([]byte(`{
		"name": "testtesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttest",
		"salary": 100.10
	}`))

	w := httptest.NewRecorder()
	req, err := http.NewRequest("PUT", uri, b)
	assert.Nil(t, err)
	router.ServeHTTP(w, req)

	assert.Equal(t, 422, w.Code)
	assert.Equal(t, "[{\"field\":\"Name\",\"message\":\"Field validation for 'Name' failed on the 'max' tag\"}]", w.Body.String())

}

func TestCustomerController_UpdateCustomer_ValidationErrorCustomerSalary(t *testing.T) {

	mock := mockCustomer{
		customer: &model.Customer{
			ID:     "1111",
			Name:   "test",
			Salary: 1111.11,
		},
		err: nil,
	}

	cc := CustomerController{mock}

	router := gin.Default()

	router.PUT("/customer/:customerId", cc.UpdateCustomer)

	uri := "/customer/11111111111111111111111111"

	b := bytes.NewReader([]byte(`{
		"name": "test",
		"salary": 0
	}`))

	w := httptest.NewRecorder()
	req, err := http.NewRequest("PUT", uri, b)
	assert.Nil(t, err)
	router.ServeHTTP(w, req)

	assert.Equal(t, 422, w.Code)
	assert.Equal(t, "[{\"field\":\"Salary\",\"message\":\"Field validation for 'Salary' failed on the 'required' tag\"}]", w.Body.String())

}

func TestCustomerController_UpdateCustomer_DuplicatedCustomerError(t *testing.T) {

	mock := mockCustomer{
		customer: &model.Customer{
			ID:     "1111",
			Name:   "test",
			Salary: 1111.11,
		},
		err: errors.DuplicatedCustomerError,
	}

	cc := CustomerController{mock}

	router := gin.Default()

	router.PUT("/customer/:customerId", cc.UpdateCustomer)

	uri := "/customer/11111111111111111111111112"

	b := bytes.NewReader([]byte(`{
		"name": "test",
		"salary": 110.10
	}`))

	w := httptest.NewRecorder()
	req, err := http.NewRequest("PUT", uri, b)
	assert.Nil(t, err)
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "[{\"message\":\"Customer já existe\"}]", w.Body.String())

}

func TestCustomerController_ListCustomer(t *testing.T) {

	cl := []model.Customer{model.Customer{
		ID:     "1111",
		Name:   "test",
		Salary: 1111.11,
	}}

	mock := mockCustomer{
		customerList: &model.CustomerList{
			Data:    cl,
			Records: int64(1),
		},
		err: nil,
	}

	cc := CustomerController{mock}

	router := gin.Default()

	router.GET("/customers", cc.ListCustomer)

	b := bytes.NewReader([]byte(`{
		"name": "test",
		"salary": 110.10
	}`))

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/customers", b)
	assert.Nil(t, err)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"records\":1,\"data\":[{\"id\":\"1111\",\"name\":\"test\",\"salary\":1111.11}]}", w.Body.String())

}
