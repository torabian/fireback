import { useT } from "@/hooks/useT";

import { CommonArchiveManager } from "@/components/entity-manager/CommonArchiveManager";
import { TemplateList } from "./TemplateList";
import { TemplateNavigationTools } from "src/sdk/{{ .SdkDir }}/modules/{{ .ModuleDir }}/xnavigation";

export const TemplateArchiveScreen = () => {
  const t = useT();

  return (
    <CommonArchiveManager
      pageTitle={t.templates.archiveTitle}
      newEntityHandler={({ locale, router }) => {
        router.push(TemplateNavigationTools.create(locale));
      }}
    >
      <TemplateList />
    </CommonArchiveManager>
  );
};
