import { enTranslations } from "@/translations/en";

export function withEssentials(
  t: typeof enTranslations,
  items: Array<any>
): Array<any> {
  return [
    {
      name: "uniqueId",
      title: t.table.uniqueId,
      width: 200,
    },
    ...items,
    {
      name: "createdFormatted",
      title: t.table.created,
      width: 200,
    },
    {
      name: "updatedFormatted",
      title: t.table.updated,
      width: 200,
    },
  ];
}
export function withExtended(
  t: typeof enTranslations,
  items: Array<any>
): Array<any> {
  return [
    {
      name: "workspaceId",
      title: t.table.workspaceId,
      width: 200,
    },
    {
      name: "userId",
      title: t.table.userId,
      width: 200,
    },

    ...withEssentials(t, items),
  ];
}
