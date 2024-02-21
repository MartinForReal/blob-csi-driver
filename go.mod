module sigs.k8s.io/blob-csi-driver

go 1.21

toolchain go1.21.4

require (
	github.com/Azure/azure-sdk-for-go v68.0.0+incompatible
	github.com/Azure/azure-sdk-for-go/sdk/azidentity v1.5.1
	github.com/Azure/go-autorest/autorest v0.11.29
	github.com/Azure/go-autorest/autorest/adal v0.9.23
	github.com/Azure/go-autorest/autorest/to v0.4.0 // indirect
	github.com/container-storage-interface/spec v1.8.0
	github.com/gofrs/uuid v4.4.0+incompatible // indirect
	github.com/golang/protobuf v1.5.3
	github.com/imdario/mergo v0.3.9 // indirect
	github.com/kubernetes-csi/csi-lib-utils v0.13.0
	github.com/onsi/gomega v1.31.1
	github.com/pborman/uuid v1.2.1
	github.com/pelletier/go-toml v1.9.5
	github.com/stretchr/testify v1.8.4
	golang.org/x/net v0.20.0
	google.golang.org/grpc v1.59.0
	google.golang.org/protobuf v1.31.0
	k8s.io/api v0.29.2
	k8s.io/apimachinery v0.29.2
	k8s.io/client-go v0.29.2
	k8s.io/component-base v0.29.0
	k8s.io/klog/v2 v2.110.1
	k8s.io/kubernetes v1.29.0
	k8s.io/mount-utils v0.29.0
	k8s.io/utils v0.0.0-20231127182322-b307cd553661
	sigs.k8s.io/cloud-provider-azure v1.27.1-0.20240220055818-b76605c9c708
	sigs.k8s.io/yaml v1.4.0
)

require (
	github.com/Azure/azure-sdk-for-go/sdk/azcore v1.9.2
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault v1.4.0
	github.com/Azure/azure-sdk-for-go/sdk/storage/azblob v1.1.0
	github.com/go-ini/ini v1.67.0
	github.com/onsi/ginkgo/v2 v2.15.0
	github.com/pkg/errors v0.9.1
	github.com/satori/go.uuid v1.2.0
	k8s.io/apiserver v0.29.0
	k8s.io/pod-security-admission v0.29.0
	sigs.k8s.io/cloud-provider-azure/pkg/azclient/configloader v0.0.0-20240218020800-ba9f211f36bf
)

require (
	github.com/distribution/reference v0.5.0 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/mxk/go-flowrate v0.0.0-20140419014527-cca7078d478f // indirect
	go.uber.org/mock v0.4.0
)

require (
	github.com/Azure/azure-sdk-for-go/sdk/internal v1.5.2 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v5 v5.5.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry v1.2.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v4 v4.7.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v4 v4.3.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/privatedns/armprivatedns v1.2.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources v1.2.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage v1.5.0 // indirect
	github.com/Azure/go-autorest v14.2.0+incompatible // indirect
	github.com/Azure/go-autorest/autorest/date v0.3.0 // indirect
	github.com/Azure/go-autorest/autorest/mocks v0.4.2 // indirect
	github.com/Azure/go-autorest/autorest/validation v0.3.1 // indirect
	github.com/Azure/go-autorest/logger v0.2.1 // indirect
	github.com/Azure/go-autorest/tracing v0.6.0 // indirect
	github.com/AzureAD/microsoft-authentication-library-for-go v1.2.1 // indirect
	github.com/NYTimes/gziphandler v1.1.1 // indirect
	github.com/antlr/antlr4/runtime/Go/antlr/v4 v4.0.0-20230305170008-8188dc5388df // indirect
	github.com/asaskevich/govalidator v0.0.0-20190424111038-f61b66f89f4a // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/blang/semver/v4 v4.0.0 // indirect
	github.com/cenkalti/backoff/v4 v4.2.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/coreos/go-semver v0.3.1 // indirect
	github.com/coreos/go-systemd/v22 v22.5.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/emicklei/go-restful/v3 v3.11.0 // indirect
	github.com/evanphx/json-patch v5.9.0+incompatible // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/go-logr/logr v1.4.1 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-openapi/jsonpointer v0.19.6 // indirect
	github.com/go-openapi/jsonreference v0.20.2 // indirect
	github.com/go-openapi/swag v0.22.4 // indirect
	github.com/go-task/slim-sprig v0.0.0-20230315185526-52ccab3ef572 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang-jwt/jwt/v4 v4.5.0 // indirect
	github.com/golang-jwt/jwt/v5 v5.2.0 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/google/cel-go v0.17.7 // indirect
	github.com/google/gnostic-models v0.6.8 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/google/pprof v0.0.0-20230602010524-ada837c32108 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.16.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/kylelemons/godebug v1.1.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
	github.com/moby/spdystream v0.2.0 // indirect
	github.com/moby/sys/mountinfo v0.6.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/selinux v1.11.0 // indirect
	github.com/pkg/browser v0.0.0-20240102092130-5ac0b6a4141c // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_golang v1.16.0 // indirect
	github.com/prometheus/client_model v0.4.0 // indirect
	github.com/prometheus/common v0.44.0 // indirect
	github.com/prometheus/procfs v0.10.1 // indirect
	github.com/spf13/cobra v1.8.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/stoewer/go-strcase v1.2.0 // indirect
	go.etcd.io/etcd/api/v3 v3.5.10 // indirect
	go.etcd.io/etcd/client/pkg/v3 v3.5.10 // indirect
	go.etcd.io/etcd/client/v3 v3.5.10 // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.46.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.46.0 // indirect
	go.opentelemetry.io/otel v1.20.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.20.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.20.0 // indirect
	go.opentelemetry.io/otel/metric v1.20.0 // indirect
	go.opentelemetry.io/otel/sdk v1.20.0 // indirect
	go.opentelemetry.io/otel/trace v1.20.0 // indirect
	go.opentelemetry.io/proto/otlp v1.0.0 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.19.0 // indirect
	golang.org/x/crypto v0.19.0 // indirect
	golang.org/x/exp v0.0.0-20230905200255-921286631fa9 // indirect
	golang.org/x/oauth2 v0.11.0 // indirect
	golang.org/x/sync v0.6.0
	golang.org/x/sys v0.17.0 // indirect
	golang.org/x/term v0.17.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	golang.org/x/time v0.5.0 // indirect
	golang.org/x/tools v0.16.1 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20231030173426-d783a09b4405 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20231016165738-49dd2c1f3d0b // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231106174013-bbf56f31fb17 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.2.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	k8s.io/apiextensions-apiserver v0.0.0 // indirect
	k8s.io/cloud-provider v0.29.0 // indirect
	k8s.io/component-helpers v0.29.0 // indirect
	k8s.io/controller-manager v0.29.0 // indirect
	k8s.io/kms v0.29.0 // indirect
	k8s.io/kube-openapi v0.0.0-20231010175941-2dd684a91f00 // indirect
	k8s.io/kubectl v0.0.0 // indirect
	k8s.io/kubelet v0.29.0 // indirect
	sigs.k8s.io/apiserver-network-proxy/konnectivity-client v0.28.0 // indirect
	sigs.k8s.io/cloud-provider-azure/pkg/azclient v0.0.0-20240218020800-ba9f211f36bf
	sigs.k8s.io/json v0.0.0-20221116044647-bc3834ca7abd // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.4.1 // indirect
)

replace (
	github.com/niemeyer/pretty => github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e
	go.etcd.io/etcd => go.etcd.io/etcd v0.0.0-20200410171415-59f5fb25a533
	k8s.io/api => k8s.io/api v0.29.0
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.29.0
	k8s.io/apimachinery => k8s.io/apimachinery v0.29.0
	k8s.io/apiserver => k8s.io/apiserver v0.29.0
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.29.0
	k8s.io/client-go => k8s.io/client-go v0.29.0
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.29.0
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.29.0
	k8s.io/code-generator => k8s.io/code-generator v0.29.0
	k8s.io/component-base => k8s.io/component-base v0.29.0
	k8s.io/component-helpers => k8s.io/component-helpers v0.29.0
	k8s.io/controller-manager => k8s.io/controller-manager v0.29.0
	k8s.io/cri-api => k8s.io/cri-api v0.29.0
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.29.0
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.29.0
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.29.0
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.29.0
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.29.0
	k8s.io/kubectl => k8s.io/kubectl v0.29.0
	k8s.io/kubelet => k8s.io/kubelet v0.29.0
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.29.0
	k8s.io/metrics => k8s.io/metrics v0.29.0
	k8s.io/mount-utils => k8s.io/mount-utils v0.29.0
	k8s.io/pod-security-admission => k8s.io/pod-security-admission v0.29.0
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.29.0
	k8s.io/sample-cli-plugin => k8s.io/sample-cli-plugin v0.29.0
	k8s.io/sample-controller => k8s.io/sample-controller v0.29.0
)
