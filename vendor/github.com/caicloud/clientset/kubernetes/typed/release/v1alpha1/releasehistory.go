/*
Copyright 2017 caicloud authors. All rights reserved.
*/

package v1alpha1

import (
	scheme "github.com/caicloud/clientset/kubernetes/scheme"
	v1alpha1 "github.com/caicloud/clientset/pkg/apis/release/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ReleaseHistoriesGetter has a method to return a ReleaseHistoryInterface.
// A group's client should implement this interface.
type ReleaseHistoriesGetter interface {
	ReleaseHistories(namespace string) ReleaseHistoryInterface
}

// ReleaseHistoryInterface has methods to work with ReleaseHistory resources.
type ReleaseHistoryInterface interface {
	Create(*v1alpha1.ReleaseHistory) (*v1alpha1.ReleaseHistory, error)
	Update(*v1alpha1.ReleaseHistory) (*v1alpha1.ReleaseHistory, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ReleaseHistory, error)
	List(opts v1.ListOptions) (*v1alpha1.ReleaseHistoryList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ReleaseHistory, err error)
	ReleaseHistoryExpansion
}

// releaseHistories implements ReleaseHistoryInterface
type releaseHistories struct {
	client rest.Interface
	ns     string
}

// newReleaseHistories returns a ReleaseHistories
func newReleaseHistories(c *ReleaseV1alpha1Client, namespace string) *releaseHistories {
	return &releaseHistories{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Create takes the representation of a releaseHistory and creates it.  Returns the server's representation of the releaseHistory, and an error, if there is any.
func (c *releaseHistories) Create(releaseHistory *v1alpha1.ReleaseHistory) (result *v1alpha1.ReleaseHistory, err error) {
	result = &v1alpha1.ReleaseHistory{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("releasehistories").
		Body(releaseHistory).
		Do().
		Into(result)
	return
}

// Update takes the representation of a releaseHistory and updates it. Returns the server's representation of the releaseHistory, and an error, if there is any.
func (c *releaseHistories) Update(releaseHistory *v1alpha1.ReleaseHistory) (result *v1alpha1.ReleaseHistory, err error) {
	result = &v1alpha1.ReleaseHistory{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("releasehistories").
		Name(releaseHistory.Name).
		Body(releaseHistory).
		Do().
		Into(result)
	return
}

// Delete takes name of the releaseHistory and deletes it. Returns an error if one occurs.
func (c *releaseHistories) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("releasehistories").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *releaseHistories) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("releasehistories").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Get takes name of the releaseHistory, and returns the corresponding releaseHistory object, and an error if there is any.
func (c *releaseHistories) Get(name string, options v1.GetOptions) (result *v1alpha1.ReleaseHistory, err error) {
	result = &v1alpha1.ReleaseHistory{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("releasehistories").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ReleaseHistories that match those selectors.
func (c *releaseHistories) List(opts v1.ListOptions) (result *v1alpha1.ReleaseHistoryList, err error) {
	result = &v1alpha1.ReleaseHistoryList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("releasehistories").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested releaseHistories.
func (c *releaseHistories) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("releasehistories").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Patch applies the patch and returns the patched releaseHistory.
func (c *releaseHistories) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ReleaseHistory, err error) {
	result = &v1alpha1.ReleaseHistory{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("releasehistories").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
