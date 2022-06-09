package store

type MockStore struct {
	WriteMock func(data interface{}) error
	ReadMock  func(data interface{}) error
}

func (ms *MockStore) Write(data interface{}) error {
	return ms.WriteMock(data)
}

func (ms *MockStore) Read(data interface{}) error {
	return ms.ReadMock(data)
}
