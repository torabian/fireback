import { useS } from "{{ .FirebackUiDir }}/hooks/useS";
import { strings } from "./strings/translations";

import { CommonArchiveManager } from "{{ .FirebackUiDir }}/components/entity-manager/CommonArchiveManager";
import { {{ .Template }}List } from "./{{ .Template }}List";
import { {{ .Template }}Entity } from "{{ .SdkDir }}/modules/{{ .ModuleDir }}/{{ .Template}}Entity";

export const {{ .Template }}ArchiveScreen = () => {
  const s = useS(strings);

  return (
    <CommonArchiveManager
      pageTitle={s.{{ .templates }}.archiveTitle}
      newEntityHandler={({ locale, router }) => {
        router.push({{ .Template }}Entity.Navigation.create(locale));
      {{ "}}" }}
    >
      <{{ .Template }}List />
    </CommonArchiveManager>
  );
};
