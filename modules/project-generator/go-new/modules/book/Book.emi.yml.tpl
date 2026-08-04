# This is an emi file which can contain different set of functionalities. Fireback embeds emi, fully without change. 
# Learn more on https://torabian.github.io/emi about the features.

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
