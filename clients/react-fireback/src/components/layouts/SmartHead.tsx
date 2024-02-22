// import Head from "next/head";
import { useT } from "../../hooks/useT";

interface PageHead {
  title: string | string[];
}

export function SmartHead(props: PageHead) {
  const t = useT();
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
      <title>{`${title} | ${t.meta.titleAffix}`}</title>
      {/* // </Head> */}
    </>
  );
}
