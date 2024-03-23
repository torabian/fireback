import { Route } from "react-router-dom";
import { DiscountCodeArchiveScreen } from "./DiscountCodeArchiveScreen";
import { DiscountCodeEntityManager } from "./DiscountCodeEntityManager";
import { DiscountCodeSingleScreen } from "./DiscountCodeSingleScreen";
import { DiscountCodeEntity } from "src/sdk/fireback/modules/shop/DiscountCodeEntity";
export function useDiscountCodeRoutes() {
  return (
    <>
      <Route
        element={<DiscountCodeEntityManager />}
        path={ DiscountCodeEntity.Navigation.Rcreate}
      />
      <Route
        element={<DiscountCodeSingleScreen />}
        path={ DiscountCodeEntity.Navigation.Rsingle}
      ></Route>
      <Route
        element={<DiscountCodeEntityManager />}
        path={ DiscountCodeEntity.Navigation.Redit}
      ></Route>
      <Route
        element={<DiscountCodeArchiveScreen />}
        path={  DiscountCodeEntity.Navigation.Rquery}
      ></Route>
    </>
  );
}