// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
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
	"os"
	"runtime"
	"text/template"

	"github.com/calvn/brokr/buildtime"
	"github.com/spf13/cobra"
)

// VersionDetails holds build-related information
type VersionDetails struct {
	Version   string
	GitCommit string
	GoVersion string
	OS        string
	Arch      string
	BuildDate string
}

var versionTemplate = `brokr:
  Version:    {{.Version}}
  Git commit: {{.GitCommit}}
  Go version: {{.GoVersion}}
  OS/Arch:    {{.OS}}/{{.Arch}}
  Built:      {{.BuildDate}}
`

func newVersionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Display detailed brokr version information",
		Long:  `Display detailed brokr version information`,
		Run:   versionCmdFunc,
	}

	return cmd
}

func versionCmdFunc(cmd *cobra.Command, args []string) {
	vd := VersionDetails{
		Version:   buildtime.Version,
		GitCommit: buildtime.GitCommit,
		GoVersion: runtime.Version(),
		OS:        runtime.GOOS,
		Arch:      runtime.GOARCH,
		BuildDate: buildtime.BuildDate,
	}

	tmpl := template.Must(template.New("").Parse(versionTemplate))

	tmpl.Execute(os.Stdout, vd)
}
