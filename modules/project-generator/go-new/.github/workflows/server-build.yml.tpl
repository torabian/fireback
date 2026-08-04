name: Build server flows
env:
  CI: false
  TERM: dumb

permissions:
  contents: write
  pages: write
  id-token: write

concurrency:
  group: "pages"
  cancel-in-progress: true

on:
  push:
    branches: 
      - "v**"
      - "main"
  workflow_dispatch:
    inputs:
      # target:
      #   type: choice
      #   options:
      #     - ubuntu-x64
      #     - windows-x64
      #     - macosx
      release_artifacts:
        description: "Create a draft release with all artifacts"
        required: false
        default: false
        type: boolean
      github_pages:
        description: "Deploy documents and projectname react to github pages"
        required: false
        default: false
        type: boolean
      projectname_windows:
        description: "Build for windows"
        required: false
        default: true
        type: boolean
      projectname_macos:
        description: "Build for macos"
        required: false
        default: true
        type: boolean
      projectname_react:
        description: "Build projectname React"
        required: false
        default: true
        type: boolean
      projectname_react_native:
        description: "Build for react native (android)"
        required: false
        default: true
        type: boolean
jobs:

  deploy_github_release:
    needs:
      - build-ubuntu
    name: deploy_github_release
    if: ${{`{{`}} inputs.release_artifacts == true {{`}}`}}
    runs-on: ubuntu-latest
    steps:
      - name: checkout
        uses: actions/checkout@v2
        with:
          fetch-depth: 0

      - name: Get the version tag
        id: get_version_tag
        run: echo "tag_name=$(basename ${GITHUB_REF})" >> $GITHUB_ENV

      - name: release
        uses: actions/create-release@v1
        id: create_release
        with:
          draft: true
          prerelease: true
          release_name: ${{`{{`}} steps.version.outputs.version {{`}}`}}
          tag_name: ${{`{{`}} env.tag_name {{`}}`}}
          body_path: CHANGELOG.md
        env:
          GITHUB_TOKEN: ${{`{{`}} github.token {{`}}`}}

      - uses: actions/download-artifact@master
        with:
          name: artifacts-all
          path: artifacts-all


      - name: upload mac amd64 zip artifact
        uses: actions/upload-release-asset@v1
        env:
          GITHUB_TOKEN: ${{`{{`}} github.token {{`}}`}}
        with:
          upload_url: ${{`{{`}} steps.create_release.outputs.upload_url {{`}}`}}
          asset_path: artifacts-all/projectname_amd64_darwin.zip
          asset_name: projectname_amd64_darwin.zip
          asset_content_type: application/zip

      - name: upload mac arm64 zip artifact
        uses: actions/upload-release-asset@v1
        env:
          GITHUB_TOKEN: ${{`{{`}} github.token {{`}}`}}
        with:
          upload_url: ${{`{{`}} steps.create_release.outputs.upload_url {{`}}`}}
          asset_path: artifacts-all/projectname_arm64_darwin.zip
          asset_name: projectname_arm64_darwin.zip
          asset_content_type: application/zip
      - name: upload linux zip amd64
        uses: actions/upload-release-asset@v1
        env:
          GITHUB_TOKEN: ${{`{{`}} github.token {{`}}`}}
        with:
          upload_url: ${{`{{`}} steps.create_release.outputs.upload_url {{`}}`}}
          asset_path: artifacts-all/projectname_amd64_linux.zip
          asset_name: projectname_amd64_linux.zip
          asset_content_type: application/zip
      - name: upload linux zip arm64
        uses: actions/upload-release-asset@v1
        env:
          GITHUB_TOKEN: ${{`{{`}} github.token {{`}}`}}
        with:
          upload_url: ${{`{{`}} steps.create_release.outputs.upload_url {{`}}`}}
          asset_path: artifacts-all/projectname_arm64_linux.zip
          asset_name: projectname_arm64_linux.zip
          asset_content_type: application/zip
      - name: upload windows zip arm64
        uses: actions/upload-release-asset@v1
        env:
          GITHUB_TOKEN: ${{`{{`}} github.token {{`}}`}}
        with:
          upload_url: ${{`{{`}} steps.create_release.outputs.upload_url {{`}}`}}
          asset_path: artifacts-all/projectname_arm64_windows.zip
          asset_name: projectname_arm64_windows.zip
          asset_content_type: application/zip

      - name: upload windows zip amd64
        uses: actions/upload-release-asset@v1
        env:
          GITHUB_TOKEN: ${{`{{`}} github.token {{`}}`}}
        with:
          upload_url: ${{`{{`}} steps.create_release.outputs.upload_url {{`}}`}}
          asset_path: artifacts-all/projectname_amd64_windows.zip
          asset_name: projectname_amd64_windows.zip
          asset_content_type: application/zip

  build-ubuntu:
    # if: ${{`{{`}} inputs.target == 'ubuntu-x64' {{`}}`}}
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3

      - name: Set up Go
        uses: actions/setup-go@v4
        with:
          go-version: 1.22
      - name: Use Node.js 18
        uses: actions/setup-node@v3
        with:
          node-version: 18
      - name: Install the dependencies
        run: cd front-end && npm i --force
      - name: Build the UI
        run: cd cmd/projectname-server && make ui2
      - name: clone latest Fireback
        run: git clone https://github.com/torabian/fireback --depth=1
      - name: Build and package deb files
        run: make everything
      - uses: actions/upload-artifact@master
        with:
          name: artifacts-all
          path: artifacts/projectname-server-all

      - name: Build for github pages.
        run: cd front-end && npm run github
      - name: Deploy 🚀
        if: ${{`{{`}} inputs.github_pages == true {{`}}`}}
        uses: JamesIves/github-pages-deploy-action@v4
        with:
          folder: ./front-end/build # The folder the action should deploy.