path: book
name: book

# Define the module entities here:

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
dtos: 
actions: