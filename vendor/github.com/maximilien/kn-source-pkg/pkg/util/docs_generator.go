// Copyright © 2020 The Knative Authors
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

package util

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

// DocGenerator is to append the usage of command to a file
func DocGenerator(cmd *cobra.Command, filepath string, prefix string) error {
	writer := bytes.NewBufferString("")
	err := docGenerator(cmd, writer)
	if err != nil {
		return err
	}

	mdFile, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer mdFile.Close()

	if _, err := io.WriteString(mdFile, prefix+"\n"); err != nil {
		return err
	}
	if _, err = io.WriteString(mdFile, "## Usage\n"); err != nil {
		return err
	}

	reader := bufio.NewReader(writer)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if strings.Contains(line, "##") {
			line = strings.Replace(line, "##", "###", 1)
		}
		_, err = io.WriteString(mdFile, line)
		if err != nil {
			return err
		}
	}
	return nil
}

func docGenerator(cmd *cobra.Command, w io.Writer) error {
	if err := doc.GenMarkdownCustom(cmd, w, linkConverter); err != nil {
		return err
	}
	for _, c := range cmd.Commands() {
		if !c.IsAvailableCommand() || c.IsAdditionalHelpTopicCommand() {
			continue
		}
		if err := docGenerator(c, w); err != nil {
			return err
		}
	}
	return nil
}

func linkConverter(ins string) string {
	outs := strings.TrimSuffix(ins, ".md")
	outs = strings.Replace(outs, "_", "-", -1)
	return "#" + outs
}
