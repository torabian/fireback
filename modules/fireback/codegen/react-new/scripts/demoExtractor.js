const fs = require("fs");
const path = require("path");

function extractSnippetsFromFile(filePath) {
  const content = fs.readFileSync(filePath, "utf8");
  const regex = /const\s+((e|E)xample\d+)\s*=\s*\(\)\s*=>\s*{/g;

  let match;
  const snippets = {};

  while ((match = regex.exec(content)) !== null) {
    const name = match[1];
    const startIndex = match.index + match[0].length;

    let braceCount = 1;
    let endIndex = startIndex;

    while (braceCount > 0 && endIndex < content.length) {
      const char = content[endIndex++];
      if (char === "{") braceCount++;
      else if (char === "}") braceCount--;
    }

    const fnBody = content.slice(match.index, endIndex);
    snippets[name] = normalizeIndent(fnBody);
  }

  const outPath = filePath.replace(/\.tsx$/, ".snippets.ts");
  const result = `export const snippets = {\n${Object.entries(snippets)
    .map(([name, fn]) => `  "${name}": \`${fn}\``)
    .join(",\n")}\n};\n`;

  fs.writeFileSync(outPath, result);
  console.log(`Snippets saved to ${outPath}`);
}
function normalizeIndent(code) {
  const lines = code.split("\n");

  // Trim empty lines from start/end
  while (lines[0]?.trim() === "") lines.shift();
  while (lines.at(-1)?.trim() === "") lines.pop();

  // Check if all non-empty lines start with more than 2 spaces
  const hasCommonExtraIndent = lines
    .filter((line) => line.trim())
    .every((line) => line.startsWith("   ")); // at least 3 spaces

  if (hasCommonExtraIndent) {
    return lines
      .map((line) =>
        line.startsWith("   ") ? line.replace(/^ {2}/, "") : line
      )
      .join("\n");
  }

  return lines.join("\n");
}

extractSnippetsFromFile("./src/apps/projectname/demo/DemoModal.tsx");
extractSnippetsFromFile("./src/apps/projectname/demo/DemoFormSelect.tsx");
extractSnippetsFromFile("./src/apps/projectname/demo/DemoFormDates.tsx");
