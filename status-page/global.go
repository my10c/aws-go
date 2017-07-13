// Copyright (c) 2017 - 2017 badassops
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
// Version		:	0.1
//
// Date			:	Jul 12, 2017
//
// History	:
// 	Date:			Author:		Info:
//	Jul 12, 2017	LIS			First release
//
// TODO:

package main

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"time"
)

const (
	width		int = 20
	height		int = 20
	imgDir		string = "images"
	awsProfile	string = "default"
	awsRegion	string = "us-east-1"
)

var (
	myVersion   = "0.1"
	now         = time.Now()
	myProgname  = path.Base(os.Args[0])
	myAuthor    = "Luc Suryo"
	myCopyright = "Copyright 2017 - " + strconv.Itoa(now.Year()) + " ©badassops"
	myLicense   = "License BSD, http://www.freebsd.org/copyright/freebsd-license.html ♥"
	myEmail     = "<luc@badassops.com>"

	// set the check info, we do not show the version as teh global version is the framework
	// version and not the check's
	myInfo = fmt.Sprintf("%s\n%s\n%s\nWritten by %s %s\n",
		myProgname, myCopyright, myLicense, myAuthor, myEmail)

	// Global variables
	defaultValues map[string]string
	confFile string
	requireKey = []string{"aws_file"}
)

func init() {
	// setup the default value, these are hardcoded.
	defaultValues = make(map[string]string)
	defaultValues["image_width"]	= strconv.Itoa(width)
	defaultValues["image_height"]	= strconv.Itoa(height)
	defaultValues["imaged_dir"]		= imgDir
	defaultValues["aws_profile"]	= awsProfile
	defaultValues["aws_region"]		= awsRegion
}
