import { useT } from "@/modules/fireback/hooks/useT";

export function MockVersionNotice() {
  const t = useT();
  return <div className="app-mock-version-notice">{t.mockNotice}</div>;
}
