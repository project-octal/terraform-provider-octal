module github.com/project-octal/terraform-provider-octal

go 1.15

require (
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.2.0
	github.com/project-octal/linkderd-cli v0.0.0-20201109211846-0f33ee424b95
)

replace github.com/wercker/stern => github.com/linkerd/stern v0.0.0-20200331220320-37779ceb2c32
