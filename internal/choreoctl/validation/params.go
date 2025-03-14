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

package validation

import (
	"fmt"

	"github.com/choreo-idp/choreo/pkg/cli/types/api"
)

// ValidateParams validates command parameters based on command and resource types
func ValidateParams(cmdType CommandType, resource ResourceType, params interface{}) error {
	switch resource {
	case ResourceProject:
		return validateProjectParams(cmdType, params)
	case ResourceComponent:
		return validateComponentParams(cmdType, params)
	case ResourceBuild:
		return validateBuildParams(cmdType, params)
	case ResourceDeployment:
		return validateDeploymentParams(cmdType, params)
	case ResourceDeploymentTrack:
		return validateDeploymentTrackParams(cmdType, params)
	case ResourceEnvironment:
		return validateEnvironmentParams(cmdType, params)
	case ResourceDeployableArtifact:
		return validateDeployableArtifactParams(cmdType, params)
	case ResourceDataPlane:
		return validateDataPlaneParams(cmdType, params)
	case ResourceOrganization:
		return validateOrganizationParams(cmdType, params)
	case ResourceEndpoint:
		return validateEndpointParams(cmdType, params)
	case ResourceLogs:
		return validateLogParams(cmdType, params)
	default:
		return fmt.Errorf("unknown resource type: %s", resource)
	}
}

// validateProjectParams validates parameters for project operations
func validateProjectParams(cmdType CommandType, params interface{}) error {
	switch cmdType {
	case CmdCreate:
		if p, ok := params.(api.CreateProjectParams); ok {
			fields := map[string]string{
				"organization": p.Organization,
				"name":         p.Name,
			}
			if !checkRequiredFields(fields) {
				return generateHelpError(cmdType, ResourceProject, fields)
			}
		}
	case CmdGet:
		if p, ok := params.(api.GetProjectParams); ok {
			fields := map[string]string{
				"organization": p.Organization,
			}
			if !checkRequiredFields(fields) {
				return generateHelpError(cmdType, ResourceProject, fields)
			}
		}
	}
	return nil
}

// validateComponentParams validates parameters for component operations
func validateComponentParams(cmdType CommandType, params interface{}) error {
	switch cmdType {
	case CmdCreate:
		if p, ok := params.(api.CreateComponentParams); ok {
			fields := map[string]string{
				"organization": p.Organization,
				"project":      p.Project,
				"name":         p.Name,
			}
			if !checkRequiredFields(fields) {
				return generateHelpError(cmdType, ResourceComponent, fields)
			}
			return ValidateGitHubURL(p.GitRepositoryURL)
		}
	case CmdGet:
		if p, ok := params.(api.GetComponentParams); ok {
			fields := map[string]string{
				"organization": p.Organization,
				"project":      p.Project,
			}
			if !checkRequiredFields(fields) {
				return generateHelpError(cmdType, ResourceComponent, fields)
			}
		}
	}
	return nil
}

// validateBuildParams validates parameters for build operations
func validateBuildParams(cmdType CommandType, params interface{}) error {
	switch cmdType {
	case CmdCreate:
		if p, ok := params.(api.CreateBuildParams); ok {
			// All required fields
			requiredFields := map[string]string{
				"organization": p.Organization,
				"project":      p.Project,
				"component":    p.Component,
				"name":         p.Name,
			}

			if !checkRequiredFields(requiredFields) {
				return generateHelpError(cmdType, ResourceBuild, requiredFields)
			}
		}

	case CmdGet:
		if p, ok := params.(api.GetBuildParams); ok {
			fields := map[string]string{
				"organization": p.Organization,
				"project":      p.Project,
				"component":    p.Component,
			}
			if !checkRequiredFields(fields) {
				return generateHelpError(cmdType, ResourceBuild, fields)
			}
		}
	}
	return nil
}

// validateDeploymentParams validates parameters for deployment operations
func validateDeploymentParams(cmdType CommandType, params interface{}) error {
	switch cmdType {
	case CmdCreate:
		if p, ok := params.(api.CreateDeploymentParams); ok {
			fields := map[string]string{
				"organization": p.Organization,
				"project":      p.Project,
				"component":    p.Component,
			}
			if !checkRequiredFields(fields) {
				return generateHelpError(cmdType, ResourceDeployment, fields)
			}
		}
	case CmdGet:
		if p, ok := params.(api.GetDeploymentParams); ok {
			fields := map[string]string{
				"organization": p.Organization,
				"project":      p.Project,
				"component":    p.Component,
			}
			if !checkRequiredFields(fields) {
				return generateHelpError(cmdType, ResourceDeployment, fields)
			}
		}
	}
	return nil
}

// validateDeploymentTrackParams validates parameters for deployment track operations
func validateDeploymentTrackParams(cmdType CommandType, params interface{}) error {
	switch cmdType {
	case CmdCreate:
		if p, ok := params.(api.CreateDeploymentTrackParams); ok {
			fields := map[string]string{
				"organization": p.Organization,
				"project":      p.Project,
				"component":    p.Component,
			}
			if !checkRequiredFields(fields) {
				return generateHelpError(cmdType, ResourceDeploymentTrack, fields)
			}
		}
	case CmdGet:
		if p, ok := params.(api.GetDeploymentTrackParams); ok {
			fields := map[string]string{
				"organization": p.Organization,
				"project":      p.Project,
				"component":    p.Component,
			}
			if !checkRequiredFields(fields) {
				return generateHelpError(cmdType, ResourceDeploymentTrack, fields)
			}
		}
	}
	return nil
}

// validateEnvironmentParams validates parameters for environment operations
func validateEnvironmentParams(cmdType CommandType, params interface{}) error {
	switch cmdType {
	case CmdCreate:
		if p, ok := params.(api.CreateEnvironmentParams); ok {
			fields := map[string]string{
				"organization": p.Organization,
				"name":         p.Name,
			}
			if !checkRequiredFields(fields) {
				return generateHelpError(cmdType, ResourceEnvironment, fields)
			}
		}
	case CmdGet:
		if p, ok := params.(api.GetEnvironmentParams); ok {
			fields := map[string]string{
				"organization": p.Organization,
			}
			if !checkRequiredFields(fields) {
				return generateHelpError(cmdType, ResourceEnvironment, fields)
			}
		}
	}
	return nil
}

// validateDeployableArtifactParams validates parameters for deployable artifact operations
func validateDeployableArtifactParams(cmdType CommandType, params interface{}) error {
	switch cmdType {
	case CmdCreate:
		if p, ok := params.(api.CreateDeployableArtifactParams); ok {
			fields := map[string]string{
				"organization": p.Organization,
				"project":      p.Project,
				"component":    p.Component,
			}
			if !checkRequiredFields(fields) {
				return generateHelpError(cmdType, ResourceDeployableArtifact, fields)
			}
		}
	case CmdGet:
		if p, ok := params.(api.GetDeployableArtifactParams); ok {
			fields := map[string]string{
				"organization": p.Organization,
				"project":      p.Project,
				"component":    p.Component,
			}
			if !checkRequiredFields(fields) {
				return generateHelpError(cmdType, ResourceDeployableArtifact, fields)
			}
		}
	}
	return nil
}

// validateLogParams validates parameters for log operations
func validateLogParams(cmdType CommandType, params interface{}) error {
	if cmdType == CmdGet {
		if p, ok := params.(api.LogParams); ok {
			fields := map[string]string{
				"organization": p.Organization,
				"project":      p.Project,
				"component":    p.Component,
			}
			if !checkRequiredFields(fields) {
				return generateHelpError(cmdType, ResourceLogs, fields)
			}
		}
	}
	return nil
}

// validateDataPlaneParams validates parameters for data plane operations
func validateDataPlaneParams(cmdType CommandType, params interface{}) error {
	switch cmdType {
	case CmdGet:
		if p, ok := params.(api.GetDataPlaneParams); ok {
			fields := map[string]string{
				"organization": p.Organization,
			}
			if !checkRequiredFields(fields) {
				return generateHelpError(cmdType, ResourceDataPlane, fields)
			}
		}
	case CmdCreate:
		if p, ok := params.(api.CreateDataPlaneParams); ok {
			fields := map[string]string{
				"organization": p.Organization,
			}
			if !checkRequiredFields(fields) {
				return generateHelpError(cmdType, ResourceDataPlane, fields)
			}
		}
	}
	return nil
}

// validateOrganizationParams validates parameters for organization operations
func validateOrganizationParams(cmdType CommandType, params interface{}) error {
	if cmdType == CmdCreate {
		if p, ok := params.(api.CreateOrganizationParams); ok {
			fields := map[string]string{
				"name": p.Name,
			}
			if !checkRequiredFields(fields) {
				return generateHelpError(cmdType, ResourceOrganization, fields)
			}
		}
	}
	return nil
}

// validateEndpointParams validates parameters for endpoint operations
func validateEndpointParams(cmdType CommandType, params interface{}) error {
	if cmdType == CmdGet {
		if p, ok := params.(api.GetEndpointParams); ok {
			fields := map[string]string{
				"organization": p.Organization,
				"project":      p.Project,
				"component":    p.Component,
			}
			if !checkRequiredFields(fields) {
				return generateHelpError(cmdType, ResourceEndpoint, fields)
			}
		}
	}
	return nil
}
