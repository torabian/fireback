import { useT } from "@/fireback/hooks/useT";
import { CommonArchiveManager } from "@/fireback/components/entity-manager/CommonArchiveManager";
import { PostList } from "./PostList";
import { PostEntity } from "src/sdk/fireback/modules/cms/PostEntity";
export const PostArchiveScreen = () => {
  const t = useT();
  return (
    <CommonArchiveManager
      pageTitle={t.posts.archiveTitle}
      newEntityHandler={({ locale, router }) => {
        router.push(PostEntity.Navigation.create(locale));
      }}
    >
      <PostList />
    </CommonArchiveManager>
  );
};
