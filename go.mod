module github.com/docker/docker

go 1.15

replace (
	github.com/containerd/containerd v1.4.0-0 => github.com/containerd/containerd v1.4.1
	github.com/docker/docker => ./
	github.com/docker/libkv => ../libkv
	github.com/docker/libnetwork => ../libnetwork
	github.com/moby/buildkit => github.com/moby/buildkit v0.7.1-0.20200718032743-4d1f260e8490
	github.com/hashicorp/go-immutable-radix => github.com/tonistiigi/go-immutable-radix v0.0.0-20170803185627-826af9ccf0fe
)

require (
	cloud.google.com/go v0.70.0
	cloud.google.com/go/logging v1.1.1
	github.com/BurntSushi/toml v0.3.1
	github.com/Graylog2/go-gelf v0.0.0-20191017102106-1550ee647df0
	github.com/Microsoft/go-winio v0.4.15-0.20200908182639-5b44b70ab3ab
	github.com/Microsoft/hcsshim v0.8.10-0.20200609165715-9dcb42f10021
	github.com/Microsoft/opengcs v0.3.10-0.20190304234800-a10967154e14
	github.com/RackSec/srslog v0.0.0-20180709174129-a4725f04ec91
	github.com/aws/aws-sdk-go v1.28.11
	github.com/bsphere/le_go v0.0.0-20170215134836-7a984a84b549
	github.com/containerd/cgroups v0.0.0-20200710171044-318312a37340
	github.com/containerd/containerd v1.4.1
	github.com/containerd/continuity v0.0.0-20200710164510-efbc4488d8fe
	github.com/containerd/fifo v0.0.0-20200410184934-f15a3290365b
	github.com/containerd/typeurl v1.0.1
	github.com/coreos/go-systemd/v22 v22.1.0
	github.com/docker/distribution v2.7.1+incompatible
	github.com/docker/go-connections v0.4.0
	github.com/docker/go-metrics v0.0.1
	github.com/docker/go-units v0.4.0
	github.com/docker/libnetwork v0.8.0-dev.2.0.20200917202933-d0951081b35f
	github.com/docker/libtrust v0.0.0-20150526203908-9cbd2a1374f4
	github.com/fluent/fluent-logger-golang v1.4.0
	github.com/fsnotify/fsnotify v1.4.9
	github.com/gogo/protobuf v1.3.1
	github.com/golang/gddo v0.0.0-20190904175337-72a348e765d2
	github.com/google/uuid v1.1.1
	github.com/gorilla/mux v1.8.0
	github.com/hashicorp/go-immutable-radix v1.0.0
	github.com/hashicorp/go-memdb v0.0.0-20161216180745-cb9a474f84cc
	github.com/imdario/mergo v0.3.9
	github.com/moby/buildkit v0.7.1-0.20200718032743-4d1f260e8490
	github.com/moby/locker v1.0.1
	github.com/moby/sys/mount v0.1.0
	github.com/moby/sys/mountinfo v0.1.3
	github.com/moby/term v0.0.0-20200915141129-7f0af18e79f2
	github.com/morikuni/aec v1.0.0
	github.com/opencontainers/go-digest v1.0.0
	github.com/opencontainers/image-spec v1.0.1
	github.com/opencontainers/runc v1.0.0-rc92
	github.com/opencontainers/runtime-spec v1.0.3-0.20200728170252-4d89ac9fbff6
	github.com/opencontainers/selinux v1.6.0
	github.com/philhofer/fwd v1.1.0 // indirect
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.6.0
	github.com/sirupsen/logrus v1.7.0
	github.com/spf13/cobra v1.0.0
	github.com/spf13/pflag v1.0.5
	github.com/syndtr/gocapability v0.0.0-20200815063812-42c35b437635
	github.com/tchap/go-patricia v2.3.0+incompatible
	github.com/tinylib/msgp v1.1.2 // indirect
	github.com/tonistiigi/fsutil v0.0.0-20200512175118-ae3a8d753069
	github.com/vbatts/tar-split v0.11.1
	github.com/vishvananda/netlink v1.1.0
	go.etcd.io/bbolt v1.3.5
	golang.org/x/net v0.0.0-20201010224723-4f7140c49acb
	golang.org/x/sync v0.0.0-20200625203802-6e8e738ad208
	golang.org/x/sys v0.0.0-20201015000850-e3ed0017c211
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0
	google.golang.org/genproto v0.0.0-20201019141844-1ed22bb0c154
	google.golang.org/grpc v1.32.0
	gotest.tools/v3 v3.0.2
)
