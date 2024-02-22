const { globSync } = require("glob");
const fs = require("fs");
const { plural } = require("pluralize");
const ejs = require("ejs");
const path = require("path");

const camelToSnakeCase = (str) =>
  str.replace(/[A-Z]/g, (letter) => `_${letter.toLowerCase()}`);

function castGetQueryFromGolangType(gotype) {
  return "useGet" + gotype.split(".")[1].replace("Entity", "") + "s";
}

function toUpper(str) {
  if (!str) {
    return "";
  }
  return `${str.substr(0, 1).toUpperCase()}${str.substr(1)}`;
}

function toLower(str) {
  if (!str) {
    return "";
  }
  return `${str.substr(0, 1).toLowerCase()}${str.substr(1)}`;
}

const pluralize = (noun) => {
  return plural(noun);
};

const DefaultFirebackFields = [
  "Visibility",
  "WorkspaceId",
  "LinkerId",
  "ParentId",
  "UniqueId",
  "UserId",
  "Rank",
  "Updated",
  "Created",
  "CreatedFormatted",
  "UpdatedFormatted",
];

function getFieldName(field) {
  if (field.includes(",")) {
    return field.split(",")[0];
  }

  return field;
}

function getPublicFieldsFromSchema(schema) {
  return schema
    .filter((item) => {
      if (item.name.substr(0, 1) !== item.name.substr(0, 1).toUpperCase()) {
        return false;
      }

      if (DefaultFirebackFields.includes(item.name)) {
        return false;
      }

      return true;
    })
    .map((item) => {
      return {
        ...item,
        name: getFieldName(item.jsonField),
      };
    });
}

function replaceTemplate(input, entityName) {
  let lowerDashed = camelToSnakeCase(entityName).replaceAll("_", "-");
  if (lowerDashed.startsWith("-")) {
    lowerDashed = lowerDashed.substr(1);
  }

  return input
    .replaceAll("Templates", pluralize(toUpper(entityName)))
    .replaceAll("templates", pluralize(entityName.toLowerCase()))
    .replaceAll("Template", toUpper(entityName))
    .replaceAll("template", toLower(entityName))
    .replaceAll("xsdk", toLower(process.env.TARGET_SDK))
    .replaceAll("TEMPLATE", entityName.toUpperCase())
    .replaceAll("xnavigation", lowerDashed + "-navigation-tools")
    .replaceAll("xtypefields", lowerDashed + "-fields")
    .replaceAll("xmodule", process.env.BACKEND_MODULE);
}

function createFromTemplate(entitySchemaName, source, dest) {
  const entityName = entitySchemaName.replace(".json", "");
  const jsonSchema = require("../../../artifacts/entity-schema/" +
    entitySchemaName);
  const files = globSync(source);

  for (const file of files) {
    let content = fs.readFileSync(file).toString();
    let destFile = path.join(
      dest,
      path.basename(file).replaceAll("Template", toUpper(entityName))
    );

    content = ejs.render(content, {
      fields: getPublicFieldsFromSchema(jsonSchema),
      castGetQueryFromGolangType,
    });

    content = replaceTemplate(content, entityName);

    fs.writeFileSync(destFile, content);
  }
}

function helpTranslations(entitySchemaName, source, dest) {
  const entityName = entitySchemaName.replace(".json", "");
  const jsonSchema = require("../../../artifacts/entity-schema/" +
    entitySchemaName);
  const files = globSync(source);
  const fields = getPublicFieldsFromSchema(jsonSchema);

  const dic = {
    ["edit" + entityName]: "Edit",
    ["new" + entityName]: "Edit",
    ["archiveTitle"]: "Edit",
  };

  for (const field of fields) {
    dic[field.name] = field.name;
    dic[field.name + "Hint"] = field.name + " Hint";
  }

  return dic;
}

function writeToTarget(entitySchemaName, target) {
  const entityName = entitySchemaName.replace(".json", "");

  const sRouteKeyword = "{/* ~ auto:useRouteJsx */}";
  const sRouteDefs = "// ~ auto:useRouteDefs";
  const sRouteImport = "// ~ auto:useRouteImport";
  const sUseMockImport = "// ~ auto:useMockImport";
  const sUseMockNew = "// ~ auto:useMocknew";
  const paths = [
    path.join("src/apps", target, "ApplicationRoutes.tsx"),
    path.join("src/apps", target, "mockServer.ts"),
  ];

  for (let file of paths) {
    let routes = fs
      .readFileSync(file)
      .toString()
      .replaceAll(
        sRouteKeyword,
        replaceTemplate(`{templateRoutes}\r\n${sRouteKeyword}`, entityName)
      )
      .replaceAll(
        sUseMockImport,
        replaceTemplate(
          `import { TemplateMockProvider } from "@/modules/xmodule/TemplateMockProvider";\r\n${sUseMockImport}`,
          entityName
        )
      )
      .replaceAll(
        sUseMockNew,
        replaceTemplate(
          `new TemplateMockProvider(),\r\n${sUseMockNew}`,
          entityName
        )
      )
      .replaceAll(
        sRouteDefs,
        replaceTemplate(
          `const templateRoutes = useTemplateRoutes();\r\n${sRouteDefs}`,
          entityName
        )
      )
      .replaceAll(
        sRouteImport,
        replaceTemplate(
          `import { useTemplateRoutes } from "@/modules/${process.env.BACKEND_MODULE}/TemplateRoutes";\r\n${sRouteImport}`,
          entityName
        )
      );
    fs.writeFileSync(file, routes);
  }
}

function helpSidebar() {
  const entityName = process.env.ENTITY_SCHEMA_NAME.replace(".json", "");

  console.log(
    replaceTemplate(
      `
  {
    href: TemplateNavigationTools.query(),
    // icon: osResources.province,
    label: t.templates.archiveTitle,
    children: [],
    // activeMatcher: /\\/template\\//, Make sure this is sneak case, test it in sub screens
    displayFn: onPermission(ROOT_TEMPLATE_QUERY),
  },
  
  `,
      entityName
    )
  );
}

createFromTemplate(
  process.env.ENTITY_SCHEMA_NAME,
  path.join("scripts/entity-template/module", "*.*"),
  `src/modules/${process.env.TAGET_MODULE_NAME}/`
);

writeToTarget(process.env.ENTITY_SCHEMA_NAME, process.env.TARGET_APP);

console.log(
  helpTranslations(
    process.env.ENTITY_SCHEMA_NAME,
    path.join("scripts/entity-template/module", "*.*"),
    `src/modules/${process.env.TAGET_MODULE_NAME}/`
  )
);

helpSidebar();
