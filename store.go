package adminaws

type AdminAwsRegion struct {
	ID string
}
type AdminAwsEc2 struct {
	ID    string
	Image string
	Ports []uint16
}

type AdminAwsStore struct {
	Ec2s    []AdminAwsEc2
	Regions []AdminAwsRegion
}

func New() *AdminAwsStore {
	s := &AdminAwsStore{}
	return s
}
