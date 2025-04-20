import { DatatableColumn } from "@/modules/fireback/definitions/definitions";
import { UserEntity } from "../../../sdk/modules/abac/UserEntity";
import { enTranslations } from "../../../translations/en";
import { GenderView } from "./GenderView";

export const columns = (t: typeof enTranslations): DatatableColumn[] => [
  {
    name: UserEntity.Fields.uniqueId,
    title: t.table.uniqueId,
    width: 100,
  },
  {
    name: "firstName",
    title: t.users.firstName,
    width: 200,
    sortable: true,
    filterable: true,
    getCellValue: (e: UserEntity) => e?.firstName,
  },

  {
    filterable: true,
    name: "lastName",
    sortable: true,
    title: t.users.lastName,
    width: 200,
    getCellValue: (e: UserEntity) => e?.lastName,
  },

  {
    name: "birthDate",
    title: "birthdate",
    width: 140,
    getCellValue: (e: UserEntity) => <>{e?.birthDate}</>,
    filterType: "date",
    filterable: true,
    sortable: true,
  },
  {
    name: "gender",
    title: "gender",
    width: 50,
    getCellValue: (e: UserEntity) => (
      <>
        <GenderView gender={e.gender} />
      </>
    ),
  },

  {
    name: "Image",
    title: "Image",
    width: 40,
    getCellValue: (e: UserEntity) => (
      <>
        {e?.photo && (
          <img src={e?.photo} style={{ width: "20px", height: "20px" }} />
        )}
      </>
    ),
  },
  {
    name: UserEntity.Fields.primaryAddress.countryCode,
    title: "Country code",
    width: 40,
    getCellValue: (e: UserEntity) => <>{e.primaryAddress?.countryCode}</>,
  },
  {
    name: UserEntity.Fields.primaryAddress.addressLine1,
    title: "Address Line 1",
    width: 180,
    getCellValue: (e: UserEntity) => <>{e.primaryAddress?.addressLine1}</>,
  },
  {
    name: UserEntity.Fields.primaryAddress.addressLine2,
    title: "Address Line 2",
    width: 180,
    getCellValue: (e: UserEntity) => <>{e.primaryAddress?.addressLine2}</>,
  },
  {
    name: UserEntity.Fields.primaryAddress.city,
    title: "City",
    width: 180,
    getCellValue: (e: UserEntity) => <>{e.primaryAddress?.city}</>,
  },
  {
    name: UserEntity.Fields.primaryAddress.postalCode,
    title: "Postal Code",
    width: 80,
    getCellValue: (e: UserEntity) => <>{e.primaryAddress?.postalCode}</>,
  },
];
