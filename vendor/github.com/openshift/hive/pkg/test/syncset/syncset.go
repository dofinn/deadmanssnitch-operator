package syncset

import (
	hivev1 "github.com/openshift/hive/pkg/apis/hive/v1alpha1"
	"github.com/openshift/hive/pkg/test/generic"
)

// Option defines a function signature for any function that wants to be passed into Build
type Option func(*hivev1.SyncSet)

// Build runs each of the functions passed in to generate the object.
func Build(opts ...Option) *hivev1.SyncSet {
	retval := &hivev1.SyncSet{}
	for _, o := range opts {
		o(retval)
	}

	return retval
}

// Generic allows common functions applicable to all objects to be used as Options to Build
func Generic(opt generic.Option) Option {
	return func(syncSet *hivev1.SyncSet) {
		opt(syncSet)
	}
}
