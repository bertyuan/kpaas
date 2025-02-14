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

package deploy

import (
	"github.com/gin-gonic/gin"
)

// Service for wizard progress

// @ID GetWizardProgress
// @Summary Get all of current deploy wizard data
// @Description Get all data, include current progress, cluster and node data. deploying progress or error.
// @Tags wizard
// @Produce application/json
// @Success 200 {object} api.GetWizardResponse
// @Router /api/v1/deploy/wizard/progresses [get]
func GetWizardProgress(c *gin.Context) {

}

// @ID ClearWizard
// @Summary Clear all of current deploy wizard data
// @Description Clear all data, include current progress, cluster and node data. deploying progress or error.
// @Tags wizard
// @Success 204
// @Router /api/v1/deploy/wizard/progresses [delete]
func ClearWizard(c *gin.Context) {

}
