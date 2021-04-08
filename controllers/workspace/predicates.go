//
// Copyright (c) 2019-2021 Red Hat, Inc.
// This program and the accompanying materials are made
// available under the terms of the Eclipse Public License 2.0
// which is available at https://www.eclipse.org/legal/epl-2.0/
//
// SPDX-License-Identifier: EPL-2.0
//
// Contributors:
//   Red Hat, Inc. - initial API and implementation
//

package controllers

import (
	"github.com/devfile/devworkspace-operator/pkg/constants"
	corev1 "k8s.io/api/core/v1"

	dw "github.com/devfile/api/v2/pkg/apis/workspaces/v1alpha2"

	"k8s.io/apimachinery/pkg/api/equality"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// predicates filters incoming events to avoid unnecessary reconciles to failed workspaces.
// If a workspace failed and its spec is changed, we trigger reconciles to allow for fixing
// issues in the workspace spec.
var predicates = predicate.Funcs{
	CreateFunc: func(_ event.CreateEvent) bool { return true },
	DeleteFunc: func(_ event.DeleteEvent) bool { return true },
	UpdateFunc: func(ev event.UpdateEvent) bool {
		newObj, ok := ev.ObjectNew.(*dw.DevWorkspace)
		if !ok {
			return true
		}
		if newObj.Status.Phase != dw.DevWorkspaceStatusFailed {
			return true
		}
		oldObj, ok := ev.ObjectOld.(*dw.DevWorkspace)
		if !ok {
			// Should never happen
			return true
		}
		// always reconcile if resource is deleted
		if newObj.GetDeletionTimestamp() != nil {
			return true
		}
		// Trigger a reconcile on failed workspaces if spec is updated.
		return !equality.Semantic.DeepEqual(oldObj.Spec, newObj.Spec)
	},
	GenericFunc: func(_ event.GenericEvent) bool { return true },
}

var podPredicate = predicate.Funcs{
	UpdateFunc: func(ev event.UpdateEvent) bool {
		newObj, ok := ev.ObjectNew.(*corev1.Pod)
		if !ok {
			return true
		}

		oldObj, ok := ev.ObjectOld.(*corev1.Pod)
		if !ok {
			// Should never happen
			return true
		}

		// If the dewworkspace label does not exist, do no reconcile
		if _, ok := newObj.Labels[constants.DevWorkspaceIDLabel]; !ok {
			return false
		}

		// Always reconcile if resource is deleted
		if newObj.GetDeletionTimestamp() != nil {
			return true
		}

		// Trigger a reconcile on failed workspaces if status is updated.
		return !equality.Semantic.DeepEqual(oldObj.Status, newObj.Status)
	},
}
