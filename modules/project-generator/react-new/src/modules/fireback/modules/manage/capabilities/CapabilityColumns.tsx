import { CapabilityDto } from "../../../sdk/abac/CapabilityDto";
import { useS } from "../../../hooks/useS";
import { strings } from "./strings/translations";
export const columns = (t: typeof strings) => [
  {
    name: "uniqueId",
    title: t.capabilities.uniqueId,
    width: 200,
  },
  {
    name: CapabilityDto.Fields.name,
    title: t.capabilities.name,
    width: 100,
  },
  {
    name: CapabilityDto.Fields.description,
    title: t.capabilities.description,
    width: 100,
  },
];
