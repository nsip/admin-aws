package adminaws

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func (s *AdminAwsStore) UpdateEc2() bool {
	s.UpdateRegions()
	states := []string{"running"}
	regions := make([]string, len(s.Regions))

	for _, region := range regions {
		sess := session.Must(session.NewSession(&aws.Config{
			Region: aws.String(region),
		}))

		ec2Svc := ec2.New(sess)
		params := &ec2.DescribeInstancesInput{
			Filters: []*ec2.Filter{
				{
					Name:   aws.String("instance-state-name"),
					Values: aws.StringSlice(states),
				},
			},
		}

		result, err := ec2Svc.DescribeInstances(params)
		if err != nil {
			// XXX move to error
			fmt.Println("Error", err)
		} else {
			// fmt.Printf("\n\n\nFetching instance details for region: %s with criteria: %s**\n ", region, instanceCriteria)
			// if len(result.Reservations) == 0 {
			// 	fmt.Printf("There is no instance for the region: %s with the matching criteria:%s  \n", region, instanceCriteria)
			// }
			for _, reservation := range result.Reservations {
				for _, instance := range reservation.Instances {
					// Find name tag
					var nt string
					for _, t := range instance.Tags {
						if *t.Key == "Name" {
							nt = *t.Value
							break
						}
					}
					// Region, Instance ID, NameTag,
					iid := &instance.InstanceId
					pip := &instance.PublicIpAddress
					fmt.Printf("%v,%v,%v,%v\n", region, iid, pip, nt)

					// fmt.Println("current State " + *instance.State.Name)
				}
			}
		}
	}
	return true
}
