// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	versioned "k8c.io/kubermatic/v2/pkg/crd/client/clientset/versioned"
	internalinterfaces "k8c.io/kubermatic/v2/pkg/crd/client/informers/externalversions/internalinterfaces"
	v1 "k8c.io/kubermatic/v2/pkg/crd/client/listers/kubermatic/v1"
	kubermaticv1 "k8c.io/kubermatic/v2/pkg/crd/kubermatic/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// MLAAdminSettingInformer provides access to a shared informer and lister for
// MLAAdminSettings.
type MLAAdminSettingInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.MLAAdminSettingLister
}

type mLAAdminSettingInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewMLAAdminSettingInformer constructs a new informer for MLAAdminSetting type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewMLAAdminSettingInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredMLAAdminSettingInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredMLAAdminSettingInformer constructs a new informer for MLAAdminSetting type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredMLAAdminSettingInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubermaticV1().MLAAdminSettings(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubermaticV1().MLAAdminSettings(namespace).Watch(context.TODO(), options)
			},
		},
		&kubermaticv1.MLAAdminSetting{},
		resyncPeriod,
		indexers,
	)
}

func (f *mLAAdminSettingInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredMLAAdminSettingInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *mLAAdminSettingInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&kubermaticv1.MLAAdminSetting{}, f.defaultInformer)
}

func (f *mLAAdminSettingInformer) Lister() v1.MLAAdminSettingLister {
	return v1.NewMLAAdminSettingLister(f.Informer().GetIndexer())
}
