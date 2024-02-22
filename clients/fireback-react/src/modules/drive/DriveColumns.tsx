import { enTranslations } from "@/translations/en";
import { Column } from "@devexpress/dx-react-grid";
import { DriveEntityFields } from "minifirma-tools/modules/drive/drive-fields";

export const columns = (t: typeof enTranslations) => [
  {
    name: DriveEntityFields.uniqueId,
    title: t.table.uniqueId,
    width: 200,
  },
  {
    name: DriveEntityFields.name,
    title: t.drive.title,
    width: 200,
  },
  {
    name: DriveEntityFields.size,
    title: t.drive.size,
    width: 100,
  },
  {
    name: DriveEntityFields.virtualpath,
    title: t.drive.virtualPath,
    width: 100,
  },
  {
    name: DriveEntityFields.type,
    title: t.drive.type,
    width: 100,
  },
];
