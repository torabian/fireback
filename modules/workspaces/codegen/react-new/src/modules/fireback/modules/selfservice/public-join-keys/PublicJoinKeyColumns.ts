import { PublicJoinKeyEntity } from "@/modules/fireback/sdk/modules/abac/PublicJoinKeyEntity";
import { enTranslations } from "@/modules/fireback/translations/en";

export const columns = (t: typeof enTranslations) => [
  {
    name: "uniqueId",
    title: t.table.uniqueId,
    width: 200,
  },
  {
    name: "role",
    title: t.table.uniqueId,
    width: 200,
    getCellValue: (entity: PublicJoinKeyEntity) => entity.role?.name,
  },
];
