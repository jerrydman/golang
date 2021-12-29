//Script to describe instance via golang
//Required Input is profile, region, and hostname
//Hostname is from a tag value and can be adjusted to anything

//Version 1
//Jerry Reid gerald.reid@gmail.com

package main

import (
	"fmt"
    "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/aws"

)

func DescribeInstances(client *ec2.EC2) (*ec2.DescribeInstancesOutput, error) {

	fmt.Println("Enter the Hostname of the server you need to describe")
	var host string
	fmt.Scanln(&host)



	result, err := client.DescribeInstances(&ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("tag:Hostname"),  //here is where you change the filter or add more
				Values: []*string{aws.String(host)},
			},
			{
				Name: aws.String("tag:PlatformOS"),
				Values: []*string{aws.String("centos")},
		     },
			},
		},
	)

	if err != nil {
		return nil, err
	}

	return result, err
}

func main() {

	fmt.Println("Enter AWS Profile name (dev,ssdev etc)") //this profile name should match ~/.aws/credentials
	var profile string
	fmt.Scanln(&profile)

	fmt.Println("Enter region (us-west-2)")
	var region string
	fmt.Scanln(&region)

	sess, err := session.NewSessionWithOptions(session.Options{
		Profile: profile,
		Config: aws.Config{
			Region: aws.String(region),
		},
	})

	if err != nil {
        fmt.Println(err)
    }

	ec2Client := ec2.New(sess)

	runningInstances, err := DescribeInstances(ec2Client)
	fmt.Println(runningInstances)

}
