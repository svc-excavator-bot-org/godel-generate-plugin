// Copyright (c) 2018 Palantir Technologies Inc. All rights reserved.
// Use of this source code is governed by the Apache License, Version 2.0
// that can be found in the LICENSE file.

package cmd

import (
	"github.com/palantir/godel/framework/pluginapi/v2/pluginapi"
	"github.com/palantir/godel/framework/verifyorder"
)

var (
	Version    = "unspecified"
	PluginInfo = pluginapi.MustNewPluginInfo(
		"com.palantir.go-generate",
		"generate-plugin",
		Version,
		pluginapi.PluginInfoUsesConfigFile(),
		pluginapi.PluginInfoGlobalFlagOptions(
			pluginapi.GlobalFlagOptionsParamDebugFlag("--"+pluginapi.DebugFlagName),
			pluginapi.GlobalFlagOptionsParamProjectDirFlag("--"+pluginapi.ProjectDirFlagName),
			pluginapi.GlobalFlagOptionsParamConfigFlag("--"+pluginapi.ConfigFlagName),
		),
		pluginapi.PluginInfoTaskInfo(
			"generate",
			"Run generate task",
			pluginapi.TaskInfoCommand("run"),
			pluginapi.TaskInfoVerifyOptions(
				pluginapi.VerifyOptionsOrdering(intVar(verifyorder.Generate)),
				pluginapi.VerifyOptionsApplyFalseArgs("--verify"),
			),
		),
		pluginapi.PluginInfoUpgradeConfigTaskInfo(
			pluginapi.UpgradeConfigTaskInfoCommand("upgrade-config"),
			pluginapi.LegacyConfigFile("generate.yml"),
		),
	)
)

func intVar(val int) *int {
	return &val
}
