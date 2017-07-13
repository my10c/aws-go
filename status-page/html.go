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
	"sort"
)

func htmlStatus(imageDir string, height, width int, statusMap map[string][]*instStatus) {
	// sort the main map
	var workList []string
	for key, _ := range statusMap {
		workList = append(workList, key)
	}
	sort.Strings(workList)
	// start of the html page generation
	htmlHeader()
	for cnt := range workList {
		serviceName := workList[cnt]
		serviceStatus := statusMap[serviceName]
		// sort the instances in this service
		sort.Slice(serviceStatus, func(i, j int) bool {
			return serviceStatus[i].id < serviceStatus[j].id
		})
		htmlBody(serviceName, imageDir, height, width, serviceStatus)
	}
	htmlFooter()
}

func htmlHeader() {
	fmt.Printf("<html>\n")
	fmt.Printf("<table>\n")
	fmt.Printf("<thead>\n")
	fmt.Printf("<tr>\n")
	fmt.Printf("<td>Instance Tag</td>\n")
	fmt.Printf("<td>instance Status</td>\n")
	fmt.Printf("<td></td>\n")
	fmt.Printf("</tr>\n")
	fmt.Printf("</thead>\n")
	fmt.Printf("<tbody>\n")
}

func htmlBody(tagBase, imageDir string, height, width int, status []*instStatus) {
	fmt.Printf("<tr>\n")
	fmt.Printf("<td>%s</td>\n", tagBase)
	fmt.Printf("<td></td>\n")
	for cnt, _ := range status {
		fmt.Printf("<td>%d></td>\n", status[cnt].id)
		switch status[cnt].status {
			case "green":
				fmt.Printf("<td><img src=\"%s/green.jpg\" height=\"%d\" width=\"%d\"/></td>\n", imageDir, height, width)
			case "yellow":
				fmt.Printf("<td><img src=\"%s/yellow.jpg\" height=\"%d\" width=\"%d\"/></td>\n", imageDir, height, width)
			case "orange-2":
				fmt.Printf("<td><img src=\"%s/orange.jpg\" height=\"%d\" width=\"%d\"/></td>\n", imageDir, height, width)
			case "red":
				fmt.Printf("<td><img src=\"%s/green.jpg\" height=\"%d\" width=\"%d\"/></td>\n", imageDir, height, width)
		}
	}
	fmt.Printf("<tr>\n")
}

func htmlFooter() {
	fmt.Printf("</tbody>\n")
	fmt.Printf("</table>\n")
	fmt.Printf("</html>\n")
}
