package nrzap

import (
	"context"
	"github.com/newrelic/go-agent/v3/integrations/logcontext"
	"github.com/newrelic/go-agent/v3/newrelic"
	"go.uber.org/zap"
)

func GetNrMetadataFields(ctx context.Context) []zap.Field {
	txn := newrelic.FromContext(ctx)
	if txn == nil {
		return []zap.Field{}
	}
	md := txn.GetLinkingMetadata()
	return []zap.Field{
		zap.String(logcontext.KeyTraceID, md.TraceID),
		zap.String(logcontext.KeySpanID, md.SpanID),
		zap.String(logcontext.KeyEntityName, md.EntityName),
		zap.String(logcontext.KeyEntityType, md.EntityType),
		zap.String(logcontext.KeyEntityGUID, md.EntityGUID),
		zap.String(logcontext.KeyHostname, md.Hostname),
	}
}