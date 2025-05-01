# Fireback React - React.js dashboard and boilerplate for desktop/web/embedded/mobile apps

A complete framework for building applications using react.

## DEMO https://torabian.github.io/fireback

* Based on react.js and typescript
* No redux crap
* Uses react-query and context api in it's core
* low dependecies, bootstrap 5 (optional)
* Batteries included, datatable, menu actions
* Build desktop apps using Go Wails ( or electron.js )
* Build mobile app with cordova
* Small output and optimized for embedded devices (500K max)


## Contribution

Please open a pull request :)


const [data, setData] = useState({
    schema:
      '{"type":"object","properties":{"newInput1":{"title":"New Input 1","type":"string"},"newInput3":{"title":"New Input 3","type":"object"},"newInput2":{"title":"New Input 2","type":"string"}},"dependencies":{},"required":[]}',
    uischema: '{"ui:order":["newInput1","newInput3","newInput2"]}',
  });


  <FormBuilder
        schema={data.schema}
        uischema={data.uischema}
        onChange={(newSchema: string, newUiSchema: string) => {
          setData({
            schema: newSchema,
            uischema: newUiSchema,
          });
        }}
      />
      <pre>{JSON.stringify(data, null, 2)}</pre>

      import { FormBuilder } from "@ginkgo-bioworks/react-json-schema-form-builder";
