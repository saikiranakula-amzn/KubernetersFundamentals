package main

import (
	"context"
	"fmt"
	"time"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

type WebApp struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WebAppSpec   `json:"spec,omitempty"`
	Status            WebAppStatus `json:"status,omitempty"`
}

type WebAppSpec struct {
	Replicas int32  `json:"replicas"`
	Image    string `json:"image"`
	Port     int32  `json:"port"`
}

type WebAppStatus struct {
	AvailableReplicas int32 `json:"availableReplicas"`
}

func (w *WebApp) DeepCopyObject() runtime.Object {
	return w.DeepCopy()
}

func (w *WebApp) DeepCopy() *WebApp {
	if w == nil {
		return nil
	}
	out := new(WebApp)
	*out = *w
	return out
}

type WebAppReconciler struct {
	client.Client
	Clientset kubernetes.Interface
}

func (r *WebAppReconciler) Reconcile(ctx context.Context, req reconcile.Request) (reconcile.Result, error) {
	webapp := &WebApp{}
	err := r.Get(ctx, req.NamespacedName, webapp)
	if err != nil {
		return reconcile.Result{}, client.IgnoreNotFound(err)
	}

	// Create Deployment
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      webapp.Name,
			Namespace: webapp.Namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &webapp.Spec.Replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{"app": webapp.Name},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{"app": webapp.Name},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "webapp",
							Image: webapp.Spec.Image,
							Ports: []corev1.ContainerPort{
								{ContainerPort: webapp.Spec.Port},
							},
						},
					},
				},
			},
		},
	}

	err = r.Create(ctx, deployment)
	if err != nil && !client.IgnoreAlreadyExists(err) != nil {
		return reconcile.Result{}, err
	}

	// Create Service
	service := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      webapp.Name + "-service",
			Namespace: webapp.Namespace,
		},
		Spec: corev1.ServiceSpec{
			Selector: map[string]string{"app": webapp.Name},
			Ports: []corev1.ServicePort{
				{
					Port:       80,
					TargetPort: intstr.FromInt(int(webapp.Spec.Port)),
				},
			},
		},
	}

	err = r.Create(ctx, service)
	if err != nil && !client.IgnoreAlreadyExists(err) != nil {
		return reconcile.Result{}, err
	}

	return reconcile.Result{RequeueAfter: time.Minute * 5}, nil
}

func main() {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err)
	}

	mgr, err := manager.New(config, manager.Options{})
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	reconciler := &WebAppReconciler{
		Client:    mgr.GetClient(),
		Clientset: clientset,
	}

	ctrl, err := controller.New("webapp-controller", mgr, controller.Options{
		Reconciler: reconciler,
	})
	if err != nil {
		panic(err)
	}

	err = ctrl.Watch(&source.Kind{Type: &WebApp{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting WebApp controller...")
	if err := mgr.Start(context.Background()); err != nil {
		panic(err)
	}
}
