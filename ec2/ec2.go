// Copyright (c) 2017 Badassops
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//	* Redistributions of source code must retain the above copyright
//	notice, this list of conditions and the following disclaimer.
//	* Redistributions in binary form must reproduce the above copyright
//	notice, this list of conditions and the following disclaimer in the
//	documentation and/or other materials provided with the distribution.
//	* Neither the name of the <organization> nor the
//	names of its contributors may be used to endorse or promote products
//	derived from this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSEcw
// ARE DISCLAIMED. IN NO EVENT SHALL <COPYRIGHT HOLDER> BE LIABLE FOR ANY
// DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
// (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
// LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
// ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
// SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
// Author		:	Luc Suryo <luc@badassops.com>
//
// Version		:	0.2
//
// Date			:	Jul 12, 2017
//
// History	:
// 	Date:			Author:		Info:
//	Jan 30, 2017	LIS			First relase
//	Jul 12, 2017	LIS			moved to seprate package and added status
//
// TODO:

package ec2

import (
	"strings"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	myUtils "github.com/my10c/utils-go"
)

type Ec2 struct {
	session		*ec2.EC2
	Info		map[string]map[string]string
	Status		map[string]map[string]string
	mu			sync.Mutex
}

// Function to create a new Ec2 object with a AWS session set
func New(awsSess *session.Session) *Ec2 {
	ec2PTR := &Ec2{
		session: ec2Session(awsSess),
	}
	ec2PTR.info()
	ec2PTR.status()
	return ec2PTR
}

// Function to get a ec2 session
func ec2Session(sess *session.Session) *ec2.EC2 {
	ec2Sess := ec2.New(sess)
	myUtils.ExitIfNill(ec2Sess)
	return ec2Sess
}

// Function set all instanace info to the Ec2 object
func (ec2Ptr *Ec2) info() {
	ec2Ptr.mu.Lock()
	defer ec2Ptr.mu.Unlock()
	infoMap := make(map[string]map[string]string)
	// Only running instances!
	params := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("instance-state-code"),
				Values: []*string{
					aws.String("16"),
				},
			},
		},
		MaxResults: aws.Int64(1024),
	}
	resp, err := ec2Ptr.session.DescribeInstances(params)
	if err != nil {
		myUtils.ExitIfError(err)
	}
	for idx := range resp.Reservations {
		for instIdx := range resp.Reservations[idx].Instances {
			instanceID := *resp.Reservations[idx].Instances[instIdx].InstanceId
			instancePrvIP := *resp.Reservations[idx].Instances[instIdx].PrivateIpAddress
			instancePubIP := *resp.Reservations[idx].Instances[instIdx].PublicIpAddress
			var instanceTag string
			for tagIdx := range resp.Reservations[idx].Instances[instIdx].Tags {
				if *resp.Reservations[idx].Instances[instIdx].Tags[tagIdx].Key == "Name" {
					instanceTag = *resp.Reservations[idx].Instances[instIdx].Tags[tagIdx].Value
				}
			}
			// trim the tag, all lowercase and remove the indenty part
			instanceTag = strings.ToLower(strings.Split(instanceTag, " ")[0])
			infoMap[instanceID] = make(map[string]string)
			infoMap[instanceID]["instanceTag"] = instanceTag
			infoMap[instanceID]["privateIP"] = instancePrvIP
			infoMap[instanceID]["publicIP"] = instancePubIP
		}
	}
	ec2Ptr.Info = infoMap
}

// Function set all instanace status to the Ec2 object
func (ec2Ptr *Ec2) status() {
	ec2Ptr.mu.Lock()
	defer ec2Ptr.mu.Unlock()
	statusMap := make(map[string]map[string]string)
	// Only running instances!
	params := &ec2.DescribeInstanceStatusInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("instance-state-code"),
				Values: []*string{
					aws.String("16"),
				},
			},
		},
		MaxResults: aws.Int64(1024),
	}
	resp, err := ec2Ptr.session.DescribeInstanceStatus(params)
	if err != nil {
		myUtils.ExitIfError(err)
	}
	for _, instanace  := range resp.InstanceStatuses {
		instanceID := *instanace.InstanceId
		instanceStatus := *instanace.InstanceStatus.Status
		systemStatus := *instanace.SystemStatus.Status
		statusMap[instanceID] = make(map[string]string)
		statusMap[instanceID]["instanceStatus"] = instanceStatus
		statusMap[instanceID]["systemStatus"] = systemStatus
	}
	ec2Ptr.Status = statusMap
}
