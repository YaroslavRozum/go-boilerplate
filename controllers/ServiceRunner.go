package controllers

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
		serviceRunner := &ServiceRunner{service}
		if serviceWithContext, ok := service.(ServiceWithContext); ok {
			serviceWithContext.SetContext(ctx)
		}
		return serviceRunner
	}
}
