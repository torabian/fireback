import { type CardComponentType } from "../../fireback-ui/components/entity-manager/FlatListMode";
import { UserDto } from "../../sdk/abac/UserDto";

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
