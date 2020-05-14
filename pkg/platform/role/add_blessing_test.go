// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package role

import (
	"errors"
	"reflect"
	"testing"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/mock"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"

	"github.com/elastic/ecctl/pkg/util"
)

func TestAddBlessing(t *testing.T) {
	type args struct {
		params AddBlessingParams
	}
	tests := []struct {
		name string
		args args
		err  error
	}{
		{
			name: "fails on parameter validation",
			args: args{},
			err: multierror.NewPrefixed("role add blessing",
				util.ErrAPIReq,
				errors.New("blessing definition cannot be empty"),
				errors.New("id cannot be empty"),
				errors.New("runner id cannot be empty"),
			),
		},
		{
			name: "fails updating the role",
			args: args{params: AddBlessingParams{
				API: api.NewMock(mock.New500Response(mock.NewStringBody(
					`{"error": "failed updating role"}`,
				))),
				Blessing: &models.Blessing{},
				RunnerID: "some",
				ID:       "one",
			}},
			err: errors.New(`{"error": "failed updating role"}`),
		},
		{
			name: "succeeds",
			args: args{params: AddBlessingParams{
				API:      api.NewMock(mock.New200Response(mock.NewStringBody(""))),
				Blessing: &models.Blessing{},
				ID:       "one",
				RunnerID: "some",
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := AddBlessing(tt.args.params); !reflect.DeepEqual(err, tt.err) {
				t.Errorf("AddBlessing() error = %v, wantErr %v", err, tt.err)
			}
		})
	}
}
