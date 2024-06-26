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

package kubernetes

import (
	"testing"

	"github.com/stretchr/testify/assert"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
)

func TestCreateOrReplaceEnv(t *testing.T) {
	containerNoEnv := &appsv1.Deployment{
		Spec: appsv1.DeploymentSpec{
			Template: v1.PodTemplateSpec{
				Spec: v1.PodSpec{
					Containers: []v1.Container{{Env: nil}},
				},
			},
		},
	}
	containerWithEnv := &appsv1.Deployment{
		Spec: appsv1.DeploymentSpec{
			Template: v1.PodTemplateSpec{
				Spec: v1.PodSpec{
					Containers: []v1.Container{{Env: []v1.EnvVar{{
						Name:  "myvar",
						Value: "myvalue",
					}}}},
				},
			},
		},
	}

	CreateOrReplaceEnv(&containerNoEnv.Spec.Template.Spec.Containers[0], "myvar", "mutated")
	assert.Equal(t, "mutated", containerNoEnv.Spec.Template.Spec.Containers[0].Env[0].Value)

	CreateOrReplaceEnv(&containerWithEnv.Spec.Template.Spec.Containers[0], "myvar", "mutated")
	assert.Equal(t, "mutated", containerWithEnv.Spec.Template.Spec.Containers[0].Env[0].Value)
}

func TestAddIfNotPresent(t *testing.T) {
	containerNoEnv := &v1.Container{Env: nil}

	wasAdded := AddEnvIfNotPresent(containerNoEnv, v1.EnvVar{Name: "var1", Value: "value1"})
	assert.True(t, wasAdded)
	assert.Equal(t, v1.EnvVar{Name: "var1", Value: "value1"}, containerNoEnv.Env[0])

	containerWithEnv := &v1.Container{Env: []v1.EnvVar{{Name: "var1", Value: "value1"}, {Name: "var2", Value: "value2"}}}
	wasAdded = AddEnvIfNotPresent(containerWithEnv, v1.EnvVar{Name: "var1", Value: "value1Changed"})
	assert.False(t, wasAdded)
	assert.Equal(t, v1.EnvVar{Name: "var1", Value: "value1"}, containerWithEnv.Env[0])
	assert.Equal(t, v1.EnvVar{Name: "var2", Value: "value2"}, containerWithEnv.Env[1])

	wasAdded = AddEnvIfNotPresent(containerWithEnv, v1.EnvVar{Name: "var3", Value: "value3"})
	assert.True(t, wasAdded)

	assert.Equal(t, v1.EnvVar{Name: "var1", Value: "value1"}, containerWithEnv.Env[0])
	assert.Equal(t, v1.EnvVar{Name: "var2", Value: "value2"}, containerWithEnv.Env[1])
	assert.Equal(t, v1.EnvVar{Name: "var3", Value: "value3"}, containerWithEnv.Env[2])
}
