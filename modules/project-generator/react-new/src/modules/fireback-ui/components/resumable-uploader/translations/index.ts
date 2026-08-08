import type { UploaderLocale, UploadTranslations } from "../types";
import { en } from "./en";
import { fa } from "./fa";

export const localeTranslations: Record<UploaderLocale, UploadTranslations> = {
  en,
  fa,
};

export type { UploaderLocale };
export { en, fa };
