/**
 * English source for internal-stats/strings-en.yml. Hand-written (rather than
 * run through the fireback language & translation manager - see
 * ui/scripts/translation-manager) since there's no $fa/$pl translation yet;
 * useS falls back to `en` for any locale without its own "$<locale>" key, so
 * this is a valid, if English-only, translations.ts. Run `npm run t` from
 * ui/ once fa/pl copy is added to strings-fa.yml/strings-pl.yml to regenerate
 * this with $fa/$pl included.
 */
export const en = {
  internalStats: {
    title: "Internal Stats",
    description:
      "Server / process health - CPU, memory, disk, network and Go runtime - refreshed automatically.",
    host: "Host",
    generatedAt: "Generated",
    metrics: "Metrics",
    snapshotSection: "Snapshot",
    chartSection: "Numeric metrics",
    loadError:
      "Could not load internal stats. This requires a root workspace session.",
    noNumericMetrics: "No numeric metrics available yet.",
    severityOk: "OK",
    severityWarn: "Warning",
    severityCritical: "Critical",
    severityInfo: "Info",
    unitPercent: "Usage (percent)",
    unitBytes: "Memory / disk / network (bytes)",
    unitSeconds: "Durations (seconds)",
    unitMhz: "Frequency (MHz)",
    unitLoad: "Load average",
    unitCount: "Counts",
    unitOther: "Other",
  },
};
export const strings = { ...en };
