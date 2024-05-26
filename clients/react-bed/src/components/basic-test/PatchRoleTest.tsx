import { useDeleteRole } from "../../sdk/modules/workspaces/useDeleteRole"
import { usePatchRole } from "../../sdk/modules/workspaces/usePatchRole"
import { usePostWorkspacePassportCheck } from "../../sdk/modules/workspaces/usePostWorkspacePassportCheck"

export function PatchRoleTest() {
    
    const { submit } = usePatchRole()
    submit({name: 'asd'}).then(res => {
        console.log(res.data?.name)
    })

    const { submit: submit2} = usePostWorkspacePassportCheck({})

    submit2({value: 'adasd'}).then(x => x.data?.exists)

    
 
    return <ul>

     </ul>

}