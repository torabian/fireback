import { PassportMethodEntity } from "@/modules/fireback/sdk/modules/workspaces/PassportMethodEntity";
import { useS } from "@/modules/fireback/hooks/useS";
import { strings } from "./strings/translations";
export const columns = (t: typeof strings) => [
  {
    name: "uniqueId",
    title: "uniqueId",
    width: 200,
  },
  {
    name: PassportMethodEntity.Fields.type,
    title: t.passportMethods.type,
    width: 100,
  },
  {
    name: PassportMethodEntity.Fields.region,
    title: t.passportMethods.region,
    width: 100,
  },
];
