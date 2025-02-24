/*
Copyright 2020 Google LLC
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    https://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package model

import (
	appsv1 "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	policyv1 "k8s.io/api/policy/v1"
)

// DesiredClusterState holds desired state of a cluster.
type DesiredClusterState struct {
	JmStatefulSet       *appsv1.StatefulSet
	JmService           *corev1.Service
	JmIngress           *networkingv1.Ingress
	TmStatefulSet       *appsv1.StatefulSet
	ConfigMap           *corev1.ConfigMap
	Job                 *batchv1.Job
	PodDisruptionBudget *policyv1.PodDisruptionBudget
}
