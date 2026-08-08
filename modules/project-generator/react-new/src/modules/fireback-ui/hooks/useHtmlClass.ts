import { useEffect } from "react";

/**
 * Adds a class to body when components adds or removes it
 */
export function useHtmlClass(className: string, selector = "body") {
  useEffect(() => {
    (document as any)?.querySelector(selector).classList.add(className);

    return () =>
      (document as any)?.querySelector(selector).classList.remove(className);
  }, []);
}

export function getOS(): string {
  if (typeof window === "undefined") {
    return "mac";
  }

  let userAgent = window?.navigator.userAgent,
    platform = window?.navigator.platform,
    macosPlatforms = ["Macintosh", "MacIntel", "MacPPC", "Mac68K"],
    windowsPlatforms = ["Win32", "Win64", "Windows", "WinCE"],
    iosPlatforms = ["iPhone", "iPad", "iPod"],
    os = "mac";

  if (macosPlatforms.indexOf(platform) !== -1) {
    os = "mac";
  } else if (iosPlatforms.indexOf(platform) !== -1) {
    os = "ios";
  } else if (windowsPlatforms.indexOf(platform) !== -1) {
    os = "windows";
  } else if (/Android/.test(userAgent)) {
    os = "android";
  } else if (!os && /Linux/.test(platform)) {
    os = "linux";
  } else {
    os = "web";
  }

  return os;
}

export function useOsClass(selector = "html") {
  useEffect(() => {
    let os = getOS();

    (document as any)?.querySelector(selector).classList.add(`${os}-theme`);

    return () =>
      (document as any)
        ?.querySelector(selector)
        .classList.remove(`${os}-theme`);
  }, []);
}
