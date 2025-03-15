import { KeyboardAction } from "../../definitions/definitions";
import { useCommonEntityManager } from "../../hooks/useCommonEntityManager";
import { useBackButton, useEditAction } from "../action-menu/ActionMenu";
import { QueryErrorView } from "../error-view/QueryError";

export const CommonSingleManager = ({
  children,
  getSingleHook,
  editEntityHandler,
  noBack,
  disableOnGetFailed,
}: {
  getSingleHook?: any;
  children?: React.ReactNode;
  editEntityHandler?: (data: { locale: string; router: any }) => void;
  noBack?: boolean;
  disableOnGetFailed?: boolean;
}) => {
  const { router, locale } = useCommonEntityManager<Partial<any>>({});

  useEditAction(
    editEntityHandler ? () => editEntityHandler({ locale, router }) : undefined,
    KeyboardAction.EditEntity
  );

  useBackButton(
    noBack !== true ? () => router.goBack() : null,
    KeyboardAction.CommonBack
  );

  return (
    <>
      <QueryErrorView query={getSingleHook.query} />

      {disableOnGetFailed === true && getSingleHook?.query?.isError ? null : (
        <>{children}</>
      )}
    </>
  );
};
