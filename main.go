package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"path/filepath"

	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	// This will load kubeconfig
	var kubeconfig string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = filepath.Join(home, ".kube", "config")
	}
	flag.StringVar(&kubeconfig, "kubeconfig", kubeconfig, "")
	flag.Parse()


	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatalf("Error loading kubeconfig: %v", err)
	}

	// This creates k8s client
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("Error creating Kubernetes client: %v", err)
	}

	// // Pod Definition
	// pod := &v1.Pod{
	// 	ObjectMeta: metav1.ObjectMeta{
	// 		Name:      "shubham-pod-go",
	// 		Namespace: "default",
	// 	},
	// 	Spec: v1.PodSpec{
	// 		Containers: []v1.Container{
	// 			{
	// 				Name:  "nginx-container",
	// 				Image: "nginx:latest",
	// 				Ports: []v1.ContainerPort{
	// 					{
	// 						ContainerPort: 80,
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }

	// // Pod Creation
	// createdPod, err := clientset.CoreV1().Pods("default").Create(context.TODO(), pod, metav1.CreateOptions{})
	// if err != nil {
	// 	log.Fatalf("Error creating pod: %v", err)
	// }
	// fmt.Printf("Pod %s created successfully in namespace %s\n", createdPod.Name, createdPod.Namespace)


		// Define Deployment
		deployment := &appsv1.Deployment{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "sbm-deployment-go",
				Namespace: "default",
			},
			Spec: appsv1.DeploymentSpec{	
				Replicas: int32Ptr(2), // 2 replicas
				Selector: &metav1.LabelSelector{
					MatchLabels: map[string]string{"app": "nginx"},
				},
				Template: v1.PodTemplateSpec{
					ObjectMeta: metav1.ObjectMeta{
						Labels: map[string]string{"app": "nginx"},
					},
					Spec: v1.PodSpec{
						Containers: []v1.Container{
							{
								Name:  "nginx",
								Image: "nginx:latest",
								Ports: []v1.ContainerPort{
									{ContainerPort: 80},
								},
							},
						},
					},
				},
			},
		}
	
		// Create Deployment
		createdDeployment, err := clientset.AppsV1().Deployments("default").Create(context.TODO(), deployment, metav1.CreateOptions{})
		if err != nil {
			log.Fatalf("Error creating Deployment: %v", err)
		}
		fmt.Printf("Deployment %s created successfully\n", createdDeployment.Name)
	
		// Define Service
		service := &v1.Service{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "sbm-service-go",
				Namespace: "default",
			},
			Spec: v1.ServiceSpec{
				Selector: map[string]string{"app": "nginx"},
				Ports: []v1.ServicePort{
					{
						Protocol:   v1.ProtocolTCP,
						Port:       80,  // Service Port
						TargetPort: intstr.FromInt(80), // Pod Port
					},
				},
				Type: v1.ServiceTypeClusterIP, // Internal service
			},
		}
	
		// Service creation
		createdService, err := clientset.CoreV1().Services("default").Create(context.TODO(), service, metav1.CreateOptions{})
		if err != nil {
			log.Fatalf("Error creating Service: %v", err)
		}
		fmt.Printf("Service %s created successfully\n", createdService.Name)
	
}

func int32Ptr(i int32) *int32 {
	return &i
}

