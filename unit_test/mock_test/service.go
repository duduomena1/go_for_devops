package myservice

type ExternalAPI interface {
	FetchData(param string) (string, error)
}

type Myservice struct {
	API ExternalAPI
}

func New(api ExternalAPI) *Myservice {
	return &Myservice{API: api}
}

func (s *Myservice) DoSomething(param string) (string, error) {
	return s.API.FetchData(param)
}
