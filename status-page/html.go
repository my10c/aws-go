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
//	Jul 13, 2017	LIS			Added some meta tag and auto-refresh
//
// TODO:

package main

import (
	"fmt"
	"sort"
	"time"
)

func htmlStatus(imageDir, cssFile string, refresh, height, width int, statusMap map[string][]*instStatus, cgiScript string) {
	// sort the main map
	var workList []string
	for key, _ := range statusMap {
		workList = append(workList, key)
	}
	sort.Strings(workList)
	// start of the html page generation
	htmlHeader(cssFile, refresh)
	for cnt := range workList {
		serviceName := workList[cnt]
		serviceStatus := statusMap[serviceName]
		// sort the instances in this service
		sort.Slice(serviceStatus, func(i, j int) bool {
			return serviceStatus[i].id < serviceStatus[j].id
		})
		htmlBody(serviceName, imageDir, height, width, serviceStatus)
	}
	htmlFooter(cgiScript)
}

func htmlHeader(cssFile string, refresh int) {
	fmt.Printf("<html>\n")
	fmt.Printf("<head>\n")
	fmt.Printf("<meta charset=\"UTF-8\">\n")
	fmt.Printf("<meta name=\"description\" content=\"Instance Health Status\">\n")
	fmt.Printf("<meta name=\"authors\" content=\"Luc Suryo\">\n")
	fmt.Printf("<meta http-equiv=\"refresh\" content=\"%d\">\n", refresh)
	// set to no cache for known browsers
	fmt.Printf("<meta http-equiv=\"cache-control\" content=\"max-age=0\"/>\n")
	fmt.Printf("<meta http-equiv=\"cache-control\" content=\"no-cache\"/>\n")
	fmt.Printf("<meta http-equiv=\"expires\" content=\"0\"/>\n")
	fmt.Printf("<meta http-equiv=\"expires\" content=\"Tue, 01 Jan 1980 1:00:00 GMT\"/>\n")
	fmt.Printf("<meta http-equiv=\"pragma\" content=\"no-cache\" />\n")
	// CSS
	fmt.Printf("<link rel=\"stylesheet\" href=\"%s\">\n", cssFile)
	fmt.Printf("</head>\n")
	fmt.Printf("<body>\n")
	// code to display last refresh
	fmt.Printf("<h5>Generated on %s\n", time.Now().Format(time.UnixDate))
	// auto refresh only if refresh was not set to 0
	if refresh > 0 {
		fmt.Printf("&nbsp;&nbsp;Last Refresh : <span id=\"date\" /><br>\n")
		fmt.Printf("<script>\n")
		fmt.Printf("\tdocument.getElementById(\"date\").innerHTML = Date();\n")
		fmt.Printf("</script></h5>\n")
	} else {
		fmt.Printf("</h5>\n")
	}
	// start of table
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
		fmt.Printf("<td>%d</td>\n", status[cnt].id)
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
	fmt.Printf("</tr>\n")
}

func htmlFooter(cgiScript string) {
	fmt.Printf("</tbody>\n")
	fmt.Printf("</table>\n")
	if len(cgiScript) > 0 {
		fmt.Printf("\n<form method=\"get\" action=\"%s\">\n", cgiScript)
		fmt.Printf("\t<input type=\"submit\" value=\"Update Now\" name=\"refresh\">\n")
		fmt.Printf("</form>\n\n")
	}
	fmt.Printf("</body>\n")
	fmt.Printf("</html>\n")
}
