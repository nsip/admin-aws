package adminaws

import "os"

type AdminAwsRegion struct {
	ID string
}

// fmt.Printf("%v,%v,%v,%v\n", region, iid, pip, nt)
type AdminAwsEc2 struct {
	ID       string
	Region   string
	Name     string
	PublicIP string
}

type AdminAwsStore struct {
	Ec2s    []AdminAwsEc2
	Regions []AdminAwsRegion
}

func New() *AdminAwsStore {
	os.Setenv("AWS_REGION", "ap-southeast-2")
	s := &AdminAwsStore{}
	return s
}
