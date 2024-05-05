import { useT } from "@/fireback/hooks/useT";
import { CommonArchiveManager } from "@/fireback/components/entity-manager/CommonArchiveManager";
import { PostTagList } from "./PostTagList";
import { PostTagEntity } from "src/sdk/fireback/modules/cms/PostTagEntity";
export const PostTagArchiveScreen = () => {
  const t = useT();
  return (
    <CommonArchiveManager
      pageTitle={t.posttags.archiveTitle}
      newEntityHandler={({ locale, router }) => {
        router.push(PostTagEntity.Navigation.create(locale));
      }}
    >
      <PostTagList />
    </CommonArchiveManager>
  );
};
