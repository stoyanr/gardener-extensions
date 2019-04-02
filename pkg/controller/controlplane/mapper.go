// Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package controlplane

import (
	"context"

	extensionscontroller "github.com/gardener/gardener-extensions/pkg/controller"

	extensionsv1alpha1 "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// NewMapper creates and returns a mapper that maps clusters and secrets to their ControlPlanes.
func NewMapper(client client.Client, predicates ...predicate.Predicate) handler.Mapper {
	return &mapper{
		client:     client,
		predicates: predicates,
	}
}

type mapper struct {
	client     client.Client
	predicates []predicate.Predicate
}

func (m *mapper) Map(obj handler.MapObject) []reconcile.Request {
	if obj.Object == nil {
		return nil
	}

	// Get object accessor
	accessor, err := meta.Accessor(obj.Object)
	if err != nil {
		return nil
	}

	// List all controlplanes in the object's namespace
	cpList := &extensionsv1alpha1.ControlPlaneList{}
	if err := m.client.List(context.TODO(), client.InNamespace(accessor.GetNamespace()), cpList); err != nil {
		return nil
	}

	var requests []reconcile.Request
	for _, cp := range cpList.Items {
		// Check predicates
		if !extensionscontroller.PredicatesMatch(m.predicates, &cp) {
			continue
		}

		// Check conditions depending on object type
		switch obj.Object.(type) {
		case *extensionsv1alpha1.Cluster:
			// For clusters, check that the cluster name matches the control plane namespace
			if accessor.GetName() == cp.Namespace {
				requests = append(requests, getRequest(&cp))
			}
		case *corev1.Secret:
			// For secrets, check that the secret name matches the referenced secret name
			if accessor.GetName() == cp.Spec.SecretRef.Name {
				requests = append(requests, getRequest(&cp))
			}
		}
	}
	return requests
}

// NewOwnerMapper creates and returns a mapper that maps objects owned by a ControlPlane to their owner.
func NewOwnerMapper(client client.Client, predicates ...predicate.Predicate) handler.Mapper {
	return &ownerMapper{
		client:     client,
		predicates: predicates,
	}
}

type ownerMapper struct {
	client     client.Client
	predicates []predicate.Predicate
}

func (m *ownerMapper) Map(obj handler.MapObject) []reconcile.Request {
	if obj.Object == nil {
		return nil
	}

	// Get object accessor
	accessor, err := meta.Accessor(obj.Object)
	if err != nil {
		return nil
	}

	// Get the ControlPlane owner name
	var name string
	for _, ref := range accessor.GetOwnerReferences() {
		if ref.APIVersion == extensionsv1alpha1.SchemeGroupVersion.String() && ref.Kind == extensionsv1alpha1.ControlPlaneResource {
			name = ref.Name
			break
		}
	}
	if name == "" {
		return nil
	}

	// Get the ControlPlane owner
	cp := &extensionsv1alpha1.ControlPlane{}
	if err := m.client.Get(context.TODO(), client.ObjectKey{Namespace: accessor.GetNamespace(), Name: name}, cp); err != nil {
		return nil
	}

	// Check predicates
	if !extensionscontroller.PredicatesMatch(m.predicates, cp) {
		return nil
	}

	return []reconcile.Request{getRequest(cp)}
}

func getRequest(cp *extensionsv1alpha1.ControlPlane) reconcile.Request {
	return reconcile.Request{
		NamespacedName: types.NamespacedName{
			Namespace: cp.Namespace,
			Name:      cp.Name,
		},
	}
}
