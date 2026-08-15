import { d as PGliteInterface } from '../pglite-DYCFVi62.js';

declare const pg_hashids: {
    name: string;
    setup: (_pg: PGliteInterface, emscriptenOpts: any) => Promise<{
        emscriptenOpts: any;
        bundlePath: URL;
    }>;
};

export { pg_hashids };
