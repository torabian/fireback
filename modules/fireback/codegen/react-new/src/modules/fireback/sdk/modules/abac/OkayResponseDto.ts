/*
*	Generated by fireback 1.2.3
*	Written by Ali Torabi.
*	Checkout the repository for licenses and contribution: https://github.com/torabian/fireback
*/
    import {
        BaseDto,
        BaseEntity,
    } from "../../core/definitions"
// In this section we have sub entities related to this object
// Class body
export type OkayResponseDtoKeys =
  keyof typeof OkayResponseDto.Fields;
export class OkayResponseDto extends BaseDto {
public static Fields = {
  ...BaseEntity.Fields,
}
  public static definition = {
  "name": "okayResponse"
}
}
