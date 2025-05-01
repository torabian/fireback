import { useT } from "@/hooks/useT";

import { CommonArchiveManager } from "@/components/entity-manager/CommonArchiveManager";
import { {{ .Template }}List } from "./{{ .Template }}List";
import { {{ .Template }}Entity } from "src/sdk/{{ .SdkDir }}/modules/{{ .ModuleDir }}/{{ .Template}}Entity";

export const {{ .Template }}ArchiveScreen = () => {
  const t = useT();

  return (
    <CommonArchiveManager
      pageTitle={t.{{ .templates }}.archiveTitle}
      newEntityHandler={({ locale, router }) => {
        router.push({{ .Template }}Entity.Navigation.create());
      {{ "}}" }}
    >
      <{{ .Template }}List />
    </CommonArchiveManager>
  );
};
