import { type DatatableColumn } from "@fireback/ui-core/types/DatatableColumn";
import { type EmiTypesFakeRow } from "./fakeRows";

// Example A's whole point: one column per emi field type (see fakeRows.ts's own
// doc comment for the type -> field mapping), rendered well enough to prove
// DataGridList/DataGridListHeaderCell handle each of them - not just strings.
export const emiTypesColumns: DatatableColumn[] = [
  {
    name: "uniqueId",
    title: "uniqueId (string)",
    width: 110,
    filterable: true,
    sortable: true,
  },
  {
    name: "title",
    title: "title (string)",
    width: 220,
    filterable: true,
    sortable: true,
  },
  {
    name: "label",
    title: "label (complex: TString)",
    width: 220,
    filterable: true,
    filterType: "tstring",
    getCellValue: (row: EmiTypesFakeRow) => Object.values(row.label).join(" · "),
  },
  {
    name: "age",
    title: "age (int)",
    width: 90,
    filterable: true,
    sortable: true,
    getCellValue: (row: EmiTypesFakeRow) => row.age,
  },
  {
    name: "views",
    title: "views (int64)",
    width: 130,
    filterable: true,
    sortable: true,
    getCellValue: (row: EmiTypesFakeRow) => row.views.toLocaleString(),
  },
  {
    name: "price",
    title: "price (float64)",
    width: 110,
    filterable: true,
    sortable: true,
    getCellValue: (row: EmiTypesFakeRow) => `$${row.price.toFixed(2)}`,
  },
  {
    name: "active",
    title: "active (bool)",
    width: 90,
    filterable: true,
    sortable: true,
    getCellValue: (row: EmiTypesFakeRow) => (row.active ? "✓" : "—"),
  },
  {
    name: "createdAt",
    title: "createdAt (complex: PlainTime)",
    width: 170,
    filterable: true,
    filterType: "date",
    sortable: true,
    getCellValue: (row: EmiTypesFakeRow) =>
      new Date(row.createdAt).toLocaleDateString(),
  },
  {
    name: "status",
    title: "status (enum)",
    width: 110,
    filterable: true,
    sortable: true,
    getCellValue: (row: EmiTypesFakeRow) => (
      <span className={`emi-demo-badge emi-demo-badge--${row.status}`}>
        {row.status}
      </span>
    ),
  },
  {
    name: "tags",
    title: "tags (slice/array<string>)",
    width: 220,
    getCellValue: (row: EmiTypesFakeRow) => row.tags.join(", "),
  },
  {
    name: "owner",
    title: "owner (one)",
    width: 160,
    getCellValue: (row: EmiTypesFakeRow) => row.owner.name,
  },
  {
    name: "collaborators",
    title: "collaborators (collection)",
    width: 220,
    getCellValue: (row: EmiTypesFakeRow) =>
      row.collaborators.length
        ? row.collaborators.map((c) => c.name).join(", ")
        : "—",
  },
  {
    name: "metadata",
    title: "metadata (object/map)",
    width: 220,
    getCellValue: (row: EmiTypesFakeRow) => JSON.stringify(row.metadata),
  },
  {
    name: "note",
    title: "note (any)",
    width: 160,
    getCellValue: (row: EmiTypesFakeRow) =>
      row.note === null || row.note === undefined ? "—" : String(row.note),
  },
];

// A leaner subset used by the two paging-behavior examples (B/C) - the point
// there is scroll behavior, not column coverage.
export const compactColumns: DatatableColumn[] = [
  { name: "uniqueId", title: "Id", width: 90, sortable: true },
  { name: "title", title: "Title", width: 260, filterable: true, sortable: true },
  {
    name: "status",
    title: "Status",
    width: 120,
    filterable: true,
    sortable: true,
    getCellValue: (row: EmiTypesFakeRow) => row.status,
  },
  {
    name: "price",
    title: "Price",
    width: 100,
    sortable: true,
    getCellValue: (row: EmiTypesFakeRow) => `$${row.price.toFixed(2)}`,
  },
];
