// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package eks

import (
	"context"
	"unique"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	inttypes "github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/internal/vcr"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) EphemeralResources(ctx context.Context) []*inttypes.ServicePackageEphemeralResource {
	return []*inttypes.ServicePackageEphemeralResource{
		{
			Factory:  newClusterAuthEphemeralResource,
			TypeName: "aws_eks_cluster_auth",
			Name:     "ClusterAuth",
			Region:   unique.Make(inttypes.ResourceRegionDefault()),
		},
	}
}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*inttypes.ServicePackageFrameworkDataSource {
	return []*inttypes.ServicePackageFrameworkDataSource{
		{
			Factory:  newClusterVersionsDataSource,
			TypeName: "aws_eks_cluster_versions",
			Name:     "Cluster Versions",
			Region:   unique.Make(inttypes.ResourceRegionDefault()),
		},
	}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*inttypes.ServicePackageFrameworkResource {
	return []*inttypes.ServicePackageFrameworkResource{
		{
			Factory:  newPodIdentityAssociationResource,
			TypeName: "aws_eks_pod_identity_association",
			Name:     "Pod Identity Association",
			Tags: unique.Make(inttypes.ServicePackageResourceTags{
				IdentifierAttribute: "association_arn",
			}),
			Region: unique.Make(inttypes.ResourceRegionDefault()),
		},
	}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*inttypes.ServicePackageSDKDataSource {
	return []*inttypes.ServicePackageSDKDataSource{
		{
			Factory:  dataSourceAccessEntry,
			TypeName: "aws_eks_access_entry",
			Name:     "Access Entry",
			Region:   unique.Make(inttypes.ResourceRegionDefault()),
		},
		{
			Factory:  dataSourceAddon,
			TypeName: "aws_eks_addon",
			Name:     "Add-On",
			Region:   unique.Make(inttypes.ResourceRegionDefault()),
		},
		{
			Factory:  dataSourceAddonVersion,
			TypeName: "aws_eks_addon_version",
			Name:     "Add-On Version",
			Region:   unique.Make(inttypes.ResourceRegionDefault()),
		},
		{
			Factory:  dataSourceCluster,
			TypeName: "aws_eks_cluster",
			Name:     "Cluster",
			Region:   unique.Make(inttypes.ResourceRegionDefault()),
		},
		{
			Factory:  dataSourceClusterAuth,
			TypeName: "aws_eks_cluster_auth",
			Name:     "Cluster Authentication Token",
			Region:   unique.Make(inttypes.ResourceRegionDefault()),
		},
		{
			Factory:  dataSourceClusters,
			TypeName: "aws_eks_clusters",
			Name:     "Clusters",
			Region:   unique.Make(inttypes.ResourceRegionDefault()),
		},
		{
			Factory:  dataSourceNodeGroup,
			TypeName: "aws_eks_node_group",
			Name:     "Node Group",
			Region:   unique.Make(inttypes.ResourceRegionDefault()),
		},
		{
			Factory:  dataSourceNodeGroups,
			TypeName: "aws_eks_node_groups",
			Name:     "Node Groups",
			Region:   unique.Make(inttypes.ResourceRegionDefault()),
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*inttypes.ServicePackageSDKResource {
	return []*inttypes.ServicePackageSDKResource{
		{
			Factory:  resourceAccessEntry,
			TypeName: "aws_eks_access_entry",
			Name:     "Access Entry",
			Tags: unique.Make(inttypes.ServicePackageResourceTags{
				IdentifierAttribute: "access_entry_arn",
			}),
			Region: unique.Make(inttypes.ResourceRegionDefault()),
		},
		{
			Factory:  resourceAccessPolicyAssociation,
			TypeName: "aws_eks_access_policy_association",
			Name:     "Access Policy Association",
			Region:   unique.Make(inttypes.ResourceRegionDefault()),
		},
		{
			Factory:  resourceAddon,
			TypeName: "aws_eks_addon",
			Name:     "Add-On",
			Tags: unique.Make(inttypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			}),
			Region: unique.Make(inttypes.ResourceRegionDefault()),
		},
		{
			Factory:  resourceCluster,
			TypeName: "aws_eks_cluster",
			Name:     "Cluster",
			Tags: unique.Make(inttypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			}),
			Region: unique.Make(inttypes.ResourceRegionDefault()),
		},
		{
			Factory:  resourceFargateProfile,
			TypeName: "aws_eks_fargate_profile",
			Name:     "Fargate Profile",
			Tags: unique.Make(inttypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			}),
			Region: unique.Make(inttypes.ResourceRegionDefault()),
		},
		{
			Factory:  resourceIdentityProviderConfig,
			TypeName: "aws_eks_identity_provider_config",
			Name:     "Identity Provider Config",
			Tags: unique.Make(inttypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			}),
			Region: unique.Make(inttypes.ResourceRegionDefault()),
		},
		{
			Factory:  resourceNodeGroup,
			TypeName: "aws_eks_node_group",
			Name:     "Node Group",
			Tags: unique.Make(inttypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			}),
			Region: unique.Make(inttypes.ResourceRegionDefault()),
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.EKS
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*eks.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))
	optFns := []func(*eks.Options){
		eks.WithEndpointResolverV2(newEndpointResolverV2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
		func(o *eks.Options) {
			if region := config[names.AttrRegion].(string); o.Region != region {
				tflog.Info(ctx, "overriding provider-configured AWS API region", map[string]any{
					"service":         p.ServicePackageName(),
					"original_region": o.Region,
					"override_region": region,
				})
				o.Region = region
			}
		},
		func(o *eks.Options) {
			if inContext, ok := conns.FromContext(ctx); ok && inContext.VCREnabled() {
				tflog.Info(ctx, "overriding retry behavior to immediately return VCR errors")
				o.Retryer = conns.AddIsErrorRetryables(cfg.Retryer().(aws.RetryerV2), retry.IsErrorRetryableFunc(vcr.InteractionNotFoundRetryableFunc))
			}
		},
		withExtraOptions(ctx, p, config),
	}

	return eks.NewFromConfig(cfg, optFns...), nil
}

// withExtraOptions returns a functional option that allows this service package to specify extra API client options.
// This option is always called after any generated options.
func withExtraOptions(ctx context.Context, sp conns.ServicePackage, config map[string]any) func(*eks.Options) {
	if v, ok := sp.(interface {
		withExtraOptions(context.Context, map[string]any) []func(*eks.Options)
	}); ok {
		optFns := v.withExtraOptions(ctx, config)

		return func(o *eks.Options) {
			for _, optFn := range optFns {
				optFn(o)
			}
		}
	}

	return func(*eks.Options) {}
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
