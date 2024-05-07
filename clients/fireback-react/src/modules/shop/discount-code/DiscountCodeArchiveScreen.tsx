import { useT } from "@/fireback/hooks/useT";
import { CommonArchiveManager } from "@/fireback/components/entity-manager/CommonArchiveManager";
import { DiscountCodeList } from "./DiscountCodeList";
import { DiscountCodeEntity } from "src/sdk/fireback/modules/shop/DiscountCodeEntity";
export const DiscountCodeArchiveScreen = () => {
  const t = useT();
  return (
    <CommonArchiveManager
      pageTitle={t.discountCodes.archiveTitle}
      newEntityHandler={({ locale, router }) => {
        router.push(DiscountCodeEntity.Navigation.create(locale));
      }}
    >
      <DiscountCodeList />
    </CommonArchiveManager>
  );
};
