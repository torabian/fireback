import { useS } from "../../fireback-ui/hooks/useS";
import { strings } from "./strings/translations";

export const DashboardScreen = () => {
  const s = useS(strings);
  return (
    <div>
      <h1 className="mt-4">{s.dashboard.title}</h1>
      <p>{s.dashboard.description}</p>
    </div>
  );
};
