export declare class ListResponse<T> {
    private ItemClass;
    items: T[];
    constructor(ItemClass: {
        new (): T;
    });
    static unserialize<T>(raw: any, ItemClass: {
        new (): T;
    }): ListResponse<T>;
    static By<T>(x: {
        new (): T;
    }): {
        new (): {};
        unserialize(raw: unknown): ListResponse<T>;
    };
}
