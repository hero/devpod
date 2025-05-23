// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	managementv1 "github.com/loft-sh/api/v4/pkg/apis/management/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
)

// IngressAuthTokenLister helps list IngressAuthTokens.
// All objects returned here must be treated as read-only.
type IngressAuthTokenLister interface {
	// List lists all IngressAuthTokens in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*managementv1.IngressAuthToken, err error)
	// Get retrieves the IngressAuthToken from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*managementv1.IngressAuthToken, error)
	IngressAuthTokenListerExpansion
}

// ingressAuthTokenLister implements the IngressAuthTokenLister interface.
type ingressAuthTokenLister struct {
	listers.ResourceIndexer[*managementv1.IngressAuthToken]
}

// NewIngressAuthTokenLister returns a new IngressAuthTokenLister.
func NewIngressAuthTokenLister(indexer cache.Indexer) IngressAuthTokenLister {
	return &ingressAuthTokenLister{listers.New[*managementv1.IngressAuthToken](indexer, managementv1.Resource("ingressauthtoken"))}
}
