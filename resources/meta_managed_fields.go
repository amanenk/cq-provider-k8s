package resources

import (
	"context"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func MetaManagedFields() *schema.Table {
	return &schema.Table{
		Name:        "k8s_meta_managed_fields",
		Description: "ManagedFieldsEntry is a workflow-id, a FieldSet and the group version of the resource that the fieldset applies to.",
		Resolver:    fetchMetaManagedFields,
		Options:     schema.TableCreationOptions{PrimaryKeys: []string{"resource_cq_id", "cq_id"}},
		Columns: []schema.Column{
			{
				Name:        "resource_cq_id",
				Description: "cq_id of parent resource",
				Type:        schema.TypeUUID,
			},
			{
				Name:        "manager",
				Description: "Manager is an identifier of the workflow managing these fields.",
				Type:        schema.TypeString,
			},
			{
				Name:        "operation",
				Description: "Operation is the type of operation which lead to this ManagedFieldsEntry being created. The only valid values for this field are 'Apply' and 'Update'.",
				Type:        schema.TypeString,
			},
			{
				Name:        "api_version",
				Description: "APIVersion defines the version of this resource that this field set applies to",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("APIVersion"),
			},
			{
				Name:        "fields_type",
				Description: "FieldsType is the discriminator for the different fields format and version. There is currently only one possible value: \"FieldsV1\"",
				Type:        schema.TypeString,
			},
			{
				Name:        "subresource",
				Description: "Subresource is the name of the subresource used to update that object, or empty string if the object was updated through the main resource",
				Type:        schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchMetaManagedFields(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	return nil
}
