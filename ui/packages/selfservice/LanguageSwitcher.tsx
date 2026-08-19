import { useLocale } from "@fireback/ui-core/hooks/useLocale";
import { useS } from "@fireback/ui-core/hooks/useS";
import { interfaceLanguages } from "./personal-settings/Langugages";
import { strings as personalSettingsStrings } from "./personal-settings/strings/translations";

/**
 * A small, always-visible language picker for screens an anonymous visitor
 * can land on before ever signing in (Welcome.screen.tsx in particular) -
 * the only other place to change interface language is
 * personal-settings/InterfaceSettings.tsx, which sits behind auth and is
 * useless to someone who hasn't gotten past the welcome screen yet.
 *
 * Reuses the same interfaceLanguages() option list InterfaceSettings.tsx
 * builds its FormSelect from, and the same setLocale() that screen calls on
 * submit - but applies on change directly (no separate Apply step): this is
 * a lightweight landing-page control, not a settings form, and picking a
 * language is the entire point of interacting with it.
 */
export const LanguageSwitcher = () => {
  const { locale, setLocale } = useLocale();
  const s = useS(personalSettingsStrings);
  const languages = interfaceLanguages(s);

  // Nothing to switch between - a build locked to one language (or a single
  // VITE_SUPPORTED_LANGUAGES entry) doesn't need this control at all.
  if (languages.length <= 1) {
    return null;
  }

  return (
    <div className="language-switcher">
      <select
        className="form-select form-select-sm"
        aria-label={s.interfaceLang.label}
        value={locale}
        onChange={(e) => setLocale(e.target.value)}
      >
        {languages.map((language) => (
          <option key={language.value} value={language.value}>
            {language.label}
          </option>
        ))}
      </select>
    </div>
  );
};
