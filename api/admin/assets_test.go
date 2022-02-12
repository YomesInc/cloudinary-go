package admin_test

import (
	"testing"

	"github.com/cloudinary/cloudinary-go/api"
	"github.com/cloudinary/cloudinary-go/api/admin"
	"github.com/cloudinary/cloudinary-go/internal/cldtest"
)

func TestAssets_AssetTypes(t *testing.T) {
	resp, err := adminAPI.AssetTypes(ctx)

	if err != nil || len(resp.AssetTypes) < 1 {
		t.Error(err, resp)
	}
}

func TestAssets_Assets(t *testing.T) {
	cldtest.UploadTestAsset(t, cldtest.PublicID)
	resp, err := adminAPI.Assets(ctx, admin.AssetsParams{Tags: true, Context: true, Moderations: true, MaxResults: 1})

	if err != nil || len(resp.Assets) != 1 {
		t.Error(err, resp)
	}
}

func TestAssets_AssetsByIDs(t *testing.T) {
	cldtest.UploadTestVideoAsset(t, cldtest.VideoPublicID)
	resp, err := adminAPI.AssetsByIDs(ctx, admin.AssetsByIDsParams{PublicIDs: []string{cldtest.PublicID}, Tags: true})

	if err != nil || len(resp.Assets) != 1 {
		t.Error(err, resp)
	}

	resp, err = adminAPI.AssetsByIDs(ctx, admin.AssetsByIDsParams{PublicIDs: []string{cldtest.VideoPublicID}, AssetType: api.Video})

	if err != nil || len(resp.Assets) != 1 {
		t.Error(err, resp)
	}
}

func TestAssets_AssetsByAssetIDs(t *testing.T) {
	asset, err := cldtest.UploadTestAsset(t, cldtest.PublicID)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := adminAPI.AssetsByAssetIDs(ctx, admin.AssetsByAssetIDsParams{AssetIDs: api.CldAPIArray{asset.AssetID}})
	if err != nil {
		t.Fatal(err)
	}

	n := len(resp.Assets)
	if n != 1 {
		t.Errorf("got %d, want 1", n)
	}

	asset2, err := cldtest.UploadTestAsset(t, cldtest.PublicID2)
	if err != nil {
		t.Fatal(err)
	}

	resp, err = adminAPI.AssetsByAssetIDs(ctx, admin.AssetsByAssetIDsParams{AssetIDs: api.CldAPIArray{asset.AssetID, asset2.AssetID}})
	if err != nil {
		t.Fatal(err)
	}

	n = len(resp.Assets)
	if n != 2 {
		t.Errorf("got %d, want 2", n)
	}

}

func TestAssets_RestoreAssets(t *testing.T) {
	resp, err := adminAPI.RestoreAssets(ctx, admin.RestoreAssetsParams{PublicIDs: []string{"api_test_restore_20891", "api_test_restore_94060"}})
	if err != nil {
		t.Error(err, resp)
	}
}

func TestAssets_DeleteAssets(t *testing.T) {
	resp, err :=
		adminAPI.DeleteAssets(ctx, admin.DeleteAssetsParams{PublicIDs: []string{"api_test_restore_20891", "api_test_restore_94060"}})
	if err != nil {
		t.Error(err, resp)
	}
}
