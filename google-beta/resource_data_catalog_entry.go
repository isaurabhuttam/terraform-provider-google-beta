// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/structure"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/verify"
)

func ResourceDataCatalogEntry() *schema.Resource {
	return &schema.Resource{
		Create: resourceDataCatalogEntryCreate,
		Read:   resourceDataCatalogEntryRead,
		Update: resourceDataCatalogEntryUpdate,
		Delete: resourceDataCatalogEntryDelete,

		Importer: &schema.ResourceImporter{
			State: resourceDataCatalogEntryImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"entry_group": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The name of the entry group this entry is in.`,
			},
			"entry_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The id of the entry to create.`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Entry description, which can consist of several sentences or paragraphs that describe entry contents.`,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `Display information such as title and description. A short name to identify the entry,
for example, "Analytics Data - Jan 2011".`,
			},
			"gcs_fileset_spec": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `Specification that applies to a Cloud Storage fileset. This is only valid on entries of type FILESET.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"file_patterns": {
							Type:     schema.TypeList,
							Required: true,
							Description: `Patterns to identify a set of files in Google Cloud Storage.
See [Cloud Storage documentation](https://cloud.google.com/storage/docs/gsutil/addlhelp/WildcardNames)
for more information. Note that bucket wildcards are currently not supported. Examples of valid filePatterns:

* gs://bucket_name/dir/*: matches all files within bucket_name/dir directory.
* gs://bucket_name/dir/**: matches all files in bucket_name/dir spanning all subdirectories.
* gs://bucket_name/file*: matches files prefixed by file in bucket_name
* gs://bucket_name/??.txt: matches files with two characters followed by .txt in bucket_name
* gs://bucket_name/[aeiou].txt: matches files that contain a single vowel character followed by .txt in bucket_name
* gs://bucket_name/[a-m].txt: matches files that contain a, b, ... or m followed by .txt in bucket_name
* gs://bucket_name/a/*/b: matches all files in bucket_name that match a/*/b pattern, such as a/c/b, a/d/b
* gs://another_bucket/a.txt: matches gs://another_bucket/a.txt`,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"sample_gcs_file_specs": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: `Sample files contained in this fileset, not all files contained in this fileset are represented here.`,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"file_path": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: `The full file path`,
									},
									"size_bytes": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: `The size of the file, in bytes.`,
									},
								},
							},
						},
					},
				},
			},
			"linked_resource": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
				Description: `The resource this metadata entry refers to.
For Google Cloud Platform resources, linkedResource is the full name of the resource.
For example, the linkedResource for a table resource from BigQuery is:
//bigquery.googleapis.com/projects/projectId/datasets/datasetId/tables/tableId
Output only when Entry is of type in the EntryType enum. For entries with userSpecifiedType,
this field is optional and defaults to an empty string.`,
			},
			"schema": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsJSON,
				StateFunc:    func(v interface{}) string { s, _ := structure.NormalizeJsonString(v); return s },
				Description: `Schema of the entry (e.g. BigQuery, GoogleSQL, Avro schema), as a json string. An entry might not have any schema
attached to it. See
https://cloud.google.com/data-catalog/docs/reference/rest/v1/projects.locations.entryGroups.entries#schema
for what fields this schema can contain.`,
			},
			"type": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: verify.ValidateEnum([]string{"FILESET", ""}),
				Description: `The type of the entry. Only used for Entries with types in the EntryType enum.
Currently, only FILESET enum value is allowed. All other entries created through Data Catalog must use userSpecifiedType. Possible values: ["FILESET"]`,
				ExactlyOneOf: []string{"type", "user_specified_type"},
			},
			"user_specified_system": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: verify.ValidateRegexp(`^[A-z_][A-z0-9_]{0,63}$`),
				Description: `This field indicates the entry's source system that Data Catalog does not integrate with.
userSpecifiedSystem strings must begin with a letter or underscore and can only contain letters, numbers,
and underscores; are case insensitive; must be at least 1 character and at most 64 characters long.`,
			},
			"user_specified_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: verify.ValidateRegexp(`^[A-z_][A-z0-9_]{0,63}$`),
				Description: `Entry type if it does not fit any of the input-allowed values listed in EntryType enum above.
When creating an entry, users should check the enum values first, if nothing matches the entry
to be created, then provide a custom value, for example "my_special_type".
userSpecifiedType strings must begin with a letter or underscore and can only contain letters,
numbers, and underscores; are case insensitive; must be at least 1 character and at most 64 characters long.`,
				ExactlyOneOf: []string{"type", "user_specified_type"},
			},
			"bigquery_date_sharded_spec": {
				Type:     schema.TypeList,
				Computed: true,
				Description: `Specification for a group of BigQuery tables with name pattern [prefix]YYYYMMDD.
Context: https://cloud.google.com/bigquery/docs/partitioned-tables#partitioning_versus_sharding.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dataset": {
							Type:     schema.TypeString,
							Computed: true,
							Description: `The Data Catalog resource name of the dataset entry the current table belongs to, for example,
projects/{project_id}/locations/{location}/entrygroups/{entryGroupId}/entries/{entryId}`,
						},
						"shard_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: `Total number of shards.`,
						},
						"table_prefix": {
							Type:     schema.TypeString,
							Computed: true,
							Description: `The table name prefix of the shards. The name of any given shard is [tablePrefix]YYYYMMDD,
for example, for shard MyTable20180101, the tablePrefix is MyTable.`,
						},
					},
				},
			},
			"bigquery_table_spec": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `Specification that applies to a BigQuery table. This is only valid on entries of type TABLE.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"table_source_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The table source type.`,
						},
						"table_spec": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: `Spec of a BigQuery table. This field should only be populated if tableSourceType is BIGQUERY_TABLE.`,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"grouped_entry": {
										Type:     schema.TypeString,
										Computed: true,
										Description: `If the table is a dated shard, i.e., with name pattern [prefix]YYYYMMDD, groupedEntry is the
Data Catalog resource name of the date sharded grouped entry, for example,
projects/{project_id}/locations/{location}/entrygroups/{entryGroupId}/entries/{entryId}.
Otherwise, groupedEntry is empty.`,
									},
								},
							},
						},
						"view_spec": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: `Table view specification. This field should only be populated if tableSourceType is BIGQUERY_VIEW.`,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"view_query": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: `The query that defines the table view.`,
									},
								},
							},
						},
					},
				},
			},
			"integrated_system": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `This field indicates the entry's source system that Data Catalog integrates with, such as BigQuery or Pub/Sub.`,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The Data Catalog resource name of the entry in URL format.
Example: projects/{project_id}/locations/{location}/entryGroups/{entryGroupId}/entries/{entryId}.
Note that this Entry and its child resources may not actually be stored in the location in this name.`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceDataCatalogEntryCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	linkedResourceProp, err := expandDataCatalogEntryLinkedResource(d.Get("linked_resource"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("linked_resource"); !isEmptyValue(reflect.ValueOf(linkedResourceProp)) && (ok || !reflect.DeepEqual(v, linkedResourceProp)) {
		obj["linkedResource"] = linkedResourceProp
	}
	displayNameProp, err := expandDataCatalogEntryDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	descriptionProp, err := expandDataCatalogEntryDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	schemaProp, err := expandDataCatalogEntrySchema(d.Get("schema"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("schema"); !isEmptyValue(reflect.ValueOf(schemaProp)) && (ok || !reflect.DeepEqual(v, schemaProp)) {
		obj["schema"] = schemaProp
	}
	typeProp, err := expandDataCatalogEntryType(d.Get("type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("type"); !isEmptyValue(reflect.ValueOf(typeProp)) && (ok || !reflect.DeepEqual(v, typeProp)) {
		obj["type"] = typeProp
	}
	userSpecifiedTypeProp, err := expandDataCatalogEntryUserSpecifiedType(d.Get("user_specified_type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("user_specified_type"); !isEmptyValue(reflect.ValueOf(userSpecifiedTypeProp)) && (ok || !reflect.DeepEqual(v, userSpecifiedTypeProp)) {
		obj["userSpecifiedType"] = userSpecifiedTypeProp
	}
	userSpecifiedSystemProp, err := expandDataCatalogEntryUserSpecifiedSystem(d.Get("user_specified_system"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("user_specified_system"); !isEmptyValue(reflect.ValueOf(userSpecifiedSystemProp)) && (ok || !reflect.DeepEqual(v, userSpecifiedSystemProp)) {
		obj["userSpecifiedSystem"] = userSpecifiedSystemProp
	}
	gcsFilesetSpecProp, err := expandDataCatalogEntryGcsFilesetSpec(d.Get("gcs_fileset_spec"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("gcs_fileset_spec"); !isEmptyValue(reflect.ValueOf(gcsFilesetSpecProp)) && (ok || !reflect.DeepEqual(v, gcsFilesetSpecProp)) {
		obj["gcsFilesetSpec"] = gcsFilesetSpecProp
	}

	url, err := ReplaceVars(d, config, "{{DataCatalogBasePath}}{{entry_group}}/entries?entryId={{entry_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Entry: %#v", obj)
	billingProject := ""

	if parts := regexp.MustCompile(`projects\/([^\/]+)\/`).FindStringSubmatch(url); parts != nil {
		billingProject = parts[1]
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Entry: %s", err)
	}
	if err := d.Set("name", flattenDataCatalogEntryName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}

	// Store the ID now
	id, err := ReplaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Entry %q: %#v", d.Id(), res)

	return resourceDataCatalogEntryRead(d, meta)
}

func resourceDataCatalogEntryRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := ReplaceVars(d, config, "{{DataCatalogBasePath}}{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	if parts := regexp.MustCompile(`projects\/([^\/]+)\/`).FindStringSubmatch(url); parts != nil {
		billingProject = parts[1]
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("DataCatalogEntry %q", d.Id()))
	}

	if err := d.Set("name", flattenDataCatalogEntryName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Entry: %s", err)
	}
	if err := d.Set("linked_resource", flattenDataCatalogEntryLinkedResource(res["linkedResource"], d, config)); err != nil {
		return fmt.Errorf("Error reading Entry: %s", err)
	}
	if err := d.Set("display_name", flattenDataCatalogEntryDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading Entry: %s", err)
	}
	if err := d.Set("description", flattenDataCatalogEntryDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading Entry: %s", err)
	}
	if err := d.Set("schema", flattenDataCatalogEntrySchema(res["schema"], d, config)); err != nil {
		return fmt.Errorf("Error reading Entry: %s", err)
	}
	if err := d.Set("type", flattenDataCatalogEntryType(res["type"], d, config)); err != nil {
		return fmt.Errorf("Error reading Entry: %s", err)
	}
	if err := d.Set("user_specified_type", flattenDataCatalogEntryUserSpecifiedType(res["userSpecifiedType"], d, config)); err != nil {
		return fmt.Errorf("Error reading Entry: %s", err)
	}
	if err := d.Set("integrated_system", flattenDataCatalogEntryIntegratedSystem(res["integratedSystem"], d, config)); err != nil {
		return fmt.Errorf("Error reading Entry: %s", err)
	}
	if err := d.Set("user_specified_system", flattenDataCatalogEntryUserSpecifiedSystem(res["userSpecifiedSystem"], d, config)); err != nil {
		return fmt.Errorf("Error reading Entry: %s", err)
	}
	if err := d.Set("gcs_fileset_spec", flattenDataCatalogEntryGcsFilesetSpec(res["gcsFilesetSpec"], d, config)); err != nil {
		return fmt.Errorf("Error reading Entry: %s", err)
	}
	if err := d.Set("bigquery_table_spec", flattenDataCatalogEntryBigqueryTableSpec(res["bigqueryTableSpec"], d, config)); err != nil {
		return fmt.Errorf("Error reading Entry: %s", err)
	}
	if err := d.Set("bigquery_date_sharded_spec", flattenDataCatalogEntryBigqueryDateShardedSpec(res["bigqueryDateShardedSpec"], d, config)); err != nil {
		return fmt.Errorf("Error reading Entry: %s", err)
	}

	return nil
}

func resourceDataCatalogEntryUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	obj := make(map[string]interface{})
	linkedResourceProp, err := expandDataCatalogEntryLinkedResource(d.Get("linked_resource"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("linked_resource"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, linkedResourceProp)) {
		obj["linkedResource"] = linkedResourceProp
	}
	displayNameProp, err := expandDataCatalogEntryDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	descriptionProp, err := expandDataCatalogEntryDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	schemaProp, err := expandDataCatalogEntrySchema(d.Get("schema"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("schema"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, schemaProp)) {
		obj["schema"] = schemaProp
	}
	userSpecifiedTypeProp, err := expandDataCatalogEntryUserSpecifiedType(d.Get("user_specified_type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("user_specified_type"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, userSpecifiedTypeProp)) {
		obj["userSpecifiedType"] = userSpecifiedTypeProp
	}
	userSpecifiedSystemProp, err := expandDataCatalogEntryUserSpecifiedSystem(d.Get("user_specified_system"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("user_specified_system"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, userSpecifiedSystemProp)) {
		obj["userSpecifiedSystem"] = userSpecifiedSystemProp
	}
	gcsFilesetSpecProp, err := expandDataCatalogEntryGcsFilesetSpec(d.Get("gcs_fileset_spec"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("gcs_fileset_spec"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, gcsFilesetSpecProp)) {
		obj["gcsFilesetSpec"] = gcsFilesetSpecProp
	}

	url, err := ReplaceVars(d, config, "{{DataCatalogBasePath}}{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Entry %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("linked_resource") {
		updateMask = append(updateMask, "linkedResource")
	}

	if d.HasChange("display_name") {
		updateMask = append(updateMask, "displayName")
	}

	if d.HasChange("description") {
		updateMask = append(updateMask, "description")
	}

	if d.HasChange("schema") {
		updateMask = append(updateMask, "schema")
	}

	if d.HasChange("user_specified_type") {
		updateMask = append(updateMask, "userSpecifiedType")
	}

	if d.HasChange("user_specified_system") {
		updateMask = append(updateMask, "userSpecifiedSystem")
	}

	if d.HasChange("gcs_fileset_spec") {
		updateMask = append(updateMask, "gcsFilesetSpec")
	}
	// updateMask is a URL parameter but not present in the schema, so ReplaceVars
	// won't set it
	url, err = transport_tpg.AddQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}
	if parts := regexp.MustCompile(`projects\/([^\/]+)\/`).FindStringSubmatch(url); parts != nil {
		billingProject = parts[1]
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating Entry %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Entry %q: %#v", d.Id(), res)
	}

	return resourceDataCatalogEntryRead(d, meta)
}

func resourceDataCatalogEntryDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := ReplaceVars(d, config, "{{DataCatalogBasePath}}{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	if parts := regexp.MustCompile(`projects\/([^\/]+)\/`).FindStringSubmatch(url); parts != nil {
		billingProject = parts[1]
	}

	log.Printf("[DEBUG] Deleting Entry %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "Entry")
	}

	log.Printf("[DEBUG] Finished deleting Entry %q: %#v", d.Id(), res)
	return nil
}

func resourceDataCatalogEntryImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)

	// current import_formats can't import fields with forward slashes in their value
	if err := ParseImportId([]string{"(?P<name>.+)"}, d, config); err != nil {
		return nil, err
	}

	name := d.Get("name").(string)
	egRegex := regexp.MustCompile("(projects/.+/locations/.+/entryGroups/.+)/entries/(.+)")

	parts := egRegex.FindStringSubmatch(name)
	if len(parts) != 3 {
		return nil, fmt.Errorf("entry name does not fit the format %s", egRegex)
	}
	if err := d.Set("entry_group", parts[1]); err != nil {
		return nil, fmt.Errorf("Error setting entry_group: %s", err)
	}
	if err := d.Set("entry_id", parts[2]); err != nil {
		return nil, fmt.Errorf("Error setting entry_id: %s", err)
	}
	return []*schema.ResourceData{d}, nil
}

func flattenDataCatalogEntryName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataCatalogEntryLinkedResource(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataCatalogEntryDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataCatalogEntryDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataCatalogEntrySchema(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	b, err := json.Marshal(v)
	if err != nil {
		// TODO: return error once https://github.com/GoogleCloudPlatform/magic-modules/issues/3257 is fixed.
		log.Printf("[ERROR] failed to marshal schema to JSON: %v", err)
	}
	return string(b)
}

func flattenDataCatalogEntryType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataCatalogEntryUserSpecifiedType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataCatalogEntryIntegratedSystem(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataCatalogEntryUserSpecifiedSystem(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataCatalogEntryGcsFilesetSpec(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["file_patterns"] =
		flattenDataCatalogEntryGcsFilesetSpecFilePatterns(original["filePatterns"], d, config)
	transformed["sample_gcs_file_specs"] =
		flattenDataCatalogEntryGcsFilesetSpecSampleGcsFileSpecs(original["sampleGcsFileSpecs"], d, config)
	return []interface{}{transformed}
}
func flattenDataCatalogEntryGcsFilesetSpecFilePatterns(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataCatalogEntryGcsFilesetSpecSampleGcsFileSpecs(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"file_path":  flattenDataCatalogEntryGcsFilesetSpecSampleGcsFileSpecsFilePath(original["filePath"], d, config),
			"size_bytes": flattenDataCatalogEntryGcsFilesetSpecSampleGcsFileSpecsSizeBytes(original["sizeBytes"], d, config),
		})
	}
	return transformed
}
func flattenDataCatalogEntryGcsFilesetSpecSampleGcsFileSpecsFilePath(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataCatalogEntryGcsFilesetSpecSampleGcsFileSpecsSizeBytes(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenDataCatalogEntryBigqueryTableSpec(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["table_source_type"] =
		flattenDataCatalogEntryBigqueryTableSpecTableSourceType(original["tableSourceType"], d, config)
	transformed["view_spec"] =
		flattenDataCatalogEntryBigqueryTableSpecViewSpec(original["viewSpec"], d, config)
	transformed["table_spec"] =
		flattenDataCatalogEntryBigqueryTableSpecTableSpec(original["tableSpec"], d, config)
	return []interface{}{transformed}
}
func flattenDataCatalogEntryBigqueryTableSpecTableSourceType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataCatalogEntryBigqueryTableSpecViewSpec(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["view_query"] =
		flattenDataCatalogEntryBigqueryTableSpecViewSpecViewQuery(original["viewQuery"], d, config)
	return []interface{}{transformed}
}
func flattenDataCatalogEntryBigqueryTableSpecViewSpecViewQuery(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataCatalogEntryBigqueryTableSpecTableSpec(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["grouped_entry"] =
		flattenDataCatalogEntryBigqueryTableSpecTableSpecGroupedEntry(original["groupedEntry"], d, config)
	return []interface{}{transformed}
}
func flattenDataCatalogEntryBigqueryTableSpecTableSpecGroupedEntry(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataCatalogEntryBigqueryDateShardedSpec(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["dataset"] =
		flattenDataCatalogEntryBigqueryDateShardedSpecDataset(original["dataset"], d, config)
	transformed["table_prefix"] =
		flattenDataCatalogEntryBigqueryDateShardedSpecTablePrefix(original["tablePrefix"], d, config)
	transformed["shard_count"] =
		flattenDataCatalogEntryBigqueryDateShardedSpecShardCount(original["shardCount"], d, config)
	return []interface{}{transformed}
}
func flattenDataCatalogEntryBigqueryDateShardedSpecDataset(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataCatalogEntryBigqueryDateShardedSpecTablePrefix(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataCatalogEntryBigqueryDateShardedSpecShardCount(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func expandDataCatalogEntryLinkedResource(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataCatalogEntryDisplayName(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataCatalogEntryDescription(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataCatalogEntrySchema(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	b := []byte(v.(string))
	if len(b) == 0 {
		return nil, nil
	}
	m := make(map[string]interface{})
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return m, nil
}

func expandDataCatalogEntryType(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataCatalogEntryUserSpecifiedType(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataCatalogEntryUserSpecifiedSystem(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataCatalogEntryGcsFilesetSpec(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedFilePatterns, err := expandDataCatalogEntryGcsFilesetSpecFilePatterns(original["file_patterns"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFilePatterns); val.IsValid() && !isEmptyValue(val) {
		transformed["filePatterns"] = transformedFilePatterns
	}

	transformedSampleGcsFileSpecs, err := expandDataCatalogEntryGcsFilesetSpecSampleGcsFileSpecs(original["sample_gcs_file_specs"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSampleGcsFileSpecs); val.IsValid() && !isEmptyValue(val) {
		transformed["sampleGcsFileSpecs"] = transformedSampleGcsFileSpecs
	}

	return transformed, nil
}

func expandDataCatalogEntryGcsFilesetSpecFilePatterns(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataCatalogEntryGcsFilesetSpecSampleGcsFileSpecs(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedFilePath, err := expandDataCatalogEntryGcsFilesetSpecSampleGcsFileSpecsFilePath(original["file_path"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedFilePath); val.IsValid() && !isEmptyValue(val) {
			transformed["filePath"] = transformedFilePath
		}

		transformedSizeBytes, err := expandDataCatalogEntryGcsFilesetSpecSampleGcsFileSpecsSizeBytes(original["size_bytes"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedSizeBytes); val.IsValid() && !isEmptyValue(val) {
			transformed["sizeBytes"] = transformedSizeBytes
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDataCatalogEntryGcsFilesetSpecSampleGcsFileSpecsFilePath(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataCatalogEntryGcsFilesetSpecSampleGcsFileSpecsSizeBytes(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
