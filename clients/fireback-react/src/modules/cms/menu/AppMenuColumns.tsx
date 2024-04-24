import { enTranslations } from "@/translations/en";
import { AppMenuEntity } from "src/sdk/fireback/modules/workspaces/AppMenuEntity";
export const columns = (t: typeof enTranslations) => [
  {
    name: AppMenuEntity.Fields.uniqueId,
    title: t.table.uniqueId,
    width: 200,
  },
  {
    name: AppMenuEntity.Fields.href,
    title: t.appMenus.href,
    width: 100,
  },
  {
    name: AppMenuEntity.Fields.icon,
    title: t.appMenus.icon,
    width: 100,
  },
  {
    name: AppMenuEntity.Fields.label,
    title: t.appMenus.label,
    width: 100,
  },
  {
    name: AppMenuEntity.Fields.activeMatcher,
    title: t.appMenus.activeMatcher,
    width: 100,
  },
  {
    name: AppMenuEntity.Fields.applyType,
    title: t.appMenus.applyType,
    width: 100,
  },
  {
    name: AppMenuEntity.Fields.capability$,
    title: t.appMenus.capability,
    getCellValue: (entity: AppMenuEntity) => entity.uniqueId,
    width: 100,
  },
];
