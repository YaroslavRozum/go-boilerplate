package controllers

type Service interface {
	Validate(interface{}) error
	Execute(interface{}) interface{}
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
	response := s.service.Execute(data)
	return response, nil
}

func NewServiceRunnerCreator(service Service) func(RunnableContext) Runnable {
	return func(ctx RunnableContext) Runnable {
		serviceRunner := &ServiceRunner{service}
		if serviceWithContext, ok := service.(ServiceWithContext); ok {
			serviceWithContext.SetContext(ctx)
		}
		return serviceRunner
	}
}
