/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"

	"github.com/go-logr/logr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	featurev1alpha1 "example/api/v1alpha1"
)

// MonitoringReconciler reconciles a Monitoring object
type MonitoringReconciler struct {
	client.Client
	Log logr.Logger
}

// +kubebuilder:rbac:groups=feature.kublr.com,resources=monitorings,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=feature.kublr.com,resources=monitorings/status,verbs=get;update;patch

func (r *MonitoringReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	_ = r.Log.WithValues("monitoring", req.NamespacedName)

	// your logic here

	return ctrl.Result{}, nil
}

func (r *MonitoringReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&featurev1alpha1.Monitoring{}).
		Complete(r)
}
