package adminaws

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func (s *AdminAwsStore) UpdateRegions() bool {
	awsSession := session.Must(session.NewSession(&aws.Config{}))

	svc := ec2.New(awsSession)
	awsRegions, err := svc.DescribeRegions(&ec2.DescribeRegionsInput{})
	if err != nil {
		log.Fatal(err)
	}

	s.Regions = nil
	for _, region := range awsRegions.Regions {
		s.Regions = append(s.Regions, AdminAwsRegion{
			ID: *region.RegionName,
		})
	}

	return true
}
