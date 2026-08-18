// Cursor encode/decode for DataGridList's infinite scroll.
//
// The backend's own cursor (modules/fireback/CrudCoreActions.go's parseCursor)
// only ever understands a single "field(numericValue)" token today (regex
// `(\w+)\((\d+)\)`, taken as "field > value"). DataGridList's cursor extends that
// same "name(value)" shape but always carries *two* tokens concatenated with "+" -
// the row to resume after, and the sort order that was in effect when the page
// carrying that row was fetched:
//
//   id(482)+sort(createdAt.desc)
//
// Sending the sort alongside the id (not just relying on "whatever sort is
// currently selected") means a page fetched under one sort order can never be
// silently continued under a different one if the user changes sort mid-scroll -
// the request just carries its own resume point end to end.
export interface DataGridListCursor {
  /** The row id to resume after (exclusive). */
  id: string;
  /** "columnName.asc" | "columnName.desc", or undefined for unsorted. */
  sort?: string;
}

const TOKEN_RE = /(\w+)\(([^)]*)\)/g;

export function encodeCursor(cursor: DataGridListCursor): string {
  const parts = [`id(${cursor.id})`];
  if (cursor.sort) {
    parts.push(`sort(${cursor.sort})`);
  }
  return parts.join("+");
}

export function decodeCursor(raw: string | undefined | null): DataGridListCursor | null {
  if (!raw) {
    return null;
  }

  const tokens: Record<string, string> = {};
  let match: RegExpExecArray | null;
  TOKEN_RE.lastIndex = 0;
  while ((match = TOKEN_RE.exec(raw))) {
    tokens[match[1]] = match[2];
  }

  if (!tokens.id) {
    return null;
  }

  return { id: tokens.id, sort: tokens.sort || undefined };
}

// sortingToString/fromString give a single canonical "column.direction" form used
// both in the cursor and as the plain `sort` query param the rest of the app
// already sends (see useDatatableFiltering's toSortString) - DataGridList only
// supports a single active sort column, unlike the older multi-column `Sorting[]`
// devexpress shape.
export interface SingleSort {
  columnName: string;
  direction: "asc" | "desc";
}

export function sortingToString(sort: SingleSort | undefined): string | undefined {
  if (!sort) {
    return undefined;
  }
  return `${sort.columnName}.${sort.direction}`;
}

export function sortingFromString(raw: string | undefined): SingleSort | undefined {
  if (!raw) {
    return undefined;
  }
  const idx = raw.lastIndexOf(".");
  if (idx === -1) {
    return undefined;
  }
  const direction = raw.slice(idx + 1);
  if (direction !== "asc" && direction !== "desc") {
    return undefined;
  }
  return { columnName: raw.slice(0, idx), direction };
}
