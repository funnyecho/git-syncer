package alioss

// Options alioss Options
type Options interface {
	Endpoint() string
	AccessKeyID() string
	AccessKeySecret() string
	Bucket() string
	Base() string
}
