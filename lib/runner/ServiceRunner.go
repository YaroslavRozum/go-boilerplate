package runner

type ServiceBuilder func() Service

func (sB ServiceBuilder) buildService() Service {
	return sB()
}

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

func NewServiceRunnerCreator(serviceBuilder ServiceBuilder) func(RunnerContext) Runner {
	return func(ctx RunnerContext) Runner {
		service := serviceBuilder.buildService()
		if serviceWithContext, ok := service.(ServiceWithContext); ok {
			serviceWithContext.SetContext(ctx)
		}
		serviceRunner := &ServiceRunner{service}
		return serviceRunner
	}
}
