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

export function Toast(
  content: ToastContent<unknown>,
  options?: ToastOptions<{}>
) {
  if (lastItem?.content == content) {
    return;
    // toast.dismiss(lastItem?.key);
    // lastItem = null;
  }

  const ref = toast(content, { hideProgressBar: true, ...options });
  lastItem = {
    content: content,
    key: ref,
  };

  setTimeout(() => {
    lastItem = null;
  }, 3000);
}
