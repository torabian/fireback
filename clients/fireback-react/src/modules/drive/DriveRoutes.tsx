import { Route } from "react-router-dom";
import { DriveArchiveScreen } from "./DriveArchiveScreen";
import { DriveFileSingleScreen } from "./DriveFileSingleScreen";

export function useDriveRoutes() {
  return (
    <>
      <Route path={"drive"} element={<DriveArchiveScreen />}></Route>
      <Route path={"drives"} element={<DriveArchiveScreen />}></Route>
      <Route path="file/:uniqueId" element={<DriveFileSingleScreen />} />
    </>
  );
}
