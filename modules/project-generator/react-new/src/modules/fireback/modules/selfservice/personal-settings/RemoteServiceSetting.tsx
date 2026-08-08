import { PageSection } from "../../../../fireback-ui/components/page-section/PageSection";
import { useS } from "../../../../fireback-ui/hooks/useS";
import React, { useCallback, useContext, useEffect, useState } from "react";
import { Formik, type FormikHelpers, type FormikProps } from "formik";
import { useCommonEntityManager } from "../../../../fireback-ui/hooks/useCommonEntityManager";
import { ErrorsView } from "../../../../fireback-ui/components/error-view/ErrorView";
import { FormText } from "../../../../fireback-ui/components/forms/form-text/FormText";
import { FormButton } from "../../../../fireback-ui/components/forms/form-button/FormButton";
import { AppConfigContext } from "../../../../fireback-ui/hooks/appConfigTools";
import { strings } from "./strings/translations";

interface AppRemoteInformation {
  remote: string;
}

function ping(url: string) {
  return fetch(url + "ping").then((response) => response.json());
}

const AppRemoteInformationFields = {
  remote: "remote",
};

const updateSettings = (
  values: Partial<AppRemoteInformation>,
  d: FormikHelpers<Partial<AppRemoteInformation>>
) => {
  if (values.remote) {
    localStorage.setItem("app_remote_address", values.remote);
  }
};

export enum RemoteState {
  Unknown = "unknown",
  Pinging = "pinging",
  Active = "active",
  Inaccessible = "inaccessible",
}

export function RemoteServiceSetting({}: {}) {
  const [remoteState, setRemoteState] = useState<RemoteState>(
    RemoteState.Unknown
  );

  const { config, patchConfig } = useContext(AppConfigContext);

  const s = useS(strings);
  const { router, uniqueId, queryClient, isEditing, locale, formik } =
    useCommonEntityManager<Partial<AppRemoteInformation>>({});

  const pingUrl = useCallback(
    (address: string) => {
      setRemoteState(RemoteState.Pinging);
      ping(address)
        .then(() => {
          setRemoteState(RemoteState.Active);
        })
        .catch(() => {
          setRemoteState(RemoteState.Inaccessible);
        });
    },
    [formik.current]
  );

  const onSubmit = (
    values: Partial<AppRemoteInformation>,
    d: FormikHelpers<Partial<AppRemoteInformation>>
  ) => {
    if (!values.remote) {
      return;
    }

    pingUrl(values.remote);
    patchConfig({ remote: values.remote });
    updateSettings(values, d);
  };

  useEffect(() => {
    if (config.remote) {
      pingUrl(config.remote);
    }
  }, []);

  useEffect(() => {
    formik.current?.setValues({ remote: config.remote });
  }, [config.remote]);

  return (
    <PageSection title={s.remoteTitle}>
      <p>{s.remoteDescription}</p>
      <Formik
        innerRef={(r) => {
          if (r) formik.current = r;
        }}
        initialValues={{}}
        onSubmit={onSubmit}
      >
        {(form: FormikProps<Partial<AppRemoteInformation>>) => (
          <form
            className="remote-service-form"
            onSubmit={(e) => e.preventDefault()}
          >
            <ErrorsView errors={form.errors} />
            <FormText
              dir="ltr"
              label={s.remoteAddress}
              onChange={(value) =>
                form.setFieldValue(
                  AppRemoteInformationFields.remote,
                  value,
                  false
                )
              }
              errorMessage={
                remoteState === RemoteState.Inaccessible
                  ? s.inaccessibleRemote
                  : ""
              }
              validMessage={
                remoteState === RemoteState.Active
                  ? s.serverConnected
                  : ""
              }
              value={form.values.remote}
            ></FormText>
            <FormButton
              disabled={
                form.values.remote === "" ||
                form.values.remote === config.remote
              }
              label={s.apply}
              onClick={() => form.submitForm()}
            />
          </form>
        )}
      </Formik>
    </PageSection>
  );
}
