import { FileEntity } from "@/sdk/fireback/modules/drive/FileEntity";
import { enTranslations } from "@/translations/en";

export const columns = (t: typeof enTranslations) => [
  {
    name: FileEntity.Fields.uniqueId,
    title: t.table.uniqueId,
    width: 200,
  },
  {
    name: FileEntity.Fields.name,
    title: t.drive.title,
    width: 200,
  },
  {
    name: FileEntity.Fields.size,
    title: t.drive.size,
    width: 100,
  },
  {
    name: FileEntity.Fields.virtualPath,
    title: t.drive.virtualPath,
    width: 100,
  },
  {
    name: FileEntity.Fields.type,
    title: t.drive.type,
    width: 100,
  },
];
