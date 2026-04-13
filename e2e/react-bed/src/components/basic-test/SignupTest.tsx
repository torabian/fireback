import { useState } from "react";

import { useClassicSigninAction, ClassicSigninActionReq } from "../../sdk/modules/abac/ClassicSignin";

export function SignupTest() {
  const [dto, setDto] = useState<ClassicSigninActionReq>(
    new ClassicSigninActionReq()
  );



  const { mutateAsync } = useClassicSigninAction({});

  const onComplete = () => {
    mutateAsync(dto)
      .then((res) => {
        alert("User has been created");
      })
      .catch((err) => {
        alert(err.toString());
      });
  };
  return (
    <div>
      <form onSubmit={onComplete}>
        <label>Email address or Phone Number</label>
        <input
          value={dto.value || ""}
          type="text"
          onChange={(e) =>
            setDto((dto: any) => {
              return {
                ...dto,
                value: e.target.value,
              };
            })
          }
        />

        <label>Password</label>
        <input
          value={dto.password || ""}
          type="password"
          onChange={(e) =>
            setDto((dto: any) => {
              return {
                ...dto,
                password: e.target.value,
              };
            })
          }
        />
        <button>Submit</button>
      </form>
    </div>
  );
}
