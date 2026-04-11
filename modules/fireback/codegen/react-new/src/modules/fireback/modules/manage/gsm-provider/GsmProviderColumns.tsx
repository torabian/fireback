import { GsmProviderEntity } from "@/modules/fireback/sdk/modules/abac/GsmProviderEntity";
import { useS } from "@/modules/fireback/hooks/useS";
import { strings } from "./strings/translations";
export const columns = (t: typeof strings) => [
  {
    name: "uniqueId",
    title: "uniqueId",
    width: 200,
  },
  {
    name: GsmProviderEntity.Fields.apiKey,
    title: t.gsmProviders.apiKey,
    width: 100,
  },
  {
    name: GsmProviderEntity.Fields.mainSenderNumber,
    title: t.gsmProviders.mainSenderNumber,
    width: 100,
  },
  {
    name: GsmProviderEntity.Fields.type,
    title: t.gsmProviders.type,
    width: 100,
  },
  {
    name: GsmProviderEntity.Fields.invokeUrl,
    title: t.gsmProviders.invokeUrl,
    width: 100,
  },
  {
    name: GsmProviderEntity.Fields.invokeBody,
    title: t.gsmProviders.invokeBody,
    width: 100,
  },
];
