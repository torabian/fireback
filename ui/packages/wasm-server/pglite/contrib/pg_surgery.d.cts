import { d as PGliteInterface } from '../pglite-DYCFVi62.cjs';

declare const pg_surgery: {
    name: string;
    setup: (_pg: PGliteInterface, _emscriptenOpts: any) => Promise<{
        bundlePath: URL;
    }>;
};

export { pg_surgery };
