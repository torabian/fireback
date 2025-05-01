package firebackgonew

import "embed"

//go:embed app.json yarn.lock README.md tsconfig.json metro.config.js package.json jest.config.js index.js Gemfile.lock Gemfile babel.config.js App.tsx app.json .yarnrc .watchmanconfig .prettierrc.js .gitignore .eslintrc.js .vscode ios/firebackrn ios/firebackrn.xcodeproj ios/firebackrn.xcworkspace ios/firebackrnTests ios/.xcode.env ios/Podfile ios/Podfile.lock src __tests__ android/app/src android/app/build.gradle android/app/debug.keystore android/app/proguard-rules.pro android/gradle/wrapper/gradle-wrapper.jar android/gradle/wrapper/gradle-wrapper.properties android/build.gradle android/gradle.properties android/gradlew android/gradlew.bat android/settings.gradle
var FbReactNativeNewTemplate embed.FS

/**
*	This directory includes a boilerplate for building react.js apps,
*	it's configured in a way that uses fireback as backend.
*   Nevertheless, it's not forced at all to use backend for fireback.
*	You can remove src/modules/fireback folder, adjust App.tsx a bit
*	and start your fully pure project in there.
 */
