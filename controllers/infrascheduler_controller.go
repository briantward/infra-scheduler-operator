/*
Copyright 2021.

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
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	infraschedulerv1alpha1 "github.com/briantward/infra-scheduler-operator/api/v1alpha1"
)

// InfraSchedulerReconciler reconciles a InfraScheduler object
type InfraSchedulerReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=infrascheduler.bward.dev,resources=infraschedulers,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=infrascheduler.bward.dev,resources=infraschedulers/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=infrascheduler.bward.dev,resources=infraschedulers/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the InfraScheduler object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.7.2/pkg/reconcile
func (r *InfraSchedulerReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = r.Log.WithValues("infrascheduler", req.NamespacedName)

	// your logic here

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *InfraSchedulerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&infraschedulerv1alpha1.InfraScheduler{}).
		Complete(r)
}

func createInfraPool() {
	// apply default manifest for infra MCP
}

func reconcileInfraPool() {
	// verify if Infra Pool exists, reconcile
}

func copyMachineSets() {
	// (default random select all AZs with replica of one or more, up to three nodes)
}

func modifyMachineSets() {
	// (adjust machine profile type for cloud spec?)
}

func createDefinedMachineSets() {
	// allow management of custom defined MachineSets
}

func assignInfraNodeLabel() {
	// for non MachineSet management
}

func removeWorkerNodeLabel() {
	// for both MachineSet management and non MachineSet management
}

func setDefaultNodeSelectorOnScheduler() {
	// update the defaultNodeSelector on the default Scheduler
}

func setWorkloadNodeSelector() {
	// update the nodeSelector on each infra workload
}

func setWorkloadTolerations() {
	// update the tolerations on each infra workload
}
