#operatorsdk

Commands used:
export GO111MODULE=on
go mod init example
operator-sdk new kublr-operator --dep-manager modules
cd kublr-operator
operator-sdk add api --api-version=features.kublr.com/v1alpha1 --kind=Logging
# change logging feature type
operator-sdk generate k8s
operator-sdk add controller --api-version=features.kublr.com/v1alpha1 --kind=Logging

Manual used:
https://github.com/operator-framework/getting-started
SDK project layout documentation https://github.com/operator-framework/operator-sdk/blob/master/doc/project_layout.md

Problems:
When no "time" in import, there is not ideas what does the following error means (even with --verbose flag):
F0608 02:01:19.078939   65239 deepcopy.go:873] Hit an unsupported type invalid type.
