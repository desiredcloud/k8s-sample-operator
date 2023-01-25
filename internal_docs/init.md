GO111MODULE=on operator-sdk init --component-config 'true' --domain 'desiredcloud.com' --fetch-deps 'true' --license 'none' \
--owner 'desiredcloud' --project-name 'k8s-sample-operator' --project-version '3' \
--repo 'github.com/desiredcloud/k8s-sample-operator' \
--plugins 'go.kubebuilder.io/v3,declarative.go.kubebuilder.io/v1,grafana.kubebuilder.io/v1-alpha' \
--verbose 'true'


DEBU[0000] Debug logging is set                         
WARN[0000] the platform of this environment (darwin/arm64) is not suppported by kustomize v3 (v3.8.7) which is used in this scaffold. You will be unable to download a binary for the kustomize version supported and used by this plugin. The currently supported platforms are: ["linux/amd64" "linux/arm64" "darwin/amd64"]
Writing kustomize manifests for you to edit...
Writing scaffold for you to edit...
Get controller runtime:
$ go get sigs.k8s.io/controller-runtime@v0.13.0
updating Dockerfile to add channels/ directory in the image
Generating Grafana manifests to visualize controller status...
Update dependencies:
$ go mod tidy
go: downloading github.com/golang-jwt/jwt/v4 v4.2.0
go: downloading golang.org/x/crypto v0.0.0-20220315160706-3147a52a75dd
Next: define a resource with:
$ operator-sdk create api
