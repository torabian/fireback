import { useEffect, useMemo, useRef } from "react";
import * as d3 from "d3";
import { useS } from "@fireback/ui-core/hooks/useS";
import {
  groupByUnit,
  normalizeSeverity,
  type InternalStatItem,
} from "./InternalStatsUtils";
import { strings } from "./strings/translations";

// Dynamic-key indexing (t.internalStats[someVar]) doesn't typecheck against
// translations.ts's plain object shape, so unit -> label goes through a real
// switch instead of a lookup table.
function unitTitle(t: typeof strings, unit: string): string {
  switch (unit) {
    case "percent":
      return t.internalStats.unitPercent;
    case "bytes":
      return t.internalStats.unitBytes;
    case "seconds":
      return t.internalStats.unitSeconds;
    case "mhz":
      return t.internalStats.unitMhz;
    case "load":
      return t.internalStats.unitLoad;
    case "count":
      return t.internalStats.unitCount;
    default:
      return t.internalStats.unitOther;
  }
}

const ROW_HEIGHT = 26;
const BAR_HEIGHT = 16;
const LABEL_WIDTH = 128;
const VALUE_WIDTH = 66;
const GUTTER = 8;
const OUTER_PADDING = 8;

// Fixed 4-way status palette (see modules/internalstats/Collector.go's
// Severity* constants) - never themed, never reused for a "series", styled
// purely via CSS custom properties (--is-status-*) so light/dark stay in one
// place. "info" is the vast majority of stats (anything Collector.go doesn't
// apply a percent threshold to) - it's deliberately the muted/neutral color,
// not a 4th alarm state.
const SEVERITY_CLASS: Record<string, string> = {
  ok: "is-severity-ok",
  warn: "is-severity-warn",
  critical: "is-severity-critical",
  info: "is-severity-info",
};

export function InternalStatsMetricsChart({
  items,
}: {
  items: InternalStatItem[];
}) {
  const t = useS(strings);
  const facets = useMemo(() => groupByUnit(items), [items]);

  if (facets.length === 0) {
    return <p className="internal-stats-empty-hint">{t.internalStats.noNumericMetrics}</p>;
  }

  return (
    <>
      <ChartLegend />
      <div className="internal-stats-chart-panels">
        {facets.map(([unit, unitItems]) => (
          <UnitFacetChart key={unit} unit={unit} items={unitItems} />
        ))}
      </div>
    </>
  );
}

function ChartLegend() {
  const t = useS(strings);
  const rows: Array<{ severity: string; label: string }> = [
    { severity: "ok", label: t.internalStats.severityOk },
    { severity: "warn", label: t.internalStats.severityWarn },
    { severity: "critical", label: t.internalStats.severityCritical },
    { severity: "info", label: t.internalStats.severityInfo },
  ];
  return (
    <div className="internal-stats-legend">
      {rows.map((row) => (
        <div className="internal-stats-legend-item" key={row.severity}>
          <span
            className={`internal-stats-legend-swatch ${SEVERITY_CLASS[row.severity]}`}
          />
          {row.label}
        </div>
      ))}
    </div>
  );
}

function axisTickFormat(unit: string): (value: d3.NumberValue) => string {
  switch (unit) {
    case "percent":
      return (v) => `${v}%`;
    case "bytes":
      return (v) => d3.format(".2~s")(v as number).replace(/G$/, "GB").replace(/M$/, "MB").replace(/k$/, "KB");
    case "seconds":
      return (v) => `${d3.format("~s")(v as number)}s`;
    case "mhz":
      return (v) => `${v} MHz`;
    default:
      return d3.format("~s");
  }
}

// UnitFacetChart is one small-multiple: every metric sharing a unit (percent,
// bytes, seconds, ...), one horizontal bar per metric, all against a single
// shared x-scale/axis for that unit - never mixed with another unit's scale
// (comparing e.g. a percent and a byte count on one axis would be meaningless).
function UnitFacetChart({
  unit,
  items,
}: {
  unit: string;
  items: InternalStatItem[];
}) {
  const t = useS(strings);
  const containerRef = useRef<HTMLDivElement | null>(null);
  const svgRef = useRef<SVGSVGElement | null>(null);

  useEffect(() => {
    const container = containerRef.current;
    const svgEl = svgRef.current;
    if (!container || !svgEl) return;

    const render = () => {
      const width = Math.max(container.clientWidth, 240);
      const barAreaX0 = OUTER_PADDING + LABEL_WIDTH + GUTTER;
      const barAreaWidth = Math.max(
        40,
        width - barAreaX0 - VALUE_WIDTH - GUTTER - OUTER_PADDING,
      );
      const axisHeight = 18;
      const height =
        axisHeight + items.length * ROW_HEIGHT + OUTER_PADDING;

      const svg = d3.select(svgEl);
      svg.selectAll("*").remove();
      svg
        .attr("width", width)
        .attr("height", height)
        .attr("viewBox", `0 0 ${width} ${height}`)
        .attr("role", "img")
        .attr("aria-label", unitTitle(t, unit));

      const maxDomain =
        unit === "percent" ? 100 : Math.max(1, d3.max(items, (d) => d.rawValue) || 1);
      const x = d3.scaleLinear().domain([0, maxDomain]).range([0, barAreaWidth]);

      // Shared axis + hairline gridlines for the whole facet - one axis per
      // chart, never per row.
      const axisG = svg
        .append("g")
        .attr("class", "internal-stats-axis")
        .attr("transform", `translate(${barAreaX0}, ${axisHeight})`);
      axisG.call(
        d3
          .axisTop(x)
          .ticks(4)
          .tickFormat(axisTickFormat(unit) as any)
          .tickSize(-(items.length * ROW_HEIGHT)),
      );
      axisG.select(".domain").remove();

      const rows = svg
        .selectAll(".internal-stats-row")
        .data(items, (d: any) => d.key)
        .join("g")
        .attr("class", "internal-stats-row")
        .attr(
          "transform",
          (_d, i) => `translate(0, ${axisHeight + i * ROW_HEIGHT})`,
        );

      rows
        .append("text")
        .attr("class", "internal-stats-row-label")
        .attr("x", OUTER_PADDING)
        .attr("y", ROW_HEIGHT / 2)
        .attr("dominant-baseline", "middle")
        .text((d) => d.label);

      const barY = (ROW_HEIGHT - BAR_HEIGHT) / 2;

      rows
        .append("rect")
        .attr("class", "internal-stats-bar-track")
        .attr("x", barAreaX0)
        .attr("y", barY)
        .attr("width", barAreaWidth)
        .attr("height", BAR_HEIGHT)
        .attr("rx", 4);

      rows
        .append("rect")
        .attr(
          "class",
          (d) =>
            `internal-stats-bar-fill ${SEVERITY_CLASS[normalizeSeverity(d.severity)]}`,
        )
        .attr("x", barAreaX0)
        .attr("y", barY)
        .attr(
          "width",
          (d) => Math.max(2, x(Math.min(d.rawValue, maxDomain))),
        )
        .attr("height", BAR_HEIGHT)
        .attr("rx", 4);

      rows
        .append("text")
        .attr("class", "internal-stats-row-value")
        .attr("x", barAreaX0 + barAreaWidth + GUTTER)
        .attr("y", ROW_HEIGHT / 2)
        .attr("dominant-baseline", "middle")
        // The server already formats this ("62.3%", "8.0 GB", ...) - see
        // InternalStatsSnapshotActionResItems.value - so the chart never
        // reformats a raw number itself and can't drift from the stat tiles
        // above it.
        .text((d) => d.value);
    };

    render();

    const observer = new ResizeObserver(() => render());
    observer.observe(container);
    return () => observer.disconnect();
  }, [items, unit, t]);

  return (
    <div className="internal-stats-facet" ref={containerRef}>
      <div className="internal-stats-facet-title">{unitTitle(t, unit)}</div>
      <svg ref={svgRef} />
    </div>
  );
}
