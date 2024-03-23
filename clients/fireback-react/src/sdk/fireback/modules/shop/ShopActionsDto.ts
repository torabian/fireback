import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
export class ConfirmPurchaseActionReqDto {
  public basketId?: string | null;
  public currencyId?: string | null;
public static Fields = {
      basketId: 'basketId',
      currencyId: 'currencyId',
}
}