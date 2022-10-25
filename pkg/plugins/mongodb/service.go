// Copyright 2022 The ILLA Authors.
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

package mongodb

import (
	"github.com/illa-family/builder-backend/pkg/plugins/common"
)

type Connector struct {
	Resource Options
	Action   Query
}

func (m *Connector) ValidateResourceOptions(resourceOptions map[string]interface{}) (common.ValidateResult, error) {
	// format resource options
	//if err := mapstructure.Decode(resourceOptions, &m.Resource); err != nil {
	//	return common.ValidateResult{Valid: false}, err
	//}
	//
	//// validate mysql options
	//validate := validator.New()
	//if err := validate.Struct(m.Resource); err != nil {
	//	return common.ValidateResult{Valid: false}, err
	//}
	return common.ValidateResult{Valid: true}, nil
}

func (m *Connector) ValidateActionOptions(actionOptions map[string]interface{}) (common.ValidateResult, error) {

	return common.ValidateResult{Valid: true}, nil
}

func (m *Connector) TestConnection(resourceOptions map[string]interface{}) (common.ConnectionResult, error) {

	return common.ConnectionResult{Success: true}, nil
}

func (m *Connector) GetMetaInfo(resourceOptions map[string]interface{}) (common.MetaInfoResult, error) {

	return common.MetaInfoResult{
		Success: true,
		Schema:  nil,
	}, nil
}

func (m *Connector) Run(resourceOptions map[string]interface{}, actionOptions map[string]interface{}) (common.RuntimeResult, error) {

	return common.RuntimeResult{Success: false}, nil
}
