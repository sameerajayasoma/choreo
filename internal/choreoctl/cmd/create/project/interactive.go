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

package project

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/wso2-enterprise/choreo-cp-declarative-api/internal/choreoctl/errors"
	"github.com/wso2-enterprise/choreo-cp-declarative-api/internal/choreoctl/interactive"
	"github.com/wso2-enterprise/choreo-cp-declarative-api/internal/choreoctl/util"
	"github.com/wso2-enterprise/choreo-cp-declarative-api/pkg/cli/types/api"
)

const (
	stateOrgSelect = iota
	stateNameInput
	stateDisplayNameInput
)

type projectModel struct {
	interactive.BaseModel // Embeds common organization selection logic
	Name                  string
	DisplayName           string
	state                 int
	selected              bool
	errorMsg              string
}

func (m projectModel) Init() tea.Cmd {
	return nil
}

func (m projectModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	keyMsg, ok := msg.(tea.KeyMsg)
	if !ok {
		return m, nil
	}

	// Quit if needed.
	if interactive.IsQuitKey(keyMsg) {
		m.selected = false
		return m, tea.Quit
	}

	switch m.state {
	case stateOrgSelect:
		// Delegate organization selection to BaseModel helper.
		if interactive.IsEnterKey(keyMsg) {
			cmd := m.UpdateOrgSelect(keyMsg)
			// Once organization selection is complete, transition to name input.
			m.state = stateNameInput
			m.errorMsg = ""
			return m, cmd
		}
		m.OrgCursor = interactive.ProcessListCursor(keyMsg, m.OrgCursor, len(m.Organizations))
	case stateNameInput:
		if interactive.IsEnterKey(keyMsg) {
			if err := util.ValidateProject(m.Name); err != nil {
				m.errorMsg = err.Error()
				return m, nil
			}
			m.state = stateDisplayNameInput
			m.errorMsg = ""
			return m, nil
		}
		m.Name, _ = interactive.EditTextInputField(keyMsg, m.Name, len(m.Name))
	case stateDisplayNameInput:
		if interactive.IsEnterKey(keyMsg) {
			m.selected = true
			m.errorMsg = ""
			return m, tea.Quit
		}
		m.DisplayName, _ = interactive.EditTextInputField(keyMsg, m.DisplayName, len(m.DisplayName))
	}

	return m, nil
}

func (m projectModel) View() string {
	// Render the progress using BaseModel helper.
	progress := m.RenderProgress()
	var view string

	switch m.state {
	case stateOrgSelect:
		view = m.RenderOrgSelection()
	case stateNameInput:
		view = interactive.RenderInputPrompt(
			"Enter project name:",
			"",
			m.Name,
			m.errorMsg,
		)
	case stateDisplayNameInput:
		view = interactive.RenderInputPrompt(
			"Enter project display name (optional):",
			"",
			m.DisplayName,
			m.errorMsg,
		)
	default:
		view = ""
	}
	return progress + view
}

func createProjectInteractive() error {
	orgs, err := util.GetOrganizationNames()
	if err != nil {
		return errors.NewError("failed to get organizations: %v", err)
	}
	if len(orgs) == 0 {
		return errors.NewError("no organizations found")
	}

	model := projectModel{
		BaseModel: interactive.BaseModel{
			Organizations: orgs,
		},
		state: stateOrgSelect,
	}

	finalModel, err := interactive.RunInteractiveModel(model)
	if err != nil {
		return err
	}

	m, ok := finalModel.(projectModel)
	if !ok || !m.selected {
		return errors.NewError("project creation cancelled")
	}

	params := api.CreateProjectParams{
		Organization: m.Organizations[m.OrgCursor],
		Name:         m.Name,
		DisplayName:  m.DisplayName,
	}

	return createProject(params)
}
