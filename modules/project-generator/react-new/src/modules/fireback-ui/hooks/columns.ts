import { strings } from "../components/strings/translations";

export function withEssentials(
  s: typeof strings,
  items: Array<any>
): Array<any> {
  return [
    {
      name: "uniqueId",
      title: s.table.uniqueId,
      width: 200,
    },
    ...items,
    {
      name: "createdFormatted",
      title: s.table.created,
      width: 200,
    },
    {
      name: "updatedFormatted",
      title: s.table.updated,
      width: 200,
    },
  ];
}
export function withExtended(
  s: typeof strings,
  items: Array<any>
): Array<any> {
  return [
    {
      name: "workspaceId",
      title: s.table.workspaceId,
      width: 200,
    },
    {
      name: "userId",
      title: s.table.userId,
      width: 200,
    },

    ...withEssentials(s, items),
  ];
}
