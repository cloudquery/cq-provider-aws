package ecr

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource repositories --config repositories.hcl --output .

func Repositories() *schema.Table {
	return &schema.Table{
		Name:          "aws_ecr_repositories",
		Description:   "An object representing a repository.",
		Resolver:      fetchEcrRepositories,
		Multiplex:     client.ServiceAccountRegionMultiplexer("api.ecr"),
		IgnoreError:   client.IgnoreCommonErrors,
		DeleteFilter:  client.DeleteAccountRegionFilter,
		Options:       schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "arn"}},
		IgnoreInTests: true,
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:        "created_at",
				Description: "The date and time, in JavaScript date format, when the repository was created.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "encryption_configuration_encryption_type",
				Description: "The encryption type to use",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("EncryptionConfiguration.EncryptionType"),
			},
			{
				Name:        "encryption_configuration_kms_key",
				Description: "If you use the KMS encryption type, specify the KMS key to use for encryption. The alias, key ID, or full ARN of the KMS key can be specified",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("EncryptionConfiguration.KmsKey"),
			},
			{
				Name:        "image_scanning_configuration_scan_on_push",
				Description: "The setting that determines whether images are scanned after being pushed to a repository",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("ImageScanningConfiguration.ScanOnPush"),
			},
			{
				Name:        "image_tag_mutability",
				Description: "The tag mutability setting for the repository.",
				Type:        schema.TypeString,
			},
			{
				Name:        "registry_id",
				Description: "The Amazon Web Services account ID associated with the registry that contains the repository.",
				Type:        schema.TypeString,
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) that identifies the repository",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("RepositoryArn"),
			},
			{
				Name:        "name",
				Description: "The name of the repository.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("RepositoryName"),
			},
			{
				Name:        "uri",
				Description: "The URI for the repository",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("RepositoryUri"),
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_ecr_repository_images",
				Description: "An object that describes an image returned by a DescribeImages operation.",
				Resolver:    fetchEcrRepositoryImages,
				Columns: []schema.Column{
					{
						Name:        "repository_cq_id",
						Description: "Unique CloudQuery ID of aws_ecr_repositories table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "account_id",
						Description: "The AWS Account ID of the resource.",
						Type:        schema.TypeString,
						Resolver:    client.ResolveAWSAccount,
					},
					{
						Name:        "region",
						Description: "The AWS Region of the resource.",
						Type:        schema.TypeString,
						Resolver:    client.ResolveAWSRegion,
					},
					{
						Name:        "artifact_media_type",
						Description: "The artifact media type of the image.",
						Type:        schema.TypeString,
					},
					{
						Name:        "image_digest",
						Description: "The sha256 digest of the image manifest.",
						Type:        schema.TypeString,
					},
					{
						Name:        "image_manifest_media_type",
						Description: "The media type of the image manifest.",
						Type:        schema.TypeString,
					},
					{
						Name:        "image_pushed_at",
						Description: "The date and time, expressed in standard JavaScript date format, at which the current image was pushed to the repository.",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "image_scan_findings_summary_finding_severity_counts",
						Description: "The image vulnerability counts, sorted by severity.",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("ImageScanFindingsSummary.FindingSeverityCounts"),
					},
					{
						Name:        "image_scan_findings_summary_image_scan_completed_at",
						Description: "The time of the last completed image scan.",
						Type:        schema.TypeTimestamp,
						Resolver:    schema.PathResolver("ImageScanFindingsSummary.ImageScanCompletedAt"),
					},
					{
						Name:        "image_scan_findings_summary_vulnerability_source_updated_at",
						Description: "The time when the vulnerability data was last scanned.",
						Type:        schema.TypeTimestamp,
						Resolver:    schema.PathResolver("ImageScanFindingsSummary.VulnerabilitySourceUpdatedAt"),
					},
					{
						Name:        "image_scan_status_description",
						Description: "The description of the image scan status.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ImageScanStatus.Description"),
					},
					{
						Name:        "image_scan_status",
						Description: "The current state of an image scan.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ImageScanStatus.Status"),
					},
					{
						Name:        "image_size_in_bytes",
						Description: "The size, in bytes, of the image in the repository",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "image_tags",
						Description: "The list of tags associated with this image.",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "last_recorded_pull_time",
						Description: "The date and time, expressed in standard JavaScript date format, when Amazon ECR recorded the last image pull",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "registry_id",
						Description: "The Amazon Web Services account ID associated with the registry to which this image belongs.",
						Type:        schema.TypeString,
					},
					{
						Name:        "repository_name",
						Description: "The name of the repository to which this image belongs.",
						Type:        schema.TypeString,
					},
				},
				Relations: []*schema.Table{
					{
						Name:        "aws_ecr_repository_image_scan_findings",
						Description: "The details of an image scan.",
						Resolver:    fetchEcrRepositoryImageScanFindings,
						Columns: []schema.Column{
							{
								Name:        "repository_image_cq_id",
								Description: "Unique CloudQuery ID of aws_ecr_repository_images table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "finding_severity_counts",
								Description: "The image vulnerability counts, sorted by severity.",
								Type:        schema.TypeJSON,
							},
							{
								Name:        "image_scan_completed_at",
								Description: "The time of the last completed image scan.",
								Type:        schema.TypeTimestamp,
							},
							{
								Name:        "vulnerability_source_updated_at",
								Description: "The time when the vulnerability data was last scanned.",
								Type:        schema.TypeTimestamp,
							},
						},
						Relations: []*schema.Table{
							{
								Name:        "aws_ecr_repository_image_scan_finding_enhanced",
								Description: "The details of an enhanced image scan",
								Resolver:    fetchEcrRepositoryImageScanFindingEnhanceds,
								Columns: []schema.Column{
									{
										Name:        "repository_image_scan_finding_cq_id",
										Description: "Unique CloudQuery ID of aws_ecr_repository_image_scan_findings table (FK)",
										Type:        schema.TypeUUID,
										Resolver:    schema.ParentIdResolver,
									},
									{
										Name:        "aws_account_id",
										Description: "The Amazon Web Services account ID associated with the image.",
										Type:        schema.TypeString,
									},
									{
										Name:        "description",
										Description: "The description of the finding.",
										Type:        schema.TypeString,
									},
									{
										Name:        "finding_arn",
										Description: "The Amazon Resource Number (ARN) of the finding.",
										Type:        schema.TypeString,
									},
									{
										Name:        "first_observed_at",
										Description: "The date and time that the finding was first observed.",
										Type:        schema.TypeTimestamp,
									},
									{
										Name:        "last_observed_at",
										Description: "The date and time that the finding was last observed.",
										Type:        schema.TypeTimestamp,
									},
									{
										Name:        "package_vulnerability_details_cvss",
										Description: "An object that contains details about the CVSS score of a finding.",
										Type:        schema.TypeJSON,
										Resolver:    resolveRepositoryImageScanFindingEnhancedsPackageVulnerabilityDetailsCvss,
									},
									{
										Name:        "package_vulnerability_details_reference_urls",
										Description: "One or more URLs that contain details about this vulnerability type.",
										Type:        schema.TypeStringArray,
										Resolver:    schema.PathResolver("PackageVulnerabilityDetails.ReferenceUrls"),
									},
									{
										Name:        "package_vulnerability_details_related_vulnerabilities",
										Description: "One or more vulnerabilities related to the one identified in this finding.",
										Type:        schema.TypeStringArray,
										Resolver:    schema.PathResolver("PackageVulnerabilityDetails.RelatedVulnerabilities"),
									},
									{
										Name:        "package_vulnerability_details_source",
										Description: "The source of the vulnerability information.",
										Type:        schema.TypeString,
										Resolver:    schema.PathResolver("PackageVulnerabilityDetails.Source"),
									},
									{
										Name:        "package_vulnerability_details_source_url",
										Description: "A URL to the source of the vulnerability information.",
										Type:        schema.TypeString,
										Resolver:    schema.PathResolver("PackageVulnerabilityDetails.SourceUrl"),
									},
									{
										Name:        "package_vulnerability_details_vendor_created_at",
										Description: "The date and time that this vulnerability was first added to the vendor's database.",
										Type:        schema.TypeTimestamp,
										Resolver:    schema.PathResolver("PackageVulnerabilityDetails.VendorCreatedAt"),
									},
									{
										Name:        "package_vulnerability_details_vendor_severity",
										Description: "The severity the vendor has given to this vulnerability type.",
										Type:        schema.TypeString,
										Resolver:    schema.PathResolver("PackageVulnerabilityDetails.VendorSeverity"),
									},
									{
										Name:        "package_vulnerability_details_vendor_updated_at",
										Description: "The date and time the vendor last updated this vulnerability in their database.",
										Type:        schema.TypeTimestamp,
										Resolver:    schema.PathResolver("PackageVulnerabilityDetails.VendorUpdatedAt"),
									},
									{
										Name:        "package_vulnerability_details_vulnerability_id",
										Description: "The ID given to this vulnerability.",
										Type:        schema.TypeString,
										Resolver:    schema.PathResolver("PackageVulnerabilityDetails.VulnerabilityId"),
									},
									{
										Name:        "package_vulnerability_details_vulnerable_packages",
										Description: "The packages impacted by this vulnerability.",
										Type:        schema.TypeJSON,
										Resolver:    resolveRepositoryImageScanFindingEnhancedsPackageVulnerabilityDetailsVulnerablePackages,
									},
									{
										Name:        "remediation_recommendation_text",
										Description: "The recommended course of action to remediate the finding.",
										Type:        schema.TypeString,
										Resolver:    schema.PathResolver("Remediation.Recommendation.Text"),
									},
									{
										Name:        "remediation_recommendation_url",
										Description: "The URL address to the CVE remediation recommendations.",
										Type:        schema.TypeString,
										Resolver:    schema.PathResolver("Remediation.Recommendation.Url"),
									},
									{
										Name:        "score",
										Description: "The Amazon Inspector score given to the finding.",
										Type:        schema.TypeFloat,
									},
									{
										Name:        "score_details_cvss_adjustments",
										Description: "An object that contains details about adjustment Amazon Inspector made to the CVSS score.",
										Type:        schema.TypeJSON,
										Resolver:    resolveRepositoryImageScanFindingEnhancedsScoreDetailsCvssAdjustments,
									},
									{
										Name:        "score_details_cvss_score",
										Description: "The CVSS score.",
										Type:        schema.TypeFloat,
										Resolver:    schema.PathResolver("ScoreDetails.Cvss.Score"),
									},
									{
										Name:        "score_details_cvss_score_source",
										Description: "The source for the CVSS score.",
										Type:        schema.TypeString,
										Resolver:    schema.PathResolver("ScoreDetails.Cvss.ScoreSource"),
									},
									{
										Name:        "score_details_cvss_scoring_vector",
										Description: "The vector for the CVSS score.",
										Type:        schema.TypeString,
										Resolver:    schema.PathResolver("ScoreDetails.Cvss.ScoringVector"),
									},
									{
										Name:        "score_details_cvss_version",
										Description: "The CVSS version used in scoring.",
										Type:        schema.TypeString,
										Resolver:    schema.PathResolver("ScoreDetails.Cvss.Version"),
									},
									{
										Name:        "severity",
										Description: "The severity of the finding.",
										Type:        schema.TypeString,
									},
									{
										Name:        "status",
										Description: "The status of the finding.",
										Type:        schema.TypeString,
									},
									{
										Name:        "title",
										Description: "The title of the finding.",
										Type:        schema.TypeString,
									},
									{
										Name:        "type",
										Description: "The type of the finding.",
										Type:        schema.TypeString,
									},
									{
										Name:        "updated_at",
										Description: "The date and time the finding was last updated at.",
										Type:        schema.TypeTimestamp,
									},
								},
								Relations: []*schema.Table{
									{
										Name:        "aws_ecr_repository_image_scan_finding_enhanced_resources",
										Description: "Details about the resource involved in a finding.",
										Resolver:    fetchEcrRepositoryImageScanFindingEnhancedResources,
										Columns: []schema.Column{
											{
												Name:        "repository_image_scan_finding_enhanced_cq_id",
												Description: "Unique CloudQuery ID of aws_ecr_repository_image_scan_finding_enhanced table (FK)",
												Type:        schema.TypeUUID,
												Resolver:    schema.ParentIdResolver,
											},
											{
												Name:        "aws_ecr_container_image_architecture",
												Description: "The architecture of the Amazon ECR container image.",
												Type:        schema.TypeString,
												Resolver:    schema.PathResolver("Details.AwsEcrContainerImage.Architecture"),
											},
											{
												Name:        "aws_ecr_container_image_author",
												Description: "The image author of the Amazon ECR container image.",
												Type:        schema.TypeString,
												Resolver:    schema.PathResolver("Details.AwsEcrContainerImage.Author"),
											},
											{
												Name:        "aws_ecr_container_image_image_hash",
												Description: "The image hash of the Amazon ECR container image.",
												Type:        schema.TypeString,
												Resolver:    schema.PathResolver("Details.AwsEcrContainerImage.ImageHash"),
											},
											{
												Name:        "aws_ecr_container_image_image_tags",
												Description: "The image tags attached to the Amazon ECR container image.",
												Type:        schema.TypeStringArray,
												Resolver:    schema.PathResolver("Details.AwsEcrContainerImage.ImageTags"),
											},
											{
												Name:        "aws_ecr_container_image_platform",
												Description: "The platform of the Amazon ECR container image.",
												Type:        schema.TypeString,
												Resolver:    schema.PathResolver("Details.AwsEcrContainerImage.Platform"),
											},
											{
												Name:        "aws_ecr_container_image_pushed_at",
												Description: "The date and time the Amazon ECR container image was pushed.",
												Type:        schema.TypeTimestamp,
												Resolver:    schema.PathResolver("Details.AwsEcrContainerImage.PushedAt"),
											},
											{
												Name:        "aws_ecr_container_image_registry",
												Description: "The registry the Amazon ECR container image belongs to.",
												Type:        schema.TypeString,
												Resolver:    schema.PathResolver("Details.AwsEcrContainerImage.Registry"),
											},
											{
												Name:        "aws_ecr_container_image_repository_name",
												Description: "The name of the repository the Amazon ECR container image resides in.",
												Type:        schema.TypeString,
												Resolver:    schema.PathResolver("Details.AwsEcrContainerImage.RepositoryName"),
											},
											{
												Name:        "id",
												Description: "The ID of the resource.",
												Type:        schema.TypeString,
											},
											{
												Name:        "tags",
												Description: "The tags attached to the resource.",
												Type:        schema.TypeJSON,
											},
											{
												Name:        "type",
												Description: "The type of resource.",
												Type:        schema.TypeString,
											},
										},
									},
								},
							},
							{
								Name:        "aws_ecr_repository_image_scan_finding_findings",
								Description: "Contains information about an image scan finding.",
								Resolver:    fetchEcrRepositoryImageScanFindingFindings,
								Columns: []schema.Column{
									{
										Name:        "repository_image_scan_finding_cq_id",
										Description: "Unique CloudQuery ID of aws_ecr_repository_image_scan_findings table (FK)",
										Type:        schema.TypeUUID,
										Resolver:    schema.ParentIdResolver,
									},
									{
										Name:        "attributes",
										Description: "A collection of attributes of the host from which the finding is generated.",
										Type:        schema.TypeJSON,
										Resolver:    resolveRepositoryImageScanFindingFindingsAttributes,
									},
									{
										Name:        "description",
										Description: "The description of the finding.",
										Type:        schema.TypeString,
									},
									{
										Name:        "name",
										Description: "The name associated with the finding, usually a CVE number.",
										Type:        schema.TypeString,
									},
									{
										Name:        "severity",
										Description: "The finding severity.",
										Type:        schema.TypeString,
									},
									{
										Name:        "uri",
										Description: "A link containing additional details about the security vulnerability.",
										Type:        schema.TypeString,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchEcrRepositories(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	maxResults := int32(1000)
	config := ecr.DescribeRepositoriesInput{
		MaxResults: &maxResults,
	}
	c := meta.(*client.Client)
	svc := c.Services().ECR
	for {
		output, err := svc.DescribeRepositories(ctx, &config, func(options *ecr.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		res <- output.Repositories
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
func fetchEcrRepositoryImages(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	maxResults := int32(1000)
	p := parent.Item.(types.Repository)
	config := ecr.DescribeImagesInput{
		RepositoryName: p.RepositoryName,
		MaxResults:     &maxResults,
	}
	c := meta.(*client.Client)
	svc := c.Services().ECR
	for {
		output, err := svc.DescribeImages(ctx, &config, func(options *ecr.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		res <- output.ImageDetails
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
func fetchEcrRepositoryImageScanFindings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	maxResults := int32(1000)
	image := parent.Item.(types.ImageDetail)
	config := ecr.DescribeImageScanFindingsInput{
		ImageId:        &types.ImageIdentifier{ImageDigest: image.ImageDigest},
		RepositoryName: image.RepositoryName,
		MaxResults:     &maxResults,
	}
	c := meta.(*client.Client)
	svc := c.Services().ECR
	for {
		output, err := svc.DescribeImageScanFindings(ctx, &config, func(options *ecr.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		if output.ImageScanFindings != nil {
			res <- output.ImageScanFindings
		}
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
func fetchEcrRepositoryImageScanFindingEnhanceds(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	res <- parent.Item.(*types.ImageScanFindings).EnhancedFindings
	return nil
}
func resolveRepositoryImageScanFindingEnhancedsPackageVulnerabilityDetailsCvss(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	f := resource.Item.(types.EnhancedImageScanFinding)
	if f.PackageVulnerabilityDetails == nil {
		return nil
	}
	b, err := json.Marshal(f.PackageVulnerabilityDetails.Cvss)
	if err != nil {
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(c.Name, b))
}
func resolveRepositoryImageScanFindingEnhancedsPackageVulnerabilityDetailsVulnerablePackages(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	f := resource.Item.(types.EnhancedImageScanFinding)
	if f.PackageVulnerabilityDetails == nil {
		return nil
	}
	b, err := json.Marshal(f.PackageVulnerabilityDetails.VulnerablePackages)
	if err != nil {
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(c.Name, b))
}
func resolveRepositoryImageScanFindingEnhancedsScoreDetailsCvssAdjustments(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	f := resource.Item.(types.EnhancedImageScanFinding)
	if f.ScoreDetails == nil || f.ScoreDetails.Cvss == nil {
		return nil
	}
	b, err := json.Marshal(f.ScoreDetails.Cvss.Adjustments)
	if err != nil {
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(c.Name, b))
}
func fetchEcrRepositoryImageScanFindingEnhancedResources(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	res <- parent.Item.(types.EnhancedImageScanFinding).Resources
	return nil
}
func fetchEcrRepositoryImageScanFindingFindings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	res <- parent.Item.(*types.ImageScanFindings).Findings
	return nil
}
func resolveRepositoryImageScanFindingFindingsAttributes(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	return diag.WrapError(resource.Set(c.Name, client.TagsToMap(resource.Item.(types.ImageScanFinding).Attributes)))
}
