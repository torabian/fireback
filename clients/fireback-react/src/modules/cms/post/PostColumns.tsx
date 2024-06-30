import { enTranslations } from "@/translations/en";
import { PostEntity } from "src/sdk/fireback/modules/cms/PostEntity";
export const columns = (t: typeof enTranslations) => [
  {
    name: PostEntity.Fields.uniqueId,
    title: t.table.uniqueId,
    width: 200,
  },
  {
    name: PostEntity.Fields.title,
    title: t.posts.title,
    width: 100,
  },    
  {
    name: PostEntity.Fields.content,
    title: t.posts.content,
    width: 100,
  },    
  {
    name: PostEntity.Fields.category,
    title: t.posts.category,
    width: 100,
  },    
  {
    name: PostEntity.Fields.tags,
    title: t.posts.tags,
    width: 100,
  },    
];