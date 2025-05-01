import { Route } from "react-router-dom";
import { {{ .Template }}ArchiveScreen } from "./{{ .Template }}ArchiveScreen";
import { {{ .Template }}EntityManager } from "./{{ .Template }}EntityManager";
import { {{ .Template }}SingleScreen } from "./{{ .Template }}SingleScreen";
import { {{ .Template }}Entity } from "{{ .SdkDir }}/modules/{{ .ModuleDir }}/{{ .Template}}Entity";

export function use{{ .Template }}Routes() {
  return (
    <>
      <Route
        element={<{{ .Template }}EntityManager />}
        path={ {{ .Template }}Entity.Navigation.Rcreate}
      />
      <Route
        element={<{{ .Template }}SingleScreen />}
        path={ {{ .Template }}Entity.Navigation.Rsingle}
      ></Route>
      <Route
        element={<{{ .Template }}EntityManager />}
        path={ {{ .Template }}Entity.Navigation.Redit}
      ></Route>
      <Route
        element={<{{ .Template }}ArchiveScreen />}
        path={  {{ .Template }}Entity.Navigation.Rquery}
      ></Route>
    </>
  );
}
