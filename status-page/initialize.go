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
	"flag"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/my10c/simpleyaml"
	"github.com/my10c/utils-go"
)

// type used for flags in initArgs
type stringFlag struct {
	value	string
	set		bool
}

// Function for the stringFlag struct, set the values
func (sf *stringFlag) Set(x string) error {
	sf.value = x
	sf.set = true
	return nil
}

// Function for the stringFlag struct, get the values
func (sf *stringFlag) String() string {
	return sf.value
}

// Function to return the yaml value, nil if error or nil if not found
func getYamlValue(yamFile *simpleyaml.Yaml, section string, key string) (string, error) {
	// Check if section exist and/or key, no point to go further if it doesn't exist
	keyExist := yamFile.GetPath(section, key)
	if keyExist.IsFound() == false {
		err := fmt.Errorf("Section %s and/or key %s not found\n", section, key)
		return "", err
	}
	// We need to ge the value and since we do not know what it is, we check
	// against the 3 supported type
	// check if value is a string
	if value, err := yamFile.Get(section).Get(key).String(); err == nil {
		return value, err
	}
	// check if value is a int
	if value, err := yamFile.Get(section).Get(key).Int(); err == nil {
		return strconv.Itoa(value), err
	}
	// check if value is a boolean
	if value, err := yamFile.Get(section).Get(key).Bool(); err == nil {
		return strconv.FormatBool(value), err
	}
	// we got here so this is neither a string, int or boolean
	err := fmt.Errorf("Unsupported value for section %s and key %s, suported are: string, int and bool\n", section, key)
	return "", err
}

// Function to get the configuration
func config(argv...string) map[string]string {
	// working variable
	var missingKeys []string
	dictCfg := make(map[string]string)
	// open given file and check that is a correct yaml file
	cfgFile, err := ioutil.ReadFile(argv[0])
	utils.ExitIfError(err)
	yamlFile, err := simpleyaml.NewYaml(cfgFile)
	utils.ExitIfError(err)
	// read the values from file, overwrite is found
	for key, _ := range defaultValues {
		if val, err := getYamlValue(yamlFile, myProgname, key); err == nil {
			dictCfg[key] = val
		} else {
			dictCfg[key] = defaultValues[key]
		}
	}
	for idx := range requireKey {
		keyName := requireKey[idx]
		if val, err := getYamlValue(yamlFile, myProgname, keyName); err == nil {
			dictCfg[keyName] = val
		} else {
			missingKeys = append(missingKeys, keyName)
		}
	}
	// make sure we have all required configs
	if len(missingKeys) != 0 {
		fmt.Printf("Following keys are missing in the configration files: %s\n", missingKeys)
		os.Exit(2)
	}
	return dictCfg
}

// Function to process the given args
func Init() map[string]string {
	var myConfigFile stringFlag
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "%s\n", myInfo)
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", myProgname)
		flag.PrintDefaults()
	}
	version := flag.Bool("version", false, "Prints current version and exit.")
	setup := flag.Bool("setup", false, "Show the setup information and exit.")
	flag.Var(&myConfigFile, "config", "Configuration file to be used.")
	flag.Parse()
	if *version {
		fmt.Printf("%s\n", myVersion)
		os.Exit(0)
	}
	if *setup {
		Setup()
	}
	// if not set we use the default
	if !myConfigFile.set{
		Help()
	}
	return config(myConfigFile.value)
}
