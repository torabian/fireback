import { PageSection } from "../../../components/page-section/PageSection";
import { useS } from "../../../hooks/useS";
import { PassportEntity } from "../../../sdk/modules/abac/PassportEntity";
import { useGetPassports } from "../../../sdk/modules/abac/useGetPassports";
import { strings } from "./strings/translations";

export const UserPassportList = ({ userId }: { userId: string }) => {
  const { items } = useGetPassports({
    query: {
      query: userId ? "user_id = " + userId : null,
    },
  });
  const s = useS(strings);

  return (
    <div>
      {/* <Link href={"/passport"}>Add passport</Link> */}
      <PageSection title={s.passports}>
        {(items || []).map((item) => {
          return <UserPassportItem passport={item} key={item.uniqueId} s={s} />;
        })}
      </PageSection>
    </div>
  );
};

function booleanToHuman(value: boolean | undefined, s: typeof strings): string {
  if (value === null || value === undefined) {
    return s.na;
  }

  if (value === true) {
    return s.yes;
  }

  if (value === false) {
    return s.no;
  }
}

const UserPassportItem = ({
  passport,
  s,
}: {
  passport: PassportEntity;
  s: typeof strings;
}) => {
  return (
    <div>
      <div className="general-entity-view ">
        <div className="entity-view-row entity-view-head">
          <div className="field-info">{s.value}</div>
          <div className="field-value">{passport.value}</div>
        </div>
        <div className="entity-view-row entity-view-head">
          <div className="field-info">{s.type}</div>
          <div className="field-value">{passport.type}</div>
        </div>
        <div className="entity-view-row entity-view-head">
          <div className="field-info">{s.confirmed}</div>
          <div className="field-value">
            {booleanToHuman(passport.confirmed, s)}
          </div>
        </div>
      </div>
    </div>
  );
};
