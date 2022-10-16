package service

import (
	"lovoo/calculator/repository"
	"lovoo/mocks"
	"reflect"
	"testing"
)

func TestNewCalculatorService(t *testing.T) {
	type args struct {
		calculatorRepository repository.CalculatorRepository
	}

	mockCalculatorRepository := new(mocks.CalculatorRepository)

	tests := []struct {
		name string
		args args
		want CalculatorService
	}{
		{
			name: "success",
			args: args{calculatorRepository: mockCalculatorRepository},
			want: calculatorService{
				calculatorRepository: mockCalculatorRepository,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCalculatorService(tt.args.calculatorRepository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCalculatorService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculatorService_Add(t *testing.T) {
	//mockCalculatorRepository := new(mocks.CalculatorRepository)
	type fields struct {
		calculatorRepository repository.CalculatorRepository
	}
	type args struct {
		n1 float32
		n2 float32
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    float32
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				calculatorRepository: repository.NewCalculatorRepository(),
			},
			args: args{
				n1: 1,
				n2: 2,
			},
			want:    3,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := calculatorService{
				calculatorRepository: tt.fields.calculatorRepository,
			}
			got, err := s.Add(tt.args.n1, tt.args.n2)
			if (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Add() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculatorService_Div(t *testing.T) {
	type fields struct {
		calculatorRepository repository.CalculatorRepository
	}
	type args struct {
		n1 float32
		n2 float32
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    float32
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				calculatorRepository: repository.NewCalculatorRepository(),
			},
			args: args{
				n1: 10,
				n2: 5,
			},
			want:    2,
			wantErr: false,
		},
		{
			name: "failed because divider cannot be zero",
			fields: fields{
				calculatorRepository: repository.NewCalculatorRepository(),
			},
			args: args{
				n1: 10,
				n2: 0,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := calculatorService{
				calculatorRepository: tt.fields.calculatorRepository,
			}
			got, err := s.Div(tt.args.n1, tt.args.n2)
			if (err != nil) != tt.wantErr {
				t.Errorf("Div() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Div() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculatorService_Multi(t *testing.T) {
	type fields struct {
		calculatorRepository repository.CalculatorRepository
	}
	type args struct {
		n1 float32
		n2 float32
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    float32
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				calculatorRepository: repository.NewCalculatorRepository(),
			},
			args: args{
				n1: 10,
				n2: 5,
			},
			want:    50,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := calculatorService{
				calculatorRepository: tt.fields.calculatorRepository,
			}
			got, err := s.Multi(tt.args.n1, tt.args.n2)
			if (err != nil) != tt.wantErr {
				t.Errorf("Multi() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Multi() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculatorService_Sub(t *testing.T) {
	type fields struct {
		calculatorRepository repository.CalculatorRepository
	}
	type args struct {
		n1 float32
		n2 float32
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    float32
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				calculatorRepository: repository.NewCalculatorRepository(),
			},
			args: args{
				n1: 10,
				n2: 5,
			},
			want:    5,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := calculatorService{
				calculatorRepository: tt.fields.calculatorRepository,
			}
			got, err := s.Sub(tt.args.n1, tt.args.n2)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sub() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Sub() got = %v, want %v", got, tt.want)
			}
		})
	}
}
