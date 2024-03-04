import { useCommonCrudActions } from "@/components/action-menu/ActionMenu";
import { ErrorsView } from "@/components/error-view/ErrorView";
import { usePageTitle } from "@/components/page-title/PageTitle";
import { useCommonEntityManager } from "@/hooks/useCommonEntityManager";

import { Formik, FormikHelpers, FormikProps } from "formik";

import { WithPermissions } from "@/components/layouts/WithPermissions";
import { useT } from "@/hooks/useT";
import { usePatchNotificationWorkspaceConfig } from "src/sdk/fireback/modules/workspaces/usePatchNotificationWorkspaceConfig";
import { useEffect } from "react";
import { WorkspaceNotificationForm } from "./WorkspaceNotificationForm";
import { NotificationConfigEntity } from "@/sdk/fireback/modules/workspaces/NotificationConfigEntity";
import { useGetNotificationConfigByUniqueId } from "@/sdk/fireback/modules/workspaces/useGetNotificationConfigByUniqueId";
import { WorkspaceEntity } from "@/sdk/fireback/modules/workspaces/WorkspaceEntity";

interface DtoEntity<T> {
  data?: T | null;
}

export const WorkspaceNotificationEntityManager = ({
  data,
}: DtoEntity<NotificationConfigEntity>) => {
  const t = useT();
  const { router, uniqueId, queryClient, isEditing, locale, formik } =
    useCommonEntityManager<Partial<NotificationConfigEntity>>({
      data,
    });

  usePageTitle(t.wokspaces.configurateWorkspaceNotification);

  const { query: getQuery } = useGetNotificationConfigByUniqueId({
    query: { uniqueId: "root", deep: true },
  });

  const { submit: patch, mutation: patchMutation } =
    usePatchNotificationWorkspaceConfig({
      queryClient,
    });

  useEffect(() => {
    if (getQuery.data?.data) {
      formik.current?.setValues(getQuery.data.data);
    }
  }, [getQuery.data]);

  useEffect(() => {
    formik.current?.setSubmitting(patchMutation.isLoading);
  }, [patchMutation.isLoading]);

  const onSubmit = (
    values: Partial<NotificationConfigEntity>,
    d: FormikHelpers<Partial<NotificationConfigEntity>>
  ) => {
    patch(values, d).then((response) => {});
  };

  useCommonCrudActions({
    onCancel() {
      router.push(WorkspaceEntity.Navigation.query(undefined, locale));
    },
    onSave() {
      formik.current?.submitForm();
    },
    access: {
      // permissions: [ROOT_WORKSPACES_CONFIG],
      permissions: [],
      onlyRoot: true,
    },
  });

  return (
    <WithPermissions onlyRoot permissions={[]}>
      <Formik
        innerRef={(r) => {
          if (r) formik.current = r;
        }}
        initialValues={{}}
        onSubmit={onSubmit}
      >
        {(form: FormikProps<Partial<NotificationConfigEntity>>) => (
          <form onSubmit={(e) => e.preventDefault()}>
            <ErrorsView errors={form.errors} />
            <WorkspaceNotificationForm form={form} />
          </form>
        )}
      </Formik>
    </WithPermissions>
  );
};
