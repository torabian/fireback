import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
export class CustomerActivityActionReqDto {
  public uniqueId?: string[] | null;
public static Fields = {
      uniqueId: 'uniqueId',
}
}