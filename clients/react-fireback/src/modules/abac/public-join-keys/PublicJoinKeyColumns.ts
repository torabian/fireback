import { enTranslations } from "@/translations/en";
import { PublicJoinKeyEntity } from "src/sdk/fireback";

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
