package asset

import (
	"infolelang/lib"
	models "infolelang/models/assets"
	"reflect"
	"testing"
	"time"

	elastic "gitlab.com/golang-package-library/elasticsearch"
	"gitlab.com/golang-package-library/logger"
)

func TestAssetRepository_GetApproval(t *testing.T) {
	type fields struct {
		db      lib.Database
		dbRaw   lib.Databases
		elastic elastic.Elasticsearch
		logger  logger.Logger
		timeout time.Duration
	}
	type args struct {
		pernr string
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		wantResponses []models.AssetsResponse
		wantErr       bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			asset := AssetRepository{
				db:      tt.fields.db,
				dbRaw:   tt.fields.dbRaw,
				elastic: tt.fields.elastic,
				logger:  tt.fields.logger,
				timeout: tt.fields.timeout,
			}
			gotResponses, err := asset.GetApproval(tt.args.pernr)
			if (err != nil) != tt.wantErr {
				t.Errorf("AssetRepository.GetApproval() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResponses, tt.wantResponses) {
				t.Errorf("AssetRepository.GetApproval() = %v, want %v", gotResponses, tt.wantResponses)
			}
		})
	}
}
