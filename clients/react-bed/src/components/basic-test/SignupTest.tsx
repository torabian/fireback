import { useState } from "react";
import { ClassicSignupActionReqDto } from "../../sdk/modules/workspaces/WorkspacesActionsDto";
import { usePostPassportsSigninClassic } from "../../sdk/modules/workspaces/usePostPassportsSigninClassic";

export function SignupTest() {
    const [dto, setDto] = useState<ClassicSignupActionReqDto>(new ClassicSignupActionReqDto())

    const { submit } = usePostPassportsSigninClassic({});

    const onComplete = () => {

        submit(dto).then(res => {
            alert("User has been created")
        }).catch(err => {
            alert(err.toString())
        })

    }
    return <div>

        <form onSubmit={onComplete}>
            <label>Email address or Phone Number</label>
            <input  value={dto.value || ''} type="text" onChange={e => setDto(dto => {
                return {
                    ...dto,
                    value: e.target.value
                }
            })} />


            <label>Password</label>
            <input value={dto.password || ''} type="password" onChange={e => setDto(dto => {
                return {
                    ...dto,
                    password: e.target.value
                }
            })} />
            <button>Submit</button>
        </form>
    </div>

}