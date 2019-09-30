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

func NewServiceRunnerCreator(service Service) func(RunnerContext) Runner {
	return func(ctx RunnerContext) Runner {
		sT := reflect.TypeOf(service)
		sE := sT.Elem()
		newService := reflect.New(sE).Interface().(Service)
		if serviceWithContext, ok := newService.(ServiceWithContext); ok {
			serviceWithContext.SetContext(ctx)
		}
		serviceRunner := &ServiceRunner{newService}
		return serviceRunner
	}
}
