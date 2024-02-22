const glob = require("glob");

function testFileLengths(dir: string) {
  glob(dir, {}, function (er: any, files: string[]) {
    if (er) {
      throw er;
    }
    for (const file of files) {
      const onDisk = file.replace("../../build/", "");
      if (onDisk.length > 30) {
        console.log(onDisk, onDisk.length);
      }
    }
  });
}

testFileLengths("../../build/**/*.*");
