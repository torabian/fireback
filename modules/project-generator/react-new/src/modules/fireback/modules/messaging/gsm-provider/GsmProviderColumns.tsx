import { GsmProviderDto } from "@/modules/fireback/sdk/messaging/GsmProviderDto";
import { useS } from "@/modules/fireback/hooks/useS";
import { strings } from "./strings/translations";
export const columns = (t: typeof strings) => [
  {
    name: "uniqueId",
    title: "uniqueId",
    width: 200,
  },
  {
    name: GsmProviderDto.Fields.apiKey,
    title: t.gsmProviders.apiKey,
    width: 100,
  },
  {
    name: GsmProviderDto.Fields.mainSenderNumber,
    title: t.gsmProviders.mainSenderNumber,
    width: 100,
  },
  {
    name: GsmProviderDto.Fields.type,
    title: t.gsmProviders.type,
    width: 100,
  },
  {
    name: GsmProviderDto.Fields.invokeUrl,
    title: t.gsmProviders.invokeUrl,
    width: 100,
  },
  {
    name: GsmProviderDto.Fields.invokeBody,
    title: t.gsmProviders.invokeBody,
    width: 100,
  },
];
