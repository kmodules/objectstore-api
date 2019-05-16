module kmodules.xyz/objectstore-api

go 1.12

require (
	cloud.google.com/go v0.38.0 // indirect
	contrib.go.opencensus.io/exporter/ocagent v0.5.0 // indirect
	github.com/Azure/azure-sdk-for-go v29.0.0+incompatible // indirect
	github.com/Azure/go-autorest v12.0.0+incompatible // indirect
	github.com/appscode/go v0.0.0-20190424183524-60025f1135c9
	github.com/appscode/osm v0.11.0
	github.com/aws/aws-sdk-go v1.19.26
	github.com/dnaeon/go-vcr v1.0.1 // indirect
	github.com/ghodss/yaml v1.0.0
	github.com/go-openapi/jsonpointer v0.19.0 // indirect
	github.com/go-openapi/jsonreference v0.19.0 // indirect
	github.com/go-openapi/spec v0.19.0
	github.com/go-openapi/swag v0.19.0 // indirect
	github.com/graymeta/stow v0.0.0-00010101000000-000000000000
	github.com/mailru/easyjson v0.0.0-20190403194419-1ea4449da983 // indirect
	github.com/ncw/swift v1.0.47 // indirect
	github.com/pkg/errors v0.8.1
	golang.org/x/sync v0.0.0-20190423024810-112230192c58 // indirect
	google.golang.org/genproto v0.0.0-20190502173448-54afdca5d873 // indirect
	k8s.io/api v0.0.0-20190503110853-61630f889b3c
	k8s.io/apimachinery v0.0.0-20190508063446-a3da69d3723c
	k8s.io/client-go v11.0.0+incompatible
	k8s.io/kube-openapi v0.0.0-20190502190224-411b2483e503
)

replace (
	github.com/graymeta/stow => github.com/appscode/stow v0.0.0-20190506085026-ca5baa008ea3
	k8s.io/api => k8s.io/api v0.0.0-20190313235455-40a48860b5ab
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.0.0-20190315093550-53c4693659ed
	k8s.io/apimachinery => github.com/kmodules/apimachinery v0.0.0-20190508045248-a52a97a7a2bf
	k8s.io/apiserver => github.com/kmodules/apiserver v0.0.0-20190508082252-8397d761d4b5
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.0.0-20190314001948-2899ed30580f
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.0.0-20190314002645-c892ea32361a
	k8s.io/klog => k8s.io/klog v0.3.0
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.0.0-20190314000639-da8327669ac5
	k8s.io/kube-openapi => k8s.io/kube-openapi v0.0.0-20190228160746-b3a7cee44a30
	k8s.io/metrics => k8s.io/metrics v0.0.0-20190314001731-1bd6a4002213
	k8s.io/utils => k8s.io/utils v0.0.0-20190221042446-c2654d5206da
)
