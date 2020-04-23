package main

import (
	"os"

	v1 "k8s.io/api/core/v1"

	"github.com/jenkins-x-labs/gsm-controller/pkg/shared"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"

	"github.com/jenkins-x/jx-logging/pkg/log"
	// Uncomment the following line to load the gcp plugin (only required to authenticate against GKE clusters).
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

func main() {

	log.Logger().Info("hey")
	o := options{}
	err := o.watch()
	if err != nil {
		log.Logger().Fatalf("failed %v", err)
		os.Exit(-1)
	}
	os.Exit(0)
}

type options struct {
}

func (o options) watch() error {
	f := shared.NewFactory()
	config, err := f.CreateKubeConfig()
	if err != nil {
		return errors.Wrap(err, "failed to get kubernetes config")
	}

	kubeclient, err := kubernetes.NewForConfig(config)
	if err != nil {
		return errors.Wrap(err, "failed to create kubernetes clientset")
	}

	namespace := shared.CurrentNamespace()
	factory := informers.NewSharedInformerFactoryWithOptions(kubeclient, 0, informers.WithNamespace(namespace))

	informer := factory.Core().V1().Services().Informer()

	stopper := make(chan struct{})
	defer close(stopper)

	defer runtime.HandleCrash()

	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    o.onAdd,
		UpdateFunc: o.onUpdate,
		DeleteFunc: o.onDelete,
	})

	go informer.Run(stopper)

	if !cache.WaitForCacheSync(stopper, informer.HasSynced) {
		return errors.New("Timed out waiting for caches to sync")
	}
	<-stopper
	return nil
}

func (o options) onAdd(obj interface{}) {
	svc := obj.(*v1.Service)
	log.Logger().Infof("added service %s", svc.Name)
}

func (o options) onUpdate(oldObj interface{}, newObj interface{}) {
	svc := newObj.(*v1.Service)
	log.Logger().Infof("updated service %s", svc.Name)
}

func (o options) onDelete(obj interface{}) {
	svc := obj.(*v1.Service)
	log.Logger().Infof("deleted service %s", svc.Name)
}
