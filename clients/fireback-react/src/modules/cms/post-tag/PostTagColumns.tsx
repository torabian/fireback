import { enTranslations } from "@/translations/en";
import { PostTagEntity } from "src/sdk/fireback/modules/cms/PostTagEntity";
export const columns = (t: typeof enTranslations) => [
  {
    name: PostTagEntity.Fields.uniqueId,
    title: t.table.uniqueId,
    width: 200,
  },
  {
    name: PostTagEntity.Fields.name,
    title: t.posttags.name,
    width: 100,
  },    
];