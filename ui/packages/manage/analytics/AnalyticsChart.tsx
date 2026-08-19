import { useEffect, useRef } from "react";
import * as d3 from "d3";
import { type AnalyticsSeries } from "./AnalyticsUtils";

const CHART_HEIGHT = 240;

// AnalyticsChart renders one column-chart panel per series (see
// AnalyticsOverviewAction's own `series` field - each one already arrives
// pre-aggregated: a monthly bucket count, or a categorical breakdown), unlike
// internal-stats/InternalStatsMetricsChart.tsx which facets a flat list of raw
// items by unit itself. There's no severity per point here (these are plain
// counts, not a threshold read on a resource) - compare-magnitude is the whole
// job, so every bar uses the same single sequential hue (see color-formula: one
// hue for magnitude) rather than a status palette.
export function AnalyticsChart({ series }: { series: AnalyticsSeries[] }) {
  return (
    <div className="analytics-chart-panels">
      {series.map((s) => (
        <SeriesChart key={s.key} series={s} />
      ))}
    </div>
  );
}

// "2026-08" (a monthlySeries bucket label - see AnalyticsActionImplementation.go)
// reads as "Aug" on the axis; anything else (a workspace type title, a role name)
// passes through unchanged.
function formatPointLabel(label: string): string {
  const match = /^(\d{4})-(\d{2})$/.exec(label);
  if (!match) {
    return label;
  }
  const date = new Date(Number(match[1]), Number(match[2]) - 1, 1);
  return date.toLocaleDateString(undefined, { month: "short" });
}

function SeriesChart({ series }: { series: AnalyticsSeries }) {
  const containerRef = useRef<HTMLDivElement | null>(null);
  const svgRef = useRef<SVGSVGElement | null>(null);
  const points = series.points.get();

  useEffect(() => {
    const container = containerRef.current;
    const svgEl = svgRef.current;
    if (!container || !svgEl || points.length === 0) return;

    const render = () => {
      const width = Math.max(container.clientWidth, 240);
      const height = CHART_HEIGHT;
      const rotateLabels = points.length > 6;
      const margin = {
        top: 24,
        right: 8,
        bottom: rotateLabels ? 46 : 28,
        left: 8,
      };
      const innerWidth = width - margin.left - margin.right;
      const innerHeight = height - margin.top - margin.bottom;

      const svg = d3.select(svgEl);
      svg.selectAll("*").remove();
      svg
        .attr("width", width)
        .attr("height", height)
        .attr("viewBox", `0 0 ${width} ${height}`)
        .attr("role", "img")
        .attr("aria-label", series.label);

      const x = d3
        .scaleBand()
        .domain(points.map((p) => p.label))
        .range([0, innerWidth])
        .padding(0.35);

      const maxValue = Math.max(1, d3.max(points, (p) => p.value) || 1);
      const y = d3
        .scaleLinear()
        .domain([0, maxValue])
        .nice()
        .range([innerHeight, 0]);

      const g = svg
        .append("g")
        .attr("transform", `translate(${margin.left},${margin.top})`);

      // Shared y-axis gridlines - hairline, recessive, one per chart (never per bar).
      g.append("g")
        .attr("class", "analytics-grid")
        .call(
          d3
            .axisLeft(y)
            .ticks(4)
            .tickSize(-innerWidth)
            .tickFormat(() => ""),
        )
        .call((sel) => sel.select(".domain").remove());

      const bars = g
        .selectAll(".analytics-bar")
        .data(points, (p: any) => p.label)
        .join("g")
        .attr("class", "analytics-bar")
        .attr("transform", (p) => `translate(${x(p.label)}, 0)`);

      bars
        .append("rect")
        .attr("class", "analytics-bar-fill")
        .attr("x", 0)
        .attr("y", (p) => y(p.value))
        .attr("width", x.bandwidth())
        .attr("height", (p) => Math.max(0, innerHeight - y(p.value)))
        .attr("rx", 4);

      // Direct value labels - only above each bar's own cap (per mark spec:
      // "Columns → value on the cap"), never one per gridline tick.
      bars
        .append("text")
        .attr("class", "analytics-bar-value")
        .attr("x", x.bandwidth() / 2)
        .attr("y", (p) => y(p.value) - 6)
        .attr("text-anchor", "middle")
        .text((p) => (p.value > 0 ? d3.format("~s")(p.value) : ""));

      const axisG = g
        .append("g")
        .attr("class", "analytics-axis-x")
        .attr("transform", `translate(0, ${innerHeight})`)
        .call(
          d3
            .axisBottom(x)
            .tickSize(0)
            .tickFormat((label) => formatPointLabel(label)),
        );
      axisG.select(".domain").remove();

      if (rotateLabels) {
        axisG
          .selectAll("text")
          .attr("text-anchor", "end")
          .attr("dx", "-0.4em")
          .attr("dy", "0.4em")
          .attr("transform", "rotate(-40)");
      }
    };

    render();

    const observer = new ResizeObserver(() => render());
    observer.observe(container);
    return () => observer.disconnect();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [series.key]);

  if (points.length === 0) {
    return null;
  }

  return (
    <div className="analytics-facet" ref={containerRef}>
      <div className="analytics-facet-title">{series.label}</div>
      <svg ref={svgRef} />
    </div>
  );
}
