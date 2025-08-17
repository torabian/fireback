export class ListResponse {
    constructor(ItemClass) {
        this.ItemClass = ItemClass;
        this.items = [];
    }
    static unserialize(raw, ItemClass) {
        const list = new ListResponse(ItemClass);
        if (Array.isArray(raw)) {
            list.items = raw.map((x) => Object.assign(new ItemClass(), x));
        }
        else if (raw.items && Array.isArray(raw.items)) {
            list.items = raw.items.map((x) => Object.assign(new ItemClass(), x));
        }
        return list;
    }
    static By(x) {
        return class {
            static unserialize(raw) {
                return ListResponse.unserialize(raw, x);
            }
        };
    }
}
