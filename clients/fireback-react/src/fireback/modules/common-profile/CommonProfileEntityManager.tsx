import { useCommonEntityManager } from "@/fireback/hooks/useCommonEntityManager";
import { usePatchCommonProfileDistinct } from "src/sdk/fireback/modules/commonprofile/usePatchCommonProfileDistinct";

import { useGetCommonProfileDistinct } from "src/sdk/fireback/modules/commonprofile/useGetCommonProfileDistinct";
import { CommonProfileEditForm } from "./CommonProfileEditForm";
import {
  CommonEntityManager,
  DtoEntity,
} from "@/fireback/components/entity-manager/CommonEntityManager";
import { EmailProviderEntity } from "@/sdk/fireback/modules/workspaces/EmailProviderEntity";

export const CommonProfileEntityManager = ({
  data,
}: DtoEntity<EmailProviderEntity>) => {
  const { t, queryClient } = useCommonEntityManager<
    Partial<EmailProviderEntity>
  >({
    data,
  });

  const getSingleHook = useGetCommonProfileDistinct({
    query: { uniqueId: "self" },
  });

  const patchHook = usePatchCommonProfileDistinct({
    queryClient,
  });

  return (
    <CommonEntityManager
      forceEdit
      getSingleHook={getSingleHook}
      patchHook={patchHook}
      Form={CommonProfileEditForm}
      onEditTitle={t.fb.commonProfile}
      data={data}
    />
  );
};
