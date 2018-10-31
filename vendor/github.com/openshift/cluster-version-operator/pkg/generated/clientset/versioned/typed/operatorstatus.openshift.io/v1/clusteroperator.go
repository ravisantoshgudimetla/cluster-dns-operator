/*
Copyright The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/openshift/cluster-version-operator/pkg/apis/operatorstatus.openshift.io/v1"
	scheme "github.com/openshift/cluster-version-operator/pkg/generated/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ClusterOperatorsGetter has a method to return a ClusterOperatorInterface.
// A group's client should implement this interface.
type ClusterOperatorsGetter interface {
	ClusterOperators(namespace string) ClusterOperatorInterface
}

// ClusterOperatorInterface has methods to work with ClusterOperator resources.
type ClusterOperatorInterface interface {
	Create(*v1.ClusterOperator) (*v1.ClusterOperator, error)
	Update(*v1.ClusterOperator) (*v1.ClusterOperator, error)
	UpdateStatus(*v1.ClusterOperator) (*v1.ClusterOperator, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.ClusterOperator, error)
	List(opts metav1.ListOptions) (*v1.ClusterOperatorList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.ClusterOperator, err error)
	ClusterOperatorExpansion
}

// clusterOperators implements ClusterOperatorInterface
type clusterOperators struct {
	client rest.Interface
	ns     string
}

// newClusterOperators returns a ClusterOperators
func newClusterOperators(c *OperatorstatusV1Client, namespace string) *clusterOperators {
	return &clusterOperators{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the clusterOperator, and returns the corresponding clusterOperator object, and an error if there is any.
func (c *clusterOperators) Get(name string, options metav1.GetOptions) (result *v1.ClusterOperator, err error) {
	result = &v1.ClusterOperator{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("clusteroperators").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ClusterOperators that match those selectors.
func (c *clusterOperators) List(opts metav1.ListOptions) (result *v1.ClusterOperatorList, err error) {
	result = &v1.ClusterOperatorList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("clusteroperators").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested clusterOperators.
func (c *clusterOperators) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("clusteroperators").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a clusterOperator and creates it.  Returns the server's representation of the clusterOperator, and an error, if there is any.
func (c *clusterOperators) Create(clusterOperator *v1.ClusterOperator) (result *v1.ClusterOperator, err error) {
	result = &v1.ClusterOperator{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("clusteroperators").
		Body(clusterOperator).
		Do().
		Into(result)
	return
}

// Update takes the representation of a clusterOperator and updates it. Returns the server's representation of the clusterOperator, and an error, if there is any.
func (c *clusterOperators) Update(clusterOperator *v1.ClusterOperator) (result *v1.ClusterOperator, err error) {
	result = &v1.ClusterOperator{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("clusteroperators").
		Name(clusterOperator.Name).
		Body(clusterOperator).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *clusterOperators) UpdateStatus(clusterOperator *v1.ClusterOperator) (result *v1.ClusterOperator, err error) {
	result = &v1.ClusterOperator{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("clusteroperators").
		Name(clusterOperator.Name).
		SubResource("status").
		Body(clusterOperator).
		Do().
		Into(result)
	return
}

// Delete takes name of the clusterOperator and deletes it. Returns an error if one occurs.
func (c *clusterOperators) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("clusteroperators").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *clusterOperators) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("clusteroperators").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched clusterOperator.
func (c *clusterOperators) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.ClusterOperator, err error) {
	result = &v1.ClusterOperator{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("clusteroperators").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
