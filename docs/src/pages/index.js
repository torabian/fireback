import clsx from "clsx";
import Link from "@docusaurus/Link";
import useDocusaurusContext from "@docusaurus/useDocusaurusContext";
import Layout from "@theme/Layout";
import HomepageFeatures from "@site/src/components/HomepageFeatures";

import Heading from "@theme/Heading";
import styles from "./index.module.css";

function HomepageHeader() {
  const { siteConfig } = useDocusaurusContext();
  return (
    <header className={clsx("hero hero--primary", styles.heroBanner)}>
      <div className="container">
        <Heading as="h1" className="hero__title">
          {siteConfig.title}
        </Heading>
        <p className="hero__subtitle">{siteConfig.tagline}</p>
        <div className={styles.buttons}>
          <Link
            className="button button--secondary button--lg"
            to="/docs/intro"
            style={{ margin: "5px" }}
          >
            Start new development life
          </Link>

          <a
            style={{ margin: "5px" }}
            className="button button--secondary button--lg"
            to="https://torabian.github.io/fireback/demo"
            target="_blank"
          >
            React.js Dashboard Demo
          </a>
        </div>
      </div>
    </header>
  );
}

export default function Home() {
  const { siteConfig } = useDocusaurusContext();
  return (
    <Layout
      title={`Fireback documentation`}
      description="Learn how to create backend in Fireback and Golang"
    >
      <HomepageHeader />
      <main>
        <div style={{ textAlign: "center" }}>
          <iframe
            style={{ margin: "100px auto" }}
            width="560"
            height="315"
            src="https://www.youtube.com/embed/G2Wjeq7ZmS0?si=dNc9igu-6Ia9-HiH"
            title="YouTube video player"
            frameborder="0"
            allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share"
            referrerpolicy="strict-origin-when-cross-origin"
            allowfullscreen
          ></iframe>
        </div>
      </main>
    </Layout>
  );
}
