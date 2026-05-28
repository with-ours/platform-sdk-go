// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package oursprivacy

import (
	"context"
	"net/http"
	"slices"

	"github.com/with-ours/platform-sdk-go/internal/apijson"
	"github.com/with-ours/platform-sdk-go/internal/requestconfig"
	"github.com/with-ours/platform-sdk-go/option"
	"github.com/with-ours/platform-sdk-go/packages/param"
	"github.com/with-ours/platform-sdk-go/packages/respjson"
)

// TagManagerAssetFolderService contains methods and other services that help with
// interacting with the ours-privacy-platform API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTagManagerAssetFolderService] method instead.
type TagManagerAssetFolderService struct {
	Options []option.RequestOption
}

// NewTagManagerAssetFolderService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTagManagerAssetFolderService(opts ...option.RequestOption) (r TagManagerAssetFolderService) {
	r = TagManagerAssetFolderService{}
	r.Options = opts
	return
}

// Assign a tag, trigger, or variable to a folder within its tag manager, or send
// `folderId: null` to remove the asset from its current folder. The assignment is
// a full replace — calling it again with a different `folderId` silently moves the
// asset. Requires scope: tagManagers:update
func (r *TagManagerAssetFolderService) New(ctx context.Context, body TagManagerAssetFolderNewParams, opts ...option.RequestOption) (res *TagManagerAssetFolderNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/tag-manager-asset-folders"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type TagManagerAssetFolderNewResponse struct {
	ID        string `json:"id" api:"required"`
	AccountID string `json:"accountId" api:"required"`
	AssetID   string `json:"assetId" api:"required"`
	// Any of "tagManagerTag", "tagManagerTrigger", "tagManagerVariable".
	AssetType    TagManagerAssetFolderNewResponseAssetType `json:"assetType" api:"required"`
	TagManagerID string                                    `json:"tagManagerId" api:"required"`
	CreatedAt    string                                    `json:"createdAt" api:"nullable"`
	FolderID     string                                    `json:"folderId" api:"nullable"`
	UpdatedAt    string                                    `json:"updatedAt" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		AccountID    respjson.Field
		AssetID      respjson.Field
		AssetType    respjson.Field
		TagManagerID respjson.Field
		CreatedAt    respjson.Field
		FolderID     respjson.Field
		UpdatedAt    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TagManagerAssetFolderNewResponse) RawJSON() string { return r.JSON.raw }
func (r *TagManagerAssetFolderNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagManagerAssetFolderNewResponseAssetType string

const (
	TagManagerAssetFolderNewResponseAssetTypeTagManagerTag      TagManagerAssetFolderNewResponseAssetType = "tagManagerTag"
	TagManagerAssetFolderNewResponseAssetTypeTagManagerTrigger  TagManagerAssetFolderNewResponseAssetType = "tagManagerTrigger"
	TagManagerAssetFolderNewResponseAssetTypeTagManagerVariable TagManagerAssetFolderNewResponseAssetType = "tagManagerVariable"
)

type TagManagerAssetFolderNewParams struct {
	// Folder UUID to assign to. Send `null` to remove the asset from its current
	// folder.
	FolderID param.Opt[string] `json:"folderId,omitzero" api:"required" format:"uuid"`
	// UUID of the tag, trigger, or variable to assign.
	AssetID string `json:"assetId" api:"required" format:"uuid"`
	// Asset type to assign. Must be one of `tagManagerTag`, `tagManagerTrigger`, or
	// `tagManagerVariable`.
	//
	// Any of "tagManagerTag", "tagManagerTrigger", "tagManagerVariable".
	AssetType TagManagerAssetFolderNewParamsAssetType `json:"assetType,omitzero" api:"required"`
	paramObj
}

func (r TagManagerAssetFolderNewParams) MarshalJSON() (data []byte, err error) {
	type shadow TagManagerAssetFolderNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TagManagerAssetFolderNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Asset type to assign. Must be one of `tagManagerTag`, `tagManagerTrigger`, or
// `tagManagerVariable`.
type TagManagerAssetFolderNewParamsAssetType string

const (
	TagManagerAssetFolderNewParamsAssetTypeTagManagerTag      TagManagerAssetFolderNewParamsAssetType = "tagManagerTag"
	TagManagerAssetFolderNewParamsAssetTypeTagManagerTrigger  TagManagerAssetFolderNewParamsAssetType = "tagManagerTrigger"
	TagManagerAssetFolderNewParamsAssetTypeTagManagerVariable TagManagerAssetFolderNewParamsAssetType = "tagManagerVariable"
)
