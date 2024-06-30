import { source } from "@/fireback/hooks/source";
import { useLocale } from "@/fireback/hooks/useLocale";

export function AboutScreenFa() {
  return (
    <div>
      <h1>نرم افزار fireback</h1>
      <p>این داشبورد مدیریت مربوط به فریم ورک fireback است</p>

      <p>نویسنده:‌ علی ترابی</p>
      <a href="https://torabian.github.io" target="_blank">
        https://torabian.github.io
      </a>

      <h2>دانلود برای ESP</h2>
      <p>
        شما میتوانید فایل های نصبی این پروژه را برای برد های ESP از طریق لینک
        زیر دانلود کنید:
      </p>
      <a href="https://github.com/torabian/fireback" target="_blank">
        https://github.com/torabian/fireback
      </a>
    </div>
  );
}

export function AboutScreenEn() {
  return (
    <div>
      <h1>fireback</h1>
      <p>
        Fireback is a major framework to build many types of software, faster.
      </p>

      <h2>Download Fireback Binaries </h2>
      <a href="https://github.com/torabian/fireback" target="_blank">
        https://github.com/torabian/fireback
      </a>

      <p>Author: Ali Torabi</p>
      <a href="https://torabian.github.io" target="_blank">
        https://torabian.github.io
      </a>
    </div>
  );
}

export function AboutScreen() {
  const { locale } = useLocale();
  return <div>{locale === "fa" ? <AboutScreenFa /> : <AboutScreenEn />}</div>;
}
