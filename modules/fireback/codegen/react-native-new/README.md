# Fireback React Native boilerplate and starter

Complete react native starter project, and tools I build over time for react native clients. Project should be able to run independently
without fireback backend. You can clone it, and use it. However, if you build backend also with fireback you would be completing project.
20x-30x times faster than writing these binding yourself.

## Core principles

**This is a boilerplate, not a library*: Only fireback backend (Go, Java, C++) are meant to be libraries as well as boilerplate. React Native for fireback, in contrary is a one time clone, if you need update, you need to check
them manually. This is much better for stability of your project, as well as our maintenance costs.


This project aims more on the typescript part of the application rather than maintaining the native libraries, however we might use some different libraries, which would be 99% used in many projects, such as 
taking photo, navigation, etc.

React native is a great framework for building wide scenario of apps, nevertheless my focus and this project
focus is on native development due to requirement of more trafficed apps.

If you don't know about fireback, and you want to start an app, read the fireback main page, it might be a
game changer for you.


## Features and screens

- [ ] Signup, using email, and phone number
- [ ] OTA based authetication
- [ ] Services to keep the authenticated setup
- [ ] User profile, update the picture setup
- [ ] Theming using fireback
- [ ] Translations using fireback tools
- [ ] Instructions for putting the app into store
- [ ] InfiniteList component
- [ ] CommonX component (CRUD + advanced items)

## Using without fireback

This repository could be starter for any type of project, even if you did not want build backend with fireback.
Try to write the hooks for http and socket manually yourself in that case, or reorgnize the structure of the project.

