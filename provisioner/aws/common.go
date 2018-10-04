// Copyright © 2018 The TK8 Authors.
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

package provisioner

import "github.com/kubernauts/tk8/internal/cluster"

type AWS struct {
}

func (p AWS) Init(args []string) {
	cluster.AWSCreate()
}

func (p AWS) Setup(args []string) {
	cluster.AWSInstall()

}

func (p AWS) Scale(args []string) {
	cluster.AWSScale()

}

func (p AWS) Reset(args []string) {
	cluster.AWSReset()

}

func (p AWS) Remove(args []string) {
	cluster.AWSRemove()

}

func (p AWS) Upgrade(args []string) {
	cluster.NotImplemented()
}

func (p AWS) Destroy(args []string) {
	cluster.AWSDestroy()
}

func NewAWS() cluster.Provisioner {
	cluster.SetClusterName()
	provisioner := new(AWS)
	return provisioner
}
