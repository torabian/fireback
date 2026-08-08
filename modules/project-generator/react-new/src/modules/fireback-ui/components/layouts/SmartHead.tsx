import { useS } from "../../hooks/useS";
import { strings } from "../strings/translations";

interface PageHead {
  title: string | string[];
}

export function SmartHead(props: PageHead) {
  const s = useS(strings);
  const title = Array.isArray(props.title)
    ? props.title
        .filter(Boolean)
        .filter((t) => `${t}`.trim())
        .join(" | ")
    : props.title;

  return (
    // <Head>
    <>
      <meta property="og:title" content={title} />
      <meta property="og:image" content="/pixelplux-spzoo-logo.png" />
      <title>{`${title} | ${s.meta.titleAffix}`}</title>
      {/* // </Head> */}
    </>
  );
}
