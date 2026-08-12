/**
 * To declare a new envelope class, you must follow these rules:
 * 1. The constructor must accept any object, which will contain the parsed JSON message
 *    from a response.
 * 2. Envelope classes must provide a function to update the payload. Since payloads are
 *    type-safe, they must be instantiated and passed to the envelope; the common constructor
 *    alone is not enough.
 * 3. Enveope must have a way to provide the content back actually, in order to create a class out of them.
 */

export type CreatorSignature<T> = (item: unknown) => T;

export interface EnvelopeClass<T> {
  setCreator(fn: CreatorSignature<T>): this;
  inject(data: unknown): this;
}
