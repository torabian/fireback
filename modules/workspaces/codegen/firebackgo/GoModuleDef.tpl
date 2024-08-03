name: {{ .name }}

# Define the module entities here:

entities:
  # This is a sample entity. You can delete it and write your own
  - name: {{ .name }}
    fields:
    - name: title
      type: string
      validate: required

dtos: 
actions: