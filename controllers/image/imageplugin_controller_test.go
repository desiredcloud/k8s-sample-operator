/*
Copyright 2023 desiredcloud.
*/

package controllers

import (
	"context"
	"fmt"
	"os"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	imagev1alpha1 "github.com/desiredcloud/k8s-sample-operator/api/v1alpha1"
)

var _ = Describe("ImagePlugin controller", func() {
	Context("ImagePlugin controller test", func() {

		const ImagePluginName = "test-imageplugin"

		ctx := context.Background()

		namespace := &corev1.Namespace{
			ObjectMeta: metav1.ObjectMeta{
				Name:      ImagePluginName,
				Namespace: ImagePluginName,
			},
		}

		typeNamespaceName := types.NamespacedName{Name: ImagePluginName, Namespace: ImagePluginName}

		BeforeEach(func() {
			By("Creating the Namespace to perform the tests")
			err := k8sClient.Create(ctx, namespace)
			Expect(err).To(Not(HaveOccurred()))

			By("Setting the Image ENV VAR which stores the Operand image")
			err = os.Setenv("IMAGEPLUGIN_IMAGE", "example.com/image:test")
			Expect(err).To(Not(HaveOccurred()))
		})

		AfterEach(func() {
			// TODO(user): Attention if you improve this code by adding other context test you MUST
			// be aware of the current delete namespace limitations. More info: https://book.kubebuilder.io/reference/envtest.html#testing-considerations
			By("Deleting the Namespace to perform the tests")
			_ = k8sClient.Delete(ctx, namespace)

			By("Removing the Image ENV VAR which stores the Operand image")
			_ = os.Unsetenv("IMAGEPLUGIN_IMAGE")
		})

		It("should successfully reconcile a custom resource for ImagePlugin", func() {
			By("Creating the custom resource for the Kind ImagePlugin")
			imageplugin := &imagev1alpha1.ImagePlugin{}
			err := k8sClient.Get(ctx, typeNamespaceName, imageplugin)
			if err != nil && errors.IsNotFound(err) {
				// Let's mock our custom resource at the same way that we would
				// apply on the cluster the manifest under config/samples
				imageplugin := &imagev1alpha1.ImagePlugin{
					ObjectMeta: metav1.ObjectMeta{
						Name:      ImagePluginName,
						Namespace: namespace.Name,
					},
					Spec: imagev1alpha1.ImagePluginSpec{
						Size: 1,
					},
				}

				err = k8sClient.Create(ctx, imageplugin)
				Expect(err).To(Not(HaveOccurred()))
			}

			By("Checking if the custom resource was successfully created")
			Eventually(func() error {
				found := &imagev1alpha1.ImagePlugin{}
				return k8sClient.Get(ctx, typeNamespaceName, found)
			}, time.Minute, time.Second).Should(Succeed())

			By("Reconciling the custom resource created")
			imagepluginReconciler := &ImagePluginReconciler{
				Client: k8sClient,
				Scheme: k8sClient.Scheme(),
			}

			_, err = imagepluginReconciler.Reconcile(ctx, reconcile.Request{
				NamespacedName: typeNamespaceName,
			})
			Expect(err).To(Not(HaveOccurred()))

			By("Checking if Deployment was successfully created in the reconciliation")
			Eventually(func() error {
				found := &appsv1.Deployment{}
				return k8sClient.Get(ctx, typeNamespaceName, found)
			}, time.Minute, time.Second).Should(Succeed())

			By("Checking the latest Status Condition added to the ImagePlugin instance")
			Eventually(func() error {
				if imageplugin.Status.Conditions != nil && len(imageplugin.Status.Conditions) != 0 {
					latestStatusCondition := imageplugin.Status.Conditions[len(imageplugin.Status.Conditions)-1]
					expectedLatestStatusCondition := metav1.Condition{Type: typeAvailableImagePlugin,
						Status: metav1.ConditionTrue, Reason: "Reconciling",
						Message: fmt.Sprintf("Deployment for custom resource (%s) with %d replicas created successfully", imageplugin.Name, imageplugin.Spec.Size)}
					if latestStatusCondition != expectedLatestStatusCondition {
						return fmt.Errorf("The latest status condition added to the imageplugin instance is not as expected")
					}
				}
				return nil
			}, time.Minute, time.Second).Should(Succeed())
		})
	})
})
