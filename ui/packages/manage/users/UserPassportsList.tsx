import { useState } from "react";
import { PageSection } from "@fireback/ui-core/components/page-section/PageSection";
import { useS } from "@fireback/ui-core/hooks/useS";
import { httpErrorHanlder } from "@fireback/ui-core/hooks/api";
import { Toast } from "@fireback/ui-core/hooks/toast";
import { useOverlay } from "@fireback/ui-core/components/overlay/OverlayProvider";
import { commonDialogs } from "@fireback/ui-core/components/overlay/CommonOverlays";
import { strings as coreStrings } from "@fireback/ui-core/components/strings/translations";
import { FormText } from "@fireback/ui-core/components/forms/form-text/FormText";
import { PassportDto } from "@fireback/manage/sdk/abac/PassportDto";
import { usePassportBrowseActionQuery } from "@fireback/manage/sdk/abac/PassportBrowseAction";
import {
  CreatePassportForUserActionReq,
  useCreatePassportForUserAction,
} from "@fireback/manage/sdk/abac/CreatePassportForUserAction";
import {
  SetPassportPasswordActionReq,
  useSetPassportPasswordAction,
} from "@fireback/manage/sdk/abac/SetPassportPasswordAction";
import {
  DisablePassportTotpActionReq,
  useDisablePassportTotpAction,
} from "@fireback/manage/sdk/abac/DisablePassportTotpAction";
import {
  SendPassportResetEmailActionReq,
  useSendPassportResetEmailAction,
} from "@fireback/manage/sdk/abac/SendPassportResetEmailAction";
import { strings } from "./strings/translations";

export const UserPassportList = ({ userId }: { userId: string }) => {
  const { data, refetch } = usePassportBrowseActionQuery({
    qs: userId
      ? new URLSearchParams({ query: "user_id = " + userId })
      : undefined,
  });
  const items = (data as any)?.data?.items as PassportDto[] | undefined;
  const s = useS(strings);
  const cs = useS(coreStrings);
  const { openDrawer } = useOverlay();
  const createPassportMutation = useCreatePassportForUserAction({});

  // Root-only (CreatePassportForUserAction - see modules/abac/
  // CreatePassportForUserActionImplementation.go): lets a manager provision a working
  // email/phone login for a user directly, password included, without the user having
  // to complete signup/email-confirmation themselves - handy when there's no mail
  // server configured yet, or the manager just wants to hand someone credentials.
  const addPassport = () => {
    openDrawer<AddPassportDrawerResult | undefined>((props) => (
      <AddPassportDrawer {...props} s={s} cs={cs} />
    ))
      .promise.then(({ type, data: form }) => {
        if (type !== "resolved" || !form) {
          return;
        }
        return createPassportMutation
          .mutateAsync(
            new CreatePassportForUserActionReq({
              userId,
              type: form.type,
              value: form.value,
              password: form.password,
            }),
          )
          .then(() => {
            Toast(s.passportCreated, { type: "success" });
            refetch();
          });
      })
      .catch((err) => httpErrorHanlder(err, cs));
  };

  return (
    <div>
      <PageSection title={s.passports}>
        <div className="mb-3">
          <button
            className="btn btn-sm btn-outline-primary"
            onClick={addPassport}
          >
            {s.addPassport}
          </button>
        </div>
        {(items || []).map((item) => {
          return (
            <UserPassportItem
              passport={item}
              key={item.uniqueId}
              s={s}
              onChanged={refetch}
            />
          );
        })}
      </PageSection>
    </div>
  );
};

type AddPassportDrawerResult = {
  type: "email" | "phone";
  value: string;
  password: string;
};

// Mirrors SetPasswordDrawer below - a minimal, non-formik drawer form (plain useState,
// resolved value read by the opener) since this is a one-off admin action, not a
// regular entity edit form.
const AddPassportDrawer = ({
  close,
  resolve,
  s,
  cs,
}: {
  close: () => void;
  resolve: (result?: AddPassportDrawerResult) => void;
  s: typeof strings;
  cs: typeof coreStrings;
}) => {
  const [type, setType] = useState<"email" | "phone">("email");
  const [value, setValue] = useState("");
  const [password, setPassword] = useState("");

  // Same minimum-length rule the backend enforces (see
  // CreatePassportForUserActionImplementation.go) - caught here too so the request
  // never round-trips just to be rejected for a password that's obviously too short.
  const canSubmit = value.trim().length > 0 && password.length >= 6;

  return (
    <div className="confirm-drawer-container p-3">
      <h2>{s.addPassport}</h2>
      <div className="mb-3">
        <label className="form-label">{s.passportTypeLabel}</label>
        <select
          className="form-select"
          value={type}
          onChange={(e) => setType(e.target.value as "email" | "phone")}
        >
          <option value="email">{s.email}</option>
          <option value="phone">{s.phone}</option>
        </select>
      </div>
      <FormText
        value={value}
        onChange={setValue}
        autoFocus
        label={type === "email" ? s.emailAddress : s.phoneNumberValue}
      />
      <FormText
        type="password"
        value={password}
        onChange={setPassword}
        label={s.newPassword}
      />
      <div>
        <button
          className="d-block w-100 btn btn-primary"
          disabled={!canSubmit}
          onClick={() => resolve({ type, value, password })}
        >
          {cs.common.save}
        </button>
        <button className="d-block w-100 btn" onClick={() => close()}>
          {cs.common.cancel}
        </button>
      </div>
    </div>
  );
};

function booleanToHuman(value: boolean | undefined, s: typeof strings): string {
  if (value === null || value === undefined) {
    return s.na;
  }

  if (value === true) {
    return s.yes;
  }

  if (value === false) {
    return s.no;
  }
}

// A minimal drawer form asking for a single new password - used by "Set password"
// below. Renders as an openDrawer render-prop component (close/resolve, same shape
// CommonOverlays.tsx's confirmDrawer/confirmModal use), resolving with the typed
// password on submit.
const SetPasswordDrawer = ({
  close,
  resolve,
  s,
  cs,
}: {
  close: () => void;
  resolve: (password?: string) => void;
  s: typeof strings;
  cs: typeof coreStrings;
}) => {
  const [password, setPassword] = useState("");

  return (
    <div className="confirm-drawer-container p-3">
      <h2>{s.setPassword}</h2>
      <FormText
        type="password"
        value={password}
        onChange={setPassword}
        autoFocus
        label={s.newPassword}
      />
      <div>
        <button
          className="d-block w-100 btn btn-primary"
          disabled={password.length < 6}
          onClick={() => resolve(password)}
        >
          {cs.common.save}
        </button>
        <button className="d-block w-100 btn" onClick={() => close()}>
          {cs.common.cancel}
        </button>
      </div>
    </div>
  );
};

const UserPassportItem = ({
  passport,
  s,
  onChanged,
}: {
  passport: PassportDto;
  s: typeof strings;
  onChanged: () => void;
}) => {
  const cs = useS(coreStrings);
  const { openDrawer } = useOverlay();
  const { confirmDrawer } = commonDialogs();

  const setPasswordMutation = useSetPassportPasswordAction({});
  const disableTotpMutation = useDisablePassportTotpAction({});
  const sendResetEmailMutation = useSendPassportResetEmailAction({});

  const setPassword = () => {
    openDrawer<string | undefined>((props) => (
      <SetPasswordDrawer {...props} s={s} cs={cs} />
    ))
      .promise.then(({ type, data: password }) => {
        if (type !== "resolved" || !password) {
          return;
        }
        return setPasswordMutation
          .mutateAsync(
            new SetPassportPasswordActionReq({
              uniqueId: passport.uniqueId,
              password,
            }),
          )
          .then(() => {
            Toast(s.passwordUpdated, { type: "success" });
          });
      })
      .catch((err) => httpErrorHanlder(err, cs));
  };

  const disableTotp = () => {
    confirmDrawer({
      title: s.disableTotp,
      description: s.disableTotpConfirm,
      confirmLabel: cs.common.yes,
      cancelLabel: cs.common.no,
    })
      .promise.then(({ type }) => {
        if (type !== "resolved") {
          return;
        }
        return disableTotpMutation
          .mutateAsync(
            new DisablePassportTotpActionReq({ uniqueId: passport.uniqueId }),
          )
          .then(() => {
            Toast(s.totpDisabled, { type: "success" });
            onChanged();
          });
      })
      .catch((err) => httpErrorHanlder(err, cs));
  };

  const sendResetEmail = () => {
    confirmDrawer({
      title: s.sendResetEmail,
      description: s.sendResetEmailConfirm.replace("{value}", passport.value),
      confirmLabel: cs.common.yes,
      cancelLabel: cs.common.no,
    })
      .promise.then(({ type }) => {
        if (type !== "resolved") {
          return;
        }
        return sendResetEmailMutation
          .mutateAsync(
            new SendPassportResetEmailActionReq({
              uniqueId: passport.uniqueId,
            }),
          )
          .then(() => {
            Toast(s.resetEmailSent, { type: "success" });
          });
      })
      .catch((err) => httpErrorHanlder(err, cs));
  };

  return (
    <div>
      <div className="general-entity-view ">
        <div className="entity-view-row entity-view-head">
          <div className="field-info">{s.value}</div>
          <div className="field-value">{passport.value}</div>
        </div>
        <div className="entity-view-row entity-view-head">
          <div className="field-info">{s.type}</div>
          <div className="field-value">{passport.type}</div>
        </div>
        <div className="entity-view-row entity-view-head">
          <div className="field-info">{s.confirmed}</div>
          <div className="field-value">
            {booleanToHuman(passport.confirmed, s)}
          </div>
        </div>
      </div>
      <div className="row mt-2">
        <div className="col-auto">
          <button
            className="btn btn-sm btn-outline-primary"
            onClick={setPassword}
          >
            {s.setPassword}
          </button>
        </div>
        {!!passport.totpSecret && (
          <div className="col-auto">
            <button
              className="btn btn-sm btn-outline-danger"
              onClick={disableTotp}
            >
              {s.disableTotp}
            </button>
          </div>
        )}
        <div className="col-auto">
          <button
            className="btn btn-sm btn-outline-secondary"
            onClick={sendResetEmail}
          >
            {s.sendResetEmail}
          </button>
        </div>
      </div>
    </div>
  );
};
