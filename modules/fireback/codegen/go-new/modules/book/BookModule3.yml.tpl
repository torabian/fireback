# This is an example of module3 definition.
# You can learn more on https://torabi.io/fireback or https://torabian.github.io/fireback
# After compiling for first time, json schema will be created in .jsonschemas folder.
# For VSCode, there is already a configuration in .vscode/settings.json
# To target all Module3.yml files with autocompletion. It requires "Redhat YAML Language Support"
# to be installed in order to show autocompletion.

path: book
name: book
entities:
  # This is a sample entity. You can delete it and write your own
  - name: book
    fields:
    - name: title
      type: string
      validate: required
      translate: true
    - name: pageCount
      type: int64
    - name: isbn
      type: string
