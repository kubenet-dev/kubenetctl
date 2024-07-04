/*
Copyright 2024 Nokia.

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

package setupcmd

import (
	"context"

	//docs "github.com/pkgserver-dev/pkgserver/internal/docs/generated/initdocs"

	"github.com/kubenet-dev/kubenetctl/pkg/run"
	"github.com/spf13/cobra"
)

func NewCommand(ctx context.Context, version string) *cobra.Command {
	return NewRunner(ctx, version).Command
}

// NewRunner returns a command runner.
func NewRunner(ctx context.Context, version string) *Runner {
	r := &Runner{}
	cmd := &cobra.Command{
		Use:  "setup [flags]",
		Args: cobra.ExactArgs(0),
		//Short:   docs.InitShort,
		//Long:    docs.InitShort + "\n" + docs.InitLong,
		//Example: docs.InitExamples,
		PreRunE: r.preRunE,
		RunE:    r.runE,
	}

	r.Command = cmd

	return r
}

type Runner struct {
	Command *cobra.Command
}

func (r *Runner) preRunE(_ *cobra.Command, _ []string) error {
	return nil
}

func (r *Runner) runE(c *cobra.Command, args []string) error {
	ctx := c.Context()
	//log := log.FromContext(ctx)
	//log.Info("create packagerevision", "src", args[0], "dst", args[1])

	x := run.NewRun("Setup kubenet Environment")

	x.Step(
		run.S("create k8s kind cluster"),
		run.S(`kind create cluster --name dummy --config=<(cat <<EOF
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
networking:
  ipFamily: dual
EOF
			)`),
	)

	x.Step(
		run.S("Allow the kind cluster to communicate with the containerlab topology (clab will be created in a later step)"),
		run.S("sudo iptables -I DOCKER-USER -o br-$(docker network inspect -f '{{ printf \"%.12s\" .ID }}' kind) -j ACCEPT"),
	)

	x.Step(
		run.S("Deploy Containerlab topology"),
		run.S("sudo containerlab deploy -t https://raw.githubusercontent.com/kubenet-dev/kubenet/v0.0.1/lab/3node.yaml --reconfigure"),
	)

	return x.Run(ctx)
}

/*

kind create cluster --name dummy --config=<(cat <<EOF
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
networking:
  ipFamily: dual
EOF
)
*/