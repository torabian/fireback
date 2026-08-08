import { type CardComponentType } from "@/modules/fireback/components/entity-manager/FlatListMode";
import { UserDto } from "@/modules/fireback/sdk/abac/UserDto";

export const UserCard: CardComponentType<UserDto> = ({ content }) => {
  return (
    <div style={{ height: "200px" }}>
      <h2>{content.firstName}</h2>
    </div>
  );
};

UserCard.getHeight = () => {
  return 230;
};
