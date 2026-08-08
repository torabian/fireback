import { useRouter } from "../../../fireback-ui/hooks/useRouter";
import { useS } from "../../../fireback-ui/hooks/useS";
import { strings } from "./strings/translations";

export const usePresenter = () => {
  const s = useS(strings);
  const { goBack, query } = useRouter();

  return {
    goBack,
    s,
  };
};
