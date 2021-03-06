# kubesphere CRD sample

## Step 1

Install `dep`, `kustomize` and `kubebuilder`. [Link](https://book.kubebuilder.io/quick_start.html#Installation)

## Step 2

Run `kubebuilder init --domain kubesphere.io --owner "The KubeSphere Authors"` to init project

## Step 3

Run `kubebuilder create api --group iam --version v1alpha2 --kind Workspace --namespaced=false` to create api.

```bash
$ kubebuilder create api --group iam --version v1alpha2 --kind Workspace --namespaced=false
Create Resource under pkg/apis [y/n]?
y
Create Controller under pkg/controller [y/n]?
y
Writing scaffold for you to edit...
pkg/apis/iam/v1alpha2/workspace_types.go
pkg/apis/iam/v1alpha2/workspace_types_test.go
pkg/controller/workspace/workspace_controller.go
pkg/controller/workspace/workspace_controller_test.go
Running make...
go generate ./pkg/... ./cmd/...
go fmt ./pkg/... ./cmd/...
go vet ./pkg/... ./cmd/...
go run vendor/sigs.k8s.io/controller-tools/cmd/controller-gen/main.go all
CRD manifests generated under '/Users/runzexia/go/src/github.com/runzexia/kubesphere-crd-sample/config/crds'
RBAC manifests generated under '/Users/runzexia/go/src/github.com/runzexia/kubesphere-crd-sample/config/rbac'
go test ./pkg/... ./cmd/... -coverprofile cover.out
?   	github.com/runzexia/kubesphere-crd-sample/pkg/apis	[no test files]
?   	github.com/runzexia/kubesphere-crd-sample/pkg/apis/iam	[no test files]
ok  	github.com/runzexia/kubesphere-crd-sample/pkg/apis/iam/v1alpha2	9.310s	coverage: 23.4% of statements
?   	github.com/runzexia/kubesphere-crd-sample/pkg/controller	[no test files]
ok  	github.com/runzexia/kubesphere-crd-sample/pkg/controller/workspace	10.674s	coverage: 67.6% of statements
?   	github.com/runzexia/kubesphere-crd-sample/pkg/webhook	[no test files]
?   	github.com/runzexia/kubesphere-crd-sample/cmd/manager	[no test files]
go build -o bin/manager github.com/runzexia/kubesphere-crd-sample/cmd/manager
```

Run `kubebuilder create api --group devops --version v1alpha2 --kind DevOpsProject --namespaced=false` to create api.


```bash
$ kubebuilder create api --group devops --version v1alpha2 --kind DevOpsProject --namespaced=false
Create Resource under pkg/apis [y/n]?
y
Create Controller under pkg/controller [y/n]?
y
Writing scaffold for you to edit...
pkg/apis/devops/v1alpha2/devopsproject_types.go
pkg/apis/devops/v1alpha2/devopsproject_types_test.go
pkg/controller/devopsproject/devopsproject_controller.go
pkg/controller/devopsproject/devopsproject_controller_test.go
Running make...
go generate ./pkg/... ./cmd/...
go fmt ./pkg/... ./cmd/...
go vet ./pkg/... ./cmd/...
go run vendor/sigs.k8s.io/controller-tools/cmd/controller-gen/main.go all
CRD manifests generated under '/Users/runzexia/go/src/github.com/runzexia/kubesphere-crd-sample/config/crds'
RBAC manifests generated under '/Users/runzexia/go/src/github.com/runzexia/kubesphere-crd-sample/config/rbac'
go test ./pkg/... ./cmd/... -coverprofile cover.out
?   	github.com/runzexia/kubesphere-crd-sample/pkg/apis	[no test files]
?   	github.com/runzexia/kubesphere-crd-sample/pkg/apis/devops	[no test files]
ok  	github.com/runzexia/kubesphere-crd-sample/pkg/apis/devops/v1alpha2	7.816s	coverage: 23.4% of statements
?   	github.com/runzexia/kubesphere-crd-sample/pkg/controller	[no test files]
ok  	github.com/runzexia/kubesphere-crd-sample/pkg/controller/devopsproject	9.255s	coverage: 67.6% of statements
?   	github.com/runzexia/kubesphere-crd-sample/pkg/webhook	[no test files]
?   	github.com/runzexia/kubesphere-crd-sample/cmd/manager	[no test files]
```

## Step 4

edit `pkg/apis/devops/v1alpha2/devopsproject_types.go` and `pkg/apis/iam/v1alpha2/workspace_types.go`, run `make` to regenerate code

```bash
$ make
go generate ./pkg/... ./cmd/...
go fmt ./pkg/... ./cmd/...
go vet ./pkg/... ./cmd/...
go run vendor/sigs.k8s.io/controller-tools/cmd/controller-gen/main.go all
CRD manifests generated under '/Users/runzexia/go/src/github.com/runzexia/kubesphere-crd-sample/config/crds' 
RBAC manifests generated under '/Users/runzexia/go/src/github.com/runzexia/kubesphere-crd-sample/config/rbac' 
go test ./pkg/... ./cmd/... -coverprofile cover.out
?       github.com/runzexia/kubesphere-crd-sample/pkg/apis      [no test files]
?       github.com/runzexia/kubesphere-crd-sample/pkg/apis/devops       [no test files]
ok      github.com/runzexia/kubesphere-crd-sample/pkg/apis/devops/v1alpha2      9.306s  coverage: 23.4% of statements
?       github.com/runzexia/kubesphere-crd-sample/pkg/apis/iam  [no test files]
ok      github.com/runzexia/kubesphere-crd-sample/pkg/apis/iam/v1alpha2 9.032s  coverage: 23.4% of statements
?       github.com/runzexia/kubesphere-crd-sample/pkg/controller        [no test files]
ok      github.com/runzexia/kubesphere-crd-sample/pkg/controller/devopsproject  10.713s coverage: 73.0% of statements
ok      github.com/runzexia/kubesphere-crd-sample/pkg/controller/workspace      11.022s coverage: 73.0% of statements
?       github.com/runzexia/kubesphere-crd-sample/pkg/webhook   [no test files]
?       github.com/runzexia/kubesphere-crd-sample/cmd/manager   [no test files]
go build -o bin/manager github.com/runzexia/kubesphere-crd-sample/cmd/manager

```

## Step 5

edit `pkg/apis/apis.go` add openapi-gen command to generate `openapi_generated.go`
`openapi_generated.go` will be used when generating `swagger.json`

```go
// Generate openapi for apis
//go:generate go run ../../vendor/k8s.io/kube-openapi/cmd/openapi-gen/openapi-gen.go -O openapi_generated -i ./iam/v1alpha2 -p github.com/runzexia/kubesphere-crd-sample/pkg/apis/iam/v1alpha2 -h ../../hack/boilerplate.go.txt
//go:generate go run ../../vendor/k8s.io/kube-openapi/cmd/openapi-gen/openapi-gen.go -O openapi_generated -i ./devops/v1alpha2 -p github.com/runzexia/kubesphere-crd-sample/pkg/apis/devops/v1alpha2 -h ../../hack/boilerplate.go.txt

```
