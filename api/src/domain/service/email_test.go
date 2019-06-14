package service

import (
	"app/domain/model"
	"app/resources/sendemail"
	"testing"
)

func TestEmailService_Send(t *testing.T) {
	type fields struct {
		Client sendemail.ClientEmail
	}
	type args struct {
		email model.Email
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := EmailService{
				Client: tt.fields.Client,
			}
			if err := e.Send(tt.args.email); (err != nil) != tt.wantErr {
				t.Errorf("EmailService.Send() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
