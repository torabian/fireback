import type { CreatorSignature, EnvelopeClass } from "../common/EnvelopeClass";

// Some responses are returning a flat array, without any other information
// regarding the pagination, etc. They start [{},{}...] style.
// For those which might have a single project which is array but none-standard,
// Use CaptureArray instead.

export class FlatArray<T> implements EnvelopeClass<T> {
  public data: T[] = [];
  creator?: CreatorSignature<T> | null = null;

  inject(data: unknown) {
    if (!Array.isArray(data)) {
      throw new Error(
        "FlatArray can only work on flat array items, such as [{},...], and doesn't accept any other type"
      );
    }

    if (typeof this.creator !== "undefined") {
      this.data = data.map((item) =>
        (this.creator as CreatorSignature<T>)(item)
      );
    } else {
      this.data = data;
    }

    return this;
  }

  setCreator(creator: CreatorSignature<T>) {
    this.creator = creator;

    return this;
  }

  // Implement thigs here, which would make the flat array to act as an array actually, how can it become
  // class instance as an array?
}
