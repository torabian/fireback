// Hand-written extension file - not touched by Emi recompiles (the JS/TS compiler
// never emits an "XEntity.ts" file for any entity, only Dto/OptionalDto/action classes -
// see UserSessionDto.ts's `passport: one? target: PassportEntity` field, which still
// needs a local `./PassportEntity` module to import). Re-exports the real generated
// PassportDto class under the name the relation field expects.
export { PassportDto as PassportEntity } from "./PassportDto";
export type { PassportDtoType as PassportEntityType } from "./PassportDto";
