import { enTranslations } from "@/translations/en";
import { PageEntity } from "src/sdk/fireback/modules/cms/PageEntity";
export const columns = (t: typeof enTranslations) => [
  {
    name: PageEntity.Fields.uniqueId,
    title: t.table.uniqueId,
    width: 200,
  },
  {
    name: PageEntity.Fields.title,
    title: t.pages.title,
    width: 100,
  },
  {
    name: "contentExcerpt",
    title: t.pages.content,
    width: 100,
  },
  {
    name: PageEntity.Fields.category,
    title: t.pages.category,
    width: 100,
  },
  {
    name: PageEntity.Fields.tags,
    title: t.pages.tags,
    width: 100,
  },
];
