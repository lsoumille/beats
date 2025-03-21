// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package cmd

import (
	"flag"

	"github.com/elastic/beats/v7/libbeat/cfgfile"
	"github.com/elastic/beats/v7/x-pack/functionbeat/function/beater"
	funcmd "github.com/elastic/beats/v7/x-pack/functionbeat/function/cmd"
)

// Name of this beat
var Name = "functionbeat"

// RootCmd to handle functionbeat
var RootCmd *funcmd.FunctionCmd

func init() {
	RootCmd = funcmd.NewFunctionCmd(Name, beater.New)
	RootCmd.PersistentFlags().AddGoFlag(flag.CommandLine.Lookup("d"))
	cfgfile.AddAllowedBackwardsCompatibleFlag("d")
	RootCmd.PersistentFlags().AddGoFlag(flag.CommandLine.Lookup("v"))
	cfgfile.AddAllowedBackwardsCompatibleFlag("v")
	RootCmd.PersistentFlags().AddGoFlag(flag.CommandLine.Lookup("e"))
	cfgfile.AddAllowedBackwardsCompatibleFlag("e")
}
