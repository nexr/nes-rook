module github.com/rook/rook

go 1.13

require (
	github.com/aws/aws-sdk-go v1.35.24
	github.com/banzaicloud/k8s-objectmatcher v1.1.0
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f
	github.com/corpix/uarand v0.1.1 // indirect
	github.com/davecgh/go-spew v1.1.1
	github.com/ghodss/yaml v1.0.1-0.20190212211648-25d852aebe32
	github.com/go-ini/ini v1.51.1
	github.com/go-logr/logr v0.3.0 // indirect
	github.com/go-sql-driver/mysql v1.5.0
	github.com/google/go-cmp v0.5.3
	github.com/google/uuid v1.1.2
	github.com/hashicorp/vault/api v1.0.5-0.20200902155336-f9d5ce5a171a
	github.com/icrowley/fake v0.0.0-20180203215853-4178557ae428
	github.com/k8snetworkplumbingwg/network-attachment-definition-client v1.1.0
	github.com/kube-object-storage/lib-bucket-provisioner v0.0.0-20210127170128-83a4fdf6edd6
	github.com/libopenstorage/secrets v0.0.0-20201006135900-af310b01fe47
	github.com/miekg/dns v1.1.34 // indirect
	github.com/openshift/cluster-api v0.0.0-20191129101638-b09907ac6668
	github.com/openshift/machine-api-operator v0.2.1-0.20190903202259-474e14e4965a
	github.com/pkg/errors v0.9.1
	github.com/prometheus-operator/prometheus-operator v0.43.0
	github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring v0.43.0
	github.com/spf13/cobra v1.1.1
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.6.1
	github.com/tevino/abool v1.2.0
	github.com/yanniszark/go-nodetool v0.0.0-20191206125106-cd8f91fa16be
	gopkg.in/ini.v1 v1.57.0
	k8s.io/api v0.20.0
	k8s.io/apiextensions-apiserver v0.20.0
	k8s.io/apimachinery v0.20.0
	k8s.io/apiserver v0.20.0
	k8s.io/client-go v12.0.0+incompatible
	k8s.io/cloud-provider v0.20.0
	k8s.io/code-generator v0.20.1
	k8s.io/kube-controller-manager v0.20.0
	k8s.io/kubernetes v1.20.0
	k8s.io/utils v0.0.0-20201110183641-67b214c5f920
	sigs.k8s.io/controller-runtime v0.6.3
	sigs.k8s.io/kustomize/kyaml v0.2.0
	sigs.k8s.io/sig-storage-lib-external-provisioner/v6 v6.1.0
)

// This looks "horrible", but is due to the Rook including k8s.io/kubernetes directly which is not recommended,
// ee https://github.com/kubernetes/kubernetes/issues/79384#issuecomment-505725449
replace (
	github.com/Knetic/govaluate => github.com/Knetic/govaluate v0.0.0-20171022003610-9aa49832a739
	github.com/boombuler/barcode => github.com/boombuler/barcode v0.0.0-20190219062509-6c824513bacc
	github.com/briankassouf/jose => github.com/briankassouf/jose v0.0.0-20180619214549-d2569464773f
	github.com/docker/docker => github.com/docker/docker v0.0.0-20200916142827-bd33bbf0497b
	github.com/dustmop/soup => github.com/dustmop/soup v0.0.0-20190516214245-38228baa104e
	github.com/fogleman/gg => github.com/fogleman/gg v0.0.0-20190220221249-0403632d5b90
	github.com/ghodss/yaml => github.com/ghodss/yaml v0.0.0-20190212211648-25d852aebe32
	github.com/googleapis/gnostic => github.com/googleapis/gnostic v0.4.1
	github.com/grpc-ecosystem/go-grpc-middleware => github.com/grpc-ecosystem/go-grpc-middleware v1.1.0
	github.com/hashicorp/go-raftchunking => github.com/hashicorp/go-raftchunking v0.0.0-20191002164813-7e9e8525653a
	github.com/hashicorp/vault/sdk => github.com/hashicorp/vault/sdk v0.2.0
	github.com/influxdata/roaring => github.com/influxdata/roaring v0.0.0-20180809181101-fc520f41fab6
	github.com/klauspost/pgzip => github.com/klauspost/pgzip v0.0.0-20170402124221-0bf5dcad4ada
	github.com/opencontainers/go-digest => github.com/opencontainers/go-digest v1.0.0
	github.com/pkg/errors => github.com/pkg/errors v0.0.0-20171018195549-f15c970de5b7
	github.com/prometheus/prometheus => github.com/prometheus/prometheus v0.0.0-20201015110737-0a7fdd3b7696
	go.etcd.io/bbolt => go.etcd.io/bbolt v1.3.5
	go.etcd.io/etcd => go.etcd.io/etcd v0.0.0-20200910180754-dd1b699fc489
	golang.org/x/mod => golang.org/x/mod v0.2.0
	golang.org/x/text => golang.org/x/text v0.3.3
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
	google.golang.org/protobuf => google.golang.org/protobuf v1.24.0
	k8s.io/api => k8s.io/api v0.20.0
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.20.0
	k8s.io/apimachinery => k8s.io/apimachinery v0.20.0
	k8s.io/apiserver => k8s.io/apiserver v0.20.0
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.20.0
	k8s.io/client-go => k8s.io/client-go v0.20.0
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.20.0
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.20.0
	k8s.io/code-generator => k8s.io/code-generator v0.20.0
	k8s.io/component-base => k8s.io/component-base v0.20.0
	k8s.io/component-helpers => k8s.io/component-helpers v0.20.0
	k8s.io/controller-manager => k8s.io/controller-manager v0.20.0
	k8s.io/cri-api => k8s.io/cri-api v0.20.0
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.20.0
	k8s.io/gengo => k8s.io/gengo v0.0.0-20201113003025-83324d819ded
	k8s.io/heapster => k8s.io/heapster v1.2.0-beta.1
	k8s.io/klog/v2 => k8s.io/klog/v2 v2.4.0
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.20.0
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.20.0
	k8s.io/kube-openapi => k8s.io/kube-openapi v0.0.0-20201113171705-d219536bb9fd
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.20.0
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.20.0
	k8s.io/kubectl => k8s.io/kubectl v0.20.0
	k8s.io/kubelet => k8s.io/kubelet v0.20.0
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.20.0
	k8s.io/metrics => k8s.io/metrics v0.20.0
	k8s.io/mount-utils => k8s.io/mount-utils v0.20.0
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.20.0
	k8s.io/system-validators => k8s.io/system-validators v1.2.0
	k8s.io/utils => k8s.io/utils v0.0.0-20201110183641-67b214c5f920
	modernc.org/cc => gitlab.com/cznic/cc v1.0.1
	modernc.org/golex => gitlab.com/cznic/golex v1.0.0
	modernc.org/internal => gitlab.com/cznic/internal v1.0.0
	modernc.org/ir => gitlab.com/cznic/ir v1.0.0
	modernc.org/mathutil => gitlab.com/cznic/mathutil v1.0.0
	modernc.org/strutil => gitlab.com/cznic/strutil v1.0.0
	modernc.org/token => gitlab.com/cznic/token v1.0.0
	modernc.org/xc => gitlab.com/cznic/xc v1.0.0
)
