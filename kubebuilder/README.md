#kubebuilder
Commands used:
export GO111MODULE=on
kubebuilder init --domain kublr.com
kubebuilder create api --group feature --version v1alpha1 --kind Logging
kubebuilder create api --group feature --version v1alpha1 --kind Monitoring
kubebuilder create api --group feature --version v1alpha1 --kind Ingress


Manual used:
https://book.kubebuilder.io/quick-start.html

Problems:
I added import from kublr common and got the following:
$ make generate
/Users/oleg/work/bin/controller-gen object:headerFile=./hack/boilerplate.go.txt paths=./api/...
177:20: missing ',' in argument list
182:20: missing ',' in argument list
187:24: missing ',' in argument list
188:1: missing ',' in argument list
189:1: missing ',' in argument list
190:9: missing ',' in composite literal
.....

$ make generate
/Users/oleg/work/bin/controller-gen object:headerFile=./hack/boilerplate.go.txt paths=./api/...
example/api/v1alpha1:-: name requested for invalid type interface{}
example/api/v1alpha1:-: name requested for invalid type interface{}
example/api/v1alpha1:-: invalid map value type interface{}
Error: not all generators ran succesfully
Usage:

It was fixed after commenting
Values *entity.ChartValues `json:"values,omitempty" yaml:"values,omitempty"`
Source:
type ChartValues map[string]interface{}
