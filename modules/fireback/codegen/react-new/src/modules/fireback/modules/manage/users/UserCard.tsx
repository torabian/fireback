import { CardComponentType } from "@/modules/fireback/components/entity-manager/FlatListMode";
import { UserEntity } from "@/modules/fireback/sdk/modules/abac/UserEntity";

export const UserCard: CardComponentType<UserEntity> = ({ content }) => {
  return (
    <div style={{ height: "200px" }}>
      <h2>{content.firstName}</h2>
    </div>
  );
};

UserCard.getHeight = () => {
  return 230;
};
