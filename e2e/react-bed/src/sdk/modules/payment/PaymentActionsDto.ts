/*
*	Generated by fireback 1.2.3
*	Written by Ali Torabi.
*	Checkout the repository for licenses and contribution: https://github.com/torabian/fireback
*/
import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
export class NotificationActionReqDto {
  /**
  The session which has been assigned during payment process initially.
  */
  public sessionId?: string | null;
  public orderId?: number | null;
public static Fields = {
      sessionId: 'sessionId',
      orderId: 'orderId',
}
}
export class VerifyTransactionActionReqDto {
  /**
  The session which has been assigned during payment process initially.
  */
  public sessionId?: string | null;
  /**
  The amount of transaction which will be payed. It's an integer, make sure you multiply the value by 100 instead of sending a float.
  */
  public amount?: number | null;
  /**
  The orderId which has been assigned by P24 and sent via notification
  */
  public orderId?: number | null;
public static Fields = {
      sessionId: 'sessionId',
      amount: 'amount',
      orderId: 'orderId',
}
}
export class RegisterTransactionActionReqDto {
  /**
  Customer email address
  */
  public email?: string | null;
  /**
  Describe the reason for the transaction
  */
  public description?: string | null;
  /**
  The amount of transaction which will be payed. It's an integer, make sure you multiply the value by 100 instead of sending a float.
  */
  public amount?: number | null;
public static Fields = {
      email: 'email',
      description: 'description',
      amount: 'amount',
}
}