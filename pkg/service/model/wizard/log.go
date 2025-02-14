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

package wizard

import (
	"bytes"
	"io"
	"io/ioutil"
	"time"

	"github.com/sony/sonyflake"
)

const (
	ServiceNodeID = 0
)

var (
	logs         map[uint64][]byte // Key Log Id, Value Log detail
	logIdCreator *sonyflake.Sonyflake
)

func init() {

	InitLogs()
}

func SetLogByReader(reader io.Reader) (logId uint64, err error) {

	logId, err = newLogId()
	if err != nil {
		return
	}
	logs[logId], err = ioutil.ReadAll(reader)
	return
}

func SetLogByString(content string) (logId uint64, err error) {

	logId, err = newLogId()
	if err != nil {
		return
	}

	logs[logId] = []byte(content)
	return
}

func GetLog(logId uint64) []byte {

	if content, exist := logs[logId]; exist {
		return content
	}

	return nil
}

func GetLogReader(logId uint64) io.ReadCloser {

	var content []byte
	var exist bool
	if content, exist = logs[logId]; !exist {
		return nil
	}

	return ioutil.NopCloser(bytes.NewReader(content))
}

func InitLogs() {

	logs = make(map[uint64][]byte)
	logIdCreator = sonyflake.NewSonyflake(
		sonyflake.Settings{
			StartTime: time.Now(),
			MachineID: func() (u uint16, e error) {
				return ServiceNodeID, nil
			},
		})
}

func newLogId() (uint64, error) {

	return logIdCreator.NextID()
}
