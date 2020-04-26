package facade

import "fmt"

// API is facade interface of facade package
type API interface {
	Test() string
}

// facade implement
type apiTmpl struct {
	
}

// NewAModuleAPI return new AModuleAPI
func NewAModuleAPI() AModuleAPI {
	return &aModuleImpl{}
}

// AModuleAPI ...
type AModuleAPI interface {
	TestA() string
}

type aModuleImpl struct {

}

func (*aModuleImpl) TestA() string {
	return "A module running"
}