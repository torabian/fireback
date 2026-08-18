// Deterministic fake row generator for the DataGridList demo - one row shape
// deliberately covering every emi field type (see modules/fireback/
// module3-json-schema.json's EmiField.type enum: string, int/int32/int64,
// float32/float64, bool, complex, slice/array, one, collection, enum,
// object/map, any) so DataGridListDemo's first example can prove each of them
// renders correctly as a column.
// complexes.TString (modules/fireback/complexes/TString.go) is a locale -> text map,
// e.g. {"en": "Home", "fa": "خانه"} - stored as one JSON object per row, not a flat
// string. Mirrored here as a plain Record rather than importing the generated SDK's
// own (currently broken - it references a `TString` type it never imports/defines,
// see AppMenuDto.ts) type.
export type TString = Record<string, string>;

export interface EmiTypesFakeRow {
  /** string */
  uniqueId: string;
  /** string */
  title: string;
  /** complex (TString) - a translated label, one value per language */
  label: TString;
  /** int */
  age: number;
  /** int64 - kept large enough to actually need 64 bits in a real system */
  views: number;
  /** float64 */
  price: number;
  /** bool */
  active: boolean;
  /** complex (PlainTime) */
  createdAt: string;
  /** slice/array of a primitive */
  tags: string[];
  /** enum */
  status: "draft" | "published" | "archived";
  /** one (a single relation) */
  owner: { uniqueId: string; name: string };
  /** collection (a to-many relation) */
  collaborators: { uniqueId: string; name: string }[];
  /** object/map */
  metadata: Record<string, unknown>;
  /** any */
  note: unknown;
}

const STATUSES: EmiTypesFakeRow["status"][] = ["draft", "published", "archived"];
const NAMES = [
  "Alex Cox",
  "Noah Stewart",
  "Zahra Cox",
  "Jason King",
  "Elena Anderson",
  "Fatima Morris",
  "Gabriel Ward",
];
const TAG_POOL = ["ui", "backend", "urgent", "design", "infra", "billing", "mobile"];

// One TString-shaped translation set per "kind" word title already picks from, so a
// row's label is recognizably related to its title (e.g. title "Item #3 - Rocket" pairs
// with label {"en": "Rocket", "fa": "موشک", "pl": "Rakieta"}) instead of an unrelated
// second batch of random words.
const KIND_TRANSLATIONS: Record<string, TString> = {
  Rocket: { en: "Rocket", fa: "موشک", pl: "Rakieta" },
  Comet: { en: "Comet", fa: "دنباله‌دار", pl: "Kometa" },
  Nebula: { en: "Nebula", fa: "سحابی", pl: "Mgławica" },
  Orbit: { en: "Orbit", fa: "مدار", pl: "Orbita" },
  Lunar: { en: "Lunar", fa: "قمری", pl: "Księżycowy" },
  Solar: { en: "Solar", fa: "خورشیدی", pl: "Słoneczny" },
};
const KINDS = Object.keys(KIND_TRANSLATIONS);

// A tiny mulberry32-style PRNG so the same id always produces the same row
// (needed so scrolling further, changing sort, or refetching a page never
// shows different data for a row you've already seen).
function seededRandom(seed: number) {
  let t = seed + 0x6d2b79f5;
  return () => {
    t = (t + 0x6d2b79f5) | 0;
    let r = Math.imul(t ^ (t >>> 15), 1 | t);
    r = (r + Math.imul(r ^ (r >>> 7), 61 | r)) ^ r;
    return ((r ^ (r >>> 14)) >>> 0) / 4294967296;
  };
}

export function generateFakeRow(id: number): EmiTypesFakeRow {
  const rand = seededRandom(id);
  const pick = <T,>(arr: T[]) => arr[Math.floor(rand() * arr.length)];

  const daysAgo = Math.floor(rand() * 400);
  const createdAt = new Date(Date.now() - daysAgo * 86400000).toISOString();

  const tagCount = 1 + Math.floor(rand() * 3);
  const tags = Array.from(new Set(Array.from({ length: tagCount }, () => pick(TAG_POOL))));

  const collaboratorCount = Math.floor(rand() * 3);
  const collaborators = Array.from({ length: collaboratorCount }, (_, i) => ({
    uniqueId: `user-${id}-${i}`,
    name: pick(NAMES),
  }));

  const kind = pick(KINDS);

  return {
    uniqueId: String(id),
    title: `Item #${id} - ${kind}`,
    label: KIND_TRANSLATIONS[kind],
    age: Math.floor(rand() * 80) + 18,
    views: Math.floor(rand() * 1_000_000_000),
    price: Math.round(rand() * 99999) / 100,
    active: rand() > 0.4,
    createdAt,
    tags,
    status: pick(STATUSES),
    owner: { uniqueId: `user-${id}-owner`, name: pick(NAMES) },
    collaborators,
    metadata: { source: pick(["import", "manual", "api"]), score: Math.round(rand() * 100) },
    note: rand() > 0.5 ? pick(["Needs review", 42, null]) : null,
  };
}
