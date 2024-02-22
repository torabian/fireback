import { enTranslations } from "@/translations/en";
import { Editor } from "tinymce";
const { renderToString } = require("katex");

declare const tinymce: any;

export const addKatexSupport = (t: typeof enTranslations) =>
  tinymce.PluginManager.add("example", function (editor: Editor, url: string) {
    var openDialog = function () {
      return editor.windowManager.open({
        title: t.katexPlugin.title,
        body: {
          type: "panel",
          items: [
            {
              type: "textarea",
              name: "title",
              label: t.katexPlugin.body,
              placeholder:
                "c = \\sum_{n=1}^{\\infty} \\pm \\sqrt[3]{(a^2 + b^3)^n}",
            },
          ],
        },
        buttons: [
          {
            type: "cancel",
            text: t.katexPlugin.cancel,
          },
          {
            type: "submit",
            text: t.katexPlugin.insert,
            primary: true,
          },
        ],
        onSubmit: function (api: any) {
          var data = api.getData();

          const m = renderToString(
            data.title
              ? data.title
              : `f(x) = \\displaystyle\\int_{-\\infty}^\\infty \\hat f(\\xi)\\,e^{2 \\pi i \\xi x} \\,d\\xi + \\lim_{h \\rightarrow 0 } \\frac{f(x+h)-f(x)}{h}`,
            { output: "mathml" }
          );

          editor.selection.setContent(
            "<span class='katex-inline-editor'>" + m + "</span>",
            { format: "raw" }
          );
          api.close();
        },
      });
    };
    /* Add a button that opens a window */
    editor.ui.registry.addButton("example", {
      text: t.katexPlugin.toolbarName,
      onAction: function () {
        /* Open window */
        openDialog();
      },
    });
    /* Adds a menu item, which can then be included in any menu via the menu/menubar configuration */
    editor.ui.registry.addMenuItem("example", {
      text: "Example plugin",
      onAction: function () {
        /* Open window */
        openDialog();
      },
    });
    /* Return the metadata for the help plugin */
    return {
      getMetadata: function () {
        return {
          name: "Example plugin",
          url: "http://exampleplugindocsurl.com",
        };
      },
    };
  });
