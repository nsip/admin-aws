package main

// Use api and Test
import (
	"fmt"

	"github.com/nsip/admin-aws"
)

func main() {
	doc := adminaws.New()
	doc.UpdateEc2()
	fmt.Printf("%+v\n", doc.Regions)
	fmt.Printf("%+v\n", doc.Ec2s)
}
