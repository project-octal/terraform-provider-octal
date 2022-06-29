module github.com/project-octal/terraform-provider-octal

go 1.15

require (
	github.com/andybalholm/crlf v0.0.0-20171020200849-670099aa064f // indirect
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.17.0
	github.com/mitchellh/go-homedir v1.1.0
	github.com/project-octal/linkderd-cli v0.0.0-20201113192458-49960f0ebf11
	k8s.io/api v0.24.2
	k8s.io/apimachinery v0.24.2
	k8s.io/client-go v0.24.2
	k8s.io/kube-aggregator v0.17.4
)

replace github.com/wercker/stern => github.com/linkerd/stern v0.0.0-20200331220320-37779ceb2c32
