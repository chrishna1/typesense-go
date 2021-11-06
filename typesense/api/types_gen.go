// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.0 DO NOT EDIT.
package api

import (
	"encoding/json"
	"fmt"
)

const (
	Api_key_headerScopes = "api_key_header.Scopes"
)

// Defines values for FieldType.
const (
	FieldTypeAuto FieldType = "auto"

	FieldTypeBool FieldType = "bool"

	FieldTypeBool1 FieldType = "bool[]"

	FieldTypeFloat FieldType = "float"

	FieldTypeFloat1 FieldType = "float[]"

	FieldTypeInt32 FieldType = "int32"

	FieldTypeInt321 FieldType = "int32[]"

	FieldTypeInt64 FieldType = "int64"

	FieldTypeInt641 FieldType = "int64[]"

	FieldTypeString FieldType = "string"

	FieldTypeString1 FieldType = "string[]"
)

// Defines values for SearchOverrideRuleMatch.
const (
	SearchOverrideRuleMatchContains SearchOverrideRuleMatch = "contains"

	SearchOverrideRuleMatchExact SearchOverrideRuleMatch = "exact"
)

// ApiKey defines model for ApiKey.
type ApiKey struct {
	// Embedded struct due to allOf(#/components/schemas/ApiKeySchema)
	ApiKeySchema `yaml:",inline"`
	// Embedded fields due to inline allOf schema
	Id          int64  `json:"id"`
	Value       string `json:"value"`
	ValuePrefix string `json:"value_prefix"`
}

// ApiKeySchema defines model for ApiKeySchema.
type ApiKeySchema struct {
	Actions     []string `json:"actions"`
	Collections []string `json:"collections"`
	Description string   `json:"description"`
	ExpiresAt   *int64   `json:"expires_at,omitempty"`
}

// ApiKeysResponse defines model for ApiKeysResponse.
type ApiKeysResponse struct {
	Keys []*ApiKey `json:"keys"`
}

// ApiResponse defines model for ApiResponse.
type ApiResponse struct {
	Message string `json:"message"`
}

// CollectionAlias defines model for CollectionAlias.
type CollectionAlias struct {
	// Name of the collection the alias mapped to
	CollectionName string `json:"collection_name"`

	// Name of the collection alias
	Name string `json:"name"`
}

// CollectionAliasSchema defines model for CollectionAliasSchema.
type CollectionAliasSchema struct {
	// Name of the collection you wish to map the alias to
	CollectionName string `json:"collection_name"`
}

// CollectionAliasesResponse defines model for CollectionAliasesResponse.
type CollectionAliasesResponse struct {
	Aliases []*CollectionAlias `json:"aliases"`
}

// CollectionResponse defines model for CollectionResponse.
type CollectionResponse struct {
	// Embedded struct due to allOf(#/components/schemas/CollectionSchema)
	CollectionSchema `yaml:",inline"`
	// Embedded fields due to inline allOf schema
	// Timestamp of when the collection was created
	CreatedAt int64 `json:"created_at"`

	// Number of documents in the collection
	NumDocuments    int64 `json:"num_documents"`
	NumMemoryShards int64 `json:"num_memory_shards"`
}

// CollectionSchema defines model for CollectionSchema.
type CollectionSchema struct {
	// The name of an int32 / float field that determines the order in which the search results are ranked when a sort_by clause is not provided during searching. This field must indicate some kind of popularity.
	DefaultSortingField *string `json:"default_sorting_field,omitempty"`

	// A list of fields for querying, filtering and faceting
	Fields []Field `json:"fields"`

	// Name of the collection
	Name string `json:"name"`
}

// Field defines model for Field.
type Field struct {
	Facet    *bool     `json:"facet,omitempty"`
	Index    *bool     `json:"index,omitempty"`
	Name     string    `json:"name"`
	Optional *bool     `json:"optional,omitempty"`
	Type     FieldType `json:"type"`
}

// FieldType defines model for Field.Type.
type FieldType string

// HealthStatus defines model for HealthStatus.
type HealthStatus struct {
	Ok bool `json:"ok"`
}

// SearchGroupedHit defines model for SearchGroupedHit.
type SearchGroupedHit struct {
	GroupKey []string `json:"group_key"`

	// The documents that matched the search query
	Hits []SearchResultHit `json:"hits"`
}

// SearchHighlight defines model for SearchHighlight.
type SearchHighlight struct {
	Field string `json:"field"`

	// The indices property will be present only for string[] fields and will contain the corresponding indices of the snippets in the search field
	Indices       []int         `json:"indices"`
	MatchedTokens []interface{} `json:"matched_tokens"`

	// Present only for (non-array) string fields
	Snippet string `json:"snippet"`

	// Present only for (array) string[] fields
	Snippets []string `json:"snippets"`
}

// SearchOverride defines model for SearchOverride.
type SearchOverride struct {
	// Embedded struct due to allOf(#/components/schemas/SearchOverrideSchema)
	SearchOverrideSchema `yaml:",inline"`
	// Embedded fields due to inline allOf schema
	Id string `json:"id"`
}

// SearchOverrideExclude defines model for SearchOverrideExclude.
type SearchOverrideExclude struct {
	// document id that should be excluded from the search results.
	Id string `json:"id"`
}

// SearchOverrideInclude defines model for SearchOverrideInclude.
type SearchOverrideInclude struct {
	// document id that should be included
	Id string `json:"id"`

	// position number where document should be included in the search results
	Position int `json:"position"`
}

// SearchOverrideRule defines model for SearchOverrideRule.
type SearchOverrideRule struct {
	// Indicates whether the match on the query term should be `exact` or `contains`. If we want to match all queries that contained the word `apple`, we will use the `contains` match instead.
	Match SearchOverrideRuleMatch `json:"match"`

	// Indicates what search queries should be overridden
	Query string `json:"query"`
}

// Indicates whether the match on the query term should be `exact` or `contains`. If we want to match all queries that contained the word `apple`, we will use the `contains` match instead.
type SearchOverrideRuleMatch string

// SearchOverrideSchema defines model for SearchOverrideSchema.
type SearchOverrideSchema struct {
	// List of document `id`s that should be excluded from the search results.
	Excludes []SearchOverrideExclude `json:"excludes"`

	// List of document `id`s that should be included in the search results with their corresponding `position`s.
	Includes []SearchOverrideInclude `json:"includes"`
	Rule     SearchOverrideRule      `json:"rule"`
}

// SearchOverridesResponse defines model for SearchOverridesResponse.
type SearchOverridesResponse struct {
	Overrides []*SearchOverride `json:"overrides"`
}

// SearchResult defines model for SearchResult.
type SearchResult struct {
	FacetCounts []int `json:"facet_counts"`

	// The number of documents found
	Found       int                `json:"found"`
	GroupedHits []SearchGroupedHit `json:"grouped_hits"`

	// The documents that matched the search query
	Hits []SearchResultHit `json:"hits"`

	// The search result page number
	Page          int `json:"page"`
	RequestParams struct {
		CollectionName string `json:"collection_name"`
		PerPage        int    `json:"per_page"`
		Q              string `json:"q"`
	} `json:"request_params"`

	// The number of milliseconds the search took
	SearchTimeMs int `json:"search_time_ms"`
}

// SearchResultHit defines model for SearchResultHit.
type SearchResultHit struct {
	// Can be any key-value pair
	Document SearchResultHit_Document `json:"document"`

	// Contains highlighted portions of the search fields
	Highlights []SearchHighlight `json:"highlights"`
}

// Can be any key-value pair
type SearchResultHit_Document struct {
	AdditionalProperties map[string]map[string]interface{} `json:"-"`
}

// SearchSynonym defines model for SearchSynonym.
type SearchSynonym struct {
	// Embedded struct due to allOf(#/components/schemas/SearchSynonymSchema)
	SearchSynonymSchema `yaml:",inline"`
	// Embedded fields due to inline allOf schema
	Id string `json:"id"`
}

// SearchSynonymSchema defines model for SearchSynonymSchema.
type SearchSynonymSchema struct {
	// For 1-way synonyms, indicates the root word that words in the `synonyms` parameter map to.
	Root string `json:"root"`

	// Array of words that should be considered as synonyms.
	Synonyms []string `json:"synonyms"`
}

// SearchSynonymsResponse defines model for SearchSynonymsResponse.
type SearchSynonymsResponse struct {
	Synonyms []*SearchSynonym `json:"synonyms"`
}

// SuccessStatus defines model for SuccessStatus.
type SuccessStatus struct {
	Success bool `json:"success"`
}

// UpsertAliasJSONBody defines parameters for UpsertAlias.
type UpsertAliasJSONBody CollectionAliasSchema

// CreateCollectionJSONBody defines parameters for CreateCollection.
type CreateCollectionJSONBody CollectionSchema

// DeleteDocumentsParams defines parameters for DeleteDocuments.
type DeleteDocumentsParams struct {
	DeleteDocumentsParameters *struct {
		// Batch size parameter controls the number of documents that should be deleted at a time. A larger value will speed up deletions, but will impact performance of other operations running on the server.
		BatchSize *int    `json:"batch_size,omitempty"`
		FilterBy  *string `json:"filter_by,omitempty"`
	} `json:"deleteDocumentsParameters,omitempty"`
}

// IndexDocumentJSONBody defines parameters for IndexDocument.
type IndexDocumentJSONBody interface{}

// IndexDocumentParams defines parameters for IndexDocument.
type IndexDocumentParams struct {
	// Additional action to perform
	Action *IndexDocumentParamsAction `json:"action,omitempty"`
}

// IndexDocumentParamsAction defines parameters for IndexDocument.
type IndexDocumentParamsAction string

// ExportDocumentsParams defines parameters for ExportDocuments.
type ExportDocumentsParams struct {
	ExportDocumentsParameters *struct {
		// List of fields from the document to exclude in the search result
		ExcludeFields []string `json:"exclude_fields"`

		// Filter conditions for refining your search results. Separate multiple conditions with &&.
		FilterBy *string `json:"filter_by,omitempty"`

		// List of fields from the document to include in the search result
		IncludeFields []string `json:"include_fields"`
	} `json:"exportDocumentsParameters,omitempty"`
}

// ImportDocumentsParams defines parameters for ImportDocuments.
type ImportDocumentsParams struct {
	ImportDocumentsParameters *struct {
		Action      *string                                                    `json:"action,omitempty"`
		BatchSize   *int                                                       `json:"batch_size,omitempty"`
		DirtyValues *ImportDocumentsParamsImportDocumentsParametersDirtyValues `json:"dirty_values,omitempty"`
	} `json:"importDocumentsParameters,omitempty"`
}

// ImportDocumentsParamsImportDocumentsParametersDirtyValues defines parameters for ImportDocuments.
type ImportDocumentsParamsImportDocumentsParametersDirtyValues string

// SearchCollectionParams defines parameters for SearchCollection.
type SearchCollectionParams struct {
	SearchParameters struct {
		// If the number of results found for a specific query is less than this number, Typesense will attempt to drop the tokens in the query until enough results are found. Tokens that have the least individual hits are dropped first. Set to 0 to disable. Default: 10
		DropTokensThreshold *int `json:"drop_tokens_threshold,omitempty"`

		// If you have some overrides defined but want to disable all of them during query time, you can do that by setting this parameter to false
		EnableOverrides *bool `json:"enable_overrides,omitempty"`

		// List of fields from the document to exclude in the search result
		ExcludeFields []string `json:"exclude_fields"`

		// A list of fields that will be used for faceting your results on. Separate multiple fields with a comma.
		FacetBy []string `json:"facet_by"`

		// Facet values that are returned can now be filtered via this parameter. The matching facet text is also highlighted. For example, when faceting by `category`, you can set `facet_query=category:shoe` to return only facet values that contain the prefix "shoe".
		FacetQuery *string `json:"facet_query,omitempty"`

		// Filter conditions for refining your search results. Separate multiple conditions with &&.
		FilterBy *string `json:"filter_by,omitempty"`

		// You can aggregate search results into groups or buckets by specify one or more `group_by` fields. Separate multiple fields with a comma. To group on a particular field, it must be a faceted field.
		GroupBy []string `json:"group_by"`

		// Maximum number of hits to be returned for every group. If the `group_limit` is set as `K` then only the top K hits in each group are returned in the response. Default: 3
		GroupLimit *int `json:"group_limit,omitempty"`

		// A list of records to unconditionally hide from search results. A list of `record_id`s to hide. Eg: to hide records with IDs 123 and 456, you'd specify `123,456`.
		// You could also use the Overrides feature to override search results based on rules. Overrides are applied first, followed by `pinned_hits` and finally `hidden_hits`.
		HiddenHits []string `json:"hidden_hits"`

		// The number of tokens that should surround the highlighted text on each side. Default: 4
		HighlightAffixNumTokens *int `json:"highlight_affix_num_tokens,omitempty"`

		// The end tag used for the highlighted snippets. Default: `</mark>`
		HighlightEndTag *string `json:"highlight_end_tag,omitempty"`

		// A list of custom fields that must be highlighted even if you don't query  for them
		HighlightFields *[]string `json:"highlight_fields,omitempty"`

		// List of fields which should be highlighted fully without snippeting
		HighlightFullFields []string `json:"highlight_full_fields"`

		// The start tag used for the highlighted snippets. Default: `<mark>`
		HighlightStartTag *string `json:"highlight_start_tag,omitempty"`

		// List of fields from the document to include in the search result
		IncludeFields []string `json:"include_fields"`

		// Maximum number of facet values to be returned.
		MaxFacetValues *int `json:"max_facet_values,omitempty"`

		// Maximum number of hits returned. Increasing this value might increase search latency. Default: 500. Use `all` to return all hits found.
		MaxHits *interface{} `json:"max_hits,omitempty"`

		// The number of typographical errors (1 or 2) that would be tolerated. Default: 2
		NumTypos *int `json:"num_typos,omitempty"`

		// Results from this specific page number would be fetched.
		Page *int `json:"page,omitempty"`

		// Number of results to fetch per page. Default: 10
		PerPage *int `json:"per_page,omitempty"`

		// A list of records to unconditionally include in the search results at specific positions. An example use case would be to feature or promote certain items on the top of search results. A list of `record_id:hit_position`. Eg: to include a record with ID 123 at Position 1 and another record with ID 456 at Position 5, you'd specify `123:1,456:5`.
		// You could also use the Overrides feature to override search results based on rules. Overrides are applied first, followed by `pinned_hits` and  finally `hidden_hits`.
		PinnedHits []string `json:"pinned_hits"`

		// You can index content from any logographic language into Typesense if you are able to segment / split the text into space-separated words yourself  before indexing and querying.
		// Set this parameter to tre to do the same
		PreSegmentedQuery *bool `json:"pre_segmented_query,omitempty"`

		// Boolean field to indicate that the last word in the query should be treated as a prefix, and not as a whole word. This is used for building autocomplete and instant search interfaces. Defaults to true.
		Prefix []bool `json:"prefix"`

		// Set this parameter to true to ensure that an exact match is ranked above the others
		PrioritizeExactMatch *bool `json:"prioritize_exact_match,omitempty"`

		// The query text to search for in the collection. Use * as the search string to return all documents. This is typically useful when used in conjunction with filter_by.
		Q string `json:"q"`

		// A list of `string` fields that should be queried against. Multiple fields are separated with a comma.
		QueryBy []string `json:"query_by"`

		// The relative weight to give each `query_by` field when ranking results. This can be used to boost fields in priority, when looking for matches. Multiple fields are separated with a comma.
		QueryByWeights []string `json:"query_by_weights"`

		// Field values under this length will be fully highlighted, instead of showing a snippet of relevant portion. Default: 30
		SnippetThreshold *int `json:"snippet_threshold,omitempty"`

		// A list of numerical fields and their corresponding sort orders that will be used for ordering your results. Up to 3 sort fields can be specified. The text similarity score is exposed as a special `_text_match` field that you can use in the list of sorting fields. If no `sort_by` parameter is specified, results are sorted by `_text_match:desc,default_sorting_field:desc`
		SortBy []string `json:"sort_by"`

		// If the number of results found for a specific query is less than this number, Typesense will attempt to look for tokens with more typos until enough results are found. Default: 100
		TypoTokensThreshold *int `json:"typo_tokens_threshold,omitempty"`
	} `json:"searchParameters"`
}

// UpdateDocumentJSONBody defines parameters for UpdateDocument.
type UpdateDocumentJSONBody interface{}

// UpsertSearchOverrideJSONBody defines parameters for UpsertSearchOverride.
type UpsertSearchOverrideJSONBody SearchOverrideSchema

// UpsertSearchSynonymJSONBody defines parameters for UpsertSearchSynonym.
type UpsertSearchSynonymJSONBody SearchSynonymSchema

// CreateKeyJSONBody defines parameters for CreateKey.
type CreateKeyJSONBody ApiKeySchema

// TakeSnapshotParams defines parameters for TakeSnapshot.
type TakeSnapshotParams struct {
	// The directory on the server where the snapshot should be saved.
	SnapshotPath string `json:"snapshot_path"`
}

// UpsertAliasJSONRequestBody defines body for UpsertAlias for application/json ContentType.
type UpsertAliasJSONRequestBody UpsertAliasJSONBody

// CreateCollectionJSONRequestBody defines body for CreateCollection for application/json ContentType.
type CreateCollectionJSONRequestBody CreateCollectionJSONBody

// IndexDocumentJSONRequestBody defines body for IndexDocument for application/json ContentType.
type IndexDocumentJSONRequestBody IndexDocumentJSONBody

// UpdateDocumentJSONRequestBody defines body for UpdateDocument for application/json ContentType.
type UpdateDocumentJSONRequestBody UpdateDocumentJSONBody

// UpsertSearchOverrideJSONRequestBody defines body for UpsertSearchOverride for application/json ContentType.
type UpsertSearchOverrideJSONRequestBody UpsertSearchOverrideJSONBody

// UpsertSearchSynonymJSONRequestBody defines body for UpsertSearchSynonym for application/json ContentType.
type UpsertSearchSynonymJSONRequestBody UpsertSearchSynonymJSONBody

// CreateKeyJSONRequestBody defines body for CreateKey for application/json ContentType.
type CreateKeyJSONRequestBody CreateKeyJSONBody

// Getter for additional properties for SearchResultHit_Document. Returns the specified
// element and whether it was found
func (a SearchResultHit_Document) Get(fieldName string) (value map[string]interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for SearchResultHit_Document
func (a *SearchResultHit_Document) Set(fieldName string, value map[string]interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for SearchResultHit_Document to handle AdditionalProperties
func (a *SearchResultHit_Document) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal map[string]interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for SearchResultHit_Document to handle AdditionalProperties
func (a SearchResultHit_Document) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}
