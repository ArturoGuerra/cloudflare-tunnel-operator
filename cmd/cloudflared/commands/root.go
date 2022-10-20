/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package commands

import (
	"github.com/spf13/cobra"

	// common imports for subcommands
	cmdgenerate "github.com/arturoguerra/cloudflare-tunnel-operator/cmd/cloudflared/commands/generate"
	cmdinit "github.com/arturoguerra/cloudflare-tunnel-operator/cmd/cloudflared/commands/init"
	cmdversion "github.com/arturoguerra/cloudflare-tunnel-operator/cmd/cloudflared/commands/version"

	// specific imports for workloads
	generateapps "github.com/arturoguerra/cloudflare-tunnel-operator/cmd/cloudflared/commands/generate/apps"
	initapps "github.com/arturoguerra/cloudflare-tunnel-operator/cmd/cloudflared/commands/init/apps"
	versionapps "github.com/arturoguerra/cloudflare-tunnel-operator/cmd/cloudflared/commands/version/apps"
	//+kubebuilder:scaffold:operator-builder:subcommands:imports
)

// CloudflaredCommand represents the base command when called without any subcommands.
type CloudflaredCommand struct {
	*cobra.Command
}

// NewCloudflaredCommand returns an instance of the CloudflaredCommand.
func NewCloudflaredCommand() *CloudflaredCommand {
	c := &CloudflaredCommand{
		Command: &cobra.Command{
			Use:   "cloudflared",
			Short: "Manage cloudflare tunnel workload",
			Long:  "Manage cloudflare tunnel workload",
		},
	}

	c.addSubCommands()

	return c
}

// Run represents the main entry point into the command
// This is called by main.main() to execute the root command.
func (c *CloudflaredCommand) Run() {
	cobra.CheckErr(c.Execute())
}

func (c *CloudflaredCommand) newInitSubCommand() {
	parentCommand := cmdinit.GetParent(c.Command)
	_ = parentCommand

	// add the init subcommands
	initapps.NewCloudflareTunnelSubCommand(parentCommand)
	//+kubebuilder:scaffold:operator-builder:subcommands:init
}

func (c *CloudflaredCommand) newGenerateSubCommand() {
	parentCommand := cmdgenerate.GetParent(c.Command)
	_ = parentCommand

	// add the generate subcommands
	generateapps.NewCloudflareTunnelSubCommand(parentCommand)
	//+kubebuilder:scaffold:operator-builder:subcommands:generate
}

func (c *CloudflaredCommand) newVersionSubCommand() {
	parentCommand := cmdversion.GetParent(c.Command)
	_ = parentCommand

	// add the version subcommands
	versionapps.NewCloudflareTunnelSubCommand(parentCommand)
	//+kubebuilder:scaffold:operator-builder:subcommands:version
}

// addSubCommands adds any additional subCommands to the root command.
func (c *CloudflaredCommand) addSubCommands() {
	c.newInitSubCommand()
	c.newGenerateSubCommand()
	c.newVersionSubCommand()
}
