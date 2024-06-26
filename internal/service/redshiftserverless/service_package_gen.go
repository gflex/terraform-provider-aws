// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package redshiftserverless

import (
	"context"

	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	redshiftserverless_sdkv2 "github.com/aws/aws-sdk-go-v2/service/redshiftserverless"
	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	endpoints_sdkv1 "github.com/aws/aws-sdk-go/aws/endpoints"
	session_sdkv1 "github.com/aws/aws-sdk-go/aws/session"
	redshiftserverless_sdkv1 "github.com/aws/aws-sdk-go/service/redshiftserverless"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{
		{
			Factory: newCustomDomainAssociationResource,
			Name:    "Custom Domain Association",
		},
	}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  dataSourceCredentials,
			TypeName: "aws_redshiftserverless_credentials",
			Name:     "Credentials",
		},
		{
			Factory:  dataSourceNamespace,
			TypeName: "aws_redshiftserverless_namespace",
			Name:     "Namespace",
		},
		{
			Factory:  dataSourceWorkgroup,
			TypeName: "aws_redshiftserverless_workgroup",
			Name:     "Workgroup",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceEndpointAccess,
			TypeName: "aws_redshiftserverless_endpoint_access",
			Name:     "Endpoint Access",
		},
		{
			Factory:  resourceNamespace,
			TypeName: "aws_redshiftserverless_namespace",
			Name:     "Namespace",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceResourcePolicy,
			TypeName: "aws_redshiftserverless_resource_policy",
			Name:     "Resource Policy",
		},
		{
			Factory:  resourceSnapshot,
			TypeName: "aws_redshiftserverless_snapshot",
			Name:     "Snapshot",
		},
		{
			Factory:  resourceUsageLimit,
			TypeName: "aws_redshiftserverless_usage_limit",
			Name:     "Usage Limit",
		},
		{
			Factory:  resourceWorkgroup,
			TypeName: "aws_redshiftserverless_workgroup",
			Name:     "Workgroup",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.RedshiftServerless
}

// NewConn returns a new AWS SDK for Go v1 client for this service package's AWS API.
func (p *servicePackage) NewConn(ctx context.Context, config map[string]any) (*redshiftserverless_sdkv1.RedshiftServerless, error) {
	sess := config[names.AttrSession].(*session_sdkv1.Session)

	cfg := aws_sdkv1.Config{}

	if endpoint := config[names.AttrEndpoint].(string); endpoint != "" {
		tflog.Debug(ctx, "setting endpoint", map[string]any{
			"tf_aws.endpoint": endpoint,
		})
		cfg.Endpoint = aws_sdkv1.String(endpoint)

		if sess.Config.UseFIPSEndpoint == endpoints_sdkv1.FIPSEndpointStateEnabled {
			tflog.Debug(ctx, "endpoint set, ignoring UseFIPSEndpoint setting")
			cfg.UseFIPSEndpoint = endpoints_sdkv1.FIPSEndpointStateDisabled
		}
	}

	return redshiftserverless_sdkv1.New(sess.Copy(&cfg)), nil
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*redshiftserverless_sdkv2.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws_sdkv2.Config))

	return redshiftserverless_sdkv2.NewFromConfig(cfg, func(o *redshiftserverless_sdkv2.Options) {
		if endpoint := config[names.AttrEndpoint].(string); endpoint != "" {
			tflog.Debug(ctx, "setting endpoint", map[string]any{
				"tf_aws.endpoint": endpoint,
			})
			o.BaseEndpoint = aws_sdkv2.String(endpoint)

			if o.EndpointOptions.UseFIPSEndpoint == aws_sdkv2.FIPSEndpointStateEnabled {
				tflog.Debug(ctx, "endpoint set, ignoring UseFIPSEndpoint setting")
				o.EndpointOptions.UseFIPSEndpoint = aws_sdkv2.FIPSEndpointStateDisabled
			}
		}
	}), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
