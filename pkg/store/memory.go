package store

type MemoryStore struct {
	WriteMock func(data interface{}) error
	ReadMock  func(data interface{}) error
}

func (ms *MemoryStore) Write(data interface{}) error {
	return ms.WriteMock(data)
}

func (ms *MemoryStore) Read(data interface{}) error {
	return ms.ReadMock(data)
}
