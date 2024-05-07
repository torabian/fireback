import { useT } from "@/fireback/hooks/useT";
import { CommonArchiveManager } from "@/fireback/components/entity-manager/CommonArchiveManager";
import { TagList } from "./TagList";
import { TagEntity } from "src/sdk/fireback/modules/shop/TagEntity";
export const TagArchiveScreen = () => {
  const t = useT();
  return (
    <CommonArchiveManager
      pageTitle={t.tags.archiveTitle}
      newEntityHandler={({ locale, router }) => {
        router.push(TagEntity.Navigation.create(locale));
      }}
    >
      <TagList />
    </CommonArchiveManager>
  );
};
