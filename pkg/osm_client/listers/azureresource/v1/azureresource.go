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

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/openservicemesh/osm/pkg/apis/azureresource/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AzureResourceLister helps list AzureResources.
type AzureResourceLister interface {
	// List lists all AzureResources in the indexer.
	List(selector labels.Selector) (ret []*v1.AzureResource, err error)
	// AzureResources returns an object that can list and get AzureResources.
	AzureResources(namespace string) AzureResourceNamespaceLister
	AzureResourceListerExpansion
}

// azureResourceLister implements the AzureResourceLister interface.
type azureResourceLister struct {
	indexer cache.Indexer
}

// NewAzureResourceLister returns a new AzureResourceLister.
func NewAzureResourceLister(indexer cache.Indexer) AzureResourceLister {
	return &azureResourceLister{indexer: indexer}
}

// List lists all AzureResources in the indexer.
func (s *azureResourceLister) List(selector labels.Selector) (ret []*v1.AzureResource, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AzureResource))
	})
	return ret, err
}

// AzureResources returns an object that can list and get AzureResources.
func (s *azureResourceLister) AzureResources(namespace string) AzureResourceNamespaceLister {
	return azureResourceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AzureResourceNamespaceLister helps list and get AzureResources.
type AzureResourceNamespaceLister interface {
	// List lists all AzureResources in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.AzureResource, err error)
	// Get retrieves the AzureResource from the indexer for a given namespace and name.
	Get(name string) (*v1.AzureResource, error)
	AzureResourceNamespaceListerExpansion
}

// azureResourceNamespaceLister implements the AzureResourceNamespaceLister
// interface.
type azureResourceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AzureResources in the indexer for a given namespace.
func (s azureResourceNamespaceLister) List(selector labels.Selector) (ret []*v1.AzureResource, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AzureResource))
	})
	return ret, err
}

// Get retrieves the AzureResource from the indexer for a given namespace and name.
func (s azureResourceNamespaceLister) Get(name string) (*v1.AzureResource, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("azureresource"), name)
	}
	return obj.(*v1.AzureResource), nil
}
