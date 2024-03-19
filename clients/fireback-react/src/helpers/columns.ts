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
