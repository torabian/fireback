import { FormCheckbox } from "../../../../fireback-ui/components/forms/form-switch/FormSwitch";
import Link from "../../../../fireback-ui/components/link/Link";
import { PageSection } from "../../../../fireback-ui/components/page-section/PageSection";
import { useT } from "../../../../fireback-ui/hooks/useT";
import { useS } from "../../../../fireback-ui/hooks/useS";
import { useUserWorkspaceBrowseActionQuery } from "../../../sdk/abac/UserWorkspaceBrowseAction";
import { useContext, useState } from "react";
import {
  RemoteQueryContext as FirebackContext,
  RemoteQueryContext,
} from "../../../sdk/core/react-tools";
import { strings } from "./strings/translations";

function UserRoleWorkspaceDebug() {
  const queryWorkspaces = useUserWorkspaceBrowseActionQuery({});
  const s = useS(strings);

  return (
    <>
      <h2>{s.debugger.userRoleWorkspaces}</h2>
      <p>{s.debugger.data}</p>
      <pre>{JSON.stringify(queryWorkspaces.data, null, 2)}</pre>
      <p>{s.debugger.error}</p>
      <pre>{JSON.stringify(queryWorkspaces.error, null, 2)}</pre>
    </>
  );
}

export function SessionDebug() {
  const fireback = useContext(RemoteQueryContext);
  const s = useS(strings);

  return (
    <>
      <h2>{s.debugger.firebackContext}</h2>
      <pre>{JSON.stringify(fireback, null, 2)}</pre>
    </>
  );
}

export function DebuggerSettings({}: {}) {
  const [debugVisible, setDebugVisible] = useState(false);
  const firebackContext = useContext(FirebackContext);

  const t = useT();
  const s = useS(strings);

  return (
    <PageSection title={t.generalSettings.debugSettings.title}>
      <p>{t.generalSettings.debugSettings.description}</p>

      <FormCheckbox
        value={debugVisible}
        label={t.debugInfo}
        onChange={() => setDebugVisible((m) => !m)}
      />
      {debugVisible && (
        <>
          <pre></pre>
          <Link href="/lalaland">{s.debugger.goToLalaland}</Link>
          <Link href="/view3d">{s.debugger.view3d}</Link>
          <UserRoleWorkspaceDebug />
          <SessionDebug />
        </>
      )}
    </PageSection>
  );
}
