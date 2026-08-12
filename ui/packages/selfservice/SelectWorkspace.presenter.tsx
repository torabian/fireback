import { useRouter } from "@fireback/ui-core/hooks/useRouter";
import { useS } from "@fireback/ui-core/hooks/useS";
import { strings } from "./strings/translations";

export const usePresenter = () => {
  const s = useS(strings);
  const { goBack, query } = useRouter();

  return {
    goBack,
    s,
  };
};
