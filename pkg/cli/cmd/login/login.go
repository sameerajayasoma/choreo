/*
 * Copyright (c) 2025, WSO2 Inc. (http://www.wso2.org) All Rights Reserved.
 *
 * WSO2 Inc. licenses this file to you under the Apache License,
 * Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package login

import (
	"github.com/spf13/cobra"

	"github.com/wso2-enterprise/choreo-cp-declarative-api/pkg/cli/common/builder"
	"github.com/wso2-enterprise/choreo-cp-declarative-api/pkg/cli/common/constants"
	"github.com/wso2-enterprise/choreo-cp-declarative-api/pkg/cli/flags"
	"github.com/wso2-enterprise/choreo-cp-declarative-api/pkg/cli/types/api"
)

// NewLoginCmd creates the login command.
func NewLoginCmd(impl api.CommandImplementationInterface) *cobra.Command {
	return (&builder.CommandBuilder{
		Command: constants.Login,
		Flags: []flags.Flag{
			flags.Kubeconfig,
			flags.KubeContext,
		},
		RunE: func(fg *builder.FlagGetter) error {
			return impl.Login(api.LoginParams{
				KubeconfigPath: fg.GetString(flags.Kubeconfig),
				Kubecontext:    fg.GetString(flags.KubeContext),
			})
		},
	}).Build()
}
