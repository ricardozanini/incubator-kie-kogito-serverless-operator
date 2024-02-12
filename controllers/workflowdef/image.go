/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package workflowdef

import (
	"github.com/apache/incubator-kie-kogito-serverless-operator/api/metadata"
	"github.com/apache/incubator-kie-kogito-serverless-operator/api/v1alpha08"
	"github.com/apache/incubator-kie-kogito-serverless-operator/version"
)

const (
	latestImageTag              = "latest"
	nightlySuffix               = "nightly"
	defaultWorkflowDevModeImage = "registry.redhat.io/openshift-serverless-1-tech-preview/logic-swf-devmode-rhel8"
	defaultWorkflowBuilderImage = "registry.redhat.io/openshift-serverless-1-tech-preview/logic-swf-builder-rhel8"
	defaultOperatorImage        = "registry.redhat.io/openshift-serverless-1-tech-preview/logic-rhel8-operator"
)

// GetWorkflowAppImageNameTag retrieve the tag for the image based on the Workflow based annotation, <workflowid>:latest otherwise
func GetWorkflowAppImageNameTag(w *v1alpha08.SonataFlow) string {
	v := w.Annotations[metadata.Version]
	if v != "" {
		return w.Name + ":" + v
	}
	return w.Name + ":" + latestImageTag
}

func GetDefaultWorkflowDevModeImageTag() string {
	return GetDefaultImageTag(defaultWorkflowDevModeImage)
}

func GetDefaultWorkflowBuilderImageTag() string {
	return GetDefaultImageTag(defaultWorkflowBuilderImage)
}

func GetDefaultOperatorImageTag() string {
	return GetDefaultImageTag(defaultOperatorImage)
}

func GetDefaultImageTag(imgTag string) string {
	imgTag += ":" + version.OperatorVersion
	return imgTag
}
