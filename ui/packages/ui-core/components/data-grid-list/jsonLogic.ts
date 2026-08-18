// Frontend counterpart of modules/fireback/jsonsql/JsonLogicToSql.go's
// ConditionToJsonLogic - builds the exact same {"and": [...]} JsonLogic
// (jsonlogic.com) tree shape from a flat list of {field, value} conditions, so a
// column-filter payload built here parses identically on the backend (the Go side
// feeds the same shape into github.com/h22rana/jsonlogic2sql).
//
// Rules mirrored 1:1 from ConditionToJsonLogic:
//  - a blank value is dropped entirely (not sent as an empty "contains")
//  - a value that looks like a JSON object/array ("{" or "[" anywhere in it) is
//    parsed as-is and used directly as the sub-expression - this is the escape
//    hatch for a power-user/advanced filter typing raw JsonLogic instead of a
//    plain contains match (e.g. {"==": [{"var":"status"},"active"]})
//  - anything else becomes {"contains": [{"var": field}, value]}, everything
//    joined with "and"
export interface JsonLogicCondition {
  field: string;
  value: string;
}

export type JsonLogicNode = Record<string, unknown>;

export function buildJsonLogic(conditions: JsonLogicCondition[]): JsonLogicNode {
  const and: unknown[] = [];

  for (const item of conditions) {
    if (!item.value) {
      continue;
    }

    if (item.value.includes("{") || item.value.includes("[")) {
      try {
        and.push(JSON.parse(item.value));
        continue;
      } catch {
        // Fall through to the plain "contains" treatment below - not valid
        // JSON, just a value that happens to contain a brace/bracket.
      }
    }

    and.push({
      contains: [{ var: item.field }, item.value],
    });
  }

  return { and };
}

// True when the tree has no real conditions left ({"and": []}) - the caller's
// signal to omit the filter query param entirely rather than sending a no-op.
export function isEmptyJsonLogic(node: JsonLogicNode): boolean {
  const and = node.and;
  return Array.isArray(and) && and.length === 0;
}

// Builds the sub-expression TStringFilterDrawer's per-locale values turn into: one
// "contains" per locale that actually has search text typed in, OR'd together - "match
// if any of the languages the user typed something into contains it", since a
// complexes.TString column (a locale -> text map, e.g. {"en": "Home", "fa": "خانه"})
// stores each locale as its own JSON key (field.<locale>, e.g. "label.en") rather than
// one flat value a plain "contains" could target.
//
// The caller (DataGridListHeaderCell) JSON.stringifies this and hands it to
// useColumnFilters' setColumnFilter the same way any other column's raw-JsonLogic power
// filter would (see buildJsonLogic's own doc comment on the "{"/"[" escape hatch) -
// there's no separate code path in useColumnFilters/buildJsonLogic for TString columns
// specifically.
export function buildTStringFilterCondition(
  field: string,
  perLocaleValues: Record<string, string>,
): JsonLogicNode | undefined {
  const conditions = Object.entries(perLocaleValues)
    .filter(([, value]) => value.trim() !== "")
    .map(([locale, value]) => ({
      contains: [{ var: `${field}.${locale}` }, value],
    }));

  if (conditions.length === 0) {
    return undefined;
  }
  if (conditions.length === 1) {
    return conditions[0];
  }
  return { or: conditions };
}
