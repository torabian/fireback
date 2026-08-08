export interface DeviceTypeInfo {
  isPhysicalPhone: boolean; // Likely a handheld phone
  isTablet: boolean; // Likely a tablet
  isDesktop: boolean; // Desktop or laptop
  isMobileView: boolean; // Based on viewport width
  isCordova?: boolean; // Running under Cordova
  viewSize: "small" | "medium" | "large"; // Viewport category
}

export const detectDeviceType = (): DeviceTypeInfo => {
  const ua = navigator.userAgent.toLowerCase();
  const isTouch = "ontouchstart" in window || navigator.maxTouchPoints > 0;
  const width = window.innerWidth || document.documentElement.clientWidth;

  const isCordova =
    !!(window as any).cordova || !!(window as any).cordovaPlatformId;

  const isPhoneRegex =
    /iphone|android.*mobile|blackberry|windows phone|opera mini|iemobile/;
  const isTabletRegex = /ipad|android(?!.*mobile)|tablet/;
  const isDesktopRegex = /windows|macintosh|linux|x11/;

  const isPhysicalPhone = isPhoneRegex.test(ua);
  const isTablet = isTabletRegex.test(ua);
  const isDesktop = !isTouch || isDesktopRegex.test(ua);

  let viewSize: DeviceTypeInfo["viewSize"] = "large";
  if (width < 600) {
    viewSize = "small";
  } else if (width < 1024) {
    viewSize = "medium";
  }

  const isMobileView = width < 1024;

  return {
    isPhysicalPhone,
    isTablet,
    isDesktop,
    isMobileView,
    isCordova,
    viewSize,
  };
};
