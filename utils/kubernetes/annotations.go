// Copyright 2023 Red Hat, Inc. and/or its affiliates
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package kubernetes

import (
	"context"
	"strings"

	operatorapi "github.com/kiegroup/kogito-serverless-operator/api/v1alpha08"

	"github.com/go-logr/logr"
	client "sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/kiegroup/kogito-serverless-operator/api/metadata"
)

const LastAppliedConfigurationAnnotation = "kubectl.kubernetes.io/last-applied-configuration"
const GenerationAnnotation = "generation"

func GetAnnotationResource(value string) string {
	stringArray := strings.Split(value, "/")
	if len(stringArray) == 2 && stringArray[0] == metadata.Domain && strings.HasPrefix(stringArray[1], "resource") {
		return stringArray[1]
	} else {
		return ""
	}
}

func getWorkflow(namespace string, name string, c client.Client, ctx context.Context, logger *logr.Logger) *operatorapi.KogitoServerlessWorkflow {
	serverlessWorkflowType := &operatorapi.KogitoServerlessWorkflow{}
	serverlessWorkflowType.Namespace = namespace
	serverlessWorkflowType.Name = name
	serverlessWorkflow := &operatorapi.KogitoServerlessWorkflow{}
	if err := c.Get(ctx, client.ObjectKeyFromObject(serverlessWorkflowType), serverlessWorkflow); err != nil {
		logger.Error(err, "Error during Get")
	}
	return serverlessWorkflow
}

func GetLastGeneration(namespace string, name string, c client.Client, ctx context.Context, logger *logr.Logger) int64 {
	workflow := getWorkflow(namespace, name, c, ctx, logger)
	return workflow.Generation
}
