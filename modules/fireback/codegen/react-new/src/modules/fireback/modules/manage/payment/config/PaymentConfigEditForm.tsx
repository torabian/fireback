import { FormText } from "@/modules/fireback/components/forms/form-text/FormText";
import { EntityFormProps } from "@/modules/fireback/definitions/definitions";
import { useS } from "@/modules/fireback/hooks/useS";
import { RemoteQueryContext } from "@/modules/fireback/sdk/core/react-tools";
import { PaymentConfigEntity } from "@/modules/fireback/sdk/modules/payment/PaymentConfigEntity";
import { useContext } from "react";
import { strings } from "./strings/translations";
import { FormCheckbox } from "@/modules/fireback/components/forms/form-switch/FormSwitch";
import { PageSection } from "@/modules/fireback/components/page-section/PageSection";
export const PaymentConfigForm = ({
  form,
  isEditing,
}: EntityFormProps<PaymentConfigEntity>) => {
  const { options } = useContext(RemoteQueryContext);
  const { values, setValues, setFieldValue, errors } = form;
  const s = useS(strings);
  return (
    <>
      <PageSection title={"Stripe configuration"}>
        <FormCheckbox
          value={values.enableStripe}
          onChange={(value) =>
            setFieldValue(PaymentConfigEntity.Fields.enableStripe, value, false)
          }
          errorMessage={errors.enableStripe}
          label={s.paymentConfigs.enableStripe}
          hint={s.paymentConfigs.enableStripeHint}
        />

        <FormText
          disabled={!values.enableStripe}
          value={values.stripeSecretKey}
          onChange={(value) =>
            setFieldValue(
              PaymentConfigEntity.Fields.stripeSecretKey,
              value,
              false
            )
          }
          errorMessage={errors.stripeSecretKey}
          label={s.paymentConfigs.stripeSecretKey}
          hint={s.paymentConfigs.stripeSecretKeyHint}
        />
        <FormText
          disabled={!values.enableStripe}
          value={values.stripeCallbackUrl}
          onChange={(value) =>
            setFieldValue(
              PaymentConfigEntity.Fields.stripeCallbackUrl,
              value,
              false
            )
          }
          errorMessage={errors.stripeCallbackUrl}
          label={s.paymentConfigs.stripeCallbackUrl}
          hint={s.paymentConfigs.stripeCallbackUrlHint}
        />
      </PageSection>
    </>
  );
};
