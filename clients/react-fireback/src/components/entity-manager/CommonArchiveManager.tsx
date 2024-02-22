import { usePageTitle } from "@/components/page-title/PageTitle";
import { KeyboardAction } from "@/definitions/definitions";
import { useExportTools } from "@/hooks/useExportTools";
import { useLocale } from "@/hooks/useLocale";
import { useRouter } from "@/Router";
import { useNewAction } from "../action-menu/ActionMenu";

export const CommonArchiveManager = ({
  children,
  newEntityHandler,
  exportPath,
  pageTitle,
}: {
  pageTitle: string;
  exportPath?: string;
  newEntityHandler?: (data: { locale: string; router: any }) => void;
  children?: React.ReactNode;
}) => {
  usePageTitle(pageTitle);

  const router = useRouter();
  const { locale } = useLocale();

  useExportTools({ path: exportPath || "" });

  useNewAction(
    newEntityHandler ? () => newEntityHandler({ locale, router }) : undefined,
    KeyboardAction.NewEntity
  );

  return <>{children}</>;
};
