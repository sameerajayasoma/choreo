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

package deployment

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/wso2-enterprise/choreo-cp-declarative-api/internal/choreoctl/errors"
	"github.com/wso2-enterprise/choreo-cp-declarative-api/internal/choreoctl/interactive"
	"github.com/wso2-enterprise/choreo-cp-declarative-api/internal/choreoctl/util"
	"github.com/wso2-enterprise/choreo-cp-declarative-api/pkg/cli/common/constants"
	"github.com/wso2-enterprise/choreo-cp-declarative-api/pkg/cli/types/api"
)

const (
	stateOrgSelect = iota
	stateProjSelect
	stateCompSelect
	stateEnvSelect
)

type deploymentListModel struct {
	interactive.BaseModel // Reuses Organizations, Projects, Components, Environments and their cursors.
	state                 int
	selected              bool
	errorMsg              string
}

func (m deploymentListModel) Init() tea.Cmd {
	return nil
}

func (m deploymentListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	keyMsg, ok := msg.(tea.KeyMsg)
	if !ok {
		return m, nil
	}

	if interactive.IsQuitKey(keyMsg) {
		m.selected = false
		return m, tea.Quit
	}

	switch m.state {
	case stateOrgSelect:
		if interactive.IsEnterKey(keyMsg) {
			projects, err := m.FetchProjects() // Reusable function from BaseModel.
			if err != nil {
				m.errorMsg = err.Error()
				return m, nil
			}
			if len(projects) == 0 {
				m.errorMsg = fmt.Sprintf("No projects found for organization: %s", m.Organizations[m.OrgCursor])
				return m, nil
			}
			m.Projects = projects
			m.state = stateProjSelect
			m.errorMsg = ""
			return m, nil
		}
		m.OrgCursor = interactive.ProcessListCursor(keyMsg, m.OrgCursor, len(m.Organizations))

	case stateProjSelect:
		if interactive.IsEnterKey(keyMsg) {
			components, err := m.FetchComponents() // Reusable function from BaseModel.
			if err != nil {
				m.errorMsg = err.Error()
				return m, nil
			}
			if len(components) == 0 {
				m.errorMsg = fmt.Sprintf("No components found for project: %s", m.Projects[m.ProjCursor])
				return m, nil
			}
			m.Components = components
			m.state = stateCompSelect
			m.errorMsg = ""
			return m, nil
		}
		m.ProjCursor = interactive.ProcessListCursor(keyMsg, m.ProjCursor, len(m.Projects))

	case stateCompSelect:
		if interactive.IsEnterKey(keyMsg) {
			environments, err := m.FetchEnvironments() // Reusable function from BaseModel.
			if err != nil {
				m.errorMsg = err.Error()
				return m, nil
			}
			if len(environments) == 0 {
				m.errorMsg = fmt.Sprintf("No environments found for organization: %s", m.Organizations[m.OrgCursor])
				return m, nil
			}
			m.Environments = environments
			m.state = stateEnvSelect
			m.errorMsg = ""
			return m, nil
		}
		m.CompCursor = interactive.ProcessListCursor(keyMsg, m.CompCursor, len(m.Components))

	case stateEnvSelect:
		if interactive.IsEnterKey(keyMsg) {
			m.selected = true
			return m, tea.Quit
		}
		m.EnvCursor = interactive.ProcessListCursor(keyMsg, m.EnvCursor, len(m.Environments))
	}

	return m, nil
}

func (m deploymentListModel) View() string {
	var progress string
	if m.state > stateOrgSelect {
		progress += fmt.Sprintf("Organization: %s\n", m.Organizations[m.OrgCursor])
	}
	if m.state > stateProjSelect {
		progress += fmt.Sprintf("Project: %s\n", m.Projects[m.ProjCursor])
	}
	if m.state > stateCompSelect {
		progress += fmt.Sprintf("Component: %s\n", m.Components[m.CompCursor])
	}

	var view string
	switch m.state {
	case stateOrgSelect:
		view = interactive.RenderListPrompt("Select organization:", m.Organizations, m.OrgCursor)
	case stateProjSelect:
		view = interactive.RenderListPrompt("Select project:", m.Projects, m.ProjCursor)
	case stateCompSelect:
		view = interactive.RenderListPrompt("Select component:", m.Components, m.CompCursor)
	case stateEnvSelect:
		view = interactive.RenderListPrompt("Select environment:", m.Environments, m.EnvCursor)
	default:
		view = ""
	}

	if m.errorMsg != "" {
		view += "\nError: " + m.errorMsg
	}

	return progress + view
}

func listDeploymentInteractive(config constants.CRDConfig) error {
	orgs, err := util.GetOrganizationNames()
	if err != nil {
		return errors.NewError("failed to get organizations: %v", err)
	}
	if len(orgs) == 0 {
		return errors.NewError("no organizations found")
	}

	model := deploymentListModel{
		BaseModel: interactive.BaseModel{
			Organizations: orgs,
		},
		state: stateOrgSelect,
	}

	finalModel, err := interactive.RunInteractiveModel(model)
	if err != nil {
		return errors.NewError("interactive mode failed: %v", err)
	}

	m, ok := finalModel.(deploymentListModel)
	if !ok || !m.selected {
		return errors.NewError("deployment listing cancelled")
	}

	params := api.ListDeploymentParams{
		Organization: m.Organizations[m.OrgCursor],
		Project:      m.Projects[m.ProjCursor],
		Component:    m.Components[m.CompCursor],
		Environment:  m.Environments[m.EnvCursor],
	}

	return listDeployments(params, config)
}
