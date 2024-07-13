import { KeyboardAction } from "@/modules/fireback/definitions/definitions";
import { useCommonEntityManager } from "@/modules/fireback/hooks/useCommonEntityManager";
import { useBackButton, useEditAction } from "../action-menu/ActionMenu";
import { QueryErrorView } from "../error-view/QueryError";

export const CommonSingleManager = ({
  children,
  getSingleHook,
  editEntityHandler,
}: {
  getSingleHook?: any;
  children?: React.ReactNode;
  editEntityHandler?: (data: { locale: string; router: any }) => void;
}) => {
  const { router, locale } = useCommonEntityManager<Partial<any>>({});

  useEditAction(
    editEntityHandler ? () => editEntityHandler({ locale, router }) : undefined,
    KeyboardAction.EditEntity
  );

  useBackButton(() => router.goBack(), KeyboardAction.CommonBack);

  return (
    <>
      <QueryErrorView query={getSingleHook.query} />

      {children}
    </>
  );
};
