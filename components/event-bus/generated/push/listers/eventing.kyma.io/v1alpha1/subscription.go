// This file was automatically generated by lister-gen

package v1alpha1

import (
	v1alpha1 "github.com/kyma-project/kyma/components/event-bus/api/push/eventing.kyma.cx/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// SubscriptionLister helps list Subscriptions.
type SubscriptionLister interface {
	// List lists all Subscriptions in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Subscription, err error)
	// Subscriptions returns an object that can list and get Subscriptions.
	Subscriptions(namespace string) SubscriptionNamespaceLister
	SubscriptionListerExpansion
}

// subscriptionLister implements the SubscriptionLister interface.
type subscriptionLister struct {
	indexer cache.Indexer
}

// NewSubscriptionLister returns a new SubscriptionLister.
func NewSubscriptionLister(indexer cache.Indexer) SubscriptionLister {
	return &subscriptionLister{indexer: indexer}
}

// List lists all Subscriptions in the indexer.
func (s *subscriptionLister) List(selector labels.Selector) (ret []*v1alpha1.Subscription, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Subscription))
	})
	return ret, err
}

// Subscriptions returns an object that can list and get Subscriptions.
func (s *subscriptionLister) Subscriptions(namespace string) SubscriptionNamespaceLister {
	return subscriptionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SubscriptionNamespaceLister helps list and get Subscriptions.
type SubscriptionNamespaceLister interface {
	// List lists all Subscriptions in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Subscription, err error)
	// Get retrieves the Subscription from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Subscription, error)
	SubscriptionNamespaceListerExpansion
}

// subscriptionNamespaceLister implements the SubscriptionNamespaceLister
// interface.
type subscriptionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Subscriptions in the indexer for a given namespace.
func (s subscriptionNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Subscription, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Subscription))
	})
	return ret, err
}

// Get retrieves the Subscription from the indexer for a given namespace and name.
func (s subscriptionNamespaceLister) Get(name string) (*v1alpha1.Subscription, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("subscription"), name)
	}
	return obj.(*v1alpha1.Subscription), nil
}
