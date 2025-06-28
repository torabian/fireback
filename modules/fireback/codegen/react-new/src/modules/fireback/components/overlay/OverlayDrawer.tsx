import Drawer from "react-modern-drawer";
import "react-modern-drawer/dist/index.css";
import {
  DrawerOpenParams,
  OverlayInstanceComponentProps,
} from "./OverlayDefinitions";

export const OverlayDrawerImp = ({
  params,
  children,
  visible,
  close,
}: OverlayInstanceComponentProps<unknown, DrawerOpenParams | undefined>) => {
  return (
    <Drawer
      open={visible}
      direction={params?.direction || "right"}
      zIndex={10000}
      onClose={close}
      size={params?.size}
    >
      {children}
    </Drawer>
  );
};
