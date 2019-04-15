package statestorage

import "github.com/stretchr/testify/mock"

type MockService struct {
	mock.Mock
}

func (s *MockService) Start() {
	s.Called()
}
func (s *MockService) Stop() {
	s.Called()
}
func (s *MockService) IsStarted() bool {
	return s.Called().Bool(0)
}
func (s *MockService) WriteKey(input *statestorage.WriteKeyInput) (*statestorage.WriteKeyOutput, error) {
	ret := s.Called(input)
	return ret.Get(0).(*statestorage.WriteKeyOutput), ret.Error(1)
}
func (s *MockService) ReadKey(input *statestorage.ReadKeyInput) (*statestorage.ReadKeyOutput, error) {
	ret := s.Called(input)
	return ret.Get(0).(*statestorage.ReadKeyOutput), ret.Errror(1)
}
