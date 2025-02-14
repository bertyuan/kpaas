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

package server

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/kpaas-io/kpaas/pkg/deploy/consts"
	pb "github.com/kpaas-io/kpaas/pkg/deploy/protos"
	"github.com/kpaas-io/kpaas/pkg/deploy/task"
)

type controller struct {
	store      task.Store
	logFileLoc string
}

func (c *controller) TestConnection(context.Context, *pb.TestConnectionRequest) (*pb.TestConnectionReply, error) {
	return nil, nil
}

func (c *controller) CheckNodes(ctx context.Context, req *pb.CheckNodesRequest) (*pb.CheckNodesReply, error) {
	logrus.Info("Begins CheckNodes request")

	taskName := getCheckNodeTaskName(req)
	taskConfig := &task.NodeCheckTaskConfig{
		NodeConfigs:     req.GetConfigs(),
		LogFileBasePath: c.logFileLoc,
	}

	nodeCheckTask, err := task.NewNodeCheckTask(taskName, taskConfig)
	if err == nil {
		// store and launch the task
		err = c.storeAndLanuchTask(nodeCheckTask)
	}
	if err != nil {
		logrus.Errorf("CheckNodes request failed: %s", err)
		return &pb.CheckNodesReply{
			Acceptd: false,
			Err: &pb.Error{
				Reason: consts.MsgRequestFailed,
				Detail: err.Error(),
			},
		}, err
	}

	logrus.Info("CheckNodes request succeeded")
	return &pb.CheckNodesReply{
		Acceptd: true,
		Err:     nil,
	}, nil
}

func (c *controller) GetCheckNodesResult(context.Context, *pb.GetCheckNodesResultRequest) (*pb.GetCheckNodesResultReply, error) {
	// TODO
	return nil, nil
}

func (c *controller) Deploy(context.Context, *pb.DeployRequest) (*pb.DeployReply, error) {
	// TODO
	return nil, nil
}

func (c *controller) GetDeployResult(context.Context, *pb.GetDeployResultRequest) (*pb.GetDeployResultReply, error) {
	// TODO
	return nil, nil
}

func (c *controller) storeTask(task task.Task) error {
	if c.store == nil {
		return fmt.Errorf("no task store")
	}

	return c.store.AddTask(task)
}

func (c *controller) storeAndLanuchTask(aTask task.Task) error {
	if c.store == nil {
		return fmt.Errorf("no task store")
	}

	// store the task
	if err := c.store.AddTask(aTask); err != nil {
		return err
	}

	// launch the task
	processor, err := task.NewProcessor(aTask.GetType())
	if err != nil {
		return err
	}
	return processor.StartTask(aTask)
}

func getCheckNodeTaskName(req *pb.CheckNodesRequest) string {
	// use a fixed name for checknode task, it may be changed in the future
	return "node-check"
}
