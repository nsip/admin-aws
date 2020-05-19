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
			fmt.Printf("%+v\n", result)
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
						ID:           aws.StringValue(instance.InstanceId),
						Region:       region.ID,
						PublicIP:     aws.StringValue(instance.PublicIpAddress),
						InstanceType: aws.StringValue(instance.InstanceType),
						Name:         nt,
					})

					/*
					   EC2S=[{ID:i-0f758a06fcfd4bb3d Region:ap-southeast-2 Name:nias2 cli PublicIP:13.211.190.184 InstanceType:}
					   {ID:i-3765f9e8 Region:ap-southeast-2 Name:files - ownCloud PublicIP:54.66.143.89 InstanceType:}
					   {ID:i-03fd17c6682d1bf83 Region:ap-southeast-2 Name:hits.beta PublicIP:3.104.151.118 InstanceType:}
					   {ID:i-0b35689af2c9d0197 Region:ap-southeast-2 Name:hits.dev - HITS Dev Test Server PublicIP:13.210.63.151 InstanceType:}
					   {ID:i-6e8dbf50 Region:ap-southeast-2 Name:hits - HITS Server PublicIP:54.66.142.11 InstanceType:}]

					*/

				}
			}
		}
	}
	return true
}
