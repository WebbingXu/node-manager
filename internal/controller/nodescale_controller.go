/*
Copyright 2023.

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

package controller

import (
	"context"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/klog/v2"
	"path/filepath"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/source"
	"sigs.k8s.io/controller-runtime/pkg/handler"

	nodev1 "github.com/log/api/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const(
	clustersKubeConfigDir = "clusters_kubeconfig/"
)

// NodeScaleReconciler reconciles a NodeScale object
type NodeScaleReconciler struct {
	client.Client
	Scheme *runtime.Scheme
	Log logr.Logger
}

//+kubebuilder:rbac:groups=node.log-platform.com,resources=nodescales,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=node.log-platform.com,resources=nodescales/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=node.log-platform.com,resources=nodescales/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the NodeScale object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.15.0/pkg/reconcile
func (r *NodeScaleReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := r.Log.WithValues("nodeScale", req.NamespacedName)

	// TODO(user): your logic here
	var nodeScale nodev1.NodeScale

	// 如果 nodeScale 资源配置被删除了，应该忽略这种事件
	if err := r.Get(ctx, req.NamespacedName, &nodeScale); err != nil {
		if errors.IsNotFound(err) {
			log.Info("nodeScale cr deleted", "nodeScale", req.Name)
			return ctrl.Result{}, nil
		}
	}
	// 获取在cr中的节点
	nodesInCR := nodeScale.Spec.Nodes

	// 获取 nodeScale 所属的集群，也就是看要对哪个集群进行扩容
	clusterName := nodeScale.Spec.ClusterName

	// 1. 获取nodeScale所属集群所有节点，判断哪些IP需要扩容，那些IP需要缩容
	// cluster 的 kubeConfig 文件通过 configmap 挂载到容器内，configmap和容器内文件名为集群名
	cli, err := CreateClient(filepath.Join(clustersKubeConfigDir, clusterName))
	if err != nil {
		return ctrl.Result{}, err
	}
	nodesInCluster, err := cli.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	if err != nil {
		klog.Errorf("list nodes from %s failed, err: %s", clusterName, err.Error())
		return ctrl.Result{}, err
	}
	saleIPs := GetSaleIPs(nodesInCR, nodesInCluster)
	shrinkIPs := GetShrinkIPs(nodesInCR, nodesInCluster)

	// 2. 调用 toc ansible 进行扩容


	// 3. 部署完后检查网络

	// 4. 打标签

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *NodeScaleReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&nodev1.NodeScale{}).
		// 对 node 的监听
		//Watches( &corev1.Node{}, &handler.EnqueueRequestForObject{}).
		Complete(r)
}
