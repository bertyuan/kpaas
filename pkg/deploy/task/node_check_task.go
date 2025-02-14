// Copyright 2019 Shanghai JingDuo Information Technology co., Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package task

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"

	pb "github.com/kpaas-io/kpaas/pkg/deploy/protos"
)

// NodeCheckTaskConfig represents the config for a node check task.
type NodeCheckTaskConfig struct {
	NodeConfigs     []*pb.NodeCheckConfig
	LogFileBasePath string
}

type nodeCheckTask struct {
	base
	nodeConfigs     []*pb.NodeCheckConfig
	logFileBasePath string
}

// NewNodeCheckTask returns a node check task based on the config.
// User should use this function to create a node check task.
func NewNodeCheckTask(taskName string, taskConfig *NodeCheckTaskConfig) (Task, error) {
	var err error
	if taskConfig == nil {
		err = fmt.Errorf("Invalid task config: nil")

	} else if len(taskConfig.NodeConfigs) == 0 {
		err = fmt.Errorf("Invalid task config: node configs is empty")

	}

	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	task := &nodeCheckTask{
		base: base{
			name:              taskName,
			taskType:          TaskTypeNodeCheck,
			status:            TaskPending,
			logFilePath:       GenTaskLogFilePath(taskConfig.LogFileBasePath, taskName),
			creationTimestamp: time.Now(),
		},
		nodeConfigs: taskConfig.NodeConfigs,
	}

	return task, nil
}
