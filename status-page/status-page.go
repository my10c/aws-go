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
// Date			:	Jul 13, 2017
//
// History	:
// 	Date:			Author:		Info:
//	Jul 12, 2017	LIS			First relase
//	Jul 13, 2017	LIS			Added refresh and css configs
//
// TODO:

package main

import (
	"os"
	"strconv"
)

type instStatus struct {
	id		int
	status	string
}

func main() {
	cfgMap := Init()
	statusMap := ec2status(cfgMap["aws_file"], cfgMap["aws_profile"], cfgMap["aws_region"])
	cfgWidth, _ := strconv.Atoi(cfgMap["image_width"])
	cfgHeight, _ := strconv.Atoi(cfgMap["image_height"])
	cfgRefresh, _ := strconv.Atoi(cfgMap["refresh_secs"])
	var cgiScript string
	if cfgMap["allow_cgi"] == "yes" {
		cgiScript = cfgMap["cgi_script"]
	}
	if cfgMap["mode"] == "text" {
		textStatus(statusMap)
	}	else {
		htmlStatus(cfgMap["imaged_dir"], cfgMap["css_file"], cfgRefresh, cfgWidth, cfgHeight, statusMap, cgiScript)
	}
	os.Exit(0)
}
