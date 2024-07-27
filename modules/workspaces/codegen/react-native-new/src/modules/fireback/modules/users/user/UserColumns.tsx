import { enTranslations } from "@/translations/en";
import { UserEntity } from "src/sdk/fireback/modules/workspaces/UserEntity";
export const columns = (t: typeof enTranslations) => [
  {
    name: UserEntity.Fields.uniqueId,
    title: t.table.uniqueId,
    width: 200,
  },
  {
    name: UserEntity.Fields.person$,
    title: t.users.person,
    getCellValue: (entity: UserEntity) => entity.uniqueId
    width: 100,
  },
  {
    name: UserEntity.Fields.avatar,
    title: t.users.avatar,
    width: 100,
  },
];
