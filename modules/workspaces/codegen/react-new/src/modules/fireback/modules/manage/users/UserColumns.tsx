import { UserEntity } from "../../../sdk/modules/workspaces/UserEntity";
import { enTranslations } from "../../../translations/en";

export const columns = (t: typeof enTranslations) => [
  {
    name: UserEntity.Fields.uniqueId,
    title: t.table.uniqueId,
    width: 100,
  },
  {
    // name: UserEntity.Fields.person.firstName,
    name: "person.firstName",
    title: t.users.firstName,
    width: 200,
    getCellValue: (e: UserEntity) => e.person?.firstName,
  },

  {
    name: "person.lastName",
    title: t.users.lastName,
    width: 200,
    getCellValue: (e: UserEntity) => e.person?.lastName,
  },
];
