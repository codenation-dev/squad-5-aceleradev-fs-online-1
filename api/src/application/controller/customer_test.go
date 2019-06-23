package controller

import (
	"net/http"
	"net/http/httptest"
	"net/textproto"
	"io"
	"os"
	"github.com/stretchr/testify/assert"
	"bytes"
	"mime/multipart"
	"app/domain/validator"
	"github.com/gin-gonic/gin"
	"app/domain/model"
	"testing"

)

type mockCustomer struct {
	customer  *model.Customer
	customerList *model.CustomerInsert
	err   error
	count int64
}

func (mk mockCustomer) Parse(file multipart.File) (*model.CustomerInsert, error){

	return mk.customerList, mk.err
}

func (mk mockCustomer) CreateCustomer(customer *model.Customer) (*model.Customer, error){

	return nil, nil
}

func (mk mockCustomer) UpdateCustomer(id string, customer *model.Customer) (*model.Customer, error){

	return nil, nil
}

func (mk mockCustomer) ListCustomer(q *validator.CustomerListRequest) (*model.CustomerList, error){

	return nil, nil
}


func TestCustomerController_UploadCustomer(t *testing.T) {

	mock := mockCustomer{
		customerList: &model.CustomerInsert{
			Success: 1,
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
    h.Set("Content-Disposition",`form-data; name=file; filename="clintes.csv"`)       
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
