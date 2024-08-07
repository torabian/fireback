module.exports = {
  presets: ['module:@react-native/babel-preset'],
  plugins: [
    [
      'module-resolver',
      {
        alias: {
          src: './src',
          '@': './src/',
        },
      },
    ],
    'react-native-reanimated/plugin',
    'transform-decorators-legacy',
  ],
};
