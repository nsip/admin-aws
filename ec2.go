package adminaws

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func (s *AdminAwsStore) UpdateEc2() bool {
	s.Ec2s = nil
	s.UpdateRegions()
	states := []string{"running"}
	for _, region := range s.Regions {
		// fmt.Printf("%+v\n", region)
		sess := session.Must(session.NewSession(&aws.Config{
			Region: aws.String(region.ID),
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
			// fmt.Println("%+v", result)
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
					// XXX - need instance type, how long running etc.
					s.Ec2s = append(s.Ec2s, AdminAwsEc2{
						ID:       aws.StringValue(instance.InstanceId),
						Region:   region.ID,
						PublicIP: aws.StringValue(instance.PublicIpAddress),
						Name:     nt,
					})
				}
			}
		}
	}
	return true
}
