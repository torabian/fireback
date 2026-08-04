import { PublicJoinKeyEntity } from "@/modules/fireback/sdk/modules/abac/PublicJoinKeyEntity";
import { strings } from "./strings/translations";

export const columns = (s: typeof strings) => [
  {
    name: "uniqueId",
    title: s.uniqueId,
    width: 200,
  },
  {
    name: "role",
    title: s.roleName,
    width: 200,
    getCellValue: (entity: PublicJoinKeyEntity) => entity.role?.name,
  },
];
