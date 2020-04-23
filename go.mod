module github.com/rawlingsj/service-controller-example

require (
	cloud.google.com/go v0.53.0
	github.com/cloudflare/cfssl v0.0.0-20190409034051-768cd563887f
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/jenkins-x-labs/gsm-controller v0.0.50 // indirect
	github.com/jenkins-x/jx-logging v0.0.1
	github.com/magiconair/properties v1.8.1
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v0.0.6
	github.com/spf13/viper v1.6.2
	github.com/stretchr/testify v1.4.0
	google.golang.org/genproto v0.0.0-20200212174721-66ed5ce911ce
	k8s.io/api v0.0.0-20191004102349-159aefb8556b
	k8s.io/apimachinery v0.0.0-20191004074956-c5d2f014d689
	k8s.io/client-go v11.0.1-0.20191004102930-01520b8320fc+incompatible
)

replace github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 v22.0.0

replace github.com/russross/blackfriday => github.com/russross/blackfriday v1.5.2

replace github.com/googleapis/gnostic => github.com/googleapis/gnostic v0.4.0

go 1.12
