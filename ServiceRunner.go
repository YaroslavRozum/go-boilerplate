package main

// Service is interface for Service Base struct
type Service interface {
	Validate(interface{}) error
	Execute(interface{}) interface{}
}

type ServiceWithContext interface {
	SetContext(interface{})
	Service
}

// ServiceRunner wraps Service
type ServiceRunner struct {
	service Service
}

// Run runs service
func (s *ServiceRunner) Run(data interface{}) (interface{}, error) {
	if err := s.service.Validate(data); err != nil {
		return nil, err
	}
	response := s.service.Execute(data)
	return response, nil
}

func NewServiceRunnerCreator(service Service) func(interface{}) Runnable {
	return func(ctx interface{}) Runnable {
		serviceRunner := &ServiceRunner{service}
		if serviceWithContext, ok := service.(ServiceWithContext); ok {
			serviceWithContext.SetContext(ctx)
		}
		return serviceRunner
	}
}
