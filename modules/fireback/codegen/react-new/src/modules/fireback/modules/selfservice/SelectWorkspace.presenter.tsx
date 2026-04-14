import { useRouter } from "../../hooks/useRouter";
import { useS } from "../../hooks/useS";
import { strings } from "./strings/translations";

export const usePresenter = () => {
  const s = useS(strings);
  const { goBack, query } = useRouter();

  return {
    goBack,
    s,
  };
};
