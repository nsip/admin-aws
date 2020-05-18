package main

// Use api and Test
import (
	"fmt"

	"github.com/nsip/admin-aws"
)

func main() {
	doc := adminaws.New()
	doc.UpdateEc2()
	fmt.Printf("REGIONS=%+v\n\n\n", doc.Regions)
	fmt.Printf("EC2S=%+v\n\n\n", doc.Ec2s)
}
