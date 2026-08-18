import { ContentAreaLoader } from "@fireback/ui-core/components/content-area-loader/ContentAreaLoader";
import { PageTitle } from "@fireback/ui-core/components/page-title/PageTitle";
import { PageSection } from "@fireback/ui-core/components/page-section/PageSection";
import { useS } from "@fireback/ui-core/hooks/useS";
import { useInternalStatsSnapshotActionQuery } from "@fireback/manage/sdk/internalstats/InternalStatsSnapshotAction";
import { InternalStatsMetricsChart } from "./InternalStatsMetricsChart";
import { groupByCategory, normalizeSeverity } from "./InternalStatsUtils";
import { strings } from "./strings/translations";

// How often the snapshot re-fetches while this screen is open - mirrors
// InternalStatsModuleConfig.Interval's own default (see
// modules/internalstats/InternalStatsModule.go's DefaultInterval), so the
// dashboard refreshes at roughly the same cadence the server would push over
// InternalStatsStream. Plain polling (rather than wiring the websocket
// stream) keeps this screen on the same GET+react-query pattern every other
// manage screen already uses.
const REFRESH_INTERVAL_MS = 2000;

export const InternalStatsDashboard = () => {
  const t = useS(strings);
  const query = useInternalStatsSnapshotActionQuery({
    refetchInterval: REFRESH_INTERVAL_MS,
  });

  const snapshot = query.data?.data?.item;
  const items = snapshot?.items?.get() ?? [];
  const categories = groupByCategory(items);

  return (
    <div className="internal-stats-page">
      <PageTitle
        title={t.internalStats.title}
        description={t.internalStats.description}
      />

      {query.isError && !snapshot && (
        <div className="alert alert-danger">{t.internalStats.loadError}</div>
      )}

      {snapshot && (
        <>
          <div className="internal-stats-meta">
            <div className="internal-stats-meta-item">
              <span className="internal-stats-meta-label">
                {t.internalStats.host}
              </span>
              <strong>{snapshot.hostname}</strong>
            </div>
            <div className="internal-stats-meta-item">
              <span className="internal-stats-meta-label">
                {t.internalStats.generatedAt}
              </span>
              <strong>
                {new Date(snapshot.generatedAt).toLocaleTimeString()}
              </strong>
            </div>
            <div className="internal-stats-meta-item">
              <span className="internal-stats-meta-label">
                {t.internalStats.metrics}
              </span>
              <strong>{items.length}</strong>
            </div>
          </div>

          <PageSection title={t.internalStats.snapshotSection}>
            <div className="internal-stats-categories">
              {categories.map(([category, categoryItems]) => (
                <div className="internal-stats-category" key={category}>
                  <h3 className="internal-stats-category-title">{category}</h3>
                  <div className="internal-stats-tiles">
                    {categoryItems.map((item) => (
                      <div className="internal-stats-tile" key={item.key}>
                        <div className="internal-stats-tile-label">
                          <span
                            className={`internal-stats-severity-dot is-severity-${normalizeSeverity(item.severity)}`}
                          />
                          {item.label}
                        </div>
                        <div className="internal-stats-tile-value">
                          {item.value}
                        </div>
                      </div>
                    ))}
                  </div>
                </div>
              ))}
            </div>
          </PageSection>

          <PageSection title={t.internalStats.chartSection}>
            <InternalStatsMetricsChart items={items} />
          </PageSection>
        </>
      )}
    </div>
  );
};
