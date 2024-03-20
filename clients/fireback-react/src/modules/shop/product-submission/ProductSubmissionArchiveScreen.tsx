import { useT } from "@/hooks/useT";
import { CommonArchiveManager } from "@/components/entity-manager/CommonArchiveManager";
import { ProductSubmissionList } from "./ProductSubmissionList";
import { ProductSubmissionEntity } from "src/sdk/fireback/modules/shop/ProductSubmissionEntity";
export const ProductSubmissionArchiveScreen = () => {
  const t = useT();
  return (
    <CommonArchiveManager
      pageTitle={t.productsubmissions.archiveTitle}
      newEntityHandler={({ locale, router }) => {
        router.push(ProductSubmissionEntity.Navigation.create(locale));
      }}
    >
      <ProductSubmissionList />
    </CommonArchiveManager>
  );
};