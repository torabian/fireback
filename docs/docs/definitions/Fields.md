---
sidebar_position: 3
---

Fireback supports fields and datatypes which are common sense in programming world. Because Fireback generates 
code primarily for Golang, and then tries to give SDK for clients such as React, Android, etc, fields 
are inspired by Golang fields.

Nevertheless, I've added some extra field types which are commonly could be used in multiple projects.

## Defining a field and it's type

Despite the fact that fireback fields could be used in Actions, Dtos, Entities, they share the same general type
and could be used in the same way. In some context, some data types or attributes, might not make any sense.

## Common structure

Usually in Fireback Module3 definition you see `fields` attribute, and it's an Array of `[]*Module2Field`
and you need to specifiy `name` and `type` minimally. Definition doesn't have a default type, and not defining type
might have weird consequences.

```
entities:
    - name: creditCard
      fields:
      - name: title
        type: string 

```

As you see, we defined a field with type string for creditCard entity.


## Type: string

This is common string type, which would become `string` in Golang, Typescript, Java, and so on. Nothing special
is generated for this field other than a pointer string.

## Type: int64

This is a general integer which would resolve to int64 as expected on Golang, Typescript will have `number` field.

## Type: bool

Same as boolean on Go and other languages, you can only set true or false on the field. Mysql might save it as 0 or 1 instead of boolean value.

## Type: date

This type is a bit special. When a field is a date, it would have some code related to date validation, on Golang it would become `XDate` struct, and have some extra functions which are available to see on `XDate.go` to see.

## Type: daterange

This is another special datatype which instead of a single date, stores on database level 2 dates as start and end.
When a field is defined as `daterange`, it would automatically makes Start and End fields into the generated 
Go code.

It's important to understand date and daterange field, will create a meta object which will be available publicly,
with some information about the days between, or some other useful details.

**Important** When date is provided less than 1500, Fireback assumes it's a Iranian calendar date, and automatically converts it to European calendar for saving in database.

## Type: array

Array, is an array of structs in Fireback definition. You can define a list of children using `fields` on array field, and would create a separate struct, prefixing the parent struct in the name.

On database level, it's important to understand that `array` in Fireback, it means there will be a separate table for it, and a one-to-many relation will be created on the parent table.

For example from license module:

```
  - name: license
    fields:
      - name: permissions
        type: array
        fields:
          - name: capability
            type: one
            target: CapabilityEntity
            module: workspaces
            allowCreate: false
```

You see that `permissions` is an array. **Important** It's important that, every array field to end
in plural english form. Some code might end up in mismatch for singular fields as a bug until 1.1.27.


## Type: object

Type on is also a nested field, which you can define `fields` very similarly in an array.
The difference is, there will be a one-2-one relation created on this type of data, besides
there will be an Addtional *Id field which would link them by a unique string id. At the moment,
Fireback is creating `id` auto incremental numeric for each table, as well as `uniqueId` as a unique
string identifier. `id` is intended to be used internally only for Fireback, and any external relationship
between data, needs to use `uniqueId` field instead (`unique_id` on database level). This might change,
as in large projects it causes large indexing tables.

For example, see array example, and instead of `type: array` use `type: object`. Other details are exactly the same.

## Type: one

The type one, is a specific type to Fireback. It means, you are targeting another table as a `one-to-one` relationship. The different with `object` is, in this type managing the data, deleting it is not depending 
on the parent entity, they are just being related over a loose relation.

## Type: arrayp

Arrayp, is an array type which would store primitives in Fireback. For example, if you decided to use array strings
in a DTO, then you can define the `type: arrayP`. Now, the primitive itself needs to be defined on a separate property called: `primitive`. For example:

```
- name: capabilities
  type: arrayP
  primitive: string

```

Will resolve to a []string slice in Golang, and similar regime in other langauges.
    