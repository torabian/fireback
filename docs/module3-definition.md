# Module3 Definition in Fireback

In this tutorial, we will learn about module tree definition. A module tree 
definition is a structure defined using YAML or, alternatively, JSON format. 
You describe what your system needs, which generally consists of a few main 
sections: actions, entities, and DTOs.

Content of the document:

0. **Overview of Module3 definition**
1. **Defining DTOs**
2. **Data Types in Fireback**
3. **Defining Entities**
4. **Defining Relations in Fireback Fields**
5. **Defining Actions**
6. **Reactive Actions**
7. **Generating Code from Definitions**
8. **Automating Code Generation in VS Code**
9. **Automating Code Generation in IntelliJ IDEA**

## Overview of Module3 definition

Using Fireback, you define the necessary components, and from that definition, 
back-end and front-end code can be generated. While you can define the module 
tree directly in a programming language like Go without using YAML, YAML 
simplifies the process and reduces the workload by providing a structured 
approach. Also, fireback is well tuned to convert these yaml to code
using cli in an streamlined process.

The module tree format primarily focuses on how the database structure should be managed,
table relationships, and how API calls are directed to them. For example, if you want 
to have a list of users, you need to define an action for retrieving the user 
list and define two DTOs: one for the criteria and one for the user itself. 
To implement this functionality, you write a function in Golang or Java 
containing the necessary logic.

DTOs (Data Transfer Objects) are simple classes. When you define a DTO, it is 
converted into a class in Java or a struct in Golang, and similarly into 
TypeScript, Swift, and other languages supported by Fireback.

Entities are more complex. When you define an entity, such as a video entity 
with fields like title, length, and total views, it creates a corresponding 
entity in the database. It also generates various actions, such as video update 
and bulk update, and handles the event system. For instance, if a video is 
updated, an event is triggered in the back end, notifying the front end of the 
update.

Entities cover many aspects of a real world entity management, from API call, CLI action,
permissions, translation. In fact, an entity defined in fireback is a ready to go production
feature, without any extra modification most of the time.

## Defining DTOs

Dto, is the simplest form of data type in most applications. DTOs are `class` or `struct` 
in different programming languages, and their job is to compose a message, and usually
they are encoded as json, or being read from a json string.

For example, in a POST message on http, your post body might represent a json with `name` and `password` field, therefor we can say you have a DTO with `email` and `password` fields,
and we might also call it `UserLoginReqDto`. Fireback adds `Dto` affix to all DTOs, to make
it clearer as a convention.

Now, let's define this dto in a yaml format, compatible with Fireback Module3 format:

```yaml
# Name the module, and explain it's path. This is how the code gen would be organized
# in different platforms.
name: test
path: test

## dto is an array, you can define as many as dtos you want.
dto:
- name: userLoginReq
  fields:
    # Set the name of the field, and then type of it as string.
  - name: email
    type: string

# Add more dto here
```

In fireback, defining all DTOs most go under `dto` item, also you need to define `name` and `path`
for each module you are creating. (For module, not for each dto).


## Data Types in Fireback

Fireback supports various data types:

- **bool**: Represents a boolean value, `true` or `false`.
- **string**: Represents text data.
- **html**: Represents html data.
- **enum**: Common enum system.
- **date**: A rich date type used for date
- **daterange**: A rich date type, for saving range of dates (from/to date)
- **int64**: Represents a 64-bit integer, suitable for large integer values.
- **float64**: Represents a 64-bit floating-point number, providing precision 
for decimal values.
- **object**: An object field, can have it's own fields, act as dependent
table to the parent table
- **array**: It's similar to the object, but it it's an array of items
depeneding
- **many2many**: It creates a many 2 many relation between another entity.
Data does not depend on connecting table.
- **one**: It's one 2 one relation between entity and another entity.


Fireback also introduces data types for representing relationships between 
data. One such type is the **Object** type, which indicates the presence of 
another object (another table in the database). Objects can contain other 
objects or arrays. There is a distinction between objects managed by the 
parent and those managed by a central system (hub).

## Defining Fields

Defining DTOs, entities, fields, and actions in Fireback involves specifying 
fields. A field is an array where you define the necessary fields, and this 
structure is used consistently across different definitions. Each field 
typically consists of a name and a type, with types being similar to Golang 
types. For example, you can use `int64` for integers and `float64` for 
floating-point numbers, which are translated to equivalent types in Java, 
TypeScript, etc.

This tutorial aims to provide you with a comprehensive understanding of module 
tree definitions in Fireback, helping you structure your system effectively. 
Stay tuned as we explore each section in detail.


## Defining Actions

Actions, are generally representing http calls, with a body, and a response type.
Fireback follows google json styleguide for errors and data.

```yaml
name: test
path: test

actions:

  # name the action
  - name: importUser

    # Set a url for http path
    url: /user/import

    # How that action would be named in cli action list
    cliName: userImport

    # similar to http methods, lower case.
    method: post

    # description, which will be appear on cli, or some other places.
    description: Imports users, and creates their passports, and all details

    # Format of the response, is an envelop, which would modify the response
    # using google json styleguide.
    # When using query, it will be returning as data.items (array)
    format: query

    # Define the request body fields. Makes sense on non-get methods
    in:
      fields:
      - name: path
        type: string

    # define the response body. Now understand the response body will 
    # be covered in a Google json styleguide by default
    # if you want to return an array, use format: query
    out:
      dto: OkayResponseDto
```

### Defining 'in' and 'out' fields

When defining actions, we can specify `in` or `out` properties. We have 3 options:

- Use `dto: Dtonamehere` for that.
- Use `entity: entitynamehere` for entities
- Defining the fields directly in, `fields: ` and normal fireback definition of fields.
Fireback will generate automatically ${Name}ActionReqDto and ${Name}ActionResDto
if necessary for all the targets.

**Note:** It does not make sense to define more than one of the options at the same time,
so entity has higher priority, and then dto, and then fields.