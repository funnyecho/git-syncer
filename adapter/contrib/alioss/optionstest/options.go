package optionstest

// MockOptions mock Options
type MockOptions struct {
	endpoint        string
	accessKeyID     string
	accessKeySecret string
	bucket          string
	base            string
}

func (m *MockOptions) Endpoint() string {
	return m.endpoint
}

func (m *MockOptions) AccessKeyID() string {
	return m.accessKeyID
}

func (m *MockOptions) AccessKeySecret() string {
	return m.accessKeySecret
}

func (m *MockOptions) Bucket() string {
	return m.bucket
}

func (m *MockOptions) Base() string {
	return m.base
}
