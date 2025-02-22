const fs = require('fs');

const tests = [
    {
        file: './src/sdk/modules/workspaces/AcceptInviteDto.ts',
        toContain: [
            'inviteUniqueId',
            'public static definition',
            'public static Fields = {'
        ]
    },
    {
        file: './src/sdk/modules/workspaces/useDeleteCapability.ts',
        toContain: [
            'export function useDeleteCapability(props?: DeleteProps) {',
            'const url = "/capability".substr(1);',
            '): Promise<IDeleteResponse> => {',
            '"*workspaces.CapabilityEntity",'
        ]
    },
    {
        file: './src/sdk/modules/workspaces/useGetPassports.ts',
        toContain: [
            'const fn = () => rpcFn("GET", computedUrl);',
            'const query$ = useQuery<any, any, IResponseList<PassportEntity>, any>(["*workspaces.PassportEntity", computedOptions, query], fn, {',
            'useGetPassports.UKEY = "*workspaces.PassportEntity"',
            'const url = "/passports".substr(1);'
        ]
    },
]

function runTest() {

    for (let test of tests) {
        // check for file, if not exists, we have a problem
        console.log(test.file, fs.existsSync(test.file))
        if (fs.existsSync(test.file) === false) {
            console.log("Unfortunately during the test file:", test.file, " was not found.") 
            process.exit(1)
        }

        if (test.toContain) {
            const body = fs.readFileSync(test.file).toString()

            if (Array.isArray(test.toContain)) {

                for (let t of test.toContain) {
                    if (body.indexOf(t) === -1) {
                        console.log("Unfortunately we expected to find '", t, "' in the file:", test.file, "but was not found")
                    }
                }
            }

        }
    }

}

runTest();