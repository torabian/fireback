path: {{ .path }}
name: {{ .name }}

# Define the module entities here:

entities:
  # This is a sample entity. You can delete it and write your own
  - name: book
    fields:
    - name: title
      type: string
      validate: required

dto: 
actions: