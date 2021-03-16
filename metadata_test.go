package nrzap

import (
	"context"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/newrelic/go-agent/v3/integrations/logcontext"
	"github.com/newrelic/go-agent/v3/newrelic"
	"go.uber.org/zap"
	"testing"
)

func TestGetNrMetadata(t *testing.T) {
	t.Parallel()
	tests := [...]struct {
		name string
		ctx func(*testing.T, string, string) context.Context
		want []zap.Field
	}{
		{
			name: "ok",
			ctx:  setup,
			want: []zap.Field{
				zap.String(logcontext.KeyTraceID, ""),
				zap.String(logcontext.KeySpanID, ""),
				zap.String(logcontext.KeyEntityName, "TestGetNrMetadata/ok"), // method/appName
				zap.String(logcontext.KeyEntityType, "SERVICE"),
				zap.String(logcontext.KeyEntityGUID, ""),
				zap.String(logcontext.KeyHostname, ""),
			},
		},
		{
			name: "EmptyContext",
			ctx:  func(*testing.T, string, string) context.Context{
				return context.Background()
			},
			want: []zap.Field{},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := GetNrMetadataFields(tt.ctx(t, t.Name(), "txn_name"))
			if diff := cmp.Diff(got, tt.want,
				cmpopts.IgnoreSliceElements(func(f zap.Field) bool {
					if f.Key == logcontext.KeyTraceID || f.Key == logcontext.KeyHostname {
						return true
					}
					return false
				}),
			); diff!= "" {
				t.Errorf("GetNrMetadataFields() differs: (-got +want)\n%s", diff)
			}
		})
	}
}


func setup(t *testing.T, appName, txnName string) context.Context {
	t.Helper()
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName(appName),
		newrelic.ConfigLicense("1234567890123456789012345678901234567890"),
		newrelic.ConfigDistributedTracerEnabled(true),
	)
	if err != nil {
		t.Fatal(err)
	}

	txn := app.StartTransaction(txnName)
	return newrelic.NewContext(context.Background(), txn)
}