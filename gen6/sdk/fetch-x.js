class TypedResponse extends Response {
    json() {
        return super.json();
    }
}
export function fetchx(input, init) {
    return fetch(input, init);
}
