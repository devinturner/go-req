// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/devinturner/go-req/request"
	"github.com/devinturner/go-req/response"
	"github.com/devinturner/go-req/util"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "perform GET request",
	Run:   get,
}

func init() {
	rootCmd.AddCommand(getCmd)

	getCmd.Flags().StringSliceP("header", "H", []string{""}, "pass in one or more custom headers as key-value pairs")
}

func get(cmd *cobra.Command, args []string) {
	headers, err := cmd.Flags().GetStringSlice("header")
	if err != nil {
		log.Fatal(err)
	}

	// Extract uri from args
	uri, err := util.GetURI(args)
	if err != nil {
		log.Fatal(err)
	}

	// Perform a GET request
	req := request.NewRequest(headers)
	raw, err := req.Get(uri)
	if err != nil {
		log.Fatal(err)
	}

	// Parse the response
	res := response.ParseResponse(raw)

	// Marshal and print response as JSON
	b, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}
