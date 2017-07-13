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
// Date			:	June 4, 2017
//
// History	:
// 	Date:			Author:		Info:
//	June 4, 2017	LIS			First Go release
//
// TODO:

package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

// Function to show how to setup the configuration file
func Setup() {
	fmt.Printf("%s", myInfo)
	fmt.Printf("Setup the configuration file:\n")
	fmt.Printf("# Create a configuration file, any name would do, as long its in yaml fornmat.\n")
	color.Green("\n\t%s:\n", myProgname)
	fmt.Printf("\t# Optional: add the following key/pair values, shown are the default values:\n")
	for key, val := range defaultValues {
		color.Green("\t  %s: %s\n", key, val)
	}
	fmt.Printf("\t# Required:\n")
	for idx := range requireKey {
		color.Yellow("\t  %s: <value>\n", requireKey[idx])
	}
	fmt.Printf("\nNOTE\n")
	color.HiRed("\t* The key must be all lowercase!\n")
	color.HiRed("\t* Any key value that contains any of these charaters: ':#[]()*' must be double quoted!\n")
	fmt.Printf("\n")
	os.Exit(0)
}

// Function to show the help information
func Help() {
	fmt.Printf("%s", myInfo)
	optionList := "<-config config file> <-setup> <-version> <-help>"
	fmt.Printf("\nUsage : %s\n\tflags: %s\n", myProgname, optionList)
	fmt.Printf("\t*config: the configuration file to use, should be full path, use --setup for more information.\n")
	fmt.Printf("\tsetup: show the setup guide.\n")
	fmt.Printf("\tversion: print %s version.\n", myProgname)
	fmt.Printf("\thelp: short version of this help page.\n")
	fmt.Printf("\n\t* == required flag.\n")
	os.Exit(1)
}
