import { CapabilityEntity } from "@/modules/fireback/sdk/modules/workspaces/CapabilityEntity";
import { useS } from "@/modules/fireback/hooks/useS";
import { strings } from "./strings/translations";
export const columns = (t: typeof strings) => [
  {
    name: "uniqueId",
    title: "uniqueId",
    width: 200,
  },
  {
    name: CapabilityEntity.Fields.name,
    title: t.capabilities.name,
    width: 100,
  },
  {
    name: CapabilityEntity.Fields.description,
    title: t.capabilities.description,
    width: 100,
  },
];
