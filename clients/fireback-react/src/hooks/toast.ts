import { Id, ToastContent, ToastOptions, toast } from "react-toastify";

// interface RefItem {
//   key: number;
//   message: string;
// }
// const toastRefs: RefItem[] = [];

let lastItem: {
  key: Id;
  content: ToastContent<unknown>;
} = null;

const TOAST_DURATION = 2500;
export function Toast(
  content: ToastContent<unknown>,
  options?: ToastOptions<{}>
) {
  if (lastItem?.content == content) {
    return;
  }

  const ref = toast(content, {
    hideProgressBar: true,
    autoClose: TOAST_DURATION,
    ...options,
  });
  lastItem = {
    content: content,
    key: ref,
  };

  setTimeout(() => {
    lastItem = null;
  }, TOAST_DURATION);
}
