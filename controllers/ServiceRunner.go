package controllers

import (
	"reflect"
)

type Service interface {
	Validate(interface{}) error
	Execute(interface{}) (interface{}, error)
}

type ServiceWithContext interface {
	SetContext(interface{})
	Service
}

type ServiceRunner struct {
	service Service
}

func (s *ServiceRunner) Run(data interface{}) (interface{}, error) {
	if err := s.service.Validate(data); err != nil {
		return nil, err
	}
	response, err := s.service.Execute(data)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func NewServiceRunnerCreator(service Service) func(RunnableContext) Runner {
	return func(ctx RunnableContext) Runner {
		sV := reflect.ValueOf(service)
		sT := sV.Elem().Type()
		newService := reflect.New(sT).Interface().(Service)
		serviceRunner := &ServiceRunner{newService}
		if serviceWithContext, ok := newService.(ServiceWithContext); ok {
			serviceWithContext.SetContext(ctx)
		}
		return serviceRunner
	}
}
