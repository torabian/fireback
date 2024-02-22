const translationDirectory = "./src/translations";
import { execSync } from "child_process";
import { writeFileSync } from "fs";
import path from "path";

import { enTranslations } from "../../src/translations/en";
import { faTranslations } from "../../src/translations/fa";
import { plTranslations } from "../../src/translations/pl";
import { uaTranslations } from "../../src/translations/ua";
import { ruTranslations } from "../../src/translations/ru";

const enCompiled = { ...enTranslations };
const faCompiled = { ...faTranslations };
const uaCompiled = { ...uaTranslations };
const ruCompiled = { ...ruTranslations };
const plCompiled = { ...plTranslations };

function wrapInTypescript(lang: string, content: any) {
  return `
 /**
 * Auto generated from en-translation with 'translation-manager'
 * You CAN edit this file, it would be read again and compared, but remember it would be sorted automatically
 */
export const ${lang}Translations = ${JSON.stringify(content, null, 2)}
    `;
}

function formatTranslations() {
  const dest = path.join("..", "..", translationDirectory);
  execSync(
    `./node_modules/.bin/prettier --print-width 1000 --write "${dest}/**/*.ts"`
  );
}

function writeLanguage(lang: string, content: any) {
  const dest = path.join("..", "..", translationDirectory, lang + ".ts");
  writeFileSync(dest, wrapInTypescript(lang, content));
  formatTranslations();
}

function syncWithPrimary(primary: any, secondary: any) {
  const keys = Object.keys(primary);

  for (const key of keys) {
    if (typeof primary[key] === "object") {
      secondary = {
        ...secondary,
        [key]: syncWithPrimary(primary[key], secondary[key] || {}),
      };
    } else if (secondary[key] === undefined) {
      secondary[key] = primary[key];
    }
  }

  const keysAfter = Object.keys(secondary).sort(function (a, b) {
    if (a < b) {
      return -1;
    }
    if (a > b) {
      return 1;
    }
    return 0;
  });

  const sorted: any = {};
  for (const key of keysAfter) {
    sorted[key] = secondary[key];
  }
  return sorted;
}

function syncLang(lang: string, content: any) {
  const synced = syncWithPrimary(enCompiled, content);
  writeLanguage(lang, synced);
}

syncLang("fa", faCompiled);
syncLang("pl", plCompiled);
syncLang("ua", uaCompiled);
syncLang("ru", ruCompiled);
syncLang("en", enCompiled);
