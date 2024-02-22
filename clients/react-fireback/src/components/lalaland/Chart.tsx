import { random } from "lodash";
import { useEffect, useRef } from "react";
var Highcharts = require("highcharts");

declare var Highcharts: any;

function previousSeconds(s: number) {
  return new Date().getTime() - s * 1000;
}

function generateMockSeries() {
  const series = [];
  for (let i = 20; i >= 1; i = i - 2) {
    series.push([previousSeconds(i), Math.floor(Math.random() * 8 + 30)]);
  }
  return series;
}

function drawChart(id: string) {
  const data = {
    title: "Bedroom Temperature",
    subTitle: "Show Today Live Statistics",
    for: "Temperature",
    unit: "°C",
    chartColor: "orange",
    series: generateMockSeries(),
  };

  return Highcharts.chart("sample_chart" + id, {
    chart: {
      events: {
        redraw: function () {
          const self: any = this;
          setTimeout(function () {
            // self.reflow();
          }, 10);
        },
      },
      backgroundColor: "rgba(255, 255, 255, 0.0)",
    },
    credits: {
      enabled: false,
    },
    title: {
      text: null,
    },
    yAxis: {
      title: {
        text: null,
      },
      visible: false,
    },
    xAxis: {
      visible: false,
      type: "datetime",
      dateTimeLabelFormats: {
        minute: "%H:%M",
      },
      //   reversed:
      //     this.globalization.getLayoutDirection() === "ltr" ? false : true,
    },
    legend: {
      layout: "horizontal",
      align: "left",
      verticalAlign: "top",
    },
    tooltip: {
      useHTML: true,
    },
    plotOptions: {
      series: {
        label: {
          connectorAllowed: false,
        },
        pointStart: new Date().getTime(),
      },
    },
    series: [
      {
        name: data.unit,
        data: data.series,
        type: "spline",
        color: data.chartColor,
        showInLegend: false,
        marker: {
          enabled: false,
        },
      },
    ],
    responsive: {
      rules: [
        {
          condition: {
            maxWidth: 500,
          },
          chartOptions: {
            legend: {
              layout: "horizontal",
              align: "center",
              verticalAlign: "bottom",
            },
          },
        },
      ],
    },
  });
}

export function Chart({ value, id }: { value: number; id: string }) {
  const chartRef = useRef<any>();

  useEffect(() => {
    const chart = drawChart(id);
    chartRef.current = chart;
  }, []);

  useEffect(() => {
    const intv = setInterval(() => {
      chartRef.current.series[0].addPoint(
        [new Date(), random(30, 300) || 100],
        true,
        true
      );
    }, 150);

    return () => clearInterval(intv);
  }, [value]);

  return (
    <>
      <div
        style={{ height: "150px", width: "100%" }}
        id={"sample_chart" + id}
        className="chart"
      ></div>
    </>
  );
}
