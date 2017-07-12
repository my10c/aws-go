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
//	Jul 12, 2017	LIS			Move to separate package/file and added more functions
//
// TODO:

package ec2

// Function to get Instance 'name' tag from the given instance ID
func (ec2Ptr *Ec2) TagFromID(instanceID string) string {
	if len(ec2Ptr.Info) == 0 {
		ec2Ptr.info()
	}
	return ec2Ptr.Info[instanceID]["instanceTag"]
}

// Function to get Tag from the given private IP
func (ec2Ptr *Ec2) TagFromPrivateIP(privateIP string) string {
	if len(ec2Ptr.Info) == 0 {
		ec2Ptr.info()
	}
	for key := range ec2Ptr.Info {
		if ec2Ptr.Info[key]["privateIP"] == privateIP {
			return  ec2Ptr.Info[key]["instanceTag"]
		}
	}
	return ""
}

// Function to get Tag from the given public IP
func (ec2Ptr *Ec2) TagFromPublicIP(publicIP string) string {
	if len(ec2Ptr.Info) == 0 {
		ec2Ptr.info()
	}
	for key := range ec2Ptr.Info {
		if ec2Ptr.Info[key]["publicIP"] == publicIP {
			return  ec2Ptr.Info[key]["instanceTag"]
		}
	}
	return ""
}
