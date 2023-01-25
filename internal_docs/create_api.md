Deploy Image Plugin (deploy-image/v1-alpha)

https://master.book.kubebuilder.io/plugins/deploy-image-plugin-v1-alpha


operator-sdk create api -h
Scaffold a Kubernetes API by writing a Resource definition and/or a Controller.

If information about whether the resource and controller should be scaffolded
was not explicitly provided, it will prompt the user if they should be.

After the scaffold is written, the dependencies will be updated and
make generate will be run.

Usage:
operator-sdk create api [flags]

Examples:
# Create a frigates API with Group: ship, Version: v1beta1 and Kind: Frigate
operator-sdk create api --group ship --version v1beta1 --kind Frigate

# Edit the API Scheme
nano api/v1beta1/frigate_types.go

# Edit the Controller
nano controllers/frigate/frigate_controller.go

# Edit the Controller Test
nano controllers/frigate/frigate_controller_test.go

# Generate the manifests
make manifests

# Install CRDs into the Kubernetes cluster using kubectl apply
make install

# Regenerate code and run against the Kubernetes cluster configured by ~/.kube/config
make run


Flags:
--controller           if set, generate the controller without prompting the user (default true)
--force                attempt to create resource even if it already exists
--group string         resource Group
-h, --help                 help for api
--kind string          resource Kind
--make make generate   if true, run make generate after generating files (default true)
--namespaced           resource is namespaced (default true)
--plural string        resource irregular plural form
--resource             if set, generate the resource without prompting the user (default true)
--version string       resource Version

Global Flags:
--plugins strings   plugin keys to be used for this subcommand execution
--verbose           Enable verbose logging
tanveeralam@Tanveers-MacBook-Pro pub % 




operator-sdk create api --group image --version v1alpha1 --kind ImagePlugin --plugins="deploy-image/v1-alpha" --image=memcached:1.4.36-alpine --image-container-command="memcached,-m=64,modern,-v" --run-as-user="1001"

updating scaffold with deploy-image/v1alpha1 plugin...
Writing scaffold for you to edit...
Writing scaffold for you to edit...
api/v1alpha1/imageplugin_types.go
controllers/imageplugin_controller.go
Writing kustomize manifests for you to edit...
api/v1alpha1/imageplugin_types.go
creating import for % github.com/desiredcloud/k8s-sample-operator/api/v1alpha1
creating import for % github.com/desiredcloud/k8s-sample-operator/api/v1alpha1
Update dependencies:
$ go mod tidy
Running make:
$ make generate
mkdir -p /Users/tanveeralam/ws/tan/k8s-sample-operator/bin
test -s /Users/tanveeralam/ws/tan/k8s-sample-operator/bin/controller-gen || GOBIN=/Users/tanveeralam/ws/tan/k8s-sample-operator/bin go install sigs.k8s.io/controller-tools/cmd/controller-gen@v0.10.0
/Users/tanveeralam/ws/tan/k8s-sample-operator/bin/controller-gen object:headerFile="hack/boilerplate.go.txt" paths="./..."
Running make:
$ make manifests
/Users/tanveeralam/ws/tan/k8s-sample-operator/bin/controller-gen rbac:roleName=manager-role crd webhook paths="./..." output:crd:artifacts:config=config/crd/bases
Next: check the implementation of your new API and controller. If you do changes in the API run the manifests with:
$ make manifests



operator-sdk create api --group group1 --version v1alpha1 --kind Kind1 --controller --make --namespaced --resource --verbose 'true'
