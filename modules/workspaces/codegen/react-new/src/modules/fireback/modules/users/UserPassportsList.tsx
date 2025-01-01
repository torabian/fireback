import { GeneralEntityView } from "../../components/general-entity-view/GeneralEntityView";
import { PageSection } from "../../components/page-section/PageSection";
import { PassportEntity } from "../../sdk/modules/workspaces/PassportEntity";
import { useGetPassports } from "../../sdk/modules/workspaces/useGetPassports";
import Link from "../../components/link/Link";

export const UserPassportList = ({ userId }: { userId: string }) => {
  const { items } = useGetPassports({
    query: {
      query: userId ? "user_id = " + userId : null,
    },
  });

  return (
    <div>
      <Link href={"/passport"}>Add passport</Link>
      <PageSection title="Passports">
        {items.map((item) => {
          return <UserPassportItem passport={item} key={item.uniqueId} />;
        })}
      </PageSection>
    </div>
  );
};

function booleanToHuman(value?: boolean): string {
  if (value === null || value === undefined) {
    return "n/a";
  }

  if (value === true) {
    return "Yes";
  }

  if (value === false) {
    return "No";
  }
}

const UserPassportItem = ({ passport }: { passport: PassportEntity }) => {
  return (
    <div>
      <div className="general-entity-view ">
        <div className="entity-view-row entity-view-head">
          <div className="field-info">Value:</div>
          <div className="field-value">{passport.value}</div>
        </div>
        <div className="entity-view-row entity-view-head">
          <div className="field-info">Type:</div>
          <div className="field-value">{passport.type}</div>
        </div>
        <div className="entity-view-row entity-view-head">
          <div className="field-info">Confirmed:</div>
          <div className="field-value">
            {booleanToHuman(passport.confirmed)}
          </div>
        </div>
      </div>
    </div>
  );
};
