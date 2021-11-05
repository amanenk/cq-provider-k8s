package resources

import (
	"context"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func MetaOwnerReferences() *schema.Table {
	return &schema.Table{
		Name:        "k8s_meta_owner_references",
		Description: "OwnerReference contains enough information to let you identify an owning object",
		Resolver:    fetchMetaOwnerReferences,
		Options:     schema.TableCreationOptions{PrimaryKeys: []string{"resource_cq_id", "uid"}},
		Columns: []schema.Column{
			{
				Name:        "resource_cq_id",
				Description: "cq_id of parent resource",
				Type:        schema.TypeUUID,
			},
			{
				Name:        "api_version",
				Description: "API version of the referent.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("APIVersion"),
			},
			{
				Name:        "kind",
				Description: "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				Type:        schema.TypeString,
			},
			{
				Name:        "name",
				Description: "Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names",
				Type:        schema.TypeString,
			},
			{
				Name:        "uid",
				Description: "UID of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#uids",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("UID"),
			},
			{
				Name:        "controller",
				Description: "If true, this reference points to the managing controller. +optional",
				Type:        schema.TypeBool,
			},
			{
				Name:        "block_owner_deletion",
				Description: "If true, AND if the owner has the \"foregroundDeletion\" finalizer, then the owner cannot be deleted from the key-value store until this reference is removed. Defaults to false. To set this field, a user needs \"delete\" permission of the owner, otherwise 422 (Unprocessable Entity) will be returned. +optional",
				Type:        schema.TypeBool,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchMetaOwnerReferences(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {

	return nil
}
