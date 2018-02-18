/*
Copyright 2018 The Kubed Authors.

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
package internalversion

import (
	kubed "github.com/appscode/kubed/apis/kubed"
	scheme "github.com/appscode/kubed/client/clientset/internalversion/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	rest "k8s.io/client-go/rest"
)

// SearchResultsGetter has a method to return a SearchResultInterface.
// A group's client should implement this interface.
type SearchResultsGetter interface {
	SearchResults(namespace string) SearchResultInterface
}

// SearchResultInterface has methods to work with SearchResult resources.
type SearchResultInterface interface {
	Get(name string, options v1.GetOptions) (*kubed.SearchResult, error)
	SearchResultExpansion
}

// searchResults implements SearchResultInterface
type searchResults struct {
	client rest.Interface
	ns     string
}

// newSearchResults returns a SearchResults
func newSearchResults(c *KubedClient, namespace string) *searchResults {
	return &searchResults{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the searchResult, and returns the corresponding searchResult object, and an error if there is any.
func (c *searchResults) Get(name string, options v1.GetOptions) (result *kubed.SearchResult, err error) {
	result = &kubed.SearchResult{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("searchresults").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}
