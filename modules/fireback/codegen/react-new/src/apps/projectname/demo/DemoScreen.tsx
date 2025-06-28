import Link from "@/modules/fireback/components/link/Link";

export function DemoScreen() {
  return (
    <div>
      <h1>Demo screen</h1>
      <p>
        Here I put some demo and example of fireback components for react.js
      </p>
      <div>
        <Link href="/demo/modals">Check modals</Link>
      </div>
      <div>
        <Link href="/demo/form-select">Check Selects</Link>
      </div>
      <div>
        <Link href="/demo/form-date">Check Date Inputs</Link>
      </div>
      <hr />
    </div>
  );
}
