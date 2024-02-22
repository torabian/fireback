interface CTX {
  params: { locale: "en" | "es" };
  locales: string;
  locale: string;
  defaultLocale: string;
}

export const getStaticPaths = async () => {
  return {
    fallback: false,
    paths: [
      {
        params: {
          locale: "en",
        },
      },
      {
        params: {
          locale: "fa",
        },
      },
    ],
  };
};

export function makeStaticProps(ns: any = {}) {
  return async function getStaticProps(ctx: CTX) {
    return {
      props: { locale: ctx.params.locale, ...ns },
    };
  };
}
