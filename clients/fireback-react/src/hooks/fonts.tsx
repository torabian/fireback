export function loadFont(name: string, url: string, weight?: string) {
  return new Promise((resolve, reject) => {
    const myFont = new FontFace(name, `url(${url})`);
    if (weight) myFont.weight = weight;
    myFont
      .load()
      .then(() => {
        document.fonts.add(myFont);
        const el = document.createElement("DIV");
        el.style.fontFamily = name;
        resolve(true);
      })
      .catch(() => reject());
  });
}

export function localizeNumber(n: string, locale: string) {
  if (locale === "fa") {
    return toFarsiNumber(n);
  }

  return n;
}
export function toFarsiNumber(n: string) {
  const farsiDigits = ["۰", "۱", "۲", "۳", "۴", "۵", "۶", "۷", "۸", "۹"];

  return n
    .toString()
    .split("")
    .map((x: any) => (farsiDigits[x] ? farsiDigits[x] : x))
    .join("");
}
