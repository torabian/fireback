type ConstructorWithArg<T = any, R = any> = new (arg: T, ...rest: any[]) => R;
export class WebSocketX<
  SendType = string | ArrayBufferLike | Blob | ArrayBufferView,
  RecieveData = string
> extends WebSocket {
  public readonly addEventListenerRaw: WebSocket["addEventListener"];
  public readonly sendRaw: WebSocket["send"];
  #factoryCls?: ConstructorWithArg<unknown>;
  constructor(
    url: string | URL,
    protocols?: string | string[],
    options?: {
      MessageFactoryClass: ConstructorWithArg<any>;
    }
  ) {
    super(url, protocols);
    this.sendRaw = super.send.bind(this);
    this.addEventListenerRaw = super.addEventListener.bind(this);
    if (options?.MessageFactoryClass) {
      this.#factoryCls = options.MessageFactoryClass;
    }
  }
  set onmessage(
    fn: ((this: WebSocket, ev: MessageEvent<RecieveData>) => any) | null
  ) {
    if (fn) {
      this.addEventListener("message", fn);
    } else {
      super.onmessage = null;
    }
  }
  get onmessage() {
    return super.onmessage;
  }
  // @ts-expect-error override to customize send
  send(data: SendType): void {
    if (
      typeof data === "string" ||
      data instanceof Blob ||
      data instanceof ArrayBuffer ||
      ArrayBuffer.isView(data)
    ) {
      super.send(data);
    } else if (data !== undefined && data !== null) {
      super.send(data.toString());
    }
  }
  addEventListener(
    type: "message",
    listener: (this: WebSocket, ev: MessageEvent<RecieveData>) => unknown,
    options?: boolean | AddEventListenerOptions
  ): void;
  // fallback overloads (other event types)
  addEventListener<K extends Exclude<keyof WebSocketEventMap, "message">>(
    type: K,
    listener: (this: WebSocket, ev: WebSocketEventMap[K]) => unknown,
    options?: boolean | AddEventListenerOptions
  ): void;
  // implementation
  addEventListener(
    type: string,
    listener: EventListenerOrEventListenerObject,
    options?: boolean | AddEventListenerOptions
  ): void {
    if (type === "message") {
      const wrapped = ((ev: MessageEvent) => {
        let parsed: unknown = ev.data;
        if (this.#factoryCls) {
          try {
            parsed = new this.#factoryCls(ev.data);
          } catch {
            // if constructor rejects (e.g. ArrayBuffer not supported), keep raw
          }
        }
        (listener as EventListener).call(
          this,
          new MessageEvent("message", { data: parsed })
        );
      }) as EventListener;
      super.addEventListener(type, wrapped, options);
    } else {
      super.addEventListener(type, listener, options);
    }
  }
}