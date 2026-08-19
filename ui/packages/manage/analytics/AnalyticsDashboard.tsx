import {
  PageTitle,
  usePageTitle,
} from "@fireback/ui-core/components/page-title/PageTitle";
import { PageSection } from "@fireback/ui-core/components/page-section/PageSection";
import { useS } from "@fireback/ui-core/hooks/useS";
import { useAnalyticsOverviewActionQuery } from "@fireback/manage/sdk/abac/AnalyticsOverviewAction";
import { AnalyticsChart } from "./AnalyticsChart";
import { groupByCategory, normalizeSeverity } from "./AnalyticsUtils";
import { strings } from "./strings/translations";

// A root-only admin dashboard, not a live-updating monitor (unlike internal-stats'
// 2s poll of a process health snapshot) - these are aggregate COUNT/GROUP BY queries
// over the whole install (see AnalyticsActionImplementation.go), not something
// worth recomputing every couple of seconds. Refetch on mount/focus (react-query's
// own default) is enough; an explicit refresh button covers "I just did something
// and want the numbers to catch up".
export const AnalyticsDashboard = () => {
  const t = useS(strings);
  // See internal-stats/InternalStatsDashboard.tsx's own comment on this same
  // pattern - PageTitle's `title` prop is unused by the component itself; the
  // <h1> it renders comes from PageTitleContext, set by this hook.
  usePageTitle(t.analytics.title);
  const query = useAnalyticsOverviewActionQuery({});

  const snapshot = query.data?.data?.item;
  const items = snapshot?.items?.get() ?? [];
  const series = snapshot?.series?.get() ?? [];
  const categories = groupByCategory(items);

  return (
    <div className="analytics-page">
      <PageTitle description={t.analytics.description} />

      <div className="mb-3">
        <button
          className="btn btn-sm btn-outline-secondary"
          disabled={query.isFetching}
          onClick={() => query.refetch()}
        >
          {query.isFetching ? t.analytics.refreshing : t.analytics.refresh}
        </button>
      </div>

      {query.isError && !snapshot && (
        <div className="alert alert-danger">{t.analytics.loadError}</div>
      )}

      {snapshot && (
        <>
          <div className="analytics-meta">
            <div className="analytics-meta-item">
              <span className="analytics-meta-label">
                {t.analytics.generatedAt}
              </span>
              <strong>
                {new Date(snapshot.generatedAt).toLocaleString()}
              </strong>
            </div>
            <div className="analytics-meta-item">
              <span className="analytics-meta-label">{t.analytics.metrics}</span>
              <strong>{items.length}</strong>
            </div>
          </div>

          <PageSection title={t.analytics.overviewSection}>
            <div className="analytics-categories">
              {categories.map(([category, categoryItems]) => (
                <div className="analytics-category" key={category}>
                  <h3 className="analytics-category-title">{category}</h3>
                  <div className="analytics-tiles">
                    {categoryItems.map((item) => (
                      <div className="analytics-tile" key={item.key}>
                        <div className="analytics-tile-label">
                          <span
                            className={`analytics-severity-dot is-severity-${normalizeSeverity(item.severity)}`}
                          />
                          {item.label}
                        </div>
                        <div className="analytics-tile-value">
                          {item.value}
                        </div>
                      </div>
                    ))}
                  </div>
                </div>
              ))}
            </div>
          </PageSection>

          <PageSection title={t.analytics.chartSection}>
            <AnalyticsChart series={series} />
          </PageSection>
        </>
      )}
    </div>
  );
};
