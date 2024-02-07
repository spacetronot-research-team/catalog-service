package logrushook

import (
	"github.com/sirupsen/logrus"
	"github.com/spacetronot-research-team/catalog-service/pkg/trace"
)

// Trace is a Logrus hook that adds request_id to entry data.
type Trace struct{}

func (*Trace) Fire(entry *logrus.Entry) error {
	if entry.Context == nil {
		return nil
	}

	if traceID, ok := entry.Context.Value(trace.Key).(string); ok {
		entry.Data["request_id"] = traceID
	}

	return nil
}

func (*Trace) Levels() []logrus.Level {
	return logrus.AllLevels
}
