import { usePureLocale } from "../../hooks/usePureLocale";
import { strings } from "../strings/translations";

export function Fallback({ error, resetErrorBoundary }: any) {
  // Call resetErrorBoundary() to reset the error boundary and retry the render.

  // Deliberately not useS()/useLocale() here: those go through useRouter(),
  // which calls useNavigate()/useParams()/useLocation() - hooks that throw
  // if there's no <Router> ancestor. This is a react-error-boundary
  // FallbackComponent, so whenever the error it's rendering for came from
  // routed content, React has already unmounted that whole subtree (Router
  // included) before mounting this in its place - there is no Router
  // ancestor to find. usePureLocale() reads the locale straight from the
  // URL instead, with no context dependency, so the fallback itself can't
  // also crash.
  const { locale } = usePureLocale();
  const s = (locale !== "en" && (strings as any)["$" + locale]) || strings;

  return (
    <div role="alert">
      <p>{s.components.somethingWentWrong}</p>
      <div style={{ color: "red", padding: "30px" }}>{error?.message}</div>
    </div>
  );
}
