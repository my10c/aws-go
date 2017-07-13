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
// Version		:	0.1
//
// Date			:	Jul 12, 2017
//
// History	:
// 	Date:			Author:		Info:
//	Jul 12, 2017	LIS			First relase
//
// TODO:

package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"github.com/my10c/aws-go/aws"
	"github.com/my10c/aws-go/ec2"
)

func ec2status(credFile, profileName, region string) map[string][]*instStatus {
	// preppare the session
	sess, err := aws.InitCredFile(credFile, profileName, region)
	if err != nil {
		fmt.Printf("Errored: %s\n", err.Error())
		os.Exit(1)
	}
	ec2client := ec2.New(sess)
	if ec2client == nil {
		fmt.Printf("Errored: in new ec2client\n")
		os.Exit(1)
	}
	// default to green
	currStatus := "green"
	statusMap := make(map[string][]*instStatus)
	// regex to separate service name and service-instance-id/cnt
	baseRegex := regexp.MustCompile("[0-9].*")
	idRegex := regexp.MustCompile("([A-Za-z]|\\-)*")
	// create the sttaus map so we can print a nice html
	for instanceID, InstanceStatus := range ec2client.Status {
		// set the current status
		if InstanceStatus["instanceStatus"] == "ok" && InstanceStatus["systemStatus"] != "ok" {
			currStatus = "yellow"
		}
		if InstanceStatus["instanceStatus"] != "ok" && InstanceStatus["systemStatus"] == "ok" {
			currStatus = "orange"
		}
		if InstanceStatus["instanceStatus"] != "ok" && InstanceStatus["systemStatus"] != "ok" {
			currStatus = "red"
		}
		tag := ec2client.TagFromID(instanceID)
		tagBase := baseRegex.ReplaceAllString(tag, "")
		tagID, err := strconv.Atoi(idRegex.ReplaceAllString(tag, ""))
		if err != nil {
			tagID = 0
		}
		currInstStatus := &instStatus{
			id:		tagID,
			status:	currStatus,
		}
		statusMap[tagBase] = append(statusMap[tagBase], currInstStatus)
	}
	return statusMap
}
