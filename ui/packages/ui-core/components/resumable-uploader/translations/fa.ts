import type { UploadTranslations } from "../types";

export const fa: UploadTranslations = {
  attachFile: "پیوست فایل",
  dropHere: "فایل‌ها را اینجا رها کنید، یا کلیک کنید تا انتخاب کنید",
  browse: "انتخاب فایل",
  queued: "در صف",
  uploading: "در حال بارگذاری",
  paused: "متوقف شده",
  offlinePaused: "متوقف شده (آفلاین)",
  completed: "تکمیل شد",
  failed: "ناموفق",
  canceled: "لغو شد",
  pause: "توقف",
  resume: "ادامه",
  retry: "تلاش مجدد",
  remove: "حذف",
  clear: "پاک کردن",
  replace: "جایگزینی",
  currentFile: "فایل فعلی",
  offlineNotice:
    "شما آفلاین هستید. بارگذاری‌ها متوقف شده‌اند و به محض بازگشت اتصال، به‌طور خودکار ادامه می‌یابند.",
  onlineResuming: "اتصال برقرار شد، ادامه بارگذاری...",
  fileTooLarge: (maxMB: number) => `حجم فایل بیش از حد مجاز است. حداکثر حجم مجاز ${maxMB} مگابایت است.`,
  fileTypeNotAllowed: "نوع فایل مجاز نیست.",
  maxSizeLabel: (size: string) => `حداکثر حجم فایل: ${size}`,
};
