import { RegionalContentEntity } from "@/modules/fireback/sdk/modules/abac/RegionalContentEntity";
import { useS } from "@/modules/fireback/hooks/useS";
import { strings } from "./strings/translations";
export const columns = (t: typeof strings) => [
  {
    name: "uniqueId",
    title: "uniqueId",
    width: 200,
  },
  {
    name: RegionalContentEntity.Fields.content,
    title: t.regionalContents.content,
    width: 100,
  },
  {
    name: RegionalContentEntity.Fields.region,
    title: t.regionalContents.region,
    width: 100,
  },
  {
    name: RegionalContentEntity.Fields.title,
    title: t.regionalContents.title,
    width: 100,
  },
  {
    name: RegionalContentEntity.Fields.languageId,
    title: t.regionalContents.languageId,
    width: 100,
  },
  {
    name: RegionalContentEntity.Fields.keyGroup,
    title: t.regionalContents.keyGroup,
    width: 100,
  },
];
