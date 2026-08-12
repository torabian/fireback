import { RegionalContentDto } from "@fireback/manage/sdk/abac/RegionalContentDto";
import { strings } from "./strings/translations";

export const columns = (t: typeof strings) => [
  {
    name: "uniqueId",
    title: t.regionalContents.uniqueId,
    width: 200,
  },
  {
    name: RegionalContentDto.Fields.content,
    title: t.regionalContents.content,
    width: 100,
  },
  {
    name: RegionalContentDto.Fields.region,
    title: t.regionalContents.region,
    width: 100,
  },
  {
    name: RegionalContentDto.Fields.title,
    title: t.regionalContents.title,
    width: 100,
  },
  {
    name: RegionalContentDto.Fields.languageId,
    title: t.regionalContents.languageId,
    width: 100,
  },
  {
    name: RegionalContentDto.Fields.keyGroup,
    title: t.regionalContents.keyGroup,
    width: 100,
  },
];
