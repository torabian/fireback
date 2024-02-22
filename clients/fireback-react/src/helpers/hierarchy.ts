/**
 * We can store the entities in database in a hierarchical order.
 * These functions would help to flatten them
 */

export interface HierarchicalEntity<T> {
  children: T[];
  entity: T;
  uniqueId: string;
}

export interface InlineRelation {
  parentId: string;
  uniqueId: string;
}

export function flattenTheNest(
  items: HierarchicalEntity<any>[],
  relation: InlineRelation[],
  parentId: string
): InlineRelation[] {
  for (const item of items) {
    relation.push({ uniqueId: item.uniqueId, parentId });

    if (item.children) {
      relation = flattenTheNest(item.children, relation, item.uniqueId);
    }
  }

  return relation;
}

export function castEntityToHierarchD<T>(
  flatItems: InlineRelation[],
  parentId?: string | null,
  level = 0
): HierarchicalEntity<T>[] {
  const structure: HierarchicalEntity<T>[] = [];

  for (let i = 0; i < flatItems.length; i++) {
    const curr = flatItems[i];

    if ((curr.parentId || "") == (parentId || "")) {
      let children: any = [];
      flatItems = flatItems.filter((_, m) => m !== i);
      i--;

      children = castEntityToHierarchD(flatItems, curr.uniqueId, ++level);
      structure.push({
        entity: curr as any,
        children,
        uniqueId: curr.uniqueId,
      });
    }
  }

  return structure;
}
