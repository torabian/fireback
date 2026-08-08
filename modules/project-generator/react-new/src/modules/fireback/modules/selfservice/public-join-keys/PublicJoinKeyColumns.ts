import { PublicJoinKeyDto } from "@/modules/fireback/sdk/abac/PublicJoinKeyDto";
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
    getCellValue: (entity: PublicJoinKeyDto) => entity.role?.name,
  },
];
